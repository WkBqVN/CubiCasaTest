package controller

import (
	"cubicasa/models"
	"cubicasa/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HubHandler struct {
	Service service.HubService
}

func (hubHandler *HubHandler) GetAll(c *gin.Context) {
	listHub, isDone := hubHandler.Service.GetAllHub()
	if isDone == true {
		c.JSON(http.StatusOK, listHub)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Service can't get list",
		})
	}
}
func (hubHandler *HubHandler) GetHubById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		hub, isDone := hubHandler.Service.GetHubById(id)
		if isDone == true {
			c.JSON(http.StatusOK, hub)
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"Message": "Can't get Hub or hub id is not exits",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (hubHandler *HubHandler) CreateHub(c *gin.Context) {
	jsonObj := models.Hub{}
	err := c.ShouldBind(&jsonObj)
	if err == nil {
		_, isDone := hubHandler.Service.CreateHub(jsonObj)
		fmt.Println(isDone)
		if isDone == true {
			c.JSON(http.StatusCreated, gin.H{
				"Message": "Hub is created",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't create hub",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Json not valid",
		})
	}
}

func (hubHandler *HubHandler) UpdateHub(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		jsonObj := models.Hub{}
		err := c.ShouldBind(&jsonObj)
		if err == nil {
			_, isDone := hubHandler.Service.UpdateHub(id, jsonObj)
			if isDone == true {
				c.JSON(http.StatusOK, "Updated Hub "+c.Param("id"))
			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"Message": "Can't update hub or hub is not exits" + c.Param("id"),
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

func (hubHandler *HubHandler) DeleteHub(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		_, isDone := hubHandler.Service.DeleteHub(id)
		if isDone == true {
			c.JSON(http.StatusOK, "Deleted hub :"+c.Param("id"))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Can't Delete hub or hub not exits" + c.Param("id"),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ID not valid",
		})
	}
}
