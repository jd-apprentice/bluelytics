package functions

import (
	consts "dolar/src/constants"
	"dolar/src/types"
	handlers "dolar/src/utils"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func InitDotEnv() {

	var environment = os.Getenv("INFISICAL_ENVIRONMENT")

	if strings.ToUpper(environment) == "DEV" {
		return
	}

	err := godotenv.Load()
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
