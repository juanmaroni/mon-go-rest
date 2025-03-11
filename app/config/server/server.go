package server

import "mon-go-rest/config/utils"

type ApiServer struct {
	Uri string
}

var Server *ApiServer

func LoadConfig() error {
	apiServerUri, err := utils.GetEnvVar("API_SERVER_URI")

	if err != nil {
		return err
	}

	Server = &ApiServer{
		Uri: apiServerUri,
	}

	return nil
}
