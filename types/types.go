package types

import (
	"encoding/json"
	"time"
)

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

type DomainListResult struct {
	Domain   string
	Duration int
	Result   bool
	ErrorMsg string
}

// Config ساختار تنظیمات نیک
type Config struct {
	EppAddress        string `json:"eppAddress"`
	Nichandle         string `json:"nichandle"`
	Token             string `json:"token"`
	AuthCode          string `json:"authCode"`
	Ns1               string `json:"ns1"`
	Ns2               string `json:"ns2"`
	PreClTRID         string `json:"preClTRID"`
	MainNicHandle     string `json:"mainNicHandle"`
	ResellerNicHandle string `json:"resellerNicHandle"`
	DefaultPeriod     int    `json:"defaultPeriod"`
}

func ConvertDomainTypeToJSONByte(dt DomainType) ([]byte, error) {
	jsonData, err := json.Marshal(dt)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func ConvertConfigTypeToJSONByte(ct Config) ([]byte, error) {
	jsonData, err := json.Marshal(ct)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func ConvertJsonByteToString(jsByte []byte) (string, error) {
	return string(jsByte), nil
}
