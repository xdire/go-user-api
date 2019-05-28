package main

import (
	"fmt"
	"github.com/xdire/go-user-api/controllers"
	data_prefill "github.com/xdire/go-user-api/data-prefill"
	conn_manager "github.com/xdire/go-user-api/manager"
	"github.com/xdire/go-user-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	// Establish database connectivity
	conn_manager.ConnectDefault()
	defer conn_manager.DisconnectDefault()
	conn, err := conn_manager.GetDefaultConnection()
	if err != nil { panic("Cannot establish default connection to database") }
	// Migrate entities
	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Group{})
	// Prefill initial set of values
	data_prefill.GenerateInitialUserSet(conn)
	data_prefill.TestInitialUserSet(conn)
	// Prefill initial set of groups
	data_prefill.GenerateInitialGroupSet(conn)
	data_prefill.TestInitialGroupSet(conn)
	// Prefill relations
	data_prefill.GenerateUserGroupSalesRelations(conn)
	data_prefill.GenerateUserGroupAuthorizedRelations(conn)
	// Start rest server serving the data
	startRestServer()
}

func assignRoutes(r *mux.Router)  {
	r.HandleFunc("/api/user/{id}", controllers.RouteCtrlGetUserById)
}

func startRestServer()  {
	// Create router
	router := mux.NewRouter()
	// Assign controllers to different routes
	assignRoutes(router)
	// Start server
	err := http.ListenAndServe(":9394", router)
	if err != nil {
		fmt.Print(err)
	}
}
