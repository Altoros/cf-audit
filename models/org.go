package models

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

// type App struct {
// 	cfclient.App
// }

func (org1 *cfclient.Org) Compare(org2 *cfclient.Org) (bool, error) {
	if org1.Name != org1.Name {
		return Collision{Elements: [org1, org2], Desctription: "name"}, nil
	}
}
