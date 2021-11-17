package model


import (

)

type User struct {
	UserId string `json:"userId"`
	DeviceId string `json:"deviceId"`
	Create_at int64 `json:"create_at"`
	Token string `json:"token"`
	Os string `json:"os"`
	Os_ver string `json:"os_ver"`
}