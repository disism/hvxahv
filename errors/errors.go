package errors

import "fmt"

const (
	ErrAccountVerification = "ACCOUNT_VERIFICATION_FAILED"
	ErrAccountAlready      = "THE_USERNAME_OR_MAIL_ALREADY_EXISTS"
	ErrAccountCreate       = "FAILED_TO_CREATE_ACCOUNT"
	ErrActorDelete         = "FAILED_TO_DELETE_ACTOR"
)

const (
	ErrAccountDatabaseCreate = "FAILED_TO_AUTOMATICALLY_CREATE_ACCOUNT_DATABASE"
)

func New(err string) error {
	return fmt.Errorf(err)
}

func Newf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
