package repository

import (
	"context"

	"github.com/Asad2730/EventManagement/connect"
	"github.com/Asad2730/EventManagement/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var tableName = "Event"

func CreateEvent(event model.Event, tableName string, av map[string]types.AttributeValue) (*dynamodb.PutItemOutput, error) {

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	res, err := connect.Client.PutItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetEvents(tableName string) (*dynamodb.ScanOutput, error) {

	scan := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	rs, err := connect.Client.Scan(context.TODO(), scan)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetEvent(id, tableName string) (*dynamodb.GetItemOutput, error) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	rs, err := connect.Client.GetItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil

}

func UpdateEvent(id, tableName string, av map[string]types.AttributeValue) (*dynamodb.UpdateItemOutput, error) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       key,
		ExpressionAttributeValues: av,
		UpdateExpression:          aws.String("SET title = :t, description = :d, EventDate = :e"),
		ReturnValues:              types.ReturnValueAllNew,
	}

	rs, err := connect.Client.UpdateItem(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func DeleteEvent(id, tableName string) (*dynamodb.DeleteItemOutput, error) {

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
