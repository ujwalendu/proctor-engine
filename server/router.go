package server

import (
	"fmt"
	"net/http"

	"github.com/gojekfarm/proctor-engine/jobs/execution"
	"github.com/gojekfarm/proctor-engine/jobs/logs"
	"github.com/gojekfarm/proctor-engine/jobs/metadata"
	"github.com/gojekfarm/proctor-engine/jobs/secrets"
	"github.com/gojekfarm/proctor-engine/kubernetes"
	"github.com/gojekfarm/proctor-engine/redis"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

	redisClient := redis.NewClient()

	kubeConfig := kubernetes.KubeConfig()
	kubeClient := kubernetes.NewClient(kubeConfig)

	metadataStore := metadata.NewStore(redisClient)
	secretsStore := secrets.NewStore(redisClient)

	jobExecutioner := execution.NewExecutioner(kubeClient, metadataStore, secretsStore)
	jobLogger := logs.NewLogger(kubeClient)
	jobMetadataHandler := metadata.NewHandler(metadataStore)
	jobSecretsHandler := secrets.NewHandler(secretsStore)

	router.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	router.HandleFunc("/jobs/execute", jobExecutioner.Handle()).Methods("POST")
	router.HandleFunc("/jobs/logs", jobLogger.Stream()).Methods("GET")
	router.HandleFunc("/jobs/metadata", jobMetadataHandler.HandleSubmission()).Methods("POST")
	router.HandleFunc("/jobs/metadata", jobMetadataHandler.HandleBulkDisplay()).Methods("GET")
	router.HandleFunc("/jobs/secrets", jobSecretsHandler.HandleSubmission()).Methods("POST")
}
