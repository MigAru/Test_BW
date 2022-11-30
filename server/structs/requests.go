package structs

type CreateTransactionRequest struct {
	UserID        uint   `json:"user_id"`
	Amount        int    `json:"amount"`
	OperationType string `json:"operation_type"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
