package db

import (
	"fmt"
	"srv/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var dbConn *gorm.DB


func ConnectDB(cfg structs.ConfigDB) error {
    dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",cfg.User,cfg.Password,cfg.Host,cfg.Port,cfg.DBname)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    dbConn = db
    if err := dbConn.AutoMigrate(&User{}); err != nil {
        return err 
    }
    if err := dbConn.AutoMigrate(&Task{}); err != nil {
        return err
    }
    return nil
}
