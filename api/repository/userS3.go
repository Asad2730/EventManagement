package repository

import (
	"context"
	"io"

	"github.com/Asad2730/EventManagement/connect"
	"github.com/Asad2730/EventManagement/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	TableName  = "User"
	BucketName = "userPic"
)

func CreateUserS3(c *gin.Context) {

	var event model.User
	event.Id = uuid.NewString()
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}

	av, err := attributevalue.MarshalMap(event)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      av,
	}

	_, err = connect.Client.PutItem(context.TODO(), input)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	_, err = connect.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(event.Id),
	})

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func GetUserS3(c *gin.Context) {

	id := c.Param("id")

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	}

	out, err := connect.Client.GetItem(context.TODO(), getItemInput)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if out == nil {
		c.JSON(404, "User Not Found")
		return
	}

	var user model.User
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		c.JSON(500, err.Error())
		return
	}

	var buf io.WriterAt

	obj, err := connect.Downloader.Download(context.TODO(), buf, &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(user.Id),
	})

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": user, "image": obj})

}
