package structs


type CreateUserRequest struct {
    Username string `json:"username" validate:"required"`
}
