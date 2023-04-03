package xmlRequest

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/irniclab/nicaction/config"
)

func domainWhoisXml(domain string, config config.Config) string {
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

func getPreClTRID(config config.Config) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100000)
	randomStr := fmt.Sprintf("%05d", randomNum)
	return config.PreClTRID + "-" + randomStr
}
