/*
Copyright © 2023 Award Malisi
In a Cobra app, typically the main.go file is very bare. It serves one purpose: initializing Cobra.


*/

package main

import (
	"cobraScratch/cmd"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	vp := viper.New()

	vp.SetConfigName("cobra")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./util")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("file changed: %s\n", in.Name)
	})

	vp.WatchConfig()

	cmd.Execute()

}
