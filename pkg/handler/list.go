package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	var input sofia.List

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list_id, err := h.services.ListOfproducts.CreateList(input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id": list_id,
	})

}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
