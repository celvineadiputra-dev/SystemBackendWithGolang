package Users

import "time"

type User struct {
	ID             int
	Name           string
	OccupationId   int
	Email          string
	Password       string
	AvatarFileName string
	RoleId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
