package handler

import (
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.authService.Login(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "login amalga oshmadi", err.Error())
		return
	}

	response.Success(c, "login muvaffaqiyatli", data)
}
