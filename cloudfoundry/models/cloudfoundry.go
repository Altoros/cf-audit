package cloudfoundry

import (
	// "errors"
	"github.com/Altoros/cf-audit/comparator"
	// cfclient "github.com/cloudfoundry-community/go-cfclient"
	//"github.com/cloudfoundry-incubator/candiedyaml"
	//"os"
	"github.com/davecgh/go-spew/spew"
)

type CloudFoundries []*CloudFoundry

type CloudFoundry struct {
	CloudFoundryItem
}


func NewCloudFoundry(name string, apiAddress string, username string, password string) (*CloudFoundry, error) {
	client, err := NewCloudFoundryClient(apiAddress, username, password)
	if err != nil {
		return nil, err
	}

	// return &CloudFoundry{meta: CloudFoundryItem{Name: name, Client: client}}, nil
	return &CloudFoundry{Name: name, Client: client}, nil
}


func (cfList CloudFoundries) FindCollisions() (comparator.Collisions, error) {
	// orgs1 := c1.Client.GetOrgs()

	// spew.Dump(cfList[0])
	orgs1, _ := cfList[0].meta.Client.ListOrgs()
	spew.Dump(orgs1)

	// OrgsWithSameNames
	return nil, nil
}
