package main

import (
	"is-public-api/application/configs"
	"is-public-api/application/routers"
	"is-public-api/cmd/api/server"
	"is-public-api/helpers/configloader"
)

func main() {
	serverConfig := new(configs.ConfigServer)
	configloader.ReadConf(serverConfig)

	server := server.NewHttpServer(serverConfig)

	router := routers.NewHttpRouter(serverConfig.Server.ContextPath)
	router.EnableCORS("*")

	server.ListenAndServe(router.Handler())
}
