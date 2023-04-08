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
	reqStr := xmlRequest.DomainRenewXml(domain, xmlRequest.FormatDateString(dt.ExpDate), period*12, conf)
	log.Printf("Request is : %s", reqStr)
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

func RenewDomainWithError(domain string, period int, conf types.Config) (bool, error) {
	dt, error := Whois(domain, conf)
	//log.Printf("Exp Date is : " + dt.ExpDate.String())
	if error != nil {
		return false, error
	}
	reqStr := xmlRequest.DomainRenewXml(domain, xmlRequest.FormatDateString(dt.ExpDate), period*12, conf)
	resp, error := xmlRequest.SendXml(reqStr, conf)
	if error != nil {
		return false, error
	}
	result, error := xmlRequest.ParseDomainRenewResponse(resp)
	if error != nil {
		return false, error
	}
	return result, nil
}

func RenewDomainList(domainList []string, period int, conf types.Config) []types.DomainRenewResult {
	var domainRenewResults []types.DomainRenewResult
	for _, dm := range domainList {
		res, error := RenewDomainWithError(dm, period, conf)
		if error != nil {
			result := types.DomainRenewResult{
				Domain:   dm,
				Duration: period,
				Result:   false,
				ErrorMsg: error.Error(),
			}
			domainRenewResults = append(domainRenewResults, result)
		} else if !res {
			result := types.DomainRenewResult{
				Domain:   dm,
				Duration: period,
				Result:   false,
				ErrorMsg: "Unknown Error.",
			}
			domainRenewResults = append(domainRenewResults, result)
		} else {
			result := types.DomainRenewResult{
				Domain:   dm,
				Duration: period,
				Result:   true,
				ErrorMsg: "",
			}
			domainRenewResults = append(domainRenewResults, result)
		}
	}
	return domainRenewResults
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
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 60)
	diff := t.Sub(today)
	return int(diff.Hours() / 24)
}
