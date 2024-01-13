package main

import (
	"fmt"
	"github.com/6a6ydoping/online-shop/internal/app"
	"github.com/6a6ydoping/online-shop/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
