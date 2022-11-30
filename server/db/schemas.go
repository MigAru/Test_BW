package db

import "srv/structs"


//User - user schema in db
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"<-:create;unique"`
	Balance   int    `gorm:"default:0"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

//Transaction - transaction schema in db
type Transaction struct {
	ID            uint  `gorm:"primaryKey"`
	UserID        uint  `gorm:"<-:create"`
	Amount        int   `gorm:"<-:create"`
	CreatedAt     int64 `gorm:"autoCreateTime"`
	TypeOperation int
}

//UserTransactions - struct for response with normalizing slice transactions
type UserTransactions struct {
    User User
    Transactions []structs.TransactionResponse
}
