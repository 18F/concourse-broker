package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/18F/concourse-broker/broker"
	"github.com/18F/concourse-broker/concourse"
	"github.com/18F/concourse-broker/config"
	"github.com/18F/concourse-broker/logger"
	"github.com/pivotal-cf/brokerapi"
)

func loadServices() ([]brokerapi.Service, error) {
	var service brokerapi.Service
	buf, err := ioutil.ReadFile("./catalog.json")
	if err != nil {
		return []brokerapi.Service{}, err
	}
	err = json.Unmarshal(buf, &service)
	if err != nil {
		return []brokerapi.Service{}, err
	}
	return []brokerapi.Service{service}, nil
}

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}
	logger, err := logger.NewLogger("concourse-broker", env)
	if err != nil {
		log.Fatalln(err)
	}
	credentials := brokerapi.BrokerCredentials{
		Username: env.BrokerUsername,
		Password: env.BrokerPassword,
	}
	services, err := loadServices()
	if err != nil {
		log.Fatalln(err)
	}

	// Update team credentials if 0th instance
	appID := os.Getenv("CF_INSTANCE_INDEX")
	if appID == "0" || appID == "" {
		concourseClient := concourse.NewClient(env, logger)
		if err := concourseClient.UpdateTeams(); err != nil {
			log.Fatalln(err)
		}
	}

	serviceBroker := broker.New(services, logger, env)
	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(fmt.Sprintf(":%s", env.Port), nil)
}
