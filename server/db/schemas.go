package db

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"<-:create;unique"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

type Task struct {
    ID            uint   `gorm:"primaryKey"`
    UserID        uint  `gorm:"<-:create"`
    Price         int   `gorm:"<-:create"`
    CreatedAt     int64 `gorm:"autoCreateTime"`
	TimeInJob     int64 
    TypeOperation int
    Status        int
	ErrorText     string
}
