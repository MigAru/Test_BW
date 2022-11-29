package db

import (
	"errors"
	"time"
)

func CreateTask(price, operation int, userID uint) (uint, error) {
	task := Task{
		UserID:        userID,
		Price:         price,
		TypeOperation: operation,
	}
	tx := dbConn.Create(&task)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return task.ID, nil
}

func UpdateStatusTask(id uint, status int) error {
	if status == TaskInJob {
		time := time.Now().Unix()
		dbConn.Model(&Task{}).Where("id=?", id).Update("time_in_job", time)
	}
	dbConn.Model(&Task{}).Where("id=?", id).Update("type_operation", status)
	return nil

}

func GetUsers() ([]User, error){
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

func GetTasks(userID int) ([]Task, error) {
	res := []Task{}
	tx := dbConn.Where("id = ?", userID).Find(&res)
	if tx.Error != nil {
		return res, tx.Error
	}
	return res, nil
}
