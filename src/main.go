package main

import (
	consts "dolar/src/constants"
	"dolar/src/functions"
	handlers "dolar/src/utils"
	"fmt"
	"os"
)

func main() {

	functions.InitDotEnv()

	var projectId = os.Getenv("INFISICAL_PROJECT_ID")
	var environment = os.Getenv("INFISICAL_ENVIRONMENT")

	secrets, err := functions.GetSecrets(projectId, environment)

	if err != nil {
		handlers.LogFatalFmt(consts.FailtedSecrets, err)
	}

	config := functions.CreateConfigMap(secrets)

	data, err := functions.GetData(config["API_ENDPOINT"])

	if err != nil {
		handlers.LogFatalFmt(consts.FailedFetch, err)
	}

	table := functions.CreateTable()
	table.Append([]string{"💵 Oficial", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.Oficial.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💵 Blue", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.Blue.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💶 Oficial_Euro", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.OficialEuro.ValueAvg, consts.ArsLabel)})
	table.Append([]string{"💶 Blue_euro", fmt.Sprintf("%s %.2f %s", consts.SackEmoji, data.BlueEuro.ValueAvg, consts.ArsLabel)})

	table.Render()
}
