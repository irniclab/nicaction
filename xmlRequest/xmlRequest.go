package xmlRequest

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/yaa110/go-jalali"

	"github.com/irniclab/nicaction/types"
)

type domainInfo struct {
	Name     string `xml:"infData>name"`
	Statuses []struct {
		Value string `xml:"s,attr"`
	} `xml:"infData>status"`
	Contacts []struct {
		Type  string `xml:"type,attr"`
		Value string `xml:",chardata"`
	} `xml:"infData>contact"`
	Ns     []string `xml:"infData>ns>hostAttr>hostName"`
	CrDate string   `xml:"infData>crDate"`
	UpDate string   `xml:"infData>upDate"`
	ExDate string   `xml:"infData>exDate"`
	Holder string   `xml:"infData>contact[type=holder]"`
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"result"`
}

func domainWhoisXml(domain string, config types.Config) string {
	xml := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
            <epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
                <command>
                    <info>
                        <domain:info xmlns:domain="%s/ns/domain-1.0">
                            <domain:name>%s</domain:name>
							<domain:authInfo>
 								<domain:pw>%s</domain:pw>
 							</domain:authInfo>
                        </domain:info>
                    </info>
                    <clTRID>%s</clTRID>
                </command>
            </epp>`
	return fmt.Sprintf(xml, config.EppAddress, domain, config.AuthCode, getPreClTRID(config))
}

func DomainRenewXml(domain string, expDate string, period int, config types.Config) string {
	xml := `<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
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
	return fmt.Sprintf(xml, domain, expDate, config.AuthCode, getPreClTRID(config))
}

func getPreClTRID(config types.Config) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100000)
	randomStr := fmt.Sprintf("%05d", randomNum)
	return config.PreClTRID + "-" + randomStr
}

func sendXml(xml, address, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", address, bytes.NewBufferString(xml))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/xml")

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

func parseDomainInfoType(xmlContent string) (*types.DomainType, error) {
	var di domainInfo
	if err := xml.Unmarshal([]byte(xmlContent), &di); err != nil {
		return nil, err
	}

	if di.Result.Code == "2502" {
		return nil, errors.New("Session limit exceeded; server closing connection")
	}

	d := &types.DomainType{
		Domain:       di.Name,
		Holder:       di.Holder,
		DomainStatus: make([]string, len(di.Statuses)),
		Ns1:          "",
		Ns2:          "",
		Ns3:          "",
		Ns4:          "",
		ExpDate:      time.Time{},
		CreateDate:   time.Time{},
		UpdateDate:   time.Time{},
	}

	copy(d.DomainStatus, func() []string {
		s := make([]string, len(di.Statuses))
		for i := range s {
			s[i] = di.Statuses[i].Value
		}
		return s
	}())

	if len(di.Ns) > 0 {
		d.Ns1 = di.Ns[0]
	}
	if len(di.Ns) > 1 {
		d.Ns2 = di.Ns[1]
	}
	if len(di.Ns) > 2 {
		d.Ns3 = di.Ns[2]
	}
	if len(di.Ns) > 3 {
		d.Ns4 = di.Ns[3]
	}

	if t, err := time.Parse("2006-01-02T15:04:05", di.ExDate); err == nil {
		d.ExpDate = t
	}
	if t, err := time.Parse("2006-01-02T15:04:05", di.CrDate); err == nil {
		d.CreateDate = t
	}
	if t, err := time.Parse("2006-01-02T15:04:05", di.UpDate); err == nil {
		d.UpdateDate = t
	}

	return d, nil
}

func formatDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func addDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func convertToJalali(t time.Time) (string, error) {
	jy, jm, jd := jalali.ToJalali(t.Year(), int(t.Month()), t.Day())
	return fmt.Sprintf("%04d/%02d/%02d", jy, jm, jd), nil
}
