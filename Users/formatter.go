package Users

import (
	"startup_be/Helper"
)

type UserFormatter struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Occupation int    `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(User User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         Helper.HashIdEncode(User.ID),
		Name:       User.Name,
		Occupation: User.OccupationId,
		Email:      User.Email,
		Token:      token,
	}
	return formatter
}
