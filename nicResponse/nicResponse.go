package nicResponse

type DomainWhoisInfoResponse struct {
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
