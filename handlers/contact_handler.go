package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	service services.ContactService
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req requests.ContactRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:	"BAD_REQUEST",
			Code:	err.Error(),
			Data: nil,
		})
		return
	}

	contact, err := h.service.CreateContact(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses,APIResponse{
			Code: "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusCreated, responses.APIResponse{
		Code: "CREATED",
		Message: "Contact Created Successfully",
		Data: responses.ContactResponseFromModel(contact)
	})
}