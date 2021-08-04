package Users

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	OccupationId int    `json:"occupationid" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}
