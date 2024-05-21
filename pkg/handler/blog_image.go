package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/lib/pq"
	"strconv"
	"github.com/spf13/viper"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

//	@Summary		Create Blog Image
//	@Description	Blog image API
//	@Tags			blog
//  @ID 			createblog api
//	@Accept			multipart/form-data
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Blog ID"
//  @Param image formData file true "Blog Images"
//	@Success		200	{object}    imageURLResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/blog/{id}/images [post]
func (h *Handler) createBlogImage(c *gin.Context) {

	var imageURL pq.StringArray

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
		return
	}

	form, err := c.MultipartForm()
	if err!=nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	images := form.File["image"]

	for i, file := range images {
		openedFile, err := file.Open()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		fileName := fmt.Sprintf("%s.jpg", randomFilename())
		path := fmt.Sprintf("%s/%d", viper.GetString("pathBlogImage"), id)

		if err := os.MkdirAll(path,0755); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		fullNamePath := fmt.Sprintf("%s/%s", path, fileName)
		targetFile, err := os.OpenFile(fullNamePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer targetFile.Close()
		imageURL = append(imageURL, fmt.Sprintf("%s/%s/%s", viper.GetString("statichost"), path, fileName))
		_, err = io.Copy(targetFile, openedFile)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if i ==0 {
			err:=h.services.BlogImage.CreatePreviewImage(imageURL[0],id)
			if err!=nil{
				newErrorResponse(c,http.StatusInternalServerError,err.Error())
				return
			}
		}
	}

	err = h.services.BlogImage.CreateImage(entity.ImageInputBlog{Blog_id: id,Image: imageURL})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, imageURLResponse{ImageURL: imageURL})
}

//	@Summary		Delete Blog Image
//	@Description	Blog image API
//	@Tags			blog
//  @ID 			deleteblog api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Blog ID"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/blog/{id}/images [delete]
func (h *Handler) deleteBlogImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
		return
	}
	err = h.services.BlogImage.DeleteImage(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
