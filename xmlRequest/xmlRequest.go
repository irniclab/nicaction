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

	"github.com/irniclab/nicaction/nicResponse"
	ptime "github.com/yaa110/go-persian-calendar"

	"github.com/irniclab/nicaction/types"
)

func DomainWhoisXml(domain string, config types.Config) string {
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

func SendXml(xml string, config types.Config) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.EppAddress, bytes.NewBufferString(xml))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+config.Token)
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

func ParseDomainInfoType(xmlContent string) (*types.DomainType, error) {
	var di nicResponse.DomainWhoisInfoResponse
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
		d.LockDate = addDays(t, 30)
		d.ReleaseDate = addDays(t, 60)
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
	pt := ptime.New(t)
	return pt, nil
}
