package main

import (
	"ripProject/internal/app/apiserver"
	"ripProject/internal/config"
)

func main() {
	cfg, err := config.ReadConfig("C:\\Users\\mrclo\\Desktop\\RIP\\configs\\config.yaml")
	if err != nil {
		panic(err)
	}
	err = apiserver.Run(cfg)
	if err != nil {
		panic(err)
	}
}
