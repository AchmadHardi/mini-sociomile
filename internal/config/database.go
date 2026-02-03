package config

import "fmt"

const (
	DBUser = "root"
	DBPass = ""
	DBHost = "127.0.0.1"
	DBPort = "3306"
	DBName = "sociomile"
)

func GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBName,
	)
}
