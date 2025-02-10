package main

import (
	"github.com/KaioAntonio/gin-rest-api/database"
	"github.com/KaioAntonio/gin-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
