package structs

// Config
type AppConfig struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

type DbConfig struct {
	User     string `json:"-"`
	Password string `json:"-"`
	Name     string `json:"db"`
	Ip       string `json:"ip"`
}

type Config struct {
	App AppConfig `json:"app"`
	Db  DbConfig  `json:"db"`
}

// END Config

// Services
type ServiceResponse struct {
	Success bool
	Status  int
	Result  interface{}
	Err     string
}

// END Services

// Validator
type IError struct {
	Field string
	Tag   string
	Value string
}

// END Validator
