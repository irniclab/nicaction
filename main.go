package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/irniclab/nicaction/config"
	"github.com/irniclab/nicaction/domainAction"
)

var period int = 0
var nicHandle = ""
var adminHandle = ""
var techHandle = ""
var billHandle = ""
var ns1 = ""
var ns2 = ""
var ns3 = ""
var ns4 = ""

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	var (
		actionFlag      = flag.String("action", "", "the action to perform")
		domainFlag      = flag.String("domain", "", "the domain to perform")
		periodFlag      = flag.Int("period", 0, "the period of the domain (required for 'register' action)")
		nicHandleFlag   = flag.String("nichandle", "", "the nicHandle for the domain (required for 'register' action)")
		configFile      = flag.String("config", "", "path to config file")
		configFlag      = flag.String("configFile", "", "path to config file")
		domainFileFlag  = flag.String("domainfile", "", "path to domain file")
		showConfig      = flag.String("showConfig", "", "Show config")
		adminHandleFlag = flag.String("adminHandle", "", "the admin nicHandle for the domain (required for 'register' action)")
		techHandleFlag  = flag.String("techHandle", "", "the tech nicHandle for the domain (required for 'register' action)")
		billHandleFlag  = flag.String("billHandle", "", "the bill nicHandle for the domain (required for 'register' action)")
		ns1Flag         = flag.String("ns1", "", "ns1 for the domain (required for 'register' action)")
		ns2Flag         = flag.String("ns2", "", "ns2 for the domain (required for 'register' action)")
		ns3Flag         = flag.String("ns3", "", "ns3 for the domain (required for 'register' action)")
		ns4Flag         = flag.String("ns4", "", "ns4 for the domain (required for 'register' action)")
	)

	flag.Parse()

	if *actionFlag == "" && *periodFlag != 0 {
		flag.Usage()
		log.Fatal("action flag is required for 'periodFlag'")
	}

	if *actionFlag == "" && *nicHandleFlag != "" {
		flag.Usage()
		log.Fatal("action flag is required for 'nicHandleFlag'")
	}

	if *actionFlag == "" && *domainFlag != "" {
		flag.Usage()
		log.Fatal("action flag is required for 'domainFlag'")
	}

	if (*actionFlag != "register" && *actionFlag != "renew" && *actionFlag != "bulkRegister" && *actionFlag != "bulkRenew") && *periodFlag != 0 {
		flag.Usage()
		log.Fatal("'periodFlag' only avialable for action register or renew or bulkRegister or bulkRenew")
	}

	if (*actionFlag != "register" && *actionFlag != "bulkRegister" && *actionFlag != "transfer") && *nicHandleFlag != "" {
		flag.Usage()
		log.Fatal("'periodFlag' only avialable for action register or renew or bulkRegister or bulkRenew")
	}

	if *periodFlag != 0 && *periodFlag != 1 && *periodFlag != 5 {
		flag.Usage()
		log.Fatal("period flag value must be either 1 or 5")
	}
	var confPath = ""
	if *configFile != "" {
		confPath = *configFile
	} else {
		confPath = *configFlag
	}
	if confPath == "" {
		log.Fatal("Error please enter config file path")
	}

	// بارگذاری فایل تنظیمات
	conf, err := config.LoadConfig(confPath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err.Error())
	}

	if *periodFlag != 0 {
		period = *periodFlag
	} else {
		period = conf.DefaultPeriod
	}

	if *nicHandleFlag != "" {
		nicHandle = *nicHandleFlag
	} else {
		nicHandle = conf.Nichandle
	}
	if *adminHandleFlag != "" {
		adminHandle = *adminHandleFlag
	} else {
		adminHandle = nicHandle
	}
	if *techHandleFlag != "" {
		techHandle = *techHandleFlag
	} else {
		techHandle = nicHandle
	}
	if *billHandleFlag != "" {
		billHandle = *billHandleFlag
	} else {
		billHandle = conf.MainNicHandle
	}

	if *ns1Flag != "" {
		ns1 = *ns1Flag
	} else {
		ns1 = conf.Ns1
	}
	if *ns2Flag != "" {
		ns2 = *ns2Flag
	} else {
		ns2 = conf.Ns2
	}
	if *ns3Flag != "" {
		ns3 = *ns3Flag
	} else {
		ns3 = conf.Ns3
	}
	if *ns4Flag != "" {
		ns4 = *ns4Flag
	} else {
		ns4 = conf.Ns4
	}
	domain := *domainFlag

	// نمایش مقادیر فعلی تنظیمات فقط در صورتی که --showConfig وارد شده باشد
	if *showConfig != "" {
		conf, err := config.LoadConfig(*configFile)
		if err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}
		fmt.Printf("Current eppAddress: %s\n", conf.EppAddress)
		fmt.Printf("Current nichandle: %s\n", conf.Nichandle)
		fmt.Printf("Current mainNicHandle: %s\n", conf.MainNicHandle)
		fmt.Printf("Current resellerNicHandle: %s\n", conf.ResellerNicHandle)
		fmt.Printf("Current nichandle: %s\n", conf.Nichandle)
		fmt.Printf("Current token: %s\n", conf.Token)
		fmt.Printf("Current ns1: %s\n", conf.Ns1)
		fmt.Printf("Current ns2: %s\n", conf.Ns2)
		fmt.Printf("Current ns3: %s\n", conf.Ns3)
		fmt.Printf("Current ns4: %s\n", conf.Ns4)
		fmt.Printf("Current pre-clTRID: %s\n", conf.PreClTRID)
		fmt.Printf("Current defaultPeriod: %d\n", conf.DefaultPeriod)
		return
	}

	// بررسی و تغییر تنظیمات در صورت وارد شدن پارامتر --config
	if *configFile != "" {
		var newValue string
		newValue = readInput(fmt.Sprintf("Enter new eppAddress (%s): ", conf.EppAddress))
		if newValue != "" {
			conf.EppAddress = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new nichandle (%s): ", conf.Nichandle))
		if newValue != "" {
			conf.Nichandle = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new mainNicHandle (%s): ", conf.MainNicHandle))
		if newValue != "" {
			conf.MainNicHandle = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new resellerNicHandle (%s): ", conf.ResellerNicHandle))
		if newValue != "" {
			conf.ResellerNicHandle = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new token (%s): ", conf.Token))
		if newValue != "" {
			conf.Token = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new authCode (%s): ", conf.AuthCode))
		if newValue != "" {
			conf.AuthCode = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new ns1 (%s): ", conf.Ns1))
		if newValue != "" {
			conf.Ns1 = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new ns2 (%s): ", conf.Ns2))
		if newValue != "" {
			conf.Ns2 = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new ns3 (%s): ", conf.Ns3))
		if newValue != "" {
			conf.Ns3 = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new ns4 (%s): ", conf.Ns4))
		if newValue != "" {
			conf.Ns4 = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new pre-clTRID (%s): ", conf.PreClTRID))
		if newValue != "" {
			conf.PreClTRID = newValue
		}
		newValue = readInput(fmt.Sprintf("Enter new defaultPeriod (%d): ", conf.DefaultPeriod))
		if newValue != "" {
			var err error
			conf.DefaultPeriod, err = strconv.Atoi(newValue)
			if err != nil {
				log.Fatalf("Invalid input: %s", err.Error())
			}
		}

		// ذخیره تغییرات در فایل تنظیمات
		err = config.SaveConfig(*configFile, conf)
		if err != nil {
			log.Fatalf("Error saving config file: %s", err.Error())
		}
	}
	if *configFile == "" {
		switch *actionFlag {
		case "whois":
			res, err := domainAction.Whois(domain, conf)
			if err != nil {
				log.Fatalf("Error is : %s", err.Error())
			}
			log.Printf("Domain: %s\nHolder: %s\nCreationDate: %s\nExpDate: %s\nns1: %s\nns2: %s\nns3: %s\nns4: %s\nDomainStatus: %s ", res.Domain, res.Holder, res.CreateDate, res.ExpDate, res.Ns1, res.Ns2, res.Ns3, res.Ns4, res.DomainStatus)
		case "renewdomainfromfile":
			result := domainAction.RenewDomainListFromPath(*domainFileFlag, period, conf)
			for _, res := range result {
				if res.Result {
					log.Printf("The domain %s has been successfully renewed for %d years.", res.Domain, period)
				} else {
					log.Printf("The domain %s has has error for renew.the error is %s", res.Domain, res.ErrorMsg)
				}

			}
		case "renew":
			if strings.Contains(domain, ",") {
				// کاراکتر , در رشته وجود دارد
				domainList := strings.Split(domain, ",")
				for _, d := range domainList {
					domain = domainAction.FixIrDomainName(d)
					result, err := domainAction.RenewDomain(domain, period, conf)
					if err != nil {
						log.Fatalf("Error is : %s", err.Error())
					}
					if result {
						log.Printf("The domain %s has been successfully renewed for %d years.", domain, period)
					}
				}
			} else {
				domain = domainAction.FixIrDomainName(domain)
				// کاراکتر , در رشته وجود ندارد
				result, err := domainAction.RenewDomain(domain, period, conf)
				if err != nil {
					log.Fatalf("Error is : %s", err.Error())
				}
				if result {
					log.Printf("The domain %s has been successfully renewed for %d years.", domain, period)
				}
			}
		case "register":
			if strings.Contains(domain, ",") {
				// کاراکتر , در رشته وجود دارد
				domainList := strings.Split(domain, ",")
				for _, d := range domainList {
					domain = domainAction.FixIrDomainName(d)
					result, err := domainAction.RenewDomain(domain, period, conf)
					if err != nil {
						log.Fatalf("Error is : %s", err.Error())
					}
					if result {
						log.Printf("The domain %s has been successfully renewed for %d years.", domain, period)
					}
				}
			} else {
				domain = domainAction.FixIrDomainName(domain)
				// کاراکتر , در رشته وجود ندارد
				result, err := domainAction.RegisterDomain(domain, period, nicHandle, adminHandle, techHandle, billHandle, ns1, ns2, ns3, ns4, conf)
				if err != nil {
					log.Fatalf("Error is : %s", err.Error())
				}
				if result {
					log.Printf("The domain %s has been successfully registered for %d years.", domain, period)
				}
			}
		case "DaysToRelease":
			result, err := domainAction.DayToRelease(domain, conf)
			if err != nil {
				log.Fatalf("Error is : %s", err.Error())
			}
			log.Printf("The domain %s has %d days to release..", domain, result)
		default:
			log.Fatalf("Invalid action parameter. Allowed values: register, renew, delete, transfer, bulkRegister, bulkRenew")
		}
	}
}

func readInput(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
