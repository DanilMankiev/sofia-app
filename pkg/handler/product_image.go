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

func (h *Handler) createImage(c *gin.Context) {

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
			}
		}
	}

	err = h.services.ProductImage.CreateImage(entity.ImageInputProduct{Product_id: id, Image: imageURL})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"imageURL": imageURL,
	})
}

	// func (h *Handler) createImage(c *gin.Context) {

	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
	// 		return
	// 	}

	// 	file, _, err := c.Request.FormFile("image")
	// 	if err != nil {
	// 		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	defer file.Close()
	// 	fileName := fmt.Sprintf("%d.jpg",time.Now().Unix())
	// 	path := fmt.Sprintf("%s/%d",viper.GetString("pathProductImage"),id)

	// 	if err := os.MkdirAll(path, 0755); err != nil {
	// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	fullNamePath := fmt.Sprintf("%s/%s",path, fileName)
	// 	targetFile, err := os.OpenFile(fullNamePath, os.O_CREATE|os.O_WRONLY, 0644)
	// 	if err != nil {
	// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	defer targetFile.Close()
	// 	imageURL:= fmt.Sprintf("%s/%s/%s",viper.GetString("statichost"), path, fileName)
	// 	_, err = io.Copy(targetFile, file)
	// 	if err != nil {
	// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	err = h.services.ProductImage.CreateImage(entity.ImageInput{Product_id: id, Image: imageURL})
	// 	if err != nil {
	// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	}

	// 	c.JSON(http.StatusOK, map[string]interface{}{
	// 		"imageURL": imageURL,
	// 	})

	// Вариант 2
	// извлекаем id продукта и создаем директорию
	// productId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	//  	newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
	//  	return
	// }
	// {
	// 	err := os.MkdirAll(fmt.Sprintf("./image/product/%d", productId), os.ModePerm) // создаем директорию с id продукта, если еще не создана
	// 	if err != nil {
	// 		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.MkdirAll error: %s", err))
	// 		return
	// 	}
	// }
	// path := fmt.Sprintf("./image/product/%d", productId)

	// // извлекаем файл из парамeтров post запроса
	// form, _ := c.MultipartForm()
	// var fileName string
	// imgExt := "jpeg"
	// // берем первое имя файла из присланного списка
	// val,_,err:=c.Request.FormFile("image")
	// val.Read()
	// for key := range form.File {
	// 	fileName = key
	// 	// извлекаем расширение файла

	// 	fmt.Println(str)
	// 	arr := strings.Split(fileName, ".")
	// 	if len(arr) > 1 {
	// 		imgExt = arr[len(arr)-1]
	// 	}
	// }
	// // извлекаем содержание присланного файла по названию файла
	// file, _, err := c.Request.FormFile(fileName)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("UploadXml c.Request.FormFile error: %s", err.Error()))
	// 	return
	// }
	// defer file.Close()

	// // читаем содержание присланного файл в []byte
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// fullFileName := fmt.Sprintf("%d.%s", time.Now().Unix(), imgExt)
	// // открываем файл для сохранения картинки
	// fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.Create err: %s", err))
	// 	return
	// }
	// defer fileOnDisk.Close()

	// _, err = fileOnDisk.Write(fileBytes)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// err = h.services.ProductImage.CreateImage(entity.ImageInput{Product_id:productId, Image: fullFileName})
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// }
	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	 	"imageURL": fullFileName,
	// 	 })



// func (h *Handler) getAllImages(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	imagePaths, err := h.services.ProductImage.GetAllImages(id)
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

// func (h *Handler) getImageById(c *gin.Context) {

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

// 	imagePath, err := h.services.ProductImage.GetImageById(id, im_id)
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

// func (h *Handler) updateImage(c *gin.Context) {

// }

func (h *Handler) deleteImage(c *gin.Context) {
	im_id, err := strconv.Atoi(c.Param("im_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid im_id param")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid im_id param")
		return
	}

	err = h.services.ProductImage.DeleteImage(im_id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
