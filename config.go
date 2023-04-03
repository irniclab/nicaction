package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config ساختار تنظیمات نیک
type Config struct {
	EppAddress string `json:"eppAddress"`
	Nichandle  string `json:"nichandle"`
	Token      string `json:"token"`
	Ns1        string `json:"ns1"`
	Ns2        string `json:"ns2"`
	PreClTRID  string `json:"preClTRID"`
}

// LoadConfig بارگذاری تنظیمات از فایل
func LoadConfig(path string) (Config, error) {
	var conf Config

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// SaveConfig ذخیره تغییرات تنظیمات در فایل
func SaveConfig(path string, conf Config) error {
	bytes, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
