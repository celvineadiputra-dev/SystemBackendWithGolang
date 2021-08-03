package Test

import "time"

type User struct {
	ID             int
	Name           string
	OccupationId   int
	Email          string
	password       string
	AvatarFileName string
	RoleId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
