package account

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/hvxahv/hvx/pkg/v"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"strconv"
)

// Accounts is a struct for account.
type Accounts struct {

	// AccountID is a unique identifier for the account.
	gorm.Model

	// Username is the primaryKey of the database which is unique,
	// in the account system of this instance is must be added during
	// the creation process to ensure the correctness of the data.
	Username string `gorm:"primaryKey;type:text;username;unique" validate:"required,min=4,max=16"`

	// Mail When registering, the user is required to provide an email,
	// and an error needs to be returned when the email is not in the
	// correct format. It is also unique in the account system.
	Mail string `gorm:"index;type:text;mail;unique" validate:"required,email"`

	// Password must be encrypted and saved. The length of the password needs to be verified
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// ActorID is used for compatibility with the ActivityPub protocol
	// to connect to the actor table by ID.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// IsPrivate sets whether the account is private or not,
	// it is a social extension that is set by the user to make the
	// account public or not.
	IsPrivate bool `gorm:"type:boolean;is_private"`
}

func (a *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ? ", in.Username).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return &pb.IsExistResponse{IsExist: ok}, nil
	}
	return &pb.IsExistResponse{IsExist: false}, nil
}

func (a *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if err := validator.New().Struct(in); err != nil {
		return nil, errors.New("FAILED_TO_VALIDATOR")
	}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACTOR_DATABASE")
	}

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACCOUNT_DATABASE")
	}

	if err := db.Debug().
		Table("accounts").
		Where("username = ? ", in.Username).Or("mail = ?", in.Mail).
		First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return nil, errors.New("THE_USERNAME_OR_MAIL_ALREADY_EXISTS")
		}
	}

	n := NewActors(in.Username, in.PublicKey, "Person")
	if err := db.Debug().Table("actors").Create(&n).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	v := NewAccounts(n.ID, in.Username, in.Mail, in.Password)
	if err := db.Debug().Table("accounts").Create(&v).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	return &pb.CreateAccountResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) GetAccountByUsername(ctx context.Context, in *pb.GetAccountByUsernameRequest) (*pb.GetAccountByUsernameResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	return &pb.GetAccountByUsernameResponse{
		AccountId: strconv.Itoa(int(a.Accounts.ID)),
		Username:  a.Accounts.Username,
		Mail:      a.Accounts.Mail,
		Password:  a.Accounts.Password,
		ActorId:   strconv.Itoa(int(a.Accounts.ActorID)),
		IsPrivate: strconv.FormatBool(a.Accounts.IsPrivate),
	}, nil
}

func (a *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	username, err := v.GetUsernameByTokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	v := NewAuthorization(username, in.Password)

	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("username = ?", username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	if err := db.Debug().Table("actors").Where("id = ?", v.ActorID).Unscoped().Delete(&Actors{}).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Table("accounts").Where("id = ?", v.ID).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return nil, err
	}

	// TODO - Delete Account related data.

	return &pb.DeleteAccountResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) EditUsername(ctx context.Context, in *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	exist, err := a.IsExist(ctx, &pb.IsExistRequest{Username: in.Username})
	if err != nil {
		return nil, err
	}

	// If the username is Exist, return error.
	if !exist.IsExist {
		return &pb.EditUsernameResponse{Code: "401", Reply: "THE_USERNAME_ALREADY_EXISTS"}, nil
	}

	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("id = ?", uint(id)).First(&a.Accounts).Update("username", in.Username).Error; err != nil {
		return &pb.EditUsernameResponse{Code: "500", Reply: err.Error()}, err
	}

	address := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), in.Username)
	inbox := fmt.Sprintf("%s/inbox", address)
	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).
		Update("preferred_username", in.Username).
		Update("inbox", inbox).
		Update("address", address).Error; err != nil {
		return &pb.EditUsernameResponse{Code: "500", Reply: err.Error()}, err
	}

	return &pb.EditUsernameResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) EditPassword(ctx context.Context, in *pb.EditPasswordRequest) (*pb.EditPasswordResponse, error) {
	v := NewAuthorization(in.Username, in.Password)

	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(in.New), bcrypt.DefaultCost)
	if err := db.Debug().Table("accounts").Where("id = ?", v.ID).Update("password", hash).Error; err != nil {
		return nil, err
	}

	// TODO - Edit Account related data.
	return &pb.EditPasswordResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) EditEmail(ctx context.Context, in *pb.EditEmailRequest) (*pb.EditEmailResponse, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.Debug().
		Table("accounts").
		Where("id = ?", id).
		Update("mail", in.Mail).
		Error; err != nil {
		return nil, err
	}

	return &pb.EditEmailResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) GetAccountCount(ctx context.Context, g *emptypb.Empty) (*pb.GetAccountCountResponse, error) {
	db := cockroach.GetDB()
	var count int64
	if err := db.Debug().Table("accounts").Count(&count).Error; err != nil {
		return nil, err
	}

	return &pb.GetAccountCountResponse{
		Code:         "200",
		AccountCount: strconv.Itoa(int(count)),
	}, nil
}

func NewAccounts(actorID uint, username, mail, password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: string(hash),
		ActorID:  actorID,
	}
}
