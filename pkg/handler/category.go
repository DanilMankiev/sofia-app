package handler

import (
	"net/http"
	"strconv"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

//	@Summary		Create category
//	@Description	Create category API
//	@Tags			category
//  @ID 			create category api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param input body entity.Category true "Create category input"
//	@Success		200	{object}    entity.Category
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/category [post]
func (h *Handler) createCategory(c *gin.Context) {
	var input entity.CreateCategory

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	Category_id, err := h.services.Category.CreateCategory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{
		ID: Category_id,
	})

}

//	@Summary		Get All category
//	@Description	Get All category API
//	@Tags			category
//  @ID 			get all category api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//	@Success		200	{object}    []entity.Category
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/category [get]
func (h *Handler) getAllCategory(c *gin.Context) {

	Categories, err := h.services.Category.GetAllCategorys()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Categories)

}

//	@Summary		Get category by id
//	@Description	Get category by id API
//	@Tags			category
//  @ID 			get category by id api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Category id"
//	@Success		200	{object}    entity.Category
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/category/{id} [get]
func (h *Handler) getCategoryById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	Category, err := h.services.Category.GetCategoryById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Category)
}

//	@Summary		Update category
//	@Description	Update category API
//	@Tags			category
//  @ID 			update category api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Category id"
//  @Param input body entity.UpdateCategoryInput true "Update category input"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/category/{id} [put]
func (h *Handler) updateCategory(c *gin.Context) {
	var input entity.UpdateCategoryInput

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.Category.UpdateCategory(id, input); err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

//	@Summary		Delete category
//	@Description	Delete category API
//	@Tags			category
//  @ID 			delete category api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Category id"
//	@Success		200	{object}   	statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/category/{id} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Category.DeleteCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
