package main

import (
	"cubicasa/controller"
	"cubicasa/service"
)

/*
I to add new route with new handler without coding straight to controller, so I changed the basic design
with func add route base on instance controller
main -> controller -> service ->  model + repo
*/

func main() {
	controllerMain := controller.GetInstance()
	controllerMain.InitController()
	controllerServiceUser := controller.UserHandler{
		Service: service.UserService{},
	}
	controllerServiceUser.Service.InitService()
	// search user by name return team and hub name
	controllerMain.AddUnauthorizedRoute("GET", "/users", "/search", controllerServiceUser.SearchUser)
	controllerMain.AddUnauthorizedRoute("GET", "/users", "", controllerServiceUser.GetAll)
	controllerMain.AddUnauthorizedRoute("GET", "/users", "/:id", controllerServiceUser.GetUserById)
	controllerMain.AddAuthorizedRoute("POST", "/users", "", controllerServiceUser.CreateUser)
	controllerMain.AddAuthorizedRoute("PUT", "/users", "/:id", controllerServiceUser.UpdateUser)
	controllerMain.AddAuthorizedRoute("DELETE", "/users", "/:id", controllerServiceUser.DeleteUser)
	// move user to new team
	controllerMain.AddAuthorizedRoute("PUT", "/users", "/move", controllerServiceUser.MoveUser)

	// team
	controllerServiceTeam := controller.TeamHandler{
		Service: service.TeamService{},
	}
	controllerServiceTeam.Service.InitService()

	controllerMain.AddUnauthorizedRoute("GET", "/teams", "", controllerServiceTeam.GetAll)
	controllerMain.AddUnauthorizedRoute("GET", "/teams", "/:id", controllerServiceTeam.GetTeamById)
	controllerMain.AddAuthorizedRoute("POST", "/teams", "", controllerServiceTeam.CreateTeam)
	controllerMain.AddAuthorizedRoute("PUT", "/teams", "/:id", controllerServiceTeam.UpdateTeam)
	controllerMain.AddAuthorizedRoute("DELETE", "/teams", "/:id", controllerServiceTeam.DeleteTeam)
	// move team to new hub
	controllerMain.AddAuthorizedRoute("PUT", "teams", "/move", controllerServiceTeam.MoveTeam)
	//hub
	controllerServiceHub := controller.HubHandler{
		Service: service.HubService{},
	}
	controllerServiceHub.Service.InitService()

	controllerMain.AddUnauthorizedRoute("GET", "/hubs", "", controllerServiceHub.GetAll)
	controllerMain.AddUnauthorizedRoute("GET", "/hubs", "/:id", controllerServiceHub.GetHubById)
	controllerMain.AddAuthorizedRoute("POST", "/hubs", "", controllerServiceHub.CreateHub)
	controllerMain.AddAuthorizedRoute("PUT", "/hubs", "/:id", controllerServiceHub.UpdateHub)
	controllerMain.AddAuthorizedRoute("DELETE", "/hubs", "/:id", controllerServiceHub.DeleteHub)

	// Run
	err := controllerMain.ControllerRouter.Run()
	if err != nil {
		return
	}
}
