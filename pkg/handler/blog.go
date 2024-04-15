package handler

import (
	"net/http"

	"strconv"
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)
func (h* Handler) createBlog(c *gin.Context){
	var input entity.CreateBlog

	if err:=c.BindJSON(&input);err!=nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id,err:= h.services.Blog.CreateBlog(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})

}

func (h *Handler) getAllBlog(c *gin.Context){

	blog,err:=h.services.Blog.GetAllBlog()
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, blog)

}
func (h *Handler) updateBlog(c *gin.Context){
	var input entity.UpdateBlog

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.services.Blog.UpdateBlog(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
func (h *Handler) deleteBlog(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
	}
	if err=h.services.Blog.DeleteBlog(id);err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"Status":"OK",
	})
}

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
