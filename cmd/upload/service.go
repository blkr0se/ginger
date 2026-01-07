package upload

import (
	"context"
	"fmt"
	"os"

	"github.com/blkr0se/shun"
)

type uploadRequest struct {
	Src string
	Dst string
}

func NewUploadRequest(src, dst string) *uploadRequest {
	return &uploadRequest{Src: src, Dst: dst}
}

type uploadService struct {
	Uploader *shun.SftpUploader
}

func NewUploadService(c *shun.SshConnector) (*uploadService, error) {
	conn, err := c.Connect()
	if err != nil {
		return nil, err
	}

	uploader := shun.NewSftpUploader(conn)

	return &uploadService{Uploader: uploader}, nil
}

func (us *uploadService) Upload(ctx context.Context, req *uploadRequest) error {
	info, err := os.Stat(req.Src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		// TODO: uploader.CompressDir(provider.Source())
		return fmt.Errorf("directory upload not implemented yet")
	}

	data, err := os.ReadFile(req.Src)
	if err != nil {
		return err
	}

	if req.Dst == "" {
		req.Dst = req.Src
	}

	return us.Uploader.Upload(data, req.Dst)
}
