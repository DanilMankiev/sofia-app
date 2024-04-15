package handler

import (
	"net/http"
	"strconv"
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createReview(c *gin.Context){
	var input entity.CreateReview

	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err:= h.services.CreateReview(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}
func (h *Handler) getAllReview(c *gin.Context){

	review,err:=h.services.Review.GetAllReview()
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, review)

}
func (h *Handler) updateReview(c *gin.Context){
	var input entity.UpdateReview

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.services.Review.UpdateReview(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
func (h *Handler) deleteReview(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
	}
	if err=h.services.Review.DeleteReview(id);err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"Status":"OK",
	})

}