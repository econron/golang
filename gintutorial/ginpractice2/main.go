package main

import (
	"github.com/gin-gonic/gin"
	controller "ginpractice2/adapter/in/web"
	service "ginpractice2/application/domain/service"
	"ginpractice2/adapter/out/persistence"
	// "log"
)

func main(){
	config := persistence.MySqlConfig{}
	conn, err := config.NewDB()
	if err != nil {
		panic("error!!")
	}
	a := persistence.UpdateUserAdapter{
		Queries: conn,
	}
	s := service.UpdateUserService{
		A: &a,
	}
	c := controller.UpdateUserController{
		S: &s,
	}

	router := gin.Default()

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	router.PUT("/user/profile", c.UpdateProfile)
	router.Run("localhost:8080")
}