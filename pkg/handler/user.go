package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get user by id 
//	@Description	Get user by id API
//	@Tags			user
//  @ID 			Get user by id api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Success		200	{object}    entity.UserDisplay
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//  @Router /api/user/ [get]
func (h *Handler) getUser(c *gin.Context){
	
	uid,err:=getUserUID(c)
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest, "failed get uid")
		return
	}

	userData,err:= h.services.User.GetUser(uid)
	if err!=nil{
		newErrorResponse(c, http.StatusBadRequest,err.Error())
		return 
	}

	c.JSON(http.StatusOK,userData)
}

//	@Summary		Add product to favorites 
//	@Description	Add product to favorites API
//	@Tags			favotites
//  @ID 			Add product to favorites api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Product ID"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//  @Router /api/user/favorites/{id} [post]
func (h *Handler) createFavorites(c *gin.Context){
	uid,err:=getUserUID(c)
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest, "failed get uid")
		return
	}
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	err=h.services.User.CreateFavorites(uid,id)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}

//	@Summary		Get all favorites 
//	@Description	Get all favorites API
//	@Tags			favotites
//  @ID 			Get all favorites api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Success		200	{object}    []entity.Product
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//  @Router /api/user/favorites/ [get]
func (h * Handler) getAllFavorites(c *gin.Context){
	uid,err:=getUserUID(c)
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest, "failed get uid")
		return
	}
	products,err:=h.services.User.GetAllFavorites(uid)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

//	@Summary		Delete favorites 
//	@Description	Delete favorites API
//	@Tags			favotites
//  @ID 			Delete favorites api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Pararm id path int true "Product ID"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//  @Router /api/user/favorites/{id} [delete]
func (h* Handler) deleteFavorites(c *gin.Context){
	uid,err:=getUserUID(c)
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest, "failed get uid")
		return
	}
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	err=h.services.User.DeleteFavorites(uid,id)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "OK"})
}