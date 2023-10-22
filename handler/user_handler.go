package handler

import (
	"ginapi-cleanarchi/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetUser(c *gin.Context)
}

type UserHandler struct {
	usecase.UserInputPort
}

func NewUserHandler(srv usecase.UserInputPort) IUserHandler {
	return &UserHandler{srv}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userid := c.Param("id")
	user, _ := h.UserInputPort.GetUser(userid)
	c.JSON(http.StatusOK, user)

}
