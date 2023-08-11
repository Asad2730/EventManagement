package controller

import (
	"context"

	"github.com/Asad2730/EventManagement/connect"
	"github.com/Asad2730/EventManagement/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var tableName = "Event"

func Create(c *gin.Context) {

	var event model.Event
	event.Id = uuid.New()
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
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = connect.Client.PutItem(context.TODO(), input)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func GetAll(c *gin.Context) {

	scan := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	rs, err := connect.Client.Scan(context.TODO(), scan)

	if err != nil {
		c.JSON(500, err.Error())
	}

	var events []model.Event

	for _, i := range rs.Items {

		var y model.Event
		if err := attributevalue.UnmarshalMap(i, &y); err != nil {
			c.JSON(500, err.Error())
		}

		events = append(events, y)
	}

	c.JSON(200, events)
}

func GetById(c *gin.Context) {

	id := c.Param("id")
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
		c.JSON(500, err.Error())
		return
	}

	var event model.Event

	err = attributevalue.UnmarshalMap(rs.Item, &event)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func Update(c *gin.Context) {

	id := c.Param("id")
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, err.Error())
		return
	}

	av, err := attributevalue.MarshalMap(event)
	if err != nil {
		c.JSON(500, err.Error())
		return
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
		c.JSON(500, err.Error())
		return
	}

	var updatedResponse model.Event
	if err := attributevalue.UnmarshalMap(rs.Attributes, &updatedResponse); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, updatedResponse)
}

func Delete(c *gin.Context) {

	id := c.Param("id")
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: id,
		},
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	_, err := connect.Client.DeleteItem(context.TODO(), input)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "Item Deleted")
}
