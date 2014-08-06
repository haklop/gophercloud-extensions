package network

type NetworkProvider interface {
	CreateNetwork(nn NewNetwork) (*Network, error)
	GetNetwork(networkId string) (*Network, error)
	DeleteNetwork(networkId string) error

	CreateSubnet(subnet NewSubnet) (*Subnet, error)
	GetSubnet(subnetId string) (*Subnet, error)
	DeleteSubnet(subnetId string) error

	CreateSecurityGroup(group NewSecurityGroup) (*SecurityGroup, error)
	GetSecurityGroup(securityGroupId string) (*SecurityGroup, error)
	DeleteSecurityGroup(securityGroupId string) error

	CreateSecurityGroupRule(rule SecurityGroupRule) (*SecurityGroupRule, error)

	CreateRouter(router NewRouter) (*Router, error)
	GetRouter(routerId string) (*Router, error)
	UpdateRouter(router Router) (*Router, error)
	AddRouterInterface(routerId string, subnetId string) (*PortId, error)
	RemoveRouterInterface(routerId string, portId string) error
	DeleteRouter(routerId string) error
	GetPorts() ([]*PortId, error)
}
