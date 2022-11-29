package structs

//Config - configuration server struct
type Config struct {
    MaxWorkers int     //if 0 or none in config = 1, default = 1 
	Port string     //example ":8080"
	Nats ConfigNats //configuration for connect nats server\servers
}

type ConfigNats struct {
	Servers  []string // example ["nats://localhost:4222", "nats://localhost:4223", ...etc]
	Username string   //if need
	Password string   //if need
}
