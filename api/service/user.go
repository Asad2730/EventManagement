package service

import (
	"github.com/Asad2730/EventManagement/model"
	"github.com/Asad2730/EventManagement/repository"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var table = "User"

func CreateUser(c *gin.Context) {

	var user model.User
	user.Id = uuid.NewString()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	output, err := repository.CreateUser(user, table, av)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, output)
}

func GetUser(c *gin.Context) {

	id := c.Param("id")

	var user model.User

	res, err := repository.GetUser(id, table)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	err = attributevalue.UnmarshalMap(res.Item, &user)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, user)
}

func GetUsers(c *gin.Context) {

	res, err := repository.GetUsers(table)

	if err != nil {
		c.JSON(500, err.Error())
	}

	var users []model.User

	for _, i := range res.Items {
		var user model.User
		if err := attributevalue.UnmarshalMap(i, &user); err != nil {
			c.JSON(500, err.Error())
			return
		}

		users = append(users, user)
	}

	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	rs, err := repository.UpdateUser(id, table, av)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, rs)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	res, err := repository.DeleteUser(id, table)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
}
