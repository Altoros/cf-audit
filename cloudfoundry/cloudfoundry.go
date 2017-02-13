package cloudfoundry

import (
	// "errors"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	//"github.com/cloudfoundry-incubator/candiedyaml"
	//"os"
)

type CloudFoundries []CloudFoundry

type CloudFoundry struct {
	Name   string
	Client *Client
	// Api      string
	// Username string
	// Password string
}

func NewCloudFoundry(name string, apiUrl string, username string, password string) (*CloudFoundry, error) {
	client, err := NewCloudFoundryClient(ApiUrl, Username, Password)
	if err != nil {
		return nil, err
	}

	return &CloudFoundry{name, client}
}

func (c *CloudFoundry) Equal(Comparable) (bool, error) {
	return nil, nil
}

func (c *CloudFoundry) Children() ([]Comparable, error) {
	return nil, nil
}

func (c *CloudFoundry) Parent() (*Comparable, error) {
	return nil, nil
}

func (c1 *CloudFoundry) Collisions(c2 *Comparable) (Collisions, error) {
	// orgs1 := c1.Client.GetOrgs()

	// OrgsWithSameNames
	return nil, nil
}
