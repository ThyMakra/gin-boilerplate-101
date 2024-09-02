package schemas

type UserSchema struct {
	ID        string `json:"id" validate:"uuid"`
	FirstName string `json:"first_name" validate:"required,lowercase"`
	LastName  string `json:"last_name" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
}
