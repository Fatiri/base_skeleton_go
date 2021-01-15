package main

import (
	"sync"

	"go.elastic.co/apm"

	"github.com/base_skeleton_go/shared/logger"

	"github.com/base_skeleton_go/config"
	"github.com/base_skeleton_go/src"
	dotenv "github.com/joho/godotenv"
)

func main() {
	apm.DefaultTracer.Flush(nil)
	err := dotenv.Load("./config/.env")
	if err != nil {
		panic(".env is not loaded properly")
	}

	// init logger
	logger.InitZap()
	cfg := config.NewConfig()
	server := src.InitServer(cfg)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	// Wait All services to end
	wg.Wait()
}
