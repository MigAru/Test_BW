package structs

//Config - configuration server struct
type Config struct {
	MaxWorkers int        `json:"max_workers" validate:"required"` //if 0 or none in config = 1, default = 1
	Port       string     `json:"port" validate:"required"`        //example ":8080"
	Nats       ConfigNats `json:"nats"`                            //configuration for connect nats server\servers
	Postgres   ConfigDB   `json:"postgres"`
}

type ConfigNats struct {
	Servers  []string `json:"servers" validate:"required"` // example ["nats://localhost:4222", "nats://localhost:4223", ...etc]
	Username string   `json:"username"`                    //if need
	Password string   `json:"password"`                    //if need
}

type ConfigDB struct {
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBname   string `json:"db_name" validate:"required"`
	Port     string `json:"port" validate:"required"` //example "5432"
}
