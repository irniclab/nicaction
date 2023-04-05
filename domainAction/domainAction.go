package domainAction

import (
	"log"
	"time"

	"github.com/irniclab/nicaction/types"
	"github.com/irniclab/nicaction/xmlRequest"
)

func registerDomain(domain string, nicHanle string, period int, ns1 string, ns2 string, preclTRID string, token string) (bool, error) {
	return true, nil
}

func RenewDomain(domain string, period int, conf types.Config) (bool, error) {
	dt, error := Whois(domain, conf)
	//log.Printf("Exp Date is : " + dt.ExpDate.String())
	if error != nil {
		log.Fatalf("Error in domain whois %s", error.Error())
	}
	reqStr := xmlRequest.DomainRenewXml(domain, xmlRequest.FormatDateString(dt.ExpDate), period, conf)
	resp, error := xmlRequest.SendXml(reqStr, conf)
	if error != nil {
		log.Fatalf("Error in renew domain from nic %s", error.Error())
	}
	result, error := xmlRequest.ParseDomainRenewResponse(resp)
	if error != nil {
		log.Fatalf("Error in renew domain from nic %s", error.Error())
	}
	return result, nil
}

func Whois(domain string, conf types.Config) (types.DomainType, error) {
	var result *types.DomainType
	reqStr := xmlRequest.DomainWhoisXml(domain, conf)
	//log.Print(reqStr)
	resStr, error := xmlRequest.SendXml(reqStr, conf)
	if error != nil {
		log.Fatalf("Error in Whois from nic %s", error.Error())
	}
	result, error = xmlRequest.ParseDomainInfoType(resStr)
	if error != nil {
		log.Fatalf("Error in Fetch result of Whois from nic %s", error.Error())
	}
	//fmt.Print("Whois domain : " + domain)
	return *result, error
}

func DayToRelease(domain string, conf types.Config) (int, error) {
	var result *types.DomainType
	reqStr := xmlRequest.DomainWhoisXml(domain, conf)
	//log.Print(reqStr)
	resStr, error := xmlRequest.SendXml(reqStr, conf)
	if error != nil {
		log.Fatalf("Error in Whois from nic %s", error.Error())
	}
	result, error = xmlRequest.ParseDomainInfoType(resStr)
	if error != nil {
		log.Fatalf("Error in Fetch result of Whois from nic %s", error.Error())
	}
	//fmt.Print("Whois domain : " + domain)
	return subtractDays(result.ExpDate), error
}

func subtractDays(t time.Time) int {
	today := time.Now()
	diff := today.Sub(t)
	return int(diff.Hours()/24) - 59
}
