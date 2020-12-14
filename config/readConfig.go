package config

import (
	"fmt"
	"github.com/P4Networking/api-gw/openapi"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

func ReadConfig() {

	fmt.Printf("Start to print the config content which has benn read.\n")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Printf("Do not found config.json. Please make sure the file exist.\n")
		} else {
			// Config file was found but another error was produced
			fmt.Errorf("Got error: %v\n", err)
			return
		}
	}

	fmt.Printf("======================== %v Data ========================\n", "Auth")
	viper.UnmarshalKey("user", &openapi.UserDB)
	for k, v := range openapi.UserDB {
		fmt.Printf("%vst data:\n", k+1)
		spew.Dump(v)
	}

	fmt.Printf("======================== %v Data ========================\n", "Subscriber")
	viper.UnmarshalKey("subscriber", &openapi.Subscribers)
	for k, v := range openapi.Subscribers {
		fmt.Printf("%vst data:\n", k+1)
		fmt.Printf("PLMN ID:%v, UE ID:%v ,Auth Method:%v \n", v.PlmnId, v.UeId, v.AuthenticationSubscription.AuthenticationMethod)

	}

	fmt.Printf("======================== %v Data ========================\n", "UE")
	viper.UnmarshalKey("ue", &openapi.UEs)
	for k, v := range openapi.UEs {
		fmt.Printf("%vst data:\n", k+1)
		spew.Dump(v)
	}

	fmt.Printf("======================== %v Data ========================\n", "RAN")
	viper.UnmarshalKey("ran", &openapi.RANs)

	for k, v := range openapi.RANs {
		fmt.Printf("%vst data:\n", k+1)
		spew.Dump(v)
	}
}
