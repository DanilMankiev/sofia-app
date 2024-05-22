package handler

import (
	"net/http"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

//	@Summary		SignUp
//	@Description	SignUp API
//	@Tags			auth
//  @ID 			signup api
//	@Accept			json
//	@Produce		json
//	@Param			input	body entity.SignUpInput true	"Account input"
//	@Success		200	{object} statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input entity.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if input.Email == "" || input.Name == "" || input.Password == "" || input.Phone == "" {
		newErrorResponse(c, http.StatusBadRequest, "Email,password,name,phone are required")
		return
	}

	err := h.services.Authorization.SignUp(input)
	if err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.HTML(http.StatusInternalServerError,"name",err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}

//	@Summary		SignIn
//	@Description	SignIn API
//	@Tags			auth
//  @ID 			signin api
//	@Accept			json
//	@Produce		json
//	@Param			input	body entity.SignInInput true	"Sign-in input"
//	@Success		200	{object}    tokenResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input entity.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if input.Email == "" || input.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "email, password are required")
		return
	}

	customToken,refreshToken, err := h.services.Authorization.SignIn(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{Token: customToken, RefreshToken: refreshToken})

}

//	@Summary		Refresh token
//	@Description	Refresh token API
//	@Tags			auth
//  @ID 			refresh tokenapi
//	@Accept			json
//	@Produce		json
//  @Security ApiKeyAuth
//  @Param  token body entity.RefreshToken true "Refresh token"
//	@Success		200	{object}    tokenResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context){
	var refreshtoken entity.RefreshToken

	if err:=c.BindJSON(&refreshtoken);err!=nil{
		newErrorResponse(c, http.StatusBadRequest,"invalid request body")
		return
	}
	accessToken,refreshTokenNew,err:=h.services.Authorization.RefreshToken(refreshtoken.Token)
	if err!=nil{
		newErrorResponse(c, http.StatusBadRequest,err.Error())
		return
	}
	c.JSON(http.StatusOK, tokenResponse{Token:accessToken,RefreshToken: refreshTokenNew})
}