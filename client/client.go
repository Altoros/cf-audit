package client

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

type CloudFoundryClient struct {
	Host     string
	Username string
	Password string
}

func NewCloudFoundryClient() (CloudFoundryClient, error) {
	return CloudFoundryClient{}
}
