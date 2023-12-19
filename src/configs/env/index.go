package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentStruct struct {
	MONGODB_URI string
	MONGODB_DB  string
}

var MyEnv EnvironmentStruct

// func getNumberEnv(key string) int {
// 	value, err := strconv.Atoi(os.Getenv(key))
// 	if err != nil {
// 		return 0
// 	}
// 	return value
// }

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("---> Error loading .env file! <---")
	}

	MyEnv = EnvironmentStruct{
		MONGODB_URI: os.Getenv("MONGODB_URI"),
		MONGODB_DB:  os.Getenv("MONGODB_DB"),
	}
}
