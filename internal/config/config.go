package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var File = "settings.json"

type Settings struct {
	Devices	[]KodiParams
}

type KodiParams struct {
	DeviceName	string
	Address		string
	User		string
	Password	string
}

func FullRead() Settings {

	var settings Settings

	jsonFile, err := os.ReadFile(File)
	if err != nil {
		fmt.Printf("Cannot Read File: %v\n", err)
		return settings
	}

	err = json.Unmarshal(jsonFile, &settings)
	if err != nil {
		fmt.Printf("Cannot Unmarshal Settings: %v\n", err)
	}

	return settings

}

(s Settings) Devices() []KodiPramas {

	devices = s.Devices

	return devices

}

// Grab the first of the Slice for default
(ks []KodiParams) Device() KodiParams {
	device = ks[0]

	return device 

}

(k KodiParams) Address() string {

	address = k.Address

	return address

}

(k KodiParams) User() string {

	user = k.User

	return user

}

(k KodiParams) Password() string {

	password = k.Password

	return password

}

(k KodiParams) DeviceName() string {

	devicName = k.DeviceName

	return deviceName

}
