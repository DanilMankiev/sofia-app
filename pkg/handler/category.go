package handler

import (
	"net/http"
	"strconv"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context) {
	var input entity.Category

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	Category_id, err := h.services.Category.CreateCategory(input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": Category_id,
	})

}

func (h *Handler) getAllCategory(c *gin.Context) {

	Categorys, err := h.services.Category.GetAllCategorys()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Categorys)

}

func (h *Handler) getCategoryById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	Category, err := h.services.Category.GetCategoryById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, Category)
}

func (h *Handler) updateCategory(c *gin.Context) {
	var input entity.UpdateCategoryInput

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.services.Category.UpdateCategory(id, input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.Category.DeleteCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
