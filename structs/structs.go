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

// Message
type Message struct {
	from int
	to   int
	text string
}

// END Message
// Services
//
//	General
type ServiceResponse struct {
	Success bool
	Status  int
	Result  interface{}
	Err     string
}

// END Services

// General
type IError struct {
	Field string
	Tag   string
	Value string
}

// END General
