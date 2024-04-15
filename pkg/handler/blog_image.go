package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"


	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBlogImage(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
		return
	}

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
	fileName := fmt.Sprintf("%d/%d.jpg",id, time.Now().Unix())
	path := os.Getenv("pathBlogImage")
	if err := os.MkdirAll(path, 0755); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	imageURL := filepath.Join(path, fileName)
	targetFile, err := os.OpenFile(imageURL, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.BlogImage.CreateImage(entity.ImageInputBlog{Blog_id: id, Image: imageURL})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"imageURL": imageURL,
	})
}

func (h *Handler) getAllBlogImages(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	imagePaths, err := h.services.BlogImage.GetAllImages(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	boundary := "boundary"

	// Для каждого пути к изображению
	for _, imagePath := range imagePaths {
		// Открываем файл изображения
		image, err := os.Open(imagePath)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer image.Close()

		// Записываем границу между частями multipart
		_, _ = c.Writer.Write([]byte("--" + boundary + "\n"))

		// _, _ = c.Writer.Write([]byte("Content-Type: image/jpeg\n\n"))

		// Копируем содержимое изображения в ResponseWriter
		if _, err := io.Copy(c.Writer, image); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// // Записываем новую строку после каждого изображения
		_, _ = c.Writer.Write([]byte("\n"))
	}

	// // Записываем последнюю границу multipart
	_, _ = c.Writer.Write([]byte("--" + boundary + "--"))

}

func (h *Handler) getBlogImageById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	im_id, err := strconv.Atoi(c.Param("im_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid im_id param")
		return
	}

	imagePath, err := h.services.BlogImage.GetImageById(id, im_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	image, err := os.Open(imagePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer image.Close()

	if _, err := io.Copy(c.Writer, image); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) updateBlogImage(c *gin.Context) {

}

func (h *Handler) deleteBlogImage(c *gin.Context) {
	im_id, err := strconv.Atoi(c.Param("im_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid im_id param")
		return
	}

	err = h.services.BlogImage.DeleteImage(im_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
