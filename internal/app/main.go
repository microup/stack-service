package app

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"stack-service/internal/config"
	"stack-service/internal/server"
	"stack-service/internal/stack"
	"stack-service/internal/types"
)

func Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("catch err: %v", err) //nolint:forbidigo
		}
	}()

	configuration := initializeConfig()
	if configuration == nil {
		log.Panic("configuration is nil")
	}

	stack := stack.New()
	err := stack.LoadStack()
	if err != nil {
		log.Panic(err)
	}

	defer stack.Save()

	server := server.New(configuration, stack)
	server.Run()
}

func initializeConfig() *config.Config {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configFile := os.Getenv("ConfigFile")

	if configFile == "." {
		configFile = "../../config"
	}

	cfg := config.New()

	if err := cfg.Load(configFile); err != nil {
		log.Fatalf("%v", err)
	}

	if resultCode := cfg.InitProxyPort(); resultCode != types.ResultOK {
		log.Fatalf("can't init proxy: %s", resultCode.ToStr())
	}

	return cfg
}
