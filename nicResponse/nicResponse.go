package nicResponse

import "encoding/xml"

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
