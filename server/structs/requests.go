package structs

//CreateTransactionRequest - struct for creating transaction
type CreateTransactionRequest struct {
	UserID        uint   `json:"user_id"`
	Amount        int    `json:"amount"`
	OperationType string `json:"operation_type"`
}

//CreateUserRequest - struct to creating user
type CreateUserRequest struct {
	Username string `json:"username"`
}
