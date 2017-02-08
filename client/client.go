package client

import (
	"code.cloudfoundry.org/cli/cf/api"
)

type CloudFoundryClient struct {
	Host     string
	Username string
	Password string
}

func NewCloudFoundryClient() (CloudFoundryClient, error) {
	return CloudFoundryClient{}
}
