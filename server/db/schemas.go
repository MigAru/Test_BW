package db

import "srv/structs"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"<-:create;unique"`
    Balance   int    `gorm:"default:0"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

type Transaction struct {
	ID            uint  `gorm:"primaryKey"`
	UserID        uint  `gorm:"<-:create"`
	Amount        int   `gorm:"<-:create"`
	CreatedAt     int64 `gorm:"autoCreateTime"`
	TypeOperation int
}

type UserTransactions struct {
    User User
    Transactions []structs.TransactionResponse
}
