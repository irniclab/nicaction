package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

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

	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --showConfig وارد شده باشد
	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --showConfig وارد شده باشد
	if *showConfig {
		fmt.Printf("Current eppAddress: %s\n", conf.EppAddress)
		fmt.Printf("Current nichandle: %s\n", conf.Nichandle)
		fmt.Printf("Current token: %s\n", conf.Token)
		fmt.Printf("Current ns1: %s\n", conf.Ns1)
		fmt.Printf("Current ns2: %s\n", conf.Ns2)
		fmt.Printf("Current pre-clTRID: %s\n", conf.PreClTRID)
		fmt.Printf("Default period for pre-registration: %d years\n", conf.DefaultPeriod)
		return // اجرای برنامه را به اتمام برسانید
	}

	// بررسی و تغییر تنظیمات در صورت وارد شدن پارامتر --config
	if flag.Arg(0) == "--config" {
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
	} else if flag.Arg(0) == "--action" {
		switch action {
		case "register":
			// در اینجا برای عملیات register کد خود را قرار دهید
			fmt.Println("Register operation")
		case "renew":
			// در اینجا برای عملیات renew کد خود را قرار دهید
			fmt.Println("Renew operation")
		case "delete":
			// در اینجا برای عملیات delete کد خود را قرار دهید
			fmt.Println("Delete operation")
		case "transfer":
			// در اینجا برای عملیات transfer کد خود را قرار دهید
			fmt.Println("Transfer operation")
		case "bulkRegister":
			// در اینجا برای عملیات bulkRegister کد خود را قرار دهید
			fmt.Println("Bulk register operation")
		case "bulkRenew":
			// در اینجا برای عملیات bulkRenew کد خود را قرار دهید
			fmt.Println("Bulk renew operation")
		case "help":
			// نمایش راهنمایی پارامترها
			fmt.Println("Available actions:")
			fmt.Println("- register")
			fmt.Println("- renew")
			fmt.Println("- delete")
			fmt.Println("- transfer")
			fmt.Println("- bulkRegister")
			fmt.Println("- bulkRenew")
			fmt.Println("- help")
		default:
			// پارامتر اشتباه وارد شده است، خطا نمایش داده شود
			fmt.Println("Invalid action parameter. Allowed parameters: register, renew, delete, transfer, bulkRegister, bulkRenew, help")
			return
		}
	}
}

func readInput(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
