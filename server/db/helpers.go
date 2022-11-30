package db

import (
	"fmt"
	"srv/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB


//ConnectDB - Connection to db with config params
func ConnectDB(cfg structs.ConfigDB) error {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dbConn = db

	//for tests
	if err := dbConn.AutoMigrate(&User{}, &Transaction{}); err != nil {
		return err
	}
	return nil
}

//NormalizeTransactions - normalizing slice transactions for response
func NormalizeTransactions(transactions []Transaction) []structs.TransactionResponse {
	res := []structs.TransactionResponse{}
	for _, transaction := range transactions {
		typeOperation := ""
        IsActive := false
		switch transaction.TypeOperation {
		case AddPrice:
			typeOperation = "add"
		case ReducePrice:
			typeOperation = "reduce"
		}
        switch transaction.IsActive {
        case ActiveTransaction:
            IsActive = true
        case NotActiveTransaction:
            IsActive = false
        }
		res = append(res, structs.TransactionResponse{
			UserID:        transaction.UserID,
			ID:            transaction.ID,
			Amount:        transaction.Amount,
			CreatedAt:     transaction.CreatedAt,
            IsActive:      IsActive,
			TypeOperation: typeOperation,
		})
	}
	return res
}

func getBalance(userID uint) int {
	user := User{}
	dbConn.First(&user)
	if user.ID == 0 {
		return -1
	}
	return user.Balance
}

//ValidateBalace - validating current balance
func ValidateBalace(userID uint, amount, typeOperation int) bool {
	balance := getBalance(userID)
	if typeOperation == ReducePrice {
		if balance-amount < 0 {
			return false
		}
	}
	return true
}
