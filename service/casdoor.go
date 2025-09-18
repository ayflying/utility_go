// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type (
	ICasdoor interface {
		Load(endpoint string, clientId string, clientSecret string, certificate string, organizationName string, applicationName string)
		New() *casdoorsdk.Client
	}
)

var (
	localCasdoor ICasdoor
)

func Casdoor() ICasdoor {
	if localCasdoor == nil {
		panic("implement not found for interface ICasdoor, forgot register?")
	}
	return localCasdoor
}

func RegisterCasdoor(i ICasdoor) {
	localCasdoor = i
}
