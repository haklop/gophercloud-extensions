package network

import (
	"github.com/racker/perigee"
	"github.com/rackspace/gophercloud"
)

// Instantiates a Network API for the provider given.
func NetworksApi(acc gophercloud.AccessProvider, criteria gophercloud.ApiCriteria) (NetworkProvider, error) {
	url := acc.FirstEndpointUrlByCriteria(criteria)
	if url == "" {
		return nil, gophercloud.ErrEndpoint
	}

	gcp := &genericNetworksProvider{
		endpoint: url,
		access:   acc,
	}

	return gcp, nil
}

type genericNetworksProvider struct {
	endpoint string
	context  *gophercloud.Context
	access   gophercloud.AccessProvider
}

func (gnp *genericNetworksProvider) CreateNetwork(nn NewNetwork) (*Network, error) {
	var n *Network

	ep := gnp.endpoint + "/v2.0/networks"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			NewNetwork *NewNetwork `json:"network"`
		}{&nn},
		Results: &struct{ Network **Network }{&n},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return n, err
}

func (gnp *genericNetworksProvider) GetNetwork(id string) (*Network, error) {
	var n *Network

	ep := gnp.endpoint + "/v2.0/networks/" + id
	err := perigee.Get(ep, perigee.Options{
		Results: &struct{ Network **Network }{&n},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return n, err
}

func (gnp *genericNetworksProvider) DeleteNetwork(id string) error {
	ep := gnp.endpoint + "/v2.0/networks/" + id
	err := perigee.Delete(ep, perigee.Options{
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{204},
	})

	return err
}

func (gnp *genericNetworksProvider) CreateSubnet(ns Subnet) (*Subnet, error) {
	var s *Subnet

	ep := gnp.endpoint + "/v2.0/subnets"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			Subnet *Subnet `json:"subnet"`
		}{&ns},
		Results: &struct{ Subnet **Subnet }{&s},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return s, err
}
