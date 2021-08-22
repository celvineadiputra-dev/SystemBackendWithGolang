package campaign

import "time"

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	Perks            string
	ShortDescription string
	Description      string
	BackerCount      int
	GoalCount        int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID int
	CampaignId int
	FileName string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}