package v1

import (
	"github.com/EnricoPDG/meli-desafio/response"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type SellerHandler struct {
	service service.SellerServiceAPI
}

func NewSellerHandler(service service.SellerServiceAPI) *SellerHandler {
	return &SellerHandler{service: service}
}

// GetSellerByID godoc
// @Summary Get seller by ID
// @Description Get details of a seller
// @Tags sellers
// @Accept json
// @Produce json
// @Param id path string true "Sellers ID"
// @Success 200 {object} response.JSONSuccessResultSeller
// @Failure 400 {object} response.JSONErrorResponse
// @Failure 404 {object} response.JSONErrorResponse
// @Failure 500 {object} response.JSONErrorResponse
// @Router /api/v1/sellers/{id} [get]
func (h *SellerHandler) GetSellerByID(c *gin.Context) {
	log.Debug("GetSellerByID called")
	strID := c.Param("id")
	id, err := uuid.Parse(strID)
	if err != nil {
		log.Error("error parsing uuid", zap.Error(err))
		response.BadRequestResponse(c, err)
		return
	}

	seller, err := h.service.GetSellerByID(id)
	if err != nil {
		log.Error("error getting sellers", zap.Error(err))
		response.ChooseErrorResponse(c, err)
		return
	}

	response.SuccessResponseSeller(c, seller)
}
