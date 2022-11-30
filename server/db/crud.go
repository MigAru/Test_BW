package db

import (
	"errors"

	"gorm.io/gorm"
)

func GetTransactionsUser(userID uint) ([]Transaction, error) {
	res := []Transaction{}
	tx := dbConn.Where("user_id = ?", userID).Find(&res)
	if tx.Error != nil {
		return res, tx.Error
	}
	return res, nil
}

func GetTransaction(id uint) (Transaction, error) {
	transaction := Transaction{}
	if dbConn.Where("id = ?", id).First(&transaction).Error != nil {
		return transaction, errors.New("transaction not found")
	}
	return transaction, nil
}

func GetFirstActiveTransaction(user_id uint) (Transaction, error) {
    transaction := Transaction{}
    if err := dbConn.Where("user_id = ? AND is_active = 0", user_id).First(&transaction).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return transaction, err
        }
        return transaction, errors.New("active transaction not found")
    }
    transaction.IsActive = NotActiveTransaction
    dbConn.Save(&transaction)
    transaction.IsActive = ActiveTransaction
    return transaction, nil
}

func CreateTransaction(amount, typeOperation int, userID uint) (uint, error) {
	user := User{}
	transaction := Transaction{
		UserID:        userID,
		Amount:        amount,
		TypeOperation: typeOperation,
	}
	tx := dbConn.Create(&transaction)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if dbConn.Where("id = ?", userID).First(&user).Error != nil {
		return 0, errors.New("user not found")
	}
	if typeOperation == 0 {
		user.Balance += amount
	} else if typeOperation == 1 {
		user.Balance -= amount
	}
	dbConn.Save(&user)
	return transaction.ID, nil
}

func GetUsers() ([]User, error) {
	users := []User{}
	if dbConn.Find(&users).Error != nil {
		return users, errors.New("users not found")
	}
	return users, nil
}

func GetUser(id uint) (User, error) {
	user := User{}
	if dbConn.Where("id = ?", id).First(&user).Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func CreateUser(username string) (uint, error) {
	user := User{
		Username: username,
	}
	tx := dbConn.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil
}
