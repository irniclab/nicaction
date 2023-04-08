package domainAction

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
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
	if hasPendingRenewStatus(dt) {
		return false, errors.New("DomainPendingRenew")
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

func RenewDomainList(domainList []string, period int, conf types.Config) []types.DomainListResult {
	var domainRenewResults []types.DomainListResult
	for _, dm := range domainList {
		res, error := RenewDomainWithError(dm, period, conf)
		if error != nil {
			result := types.DomainListResult{
				Domain:   dm,
				Duration: period,
				Result:   false,
				ErrorMsg: error.Error(),
			}
			domainRenewResults = append(domainRenewResults, result)
		} else if !res {
			result := types.DomainListResult{
				Domain:   dm,
				Duration: period,
				Result:   false,
				ErrorMsg: "Unknown Error.",
			}
			domainRenewResults = append(domainRenewResults, result)
		} else {
			result := types.DomainListResult{
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

func RenewDomainListFromPath(filePath string, period int, conf types.Config) []types.DomainListResult {
	var domainRenewResults []types.DomainListResult
	var domainPendingRenew []string
	domainList, error := readDomainListFromFile(filePath)
	if error != nil {
		log.Fatalf("Error in reading files %s", error.Error())
	}
	for _, dm := range domainList {
		if dm != "" && len(dm) >= 3 {
			res, error := RenewDomainWithError(FixIrDomainName(dm), period, conf)
			if error != nil {
				result := types.DomainListResult{
					Domain:   dm,
					Duration: period,
					Result:   false,
					ErrorMsg: error.Error(),
				}
				domainRenewResults = append(domainRenewResults, result)
				if error.Error() == "DomainPendingRenew" {
					domainPendingRenew = append(domainPendingRenew, dm)
				}
			} else if !res {
				result := types.DomainListResult{
					Domain:   dm,
					Duration: period,
					Result:   false,
					ErrorMsg: "Unknown Error.",
				}
				domainRenewResults = append(domainRenewResults, result)
			} else {
				result := types.DomainListResult{
					Domain:   dm,
					Duration: period,
					Result:   true,
					ErrorMsg: "",
				}
				domainRenewResults = append(domainRenewResults, result)
			}
		}

	}
	successList := getSuccessListFromListResult(domainRenewResults)
	remainList := FilterSlice(domainList, successList)
	remainList = FilterSlice(remainList, domainPendingRenew)
	writeDomainListToFile(filePath, remainList)
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

func readDomainListFromFile(filePath string) ([]string, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileContent), "\n")
	var validLines []string
	for _, line := range lines {
		if line != "" && len(line) >= 3 {
			validLines = append(validLines, line)
		}
	}
	return validLines, nil
}

func writeDomainListToFile(filePath string, domainList []string) error {
	outputContent := strings.Join(domainList, "\n")
	err := ioutil.WriteFile(filePath, []byte(outputContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func getSuccessListFromListResult(dominwResultList []types.DomainListResult) []string {
	result := []string{}
	for _, dr := range dominwResultList {
		if dr.Result {
			result = append(result, dr.Domain)
		}
	}
	return result
}

func FilterSlice(s1 []string, s2 []string) []string {
	result := []string{}
	for _, str := range s1 {
		if !containsString(s2, str) {
			result = append(result, str)
		}
	}
	return result
}

func FixIrDomainName(domain string) string {
	domain = strings.ToLower(strings.TrimSpace(domain))
	if !strings.HasSuffix(domain, ".ir") {
		domain = domain + ".ir"
	}
	return domain
}

func hasPendingRenewStatus(domain types.DomainType) bool {
	for _, status := range domain.DomainStatus {
		if status == "pendingRenew" {
			return true
		}
	}
	return false
}
