package xmlRequest

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/yaa110/go-jalali"

	"github.com/irniclab/nicaction/types"
)

type hostName struct {
	HostName string `xml:"hostName"`
}

type hostAttr struct {
	HostNames []hostName `xml:"hostName"`
}

type ns struct {
	HostAttrs []hostAttr `xml:"hostAttr"`
}

type infData struct {
	Name     string `xml:"name"`
	Statuses []struct {
		Value string `xml:"s,attr"`
	} `xml:"status"`
	Contacts []struct {
		Type  string `xml:"type,attr"`
		Value string `xml:",chardata"`
	} `xml:"contact"`
	Ns     ns     `xml:"ns"`
	CrDate string `xml:"crDate"`
	UpDate string `xml:"upDate"`
	ExDate string `xml:"exDate"`
}

type resData struct {
	InfData infData `xml:"infData"`
}

type response struct {
	ResData resData `xml:"resData"`
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

func parseDomainType(xmlContent string) (*types.DomainType, error) {
	d := &types.DomainType{}
	var resp response
	if err := xml.Unmarshal([]byte(xmlContent), &resp); err != nil {
		return nil, err
	}

	d.Domain = resp.ResData.InfData.Name
	d.Holder = resp.ResData.InfData.Contacts[0].Value
	for _, status := range resp.ResData.InfData.Statuses {
		d.DomainStatus = append(d.DomainStatus, status.Value)
	}

	for _, hostAttr := range resp.ResData.InfData.Ns.HostAttrs {
		for _, hostName := range hostAttr.HostNames {
			if len(d.Ns1) == 0 {
				d.Ns1 = hostName.HostName
			} else if len(d.Ns2) == 0 {
				d.Ns2 = hostName.HostName
			} else if len(d.Ns3) == 0 {
				d.Ns3 = hostName.HostName
			} else if len(d.Ns4) == 0 {
				d.Ns4 = hostName.HostName
			}
		}
	}

	if t, err := time.Parse("2006-01-02T15:04:05", resp.ResData.InfData.ExDate); err == nil {
		d.ExpDate = t
	}
	if t, err := time.Parse("2006-01-02T15:04:05", resp.ResData.InfData.CrDate); err == nil {
		d.CreateDate = t
	}
	if t, err := time.Parse("2006-01-02T15:04:05", resp.ResData.InfData.UpDate); err == nil {
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
