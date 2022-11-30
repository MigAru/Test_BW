package structs


type TransactionResponse struct {
    ID            uint
	UserID        uint
	Amount        int
	CreatedAt     int64
	TypeOperation string
}

//MessageResponse - default or not success response
type MessageResponse struct {
    Message string `json:"message"`
}
