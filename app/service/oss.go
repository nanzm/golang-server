package service

import (
	"dora/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type OssService interface {
	Put(file io.Reader, fileName string) error
}

type ossService struct {
	cof *config.AliyunConfig
}

func NewOssService(conf *config.AliyunConfig) OssService {
	return &ossService{
		cof: conf,
	}
}

func (o *ossService) Put(file io.Reader, fileName string) error {
	client, err := oss.New(o.cof.Endpoint, o.cof.AccessKey, o.cof.Secret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(o.cof.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(fileName, file)
	if err != nil {
		return err
	}

	return nil
}
