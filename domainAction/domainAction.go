package domainAction

import (
	"log"

	"github.com/irniclab/nicaction/types"
	"github.com/irniclab/nicaction/xmlRequest"
)

func registerDomain(domain string, nicHanle string, period int, ns1 string, ns2 string, preclTRID string, token string) (bool, error) {
	return true, nil
}

func Whois(domain string, conf types.Config) (types.DomainType, error) {
	var result types.DomainType
	reqStr := xmlRequest.DomainWhoisXml(string, conf)
	resStr, error := xmlRequest.SendXml(reqStr, conf)
	if error != nil {
		log.Fatalf("Error in Whois from nic %s", error.Error())
	}
	result, error = xmlRequest.ParseDomainInfoType(resStr)
	if error != nil {
		log.Fatalf("Error in Fetch result of Whois from nic %s", error.Error())
	}
	//fmt.Print("Whois domain : " + domain)
	return result, error
}
