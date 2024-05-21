package handler

import (
	"net/http"

	"strconv"
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)
//	@Summary		Create blog
//	@Description	Create blog API
//	@Tags			blog
//  @ID 			create blog api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Param			input	body entity.CreateBlog true	"Blog input"
//	@Success		200	{object}    idResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/blog [post]
func (h* Handler) createBlog(c *gin.Context){
	var input entity.CreateBlog

	if err:=c.BindJSON(&input);err!=nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	id,err:= h.services.Blog.CreateBlog(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,idResponse{ID: id})

}

//	@Summary		Get All blog
//	@Description	Get all blog API
//	@Tags			blog
//  @ID 			get all blog api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Success		200	{object}    []entity.Blog
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/blog [get]
func (h *Handler) getAllBlog(c *gin.Context){

	blog,err:=h.services.Blog.GetAllBlog()
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, blog)

}

//	@Summary		Update blog
//	@Description	Update blog API
//	@Tags			blog
//  @ID 			update blog api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Blog ID"
//	@Param			input	body entity.UpdateBlog true	"Update Blog input"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/blog/{id} [put]

func (h *Handler) updateBlog(c *gin.Context){
	var input entity.UpdateBlog

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	err=h.services.Blog.UpdateBlog(id,input)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}


//	@Summary		Delete blog
//	@Description	Delete blog API
//	@Tags			blog
//  @ID 			delete blog api
//	@Accept			json
//	@Produce		json
//  @Param id path int true "Blog ID"
// @Security ApiKeyAuth
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/blog/{id} [delete]
func (h *Handler) deleteBlog(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	if err=h.services.Blog.DeleteBlog(id);err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,statusResponse{
		Status:"OK",
	})
}


//	@Summary		Get blog by ID
//	@Description	Get blog by ID API
//	@Tags			blog
//  @ID 			Get blog by ID api
//	@Accept			json
//	@Produce		json
//  @Param id path int true "Blog ID"
// @Security ApiKeyAuth
//	@Success		200	{object}    entity.Blog
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/blog/{id} [get]
func (h* Handler) getBlogById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "ivalid param")
		return
	}

	blog, err := h.services.Blog.GetBlogById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, blog)
}
