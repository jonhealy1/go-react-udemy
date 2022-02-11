package main

import "flag"
import "fmt"
import "net/http"
import "log"
import "encoding/json"

const version = "1.0.0"

type config struct {
	port int
	env string
}

type AppStatus struct {
	Status string
	Environment string
	Version string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listne on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse()

	fmt.Println("Running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status: "Available",
			Environment: cfg.env,
			Version: version,
		}
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}

}