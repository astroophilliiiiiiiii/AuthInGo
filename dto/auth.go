package dto

type LoginUserRequestDTO struct {
	// Jab JSON read ya send hoga, is field ka naam email hoga.
	//Validator ke rules hain. required = field empty nhi honi chahiye. email = valid email format hona chahiye.
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserRequestDto struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
