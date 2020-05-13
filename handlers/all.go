package handlers

import (
	"net/http"

	"github.com/a-h/aws-go-api/db"
	"github.com/a-h/aws-go-api/handlers/organisation"
	"github.com/a-h/aws-go-api/handlers/user"
	"github.com/a-h/aws-go-api/log"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Configuration struct {
	Region           string
	TableName        string
	UseDynamoDBLocal bool
}

func All(conf Configuration) http.Handler {
	organisationStore, err := db.NewOrganisationStore(conf.Region, conf.TableName)
	if err != nil {
		log.Default.Fatal("failed to create organisation store", zap.Error(err))
	}
	userStore, err := db.NewUserStore(conf.Region, conf.TableName)
	if err != nil {
		log.Default.Fatal("failed to create user store", zap.Error(err))
	}

	if conf.UseDynamoDBLocal {
		organisationStore.Client.Endpoint = "http://localhost:8000"
		userStore.Client.Endpoint = "http://localhost:8000"
	}

	user := user.NewHandler(userStore, mux.Vars)
	organisation := organisation.NewHandler(organisationStore, mux.Vars)

	//TODO: Add authentication.
	m := mux.NewRouter()
	m.HandleFunc("/organisation/{id}/details", organisation.DetailsGet).Methods("GET")
	m.HandleFunc("/user/{id}/details", user.DetailsGet).Methods("GET")

	return log.Responses(m)
}
