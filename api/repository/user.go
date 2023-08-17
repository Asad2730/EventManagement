package repository

import (
	"context"

	"github.com/Asad2730/EventManagement/connect"
	"github.com/Asad2730/EventManagement/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateUser(user model.User, tableName string, av map[string]types.AttributeValue) (*dynamodb.PutItemOutput, error) {

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	output, err := connect.Client.PutItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetUser(id, tableName string) (*dynamodb.GetItemOutput, error) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	res, err := connect.Client.GetItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func GetUsers(tableName string) (*dynamodb.ScanOutput, error) {

	scan := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	res, err := connect.Client.Scan(context.TODO(), scan)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateUser(id, tableName string, av map[string]types.AttributeValue) (*dynamodb.UpdateItemOutput, error) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       key,
		ExpressionAttributeValues: av,
		UpdateExpression:          aws.String("SET email = :e, name = :n, eventId = :eid"),
		ReturnValues:              types.ReturnValueAllNew,
	}

	res, err := connect.Client.UpdateItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteUser(id, tableName string) (*dynamodb.DeleteItemOutput, error) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	res, err := connect.Client.DeleteItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return res, nil
}
