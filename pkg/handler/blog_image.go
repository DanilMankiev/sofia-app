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

func (h *Handler) createBlogImage(c *gin.Context) {

	var imageURL pq.StringArray

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
		return
	}

	form, err := c.MultipartForm()
	images := form.File["image"]

	for i, file := range images {
		openedFile, err := file.Open()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Cant open image/file")
		}
		fileName := fmt.Sprintf("%s.jpg", randomFilename())
		path := fmt.Sprintf("%s/%d", viper.GetString("pathBlogImage"), id)

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
			err:=h.services.BlogImage.CreatePreviewImage(imageURL[0],id)
			if err!=nil{
				newErrorResponse(c,http.StatusInternalServerError,err.Error())
			}
		}
	}

	err = h.services.BlogImage.CreateImage(entity.ImageInputBlog{Blog_id: id,Image: imageURL})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"imageURL": imageURL,
	})
}

// func (h *Handler) getAllBlogImages(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	imagePaths, err := h.services.BlogImage.GetAllImages(id)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 	}

// 	boundary := "boundary"

// 	// Для каждого пути к изображению
// 	for _, imagePath := range imagePaths {
// 		// Открываем файл изображения
// 		image, err := os.Open(imagePath)
// 		if err != nil {
// 			newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 		defer image.Close()

// 		// Записываем границу между частями multipart
// 		_, _ = c.Writer.Write([]byte("--" + boundary + "\n"))

// 		// _, _ = c.Writer.Write([]byte("Content-Type: image/jpeg\n\n"))

// 		// Копируем содержимое изображения в ResponseWriter
// 		if _, err := io.Copy(c.Writer, image); err != nil {
// 			newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		// // Записываем новую строку после каждого изображения
// 		_, _ = c.Writer.Write([]byte("\n"))
// 	}

// 	// // Записываем последнюю границу multipart
// 	_, _ = c.Writer.Write([]byte("--" + boundary + "--"))

// }

// func (h *Handler) getBlogImageById(c *gin.Context) {

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	im_id, err := strconv.Atoi(c.Param("im_id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid im_id param")
// 		return
// 	}

// 	imagePath, err := h.services.BlogImage.GetImageById(id, im_id)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	image, err := os.Open(imagePath)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	defer image.Close()

// 	if _, err := io.Copy(c.Writer, image); err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// }

// func (h *Handler) updateBlogImage(c *gin.Context) {

// }

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
