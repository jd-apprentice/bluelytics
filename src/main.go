package main

import (
	consts "dolar/src/constants"
	fn "dolar/src/functions"
	handlers "dolar/src/utils"
	"fmt"
	"os"
)

func main() {

	fn.InitDotEnv()

	var projectId = os.Getenv("INFISICAL_WORKSPACE_ID")
	var environment = os.Getenv("INFISICAL_ENVIRONMENT")

	if environment == "" {
		environment = "dev"
	}

	secrets, err := fn.GetSecrets(projectId, environment)

	if err != nil {
		handlers.LogFatalFmt(consts.FailtedSecrets, err)
	}

	config := fn.CreateConfigMap(secrets)

	data, err := fn.GetData(config["API_ENDPOINT"])

	if err != nil {
		handlers.LogFatalFmt(consts.FailedFetch, err)
	}

	table := fn.CreateTable()
	table.Append([]string{"💵 Oficial", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.Oficial.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💵 Blue", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.Blue.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💶 Oficial_Euro", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.OficialEuro.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💶 Blue_euro", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.BlueEuro.ValueAvg, consts.ArsLabel)})

	table.Render()
}
