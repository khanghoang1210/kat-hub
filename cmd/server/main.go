package main

import (
	_ "kathub/docs"
	"kathub/internal/routers"
	"kathub/internal/wire"
)

// @title   Kathub API
// @version	1.0
// @description Kathub API

// @host 	localhost:8080
// @BasePath /api/v1
func main() {

	// initial user instance
	uc := wire.InitUserRouterHandler()

	// initial account instance
	ac := wire.InitAccountRouterHandler()

	// initial account instance
	pc := wire.InitPostRouterHandler()

	r := routers.NewRouter(uc, ac, pc)
	r.Run()
}
