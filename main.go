package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/irniclab/nicaction/config"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to the config file")
	showConfig := flag.String("showConfig", "", "Show application config")
	flag.Parse()

	// بارگذاری فایل تنظیمات
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())
	}

	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --config وارد شده باشد
	if flag.Arg(0) == "--config" {
		fmt.Printf("Current eppAddress: %s\n", conf.EppAddress)
		fmt.Printf("Current nichandle: %s\n", conf.Nichandle)
		fmt.Printf("Current token: %s\n", conf.Token)
		fmt.Printf("Current ns1: %s\n", conf.Ns1)
		fmt.Printf("Current ns2: %s\n", conf.Ns2)
		fmt.Printf("Current pre-clTRID: %s\n", conf.PreClTRID)

		// بررسی و تغییر تنظیمات
		var newValue string
		switch flag.Arg(1) {
		case "eppAddress":
			newValue = readInput(fmt.Sprintf("Enter new eppAddress (%s): ", conf.EppAddress))
			if newValue != "" {
				conf.EppAddress = newValue
			}
		case "nichandle":
			newValue = readInput(fmt.Sprintf("Enter new nichandle (%s): ", conf.Nichandle))
			if newValue != "" {
				conf.Nichandle = newValue
			}
		case "token":
			newValue = readInput(fmt.Sprintf("Enter new token (%s): ", conf.Token))
			if newValue != "" {
				conf.Token = newValue
			}
		case "ns1":
			newValue = readInput(fmt.Sprintf("Enter new ns1 (%s): ", conf.Ns1))
			if newValue != "" {
				conf.Ns1 = newValue
			}
		case "ns2":
			newValue = readInput(fmt.Sprintf("Enter new ns2 (%s): ", conf.Ns2))
			if newValue != "" {
				conf.Ns2 = newValue
			}
		case "pre-clTRID":
			newValue = readInput(fmt.Sprintf("Enter new pre-clTRID (%s): ", conf.PreClTRID))
			if newValue != "" {
				conf.PreClTRID = newValue
			}
		default:
			log.Fatal("Invalid config option")
		}

		// ذخیره تغییرات در فایل تنظیمات
		err = config.SaveConfig(*configPath, conf)
		if err != nil {
			log.Fatalf("Error saving config file: %s", err.Error())
		}
	}
}

func readInput(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
