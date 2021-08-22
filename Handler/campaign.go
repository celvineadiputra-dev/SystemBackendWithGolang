package Handler

import (
	"github.com/gin-gonic/gin"
	"startup_be/campaign"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewHandler(campaignService campaign.Service) *campaignHandler{
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaigns(c *gin.Context){
	
}