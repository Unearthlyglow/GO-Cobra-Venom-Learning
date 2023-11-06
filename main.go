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


/*
1. Clear is better than clever
2. A little copying is better than a little dependency 
3. Concurrency is not paralleism
4. Channels orchestrate; mutexes serialize 
5. Don’t communicate by sharing memory, share memory by communicating 
6. Make the zero value useful 
7. The bigger the interface, the weaker the abstraction
8. Errors are values
9. Don’t just check errors, handle them gracefully.
*/