package connect

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	Client     *dynamodb.Client
	Uploader   *manager.Uploader
	Downloader *manager.Downloader
)

func ConnectAws() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Error connecting to AWS:", err.Error())
	}

	client := s3.NewFromConfig(cfg)

	db := dynamodb.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)

	Client = db
	Uploader = uploader
	Downloader = downloader
}
