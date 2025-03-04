package main

import (
	"fmt"
	"gator/internal/config"
)

func main()  {
	configFile := config.Read()
	configFile.SetUser("musasda")
	
	updatedConfigFile := config.Read()

	fmt.Printf("'%v'\n", updatedConfigFile.DBUrl)
}