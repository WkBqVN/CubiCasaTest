package controller

import (
	"cubicasa/models"
	"cubicasa/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	Service service.TeamService
}

func (teamHandler *TeamHandler) GetAll(c *gin.Context) {
	listTeam, isDone := teamHandler.Service.GetAllTeam()
	if isDone == true {
		c.JSON(http.StatusOK, listTeam)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Service can't get list",
		})
	}
}

func (teamHandler *TeamHandler) GetTeamById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		team, isDone := teamHandler.Service.GetTeamById(id)
		if isDone == true {
			c.JSON(http.StatusOK, team)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't get team or team id is not exits",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (teamHandler *TeamHandler) CreateTeam(c *gin.Context) {
	jsonObj := models.Team{}
	err := c.ShouldBind(&jsonObj)
	if err == nil {
		_, isDone := teamHandler.Service.CreateTeam(jsonObj)
		if isDone == true {
			c.JSON(http.StatusCreated, "Team is created")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't create team",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Json not valid",
		})
	}
}

func (teamHandler *TeamHandler) UpdateTeam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		jsonObj := models.Team{}
		err := c.ShouldBind(&jsonObj)
		if err == nil {
			_, isDone := teamHandler.Service.UpdateTeam(id, jsonObj)
			if isDone == true {
				c.JSON(http.StatusOK, "Updated Team: "+c.Param("id"))
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Can't update team " + c.Param("id"),
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

func (teamHandler *TeamHandler) DeleteTeam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		_, isDone := teamHandler.Service.DeleteTeam(id)
		if isDone == true {
			c.JSON(http.StatusOK, "Deleted Team "+c.Param("id"))
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Can't Delete user or user not exits" + c.Param("id"),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "ID not valid",
		})
	}
}

func (teamHandler *TeamHandler) MoveTeam(c *gin.Context) {
	isValid := true
	teamId, err := strconv.Atoi(c.Query("teamId"))
	if err != nil {
		isValid = false
	}
	hubId, err := strconv.Atoi(c.Query("hubId"))
	if err != nil {
		isValid = false
	}
	if isValid {
		isMove := teamHandler.Service.MoveTeam(teamId, hubId)
		if isMove {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Team " + c.Query("teamId") + "is moved to hub" + c.Query("hubId"),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Can't move team " + c.Query("teamId") + " to new hub " + c.Query("hubId"),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "hub id or team id is not valid",
		})
	}
}
