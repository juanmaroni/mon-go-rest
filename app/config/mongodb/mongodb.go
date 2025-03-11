package mongodb

import "mon-go-rest/config/utils"

type Mongo struct {
	Uri string
}

var MongoDb *Mongo

func LoadConfig() error {
	mongodbUri, err := utils.GetEnvVar("MONGODB_URI")

	if err != nil {
		return err
	}

	MongoDb = &Mongo {
		Uri: mongodbUri,
	}

	return nil
}
