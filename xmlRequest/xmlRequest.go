package xmlRequest

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/irniclab/nicaction/nicResponse"
	ptime "github.com/yaa110/go-persian-calendar"

	"github.com/irniclab/nicaction/types"
)

type DomainWhoisResponse struct {
	XMLName  xml.Name `xml:"epp"`
	Response struct {
		XMLName xml.Name `xml:"response"`
		Result  struct {
			Code string `xml:"code,attr"`
			Msg  string `xml:"msg"`
		} `xml:"result"`
		ResData struct {
			XMLName xml.Name `xml:"resData"`
			InfData struct {
				XMLName  xml.Name `xml:"infData"`
				Name     string   `xml:"name"`
				Statuses []struct {
					Value string `xml:"s,attr"`
				} `xml:"status"`
				Contacts []struct {
					Type  string `xml:"type,attr"`
					Value string `xml:",chardata"`
				} `xml:"contact"`
				Ns     []string `xml:"ns>hostAttr>hostName"`
				CrDate string   `xml:"crDate"`
				UpDate string   `xml:"upDate"`
				ExDate string   `xml:"exDate"`
				Holder string   `xml:"contact[type=holder]"`
			} `xml:"infData"`
		} `xml:"resData"`
	} `xml:"response"`
}

func DomainWhoisXml(domain string, config types.Config) string {
	xml := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
				<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
                	<command>
                    	<info>
						<domain:info xmlns:domain="http://epp.nic.ir/ns/domain-1.0">
                            	<domain:name>%s</domain:name>
								<domain:authInfo>
 									<domain:pw>%s</domain:pw>
 								</domain:authInfo>
                        	</domain:info>
                    	</info>
                    	<clTRID>%s</clTRID>
                	</command>
            	</epp>`
	return fmt.Sprintf(xml, domain, config.AuthCode, getPreClTRID(config))
}

func DomainRenewXml(domain string, expDate string, period int, config types.Config) string {
	xml := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
				<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
					<command>
						<renew>
							<domain:renew xmlns:domain="http://epp.nic.ir/ns/domain-1.0">
								<domain:name>%s</domain:name>
								<domain:curExpDate>%s</domain:curExpDate>
								<domain:period unit="m">%d</domain:period>
								<domain:authInfo>
									<domain:pw>%s</domain:pw>
								</domain:authInfo>
							</domain:renew>
						</renew>
						<clTRID>%s</clTRID>
					</command>
   				</epp>`
	return fmt.Sprintf(xml, domain, expDate, period, config.AuthCode, getPreClTRID(config))
}

func getPreClTRID(config types.Config) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100000)
	randomStr := fmt.Sprintf("%05d", randomNum)
	return config.PreClTRID + "-" + randomStr
}

func SendXml(xml string, config types.Config) (string, error) {
	//log.Printf("Raw Result is : %s", xml)
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.EppAddress, bytes.NewBufferString(xml))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+config.Token)
	//req.Header.Set("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func ParseDomainInfoType(xmlContent string) (*types.DomainType, error) {
	var di DomainWhoisResponse
	//log.Printf("Raw Result is : %s", xmlContent)
	if err := xml.Unmarshal([]byte(xmlContent), &di); err != nil {
		return nil, err
	}

	if di.Response.Result.Code == "2502" {
		return nil, errors.New("Session limit exceeded; server closing connection")
	}
	log.Printf("di.Response.ResData.InfData.ExDate : %s", di.Response.ResData.InfData.ExDate)

	d := &types.DomainType{
		Domain:       di.Response.ResData.InfData.Name,
		Holder:       di.Response.ResData.InfData.Holder,
		DomainStatus: make([]string, len(di.Response.ResData.InfData.Statuses)),
		Ns1:          "",
		Ns2:          "",
		Ns3:          "",
		Ns4:          "",
		ExpDate:      time.Time{},
		CreateDate:   time.Time{},
		UpdateDate:   time.Time{},
	}

	copy(d.DomainStatus, func() []string {
		s := make([]string, len(di.Response.ResData.InfData.Statuses))
		for i := range s {
			s[i] = di.Response.ResData.InfData.Statuses[i].Value
		}
		return s
	}())

	if len(di.Response.ResData.InfData.Ns) > 0 {
		d.Ns1 = di.Response.ResData.InfData.Ns[0]
	}
	if len(di.Response.ResData.InfData.Ns) > 1 {
		d.Ns2 = di.Response.ResData.InfData.Ns[1]
	}
	if len(di.Response.ResData.InfData.Ns) > 2 {
		d.Ns3 = di.Response.ResData.InfData.Ns[2]
	}
	if len(di.Response.ResData.InfData.Ns) > 3 {
		d.Ns4 = di.Response.ResData.InfData.Ns[3]
	}

	if t, err := time.Parse("2006-01-02T15:04:05", di.Response.ResData.InfData.ExDate); err == nil {
		d.ExpDate = t
		d.LockDate = addDays(t, 30)
		d.ReleaseDate = addDays(t, 60)
	}
	if t, err := time.Parse("2006-01-02T15:04:05", di.Response.ResData.InfData.CrDate); err == nil {
		d.CreateDate = t
	}
	if t, err := time.Parse("2006-01-02T15:04:05", di.Response.ResData.InfData.UpDate); err == nil {
		d.UpdateDate = t
	}
	log.Printf("d.ExpDate format : %s", d.ExpDate.Format("2006-01-02"))
	log.Printf("d.ExpDate String : %s", d.ExpDate.String())
	return d, nil
}

func ParseDomainRenewResponse(xmlContent string) (bool, error) {
	var di nicResponse.GeneralResponse
	if err := xml.Unmarshal([]byte(xmlContent), &di); err != nil {
		return false, err
	}

	if di.Result.Code == "1001" {
		return true, nil
	}
	if di.Result.Code == "2502" {
		return false, errors.New("Session limit exceeded; server closing connection")
	} else if di.Result.Code == "2000" {
		return false, errors.New("Request has a wrong format")
	} else if di.Result.Code == "2001" {
		return false, errors.New("Request has a wrong command")
	} else if di.Result.Code == "2003" {
		return false, errors.New("mandatory fields in the request is missing")
	} else if di.Result.Code == "2004" {
		return false, errors.New("The value sent in the request is out of the acceptable range")
	} else if di.Result.Code == "2005" {
		return false, errors.New("The value sent in the request has an incorrect format")
	} else if di.Result.Code == "2101" {
		return false, errors.New("The request sent is incorrect")
	} else if di.Result.Code == "2104" {
		return false, errors.New("Billing failure")
	} else if di.Result.Code == "2105" {
		return false, errors.New("Object is not eligible for renewal")
	} else if di.Result.Code == "2200" {
		return false, errors.New("Authentication error")
	} else if di.Result.Code == "2201" {
		return false, errors.New("Authorization error")
	} else if di.Result.Code == "2202" {
		return false, errors.New("Invalid authorization information")
	} else if di.Result.Code == "2303" {
		return false, errors.New("Object not does exist")
	} else if di.Result.Code == "2304" {
		return false, errors.New("Object status prohibits operation")
	} else if di.Result.Code == "2400" {
		return false, errors.New("Command failed")
	} else {
		return false, errors.New("An unknown error has occurred")
	}
}

func FormatDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func addDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func convertToJalali(t time.Time) (ptime.Time, error) {
	pt := ptime.New(t)
	return pt, nil
}
