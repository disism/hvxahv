package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/fs"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/fs"
	"github.com/hvxahv/hvx/microsvc"
	"log"
)

// AvatarHandler Uploads an avatar to the object store and returns the address (name) of the image.
// The size and format of the image needs to be calibrated.
// Only PNG/JPEG format images are accepted as avatars.
// After uploading, the image address and uploader account ID are sent to the fs server
// for storage and used as file ownership verification.
func AvatarHandler(c *gin.Context) {
	parse, _ := ParseAuthorization(c.Request.Header.Get("Authorization"))
	f, err := c.FormFile("avatar")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	}
	if f.Size > 1024*1024*2 {
		c.JSON(403, errors.NewHandler("403", errors.ErrFilesSize))
		return
	}
	var (
		ft = f.Header.Get("Content-Type")
		fn = fmt.Sprintf("%s-%s", uuid.New().String(), f.Filename)
	)
	if ft != "image/png" && ft != "image/jpeg" {
		c.JSON(403, errors.NewHandler("403", errors.ErrFilesAllowedPNGAndJPEG))
		return
	}

	reader, _ := f.Open()
	put, err := fs.NewFsPut("minio", "avatar", fn, ft, reader, f.Size).Put()
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut+err.Error()))
		return
	}

	client, err := clientv1.New(c, microsvc.NewGRPCAddress("fs").Get())
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrConnectFsRPCServer))
		return
	}
	defer client.Close()
	create, err := pb.NewFsClient(client.Conn).Create(c, &pb.CreateRequest{
		AccountId: parse.AccountId,
		FileName:  fn,
		Address:   put,
	})
	if err != nil && create.Code != "200" {
		c.JSON(500, errors.NewHandler("500", errors.ErrInternalServer))
		return
	}
	c.JSON(200, gin.H{
		"code":   create.Code,
		"name":   fn,
		"avatar": put,
	})
}

type Attas struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

// AttachHandler The upload of attachment files is designed in the same way as the avatar handler,
// but the method allows uploading multiple files at the same time and not only in image format,
// and is designed for a maximum of 9 files per upload.
func AttachHandler(c *gin.Context) {
	parse, _ := ParseAuthorization(c.Request.Header.Get("Authorization"))
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
		return
	}
	var idx int
	for i, _ := range form.File["attach"] {
		idx = i + 1
	}
	if idx > 9 {
		c.JSON(403, errors.NewHandler("403", errors.ErrFileMaximum))
		return
	}
	client, err := clientv1.New(c, microsvc.NewGRPCAddress("fs").Get())
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrConnectFsRPCServer))
		return
	}
	defer client.Close()

	var a []*Attas
	for _, f := range form.File["attach"] {
		var (
			ft = f.Header.Get("Content-Type")
			fn = fmt.Sprintf("%s-%s", uuid.New().String(), f.Filename)
		)
		reader, _ := f.Open()
		put, err := fs.NewFsPut("minio", "attach", fn, ft, reader, f.Size).Put()
		if err != nil {
			c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
			return
		}
		create, err := pb.NewFsClient(client.Conn).Create(c, &pb.CreateRequest{
			AccountId: parse.AccountId,
			FileName:  fn,
			Address:   put,
		})
		if err != nil && create.Code != "200" {
			c.JSON(500, errors.NewHandler("500", errors.ErrInternalServer))
			return
		}
		a = append(a, &Attas{
			Name:    fn,
			Address: put,
		})
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"attas": a,
	})
}

// DeleteFsHandler The handler to delete the file.
// It will delete the file from the object storage server with the caller's account ID and file name,
// and then send it to the fs server to delete the data from the database.
func DeleteFsHandler(c *gin.Context) {
	parse, _ := ParseAuthorization(c.Request.Header.Get("Authorization"))
	fn := c.PostForm("name")

	if err := fs.NewFsDelete("minio", "attach", fn).Delete(); err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrFilesPut))
		return
	}

	client, err := clientv1.New(c, microsvc.NewGRPCAddress("fs").Get())
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrConnectFsRPCServer))
		return
	}
	defer client.Close()
	a, err := pb.NewFsClient(client.Conn).Delete(c, &pb.DeleteRequest{
		AccountId: parse.AccountId,
		FileName:  fn,
	})
	if err != nil && a.Code != "200" {
		c.JSON(500, errors.NewHandler("500", errors.ErrFileDelete+err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"code":   a.Code,
		"status": "ok",
	})
}

func GetFsAddressHandler(c *gin.Context) {
	parse, _ := ParseAuthorization(c.Request.Header.Get("Authorization"))
	fn := c.Param("name")
	client, err := clientv1.New(c, microsvc.NewGRPCAddress("fs").Get())
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrConnectFsRPCServer))
		return
	}
	defer client.Close()
	f, err := pb.NewFsClient(client.Conn).Get(c, &pb.GetRequest{
		AccountId: parse.AccountId,
		FileName:  fn,
	})
	if err != nil && f.Code != "200" {
		c.JSON(500, errors.NewHandler("500", errors.ErrFileDelete+err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"code":    f.Code,
		"name":    f.FileName,
		"address": f.Address,
	})
}
