package handler

import (
	"net/http"
	"strconv"
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)


//	@Summary		Create review 
//	@Description	Create review API
//	@Tags			review
//  @ID 			create review api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param input body entity.Review true "Review input"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/review [post]
func (h *Handler) createReview(c *gin.Context){
	var input entity.CreateReview

	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid request body")
		return
	}

	id,err:= h.services.CreateReview(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{ID:id})
}

//	@Summary		Get all review 
//	@Description	Get all review API
//	@Tags			review
//  @ID 			get all review api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Success		200	{object}    []entity.Review
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/review [get]
func (h *Handler) getAllReview(c *gin.Context){

	review,err:=h.services.Review.GetAllReview()
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, review)
}

//	@Summary		Update review 
//	@Description	Update review API
//	@Tags			review
//  @ID 			update review api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Review id"
//  @Param input body entity.Review true "Review input"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/review/{id} [put]
func (h *Handler) updateReview(c *gin.Context){
	var input entity.UpdateReview

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = h.services.Review.UpdateReview(id,input); err!=nil{
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

//	@Summary		Delete review 
//	@Description	Delete review API
//	@Tags			review
//  @ID 			delete review api
//	@Accept			json
//	@Produce		json
//  @Param id path int true "Review id"
// @Security ApiKeyAuth
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/review/{id} [delete]
func (h *Handler) deleteReview(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	if err=h.services.Review.DeleteReview(id);err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,statusResponse{Status: "OK"})
}