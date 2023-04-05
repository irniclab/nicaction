package nicResponse

import "time"

type DomainWhoisInfoResponse struct {
	Name     string `xml:"domain:infData>name"`
	Statuses []struct {
		Value string `xml:"s,attr"`
	} `xml:"domain:infData>status"`
	Contacts []struct {
		Type  string `xml:"type,attr"`
		Value string `xml:",chardata"`
	} `xml:"domain:infData>contact"`
	Ns     []string  `xml:"domain:infData>ns>hostAttr>hostName"`
	CrDate time.Time `xml:"domain:infData>crDate"`
	UpDate string    `xml:"domain:infData>upDate"`
	ExDate string    `xml:"domain:infData>exDate"`
	Holder string    `xml:"domain:infData>contact[type=holder]"`
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"result"`
}

type DomainRenewResponse struct {
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"result"`
	ResData struct {
		RenData struct {
			Name string `xml:"http://epp.nic.ir/ns/domain-1.0 name"`
		} `xml:"http://epp.nic.ir/ns/domain-1.0 renData"`
	} `xml:"resData"`
	TrID struct {
		ClTRID string `xml:"clTRID"`
		SvTRID string `xml:"svTRID"`
	} `xml:"trID"`
}

type DomainRegisterResponse struct {
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"result"`
	ResData struct {
		CreData struct {
			Name   string `xml:"http://epp.nic.ir/ns/domain-1.0 name"`
			CrDate string `xml:"http://epp.nic.ir/ns/domain-1.0 crDate"`
			ExDate string `xml:"http://epp.nic.ir/ns/domain-1.0 exDate"`
		} `xml:"http://epp.nic.ir/ns/domain-1.0 creData"`
	} `xml:"resData"`
	TrID struct {
		ClTRID string `xml:"clTRID"`
		SvTRID string `xml:"svTRID"`
	} `xml:"trID"`
}

type GeneralResponse struct {
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"response>result"`
	TrID struct {
		ClTRID string `xml:"clTRID"`
		SvTRID string `xml:"svTRID"`
	} `xml:"response>trID"`
}
