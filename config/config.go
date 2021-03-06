package config

import "os"

type Configset struct {
	Port string
}

func InitConfiguration() (conf *Configset) {
	conf = &Configset{
		Port: port(),
	}
	return
}

// Good practice to microservices to do not hardcore the configurable values
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8083"
	}
	return ":" + port
}
