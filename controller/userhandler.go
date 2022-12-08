package controller

import (
	"cubicasa/models"
	"cubicasa/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service service.UserService
}

func (userHandler *UserHandler) GetAll(c *gin.Context) {
	listUser, isDone := userHandler.Service.GetAllUser()
	if isDone == true {
		c.JSON(http.StatusOK, listUser)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Service can't get list",
		})
	}
}
func (userHandler *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		user, isDone := userHandler.Service.GetUserById(id)
		if isDone == true {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't get User or user id not exits " + c.Param("id"),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (userHandler *UserHandler) CreateUser(c *gin.Context) {
	jsonObj := models.User{}
	err := c.ShouldBind(&jsonObj)
	if err == nil {
		_, isDone := userHandler.Service.CreateUser(jsonObj)
		if isDone == true {
			c.JSON(http.StatusCreated, gin.H{
				"Message": "User Created",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't create user",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Json is not valid",
		})
	}
}

func (userHandler *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		jsonObj := models.User{}
		err := c.ShouldBind(&jsonObj)
		if err == nil {
			_, isDone := userHandler.Service.UpdateUser(id, jsonObj)
			if isDone == true {
				c.JSON(http.StatusOK, gin.H{
					"Message": "Updated User: " + c.Param("id"),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Can't update user or user not exits " + c.Param("id"),
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Json not valid",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (userHandler *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		_, isDone := userHandler.Service.DeleteUser(id)
		fmt.Println(isDone)
		if isDone == true {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Deleted User: " + c.Param("id"),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't Delete user or user not exits" + c.Param("id"),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (userHandler *UserHandler) MoveUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	isValid := true
	if err != nil {
		isValid = false
	}
	teamId, err := strconv.Atoi(c.Query("teamId"))
	if err != nil {
		isValid = false
	}
	if isValid {
		isMoved := userHandler.Service.MoveUser(userId, teamId)
		if isMoved == true {
			c.JSON(http.StatusOK, gin.H{
				"Message": "User " + c.Query("userId") + " is moved to team " + c.Query("teamId"),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't move  user " + c.Query("userId") + " to new team" + c.Param("teamId"),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "User id or Team id is not valid",
		})
	}
}

func (userHandler *UserHandler) SearchUser(c *gin.Context) {
	userName := c.Query("name")
	result, isDone := userHandler.Service.SearchUser(userName)
	if isDone == true {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "name you search are not valid or name not exits in database",
		})
	}
}
