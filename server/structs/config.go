package structs

//Config - configuration server struct
type Config struct {
    Port       string     `json:"port" validate:"required"`        //example ":8080"
    Postgres   ConfigDB   `json:"postgres"`
}

type ConfigDB struct {
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBname   string `json:"db_name" validate:"required"`
	Port     string `json:"port" validate:"required"` //example "5432"
}
