package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup_be/Helper"
	"startup_be/campaign"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewHandler(campaignService campaign.Service) *campaignHandler{
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaigns(c *gin.Context){
	userID := Helper.HashIdDecode(c.Query("user_id"))
	campaigns, err := h.campaignService.FindCampaigns(userID)
	if err != nil{
		response := Helper.APIResponse("Get campaign Failed", http.StatusUnprocessableEntity, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
	}
	response := Helper.APIResponse("List of campaign", http.StatusOK, "Success", campaigns)
	c.JSON(http.StatusOK, response)
}