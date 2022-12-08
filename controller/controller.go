package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type Controller struct {
	ControllerRouter *gin.Engine
}

var once sync.Once
var controller *Controller

func GetInstance() *Controller {
	once.Do(func() {
		controller = &Controller{}
	})
	return controller
}

func (controller *Controller) InitController() {
	controller.ControllerRouter = gin.Default()
}

// AddAuthorizedRoute AddRoute basic method is GET, POST , PUT and DELETE other case need more time to code
// GET is unauthorized
// PUT, POST , DELETE need authorized to execute
func (controller *Controller) AddAuthorizedRoute(stringType string, srcPath string, desPath string, handlerFunc gin.HandlerFunc) {
	// bind default user for test
	routerGroup := controller.ControllerRouter.Group(srcPath, gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	if stringType == "POST" {
		routerGroup.POST(desPath, handlerFunc)
	} else if stringType == "PUT" {
		routerGroup.PUT(desPath, handlerFunc)
	} else {
		routerGroup.DELETE(desPath, handlerFunc)
	}
}

func (controller *Controller) AddUnauthorizedRoute(stringType string, srcPath string, desPath string, handlerFunc gin.HandlerFunc) {
	if stringType == "GET" {
		controller.ControllerRouter.GET(srcPath+desPath, handlerFunc)
	} else {
		log.Println("Other type not allow Unauthorized")
	}
}
