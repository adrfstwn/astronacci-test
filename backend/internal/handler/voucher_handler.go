package handler

import (
	"errors"
	"io"
	"net/http"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func bindError(err error) string {
	if errors.Is(err, io.EOF) {
		return "request body is required"
	}
	return err.Error()
}

type VoucherHandler struct {
	service *service.VoucherService
}

func NewVoucherHandler(s *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{service: s}
}

func (h *VoucherHandler) Check(c *gin.Context) {
	var req dto.CheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindError(err)})
		return
	}

	exists, err := h.service.CheckExists(req.FlightNumber, req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check voucher"})
		return
	}

	c.JSON(http.StatusOK, dto.CheckResponse{Exists: exists})
}

func (h *VoucherHandler) Generate(c *gin.Context) {
	var req dto.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindError(err)})
		return
	}

	seats, err := h.service.GenerateVoucher(req)
	if err != nil {
		switch err {
		case service.ErrVoucherExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case service.ErrInvalidAircraft:
			c.JSON(http.StatusBadRequest, gin.H{"error": bindError(err)})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate voucher"})
		}
		return
	}

	c.JSON(http.StatusOK, dto.GenerateResponse{Success: true, Seats: seats})
}
