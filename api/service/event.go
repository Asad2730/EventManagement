package service

import (
	"github.com/Asad2730/EventManagement/model"
	"github.com/Asad2730/EventManagement/repository"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var tableName = "Event"

func CreateEvent(c *gin.Context) {

	var event model.Event
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

	output, err := repository.CreateEvent(event, tableName, av)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, output)
}

func GetEvent(c *gin.Context) {

	id := c.Param("id")

	var event model.Event

	res, err := repository.GetEvent(id, table)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	err = attributevalue.UnmarshalMap(res.Item, &event)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, event)
}

func GetEvents(c *gin.Context) {

	res, err := repository.GetEvents(table)

	if err != nil {
		c.JSON(500, err.Error())
	}

	var events []model.Event

	for _, i := range res.Items {
		var event model.Event
		if err := attributevalue.UnmarshalMap(i, &event); err != nil {
			c.JSON(500, err.Error())
			return
		}

		events = append(events, event)
	}

	c.JSON(200, events)
}

func UpdateEvent(c *gin.Context) {

	id := c.Param("id")
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

	rs, err := repository.UpdateEvent(id, table, av)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, rs)
}

func DeleteEvent(c *gin.Context) {

	id := c.Param("id")
	res, err := repository.DeleteEvent(id, table)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
}
