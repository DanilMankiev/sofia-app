package handler

import (
	"net/http"
	"strconv"

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/gin-gonic/gin"
)

//	@Summary		Get all product 
//	@Description	Get all product API
//	@Tags			product
//  @ID 			get all products api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Category id"
//	@Success		200	{object}    []entity.Product
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/category/{id}/products [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	category_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "inavlid id param")
		return
	}

	products, err := h.services.Product.GetAllItems(category_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)

}


//	@Summary		Create product 
//	@Description	Create product API
//	@Tags			product
//  @ID 			create products api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path integer true "Category id"
//  @Param input body entity.CreateProduct true "Create product input"
//	@Success		200	{object}    idResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/category/{id}/products [post]
func (h *Handler) createProduct(c *gin.Context) {

	category_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "ivalid param")
		return
	}

	var input entity.CreateProduct
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	id, err := h.services.Product.CreateProduct(category_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, idResponse{ID :id})
}

//	@Summary		Get product by id 
//	@Description	Get product by id API
//	@Tags			product
//  @ID 			get product by id api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Product id"
//	@Success		200	{object}    entity.Product
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/products/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	product_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "ivalid param")
		return
	}

	product, err := h.services.Product.GetItemByid(product_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}


//	@Summary		Update product 
//	@Description	Update product API
//	@Tags			product
//  @ID 			update product api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Product id"
//  @Param input body entity.UpdateProductInput true "Product input"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/products/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	var input entity.UpdateProductInput

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = h.services.Product.UpdateItem(id, input); err!=nil{
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}

//	@Summary		Delete product 
//	@Description	Delete product API
//	@Tags			product
//  @ID 			delete product api
//	@Accept			json
//	@Produce		json
// @Security ApiKeyAuth
//  @Param id path int true "Product id"
//	@Success		200	{object}    statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
// @Router /api/products/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	product_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Product.DeleteItem(product_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}
