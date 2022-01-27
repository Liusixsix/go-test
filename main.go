package main

import (
	"gin-demo/common"
	"gin-demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	panic(r.Run(":8081"))
}
