package nicResponse

type DomainWhoisInfoResponse struct {
	Result struct {
		Code string `xml:"code,attr"`
		Msg  string `xml:"msg"`
	} `xml:"result"`
	Data struct {
		InfData struct {
			Name     string `xml:"http://epp.nic.ir/ns/domain-1.0 name"`
			Roid     string `xml:"http://epp.nic.ir/ns/domain-1.0 roid"`
			Statuses []struct {
				Value string `xml:"s,attr"`
			} `xml:"http://epp.nic.ir/ns/domain-1.0 status"`
			Contacts []struct {
				Type  string `xml:"type,attr"`
				Value string `xml:",chardata"`
			} `xml:"http://epp.nic.ir/ns/domain-1.0 contact"`
			Ns     []string `xml:"http://epp.nic.ir/ns/domain-1.0 ns>hostAttr>hostName"`
			CrDate string   `xml:"http://epp.nic.ir/ns/domain-1.0 crDate"`
			UpDate string   `xml:"http://epp.nic.ir/ns/domain-1.0 upDate"`
			ExDate string   `xml:"http://epp.nic.ir/ns/domain-1.0 exDate"`
		} `xml:"http://epp.nic.ir/ns/domain-1.0 infData"`
	} `xml:"resData"`
	TRID struct {
		ClTRID string `xml:"clTRID"`
		SvTRID string `xml:"svTRID"`
	} `xml:"trID"`
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
