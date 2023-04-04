package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/irniclab/nicaction/types"
)

func generateConfigFile(path string) error {
	defaultConfig := types.Config{
		EppAddress:    "https://epp.nic.ir",
		Nichandle:     "rb994-irnic",
		Token:         "my-token",
		AuthCode:      "my-authcode",
		Ns1:           "ns1.iranns.ir",
		Ns2:           "ns2.iranns.ir",
		PreClTRID:     "viraarvand",
		DefaultPeriod: 1,
	}
	SaveConfig(path, defaultConfig)
}

// LoadConfig بارگذاری تنظیمات از فایل
func LoadConfig(path string) (types.Config, error) {
	var conf types.Config
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		generateConfigFile(path)
	}

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
func SaveConfig(path string, conf types.Config) error {
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
