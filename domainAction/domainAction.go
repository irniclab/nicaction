package domainAction

import (
	"fmt"

	"github.com/irniclab/nicaction/types"
)

func registerDomain(domain string, nicHanle string, period int, ns1 string, ns2 string, preclTRID string, token string) (bool, error) {
	return true, nil
}

func Whois(domain string, conf types.Config) (string, error) {
	fmt.Print("Whois domain : " + domain)
	return "", nil
}
