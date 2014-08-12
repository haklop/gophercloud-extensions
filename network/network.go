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

func (gnp *genericNetworksProvider) GetNetworks() ([]Network, error) {
	var n []Network

	ep := gnp.endpoint + "/v2.0/networks"
	err := perigee.Get(ep, perigee.Options{
		Results: &struct {
			Network *[]Network `json:"networks"`
		}{&n},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return n, err
}

func (gnp *genericNetworksProvider) UpdateNetwork(networkId string, updatedNetwork UpdatedNetwork) (*Network, error) {
	var n *Network

	ep := gnp.endpoint + "/v2.0/networks/" + networkId
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			UpdatedNetwork *UpdatedNetwork `json:"network"`
		}{&updatedNetwork},
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

func (gnp *genericNetworksProvider) CreateSubnet(ns NewSubnet) (*Subnet, error) {
	var s *Subnet

	ep := gnp.endpoint + "/v2.0/subnets"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			NewSubnet *NewSubnet `json:"subnet"`
		}{&ns},
		Results: &struct{ Subnet **Subnet }{&s},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return s, err
}

func (gnp *genericNetworksProvider) GetSubnet(id string) (*Subnet, error) {
	var s *Subnet

	ep := gnp.endpoint + "/v2.0/subnets/" + id
	err := perigee.Get(ep, perigee.Options{
		Results: &struct{ Subnet **Subnet }{&s},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return s, err
}

func (gnp *genericNetworksProvider) UpdateSubnet(subnetId string, updatedSubnet UpdatedSubnet) (*Subnet, error) {
	var s *Subnet

	ep := gnp.endpoint + "/v2.0/subnets/" + subnetId
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			UpdatedSubnet *UpdatedSubnet `json:"subnet"`
		}{&updatedSubnet},
		Results: &struct{ Subnet **Subnet }{&s},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return s, err
}

func (gnp *genericNetworksProvider) DeleteSubnet(id string) error {
	ep := gnp.endpoint + "/v2.0/subnets/" + id
	err := perigee.Delete(ep, perigee.Options{
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{204},
	})

	return err
}

func (gnp *genericNetworksProvider) CreateSecurityGroup(nsg NewSecurityGroup) (*SecurityGroup, error) {
	var csg *SecurityGroup
	ep := gnp.endpoint + "/v2.0/security-groups"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			NewSecurityGroup *NewSecurityGroup `json:"security_group"`
		}{&nsg},
		Results: &struct {
			SecurityGroup **SecurityGroup `json:"security_group"`
		}{&csg},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return csg, err
}

func (gnp *genericNetworksProvider) GetSecurityGroup(id string) (*SecurityGroup, error) {
	var s *SecurityGroup

	ep := gnp.endpoint + "/v2.0/security-groups/" + id
	err := perigee.Get(ep, perigee.Options{
		Results: &struct{ SecurityGroup **SecurityGroup }{&s},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return s, err
}

func (gnp *genericNetworksProvider) DeleteSecurityGroup(id string) error {
	ep := gnp.endpoint + "/v2.0/security-groups/" + id
	err := perigee.Delete(ep, perigee.Options{
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{204},
	})

	return err
}

func (gnp *genericNetworksProvider) CreateSecurityGroupRule(nsgr SecurityGroupRule) (*SecurityGroupRule, error) {
	var sgr *SecurityGroupRule

	ep := gnp.endpoint + "/v2.0/security-group-rules"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			SecurityGroupRule *SecurityGroupRule `json:"security_group_rule"`
		}{&nsgr},
		Results: &struct{ SecurityGroupRule **SecurityGroupRule }{&sgr},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return sgr, err
}

func (gnp *genericNetworksProvider) CreateRouter(newRouter NewRouter) (*Router, error) {
	var router *Router

	ep := gnp.endpoint + "/v2.0/routers"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			NewRouter *NewRouter `json:"router"`
		}{&newRouter},
		Results: &struct{ Router **Router }{&router},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return router, err
}

func (gnp *genericNetworksProvider) GetRouter(routerId string) (*Router, error) {
	var r *Router

	ep := gnp.endpoint + "/v2.0/routers/" + routerId
	err := perigee.Get(ep, perigee.Options{
		Results: &struct{ Router **Router }{&r},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return r, err
}

func (gnp *genericNetworksProvider) DeleteRouter(id string) error {
	ep := gnp.endpoint + "/v2.0/routers/" + id
	err := perigee.Delete(ep, perigee.Options{
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{204},
	})

	return err
}

func (gnp *genericNetworksProvider) AddRouterInterface(routerId string, subnetId string) (*PortId, error) {
	var portId *PortId

	ep := gnp.endpoint + "/v2.0/routers/" + routerId + "/add_router_interface"
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			SubnetId string `json:"subnet_id"`
		}{subnetId},
		Results: &struct{ PortId **PortId }{&portId},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return portId, err
}

func (gnp *genericNetworksProvider) RemoveRouterInterface(routerId string, portId string) error {

	ep := gnp.endpoint + "/v2.0/routers/" + routerId + "/remove_router_interface"
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			PortId string `json:"port_id"`
		}{portId},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return err
}

func (gnp *genericNetworksProvider) GetPorts() ([]*PortId, error) {
	var ports []*PortId

	ep := gnp.endpoint + "/v2.0/ports"
	err := perigee.Get(ep, perigee.Options{
		Results: &struct {
			PortId *[]*PortId `json:"ports"`
		}{&ports},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return ports, err
}

func (gnp *genericNetworksProvider) UpdateRouter(router Router) (*Router, error) {
	var r *Router

	ep := gnp.endpoint + "/v2.0/routers/" + router.Id
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			router Router `json:"router"`
		}{router},
		Results: &struct{ Router **Router }{&r},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return r, err
}

func (gnp *genericNetworksProvider) CreatePool(newPool NewPool) (*Pool, error) {
	var pool *Pool

	ep := gnp.endpoint + "/v2.0/lb/pools"
	err := perigee.Post(ep, perigee.Options{
		ReqBody: &struct {
			NewPool *NewPool `json:"pool"`
		}{&newPool},
		Results: &struct{ Pool **Pool }{&pool},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{201},
	})

	return pool, err
}

func (gnp *genericNetworksProvider) GetPool(poolId string) (*Pool, error) {
	var pool *Pool

	ep := gnp.endpoint + "/v2.0/lb/pools/" + poolId
	err := perigee.Get(ep, perigee.Options{
		Results: &struct{ Pool **Pool }{&pool},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return pool, err
}

func (gnp *genericNetworksProvider) UpdatePool(pool Pool) (*Pool, error) {
	var updatedPool *Pool

	ep := gnp.endpoint + "/v2.0/lb/pools/" + pool.Id

	pool.Id = "" // remove read-only attribute
	err := perigee.Put(ep, perigee.Options{
		ReqBody: &struct {
			Pool *Pool `json:"pool"`
		}{&pool},
		Results: &struct{ Pool **Pool }{&updatedPool},
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{200},
	})

	return updatedPool, err
}

func (gnp *genericNetworksProvider) DeletePool(poolId string) error {
	ep := gnp.endpoint + "/v2.0/lb/pools/" + poolId
	err := perigee.Delete(ep, perigee.Options{
		MoreHeaders: map[string]string{
			"X-Auth-Token": gnp.access.AuthToken(),
		},
		OkCodes: []int{204},
	})

	return err
}
