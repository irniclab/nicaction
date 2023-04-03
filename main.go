package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/irniclab/nicaction/config"
)

func main() {
	configFile = flag.String("config", "config.json", "Config file for nic action")
	showConfig = flag.Bool("showConfig", false, "Show current config values")
	defaultPeriod = flag.Int("defaultPeriod", 1, "Default period for pre-registration (in years)")
	action = flag.String("action", "", "Action to perform (register, renew, delete, transfer, bulkRegister, bulkRenew)")

	// بارگذاری فایل تنظیمات
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())
	}

	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --showConfig وارد شده باشد
	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --showConfig وارد شده باشد
	if *showConfig {
		conf, err := readConfig(*configFile)
		if err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}
		fmt.Printf("Current eppAddress: %s\n", conf.EppAddress)
		fmt.Printf("Current nichandle: %s\n", conf.Nichandle)
		fmt.Printf("Current token: %s\n", conf.Token)
		fmt.Printf("Current ns1: %s\n", conf.Ns1)
		fmt.Printf("Current ns2: %s\n", conf.Ns2)
		fmt.Printf("Current pre-clTRID: %s\n", conf.PreClTRID)
		fmt.Printf("Current defaultPeriod: %d\n", conf.DefaultPeriod)
		return
	}

	// بررسی و تغییر تنظیمات در صورت وارد شدن پارامتر --config
	if *configFile {
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
		case "defaultPeriod":
			newValue = readInput(fmt.Sprintf("Enter new defaultPeriod (%d): ", conf.DefaultPeriod))
			if newValue != "" {
				var err error
				conf.DefaultPeriod, err = strconv.Atoi(newValue)
				if err != nil {
					log.Fatalf("Invalid input: %s", err.Error())
				}
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
	switch *action {
	case "register":
		registerDomain(config)
	case "renew":
		renewDomain(config)
	case "delete":
		deleteDomain(config)
	case "transfer":
		transferDomain(config)
	case "bulkRegister":
		bulkRegister(config, *defaultPeriod)
	case "bulkRenew":
		bulkRenew(config, *defaultPeriod)
	default:
		log.Fatalf("Invalid action parameter. Allowed values: register, renew, delete, transfer, bulkRegister, bulkRenew")
	}
}

func readInput(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
