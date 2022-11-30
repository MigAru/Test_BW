package structs


type TransactionResponse struct {
    ID            uint
	UserID        uint
	Amount        int
	CreatedAt     int64
	TypeOperation string
}

type MessageResponse struct {
    Message string `json:"message"`
}
