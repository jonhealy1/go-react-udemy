package main

import "flag"
import "fmt"

const version = "1.0.0"

type config struct {
	port int
	env string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listne on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse()

	fmt.Println("Running")
}