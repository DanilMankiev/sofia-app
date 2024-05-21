package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lib/pq"
	"strconv"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	
)

//	@Summary		Create product image
//	@Description	Create product image API
//	@Tags			product
//  @ID 			create product image api
//	@Accept			multipart/form-data
//	@Produce		json
// @Security ApiKeyAuth
//  @Param image formData file true "Product images"
// @Param id path int true "Product ID"
//	@Success		200	{object}    imageURLResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/products/{id}/images [post]
func (h *Handler) createImage(c *gin.Context) {

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
		}
		fileName := fmt.Sprintf("%s.jpg", randomFilename())
		path := fmt.Sprintf("%s/%d", viper.GetString("pathProductImage"), id)

		if err := os.MkdirAll(path, 0755); err != nil {
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
			err:=h.services.ProductImage.CreatePreviewImage(imageURL[0],id)
			if err!=nil{
				newErrorResponse(c,http.StatusInternalServerError,err.Error())
				return
			}
		}
	}

	err = h.services.ProductImage.CreateImage(entity.ImageInputProduct{Product_id: id, Image: imageURL})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, imageURLResponse{
		ImageURL: imageURL,
	})
}

//	@Summary		Delete product image
//	@Description	Delete product image API
//	@Tags			product
//  @ID 			delete product image api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Product id"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/products/{id}/images [delete]
func (h *Handler) deleteImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.ProductImage.DeleteImage(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
