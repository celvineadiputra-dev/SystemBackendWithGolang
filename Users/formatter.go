package Users

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation int    `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(User User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         User.ID,
		Name:       User.Name,
		Occupation: User.OccupationId,
		Email:      User.Email,
		Token:      token,
	}
	return formatter
}
