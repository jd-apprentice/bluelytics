package functions

import (
	consts "dolar/src/constants"
	"dolar/src/types"
	handlers "dolar/src/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getUserHome() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return home
}

func InitDotEnv() {

	home := getUserHome()

	err := godotenv.Load(home + consts.EnvPath)
	if err != nil {
		handlers.LogFatalFmt(consts.FailedFetch, err)
	}
}

func CreateConfigMap(secrets []types.Secret) map[string]string {
	config := make(map[string]string)
	for _, secret := range secrets {
		config[secret.SecretKey] = secret.SecretValue
	}
	return config
}
