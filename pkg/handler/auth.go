package handler

import (
	
	"net/http"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
	
)




func (h *Handler) signUp(c *gin.Context) {
	var input entity.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if input.Email == "" || input.Name == "" || input.Password == ""  || input.Phone == "" {
		newErrorResponse(c,http.StatusBadRequest, "Email,password,name,phone are required")
	}

	customToken, err := h.services.Authorization.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"token": customToken})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h * Handler) signIn(c *gin.Context){
	var input entity.SignInInput

	if err:=c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	if input.Email =="" || input.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "email, password are required")
		return
	}

	customToken,err:=h.services.Authorization.SignIn(input)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK,gin.H{"token":customToken})

} 