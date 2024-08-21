package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginpractice2/application/domain/service"
	port "ginpractice2/application/port/in"
)

type updateUserRequest struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type UpdateUserController struct {
	S *service.UpdateUserService
}

func (c *UpdateUserController) UpdateProfile (gc *gin.Context) {
	var request updateUserRequest
	if err := gc.BindJSON(&request); err != nil {
		return
	}
	command := port.UpdateUserCommand{
		UserID: request.ID,
		Name: request.Name,
	}
	ret := c.S.UpdateUserById(&command)

	if !ret {
		gc.IndentedJSON(http.StatusBadRequest, "karioki error")
	}
	gc.IndentedJSON(http.StatusOK, "suceeded")
}