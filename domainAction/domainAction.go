package domainAction
import {
	"github.com/irniclab/nicaction/xmlRequest"
	"github.com/irniclab/nicaction/types"
}

var a = 1

func registerDomain(domain string, nicHanle string, period int, ns1 string, ns2 string, preclTRID string, token string) (bool, error) {
	return true, nil
}

func whois(domain string, conf types.Config) (string, error) {
	fmt.Print("Whois domain : " + domain)
	return "", nil
}
