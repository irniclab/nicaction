package types

import "time"

type DomainType struct {
	ExpDate      time.Time
	CreateDate   time.Time
	UpdateDate   time.Time
	LastUpdate   time.Time
	LockDate     time.Time
	ReleaseDate  time.Time
	Domain       string
	Holder       string
	Ns1          string
	Ns2          string
	Ns3          string
	Ns4          string
	DomainStatus []string
}

// Config ساختار تنظیمات نیک
type Config struct {
	EppAddress    string `json:"eppAddress"`
	Nichandle     string `json:"nichandle"`
	Token         string `json:"token"`
	AuthCode      string `json:"authCode"`
	Ns1           string `json:"ns1"`
	Ns2           string `json:"ns2"`
	PreClTRID     string `json:"preClTRID"`
	DefaultPeriod int    `json:"defaultPeriod"`
}
