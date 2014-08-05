package network

type NetworkProvider interface {
	CreateNetwork(nn NewNetwork) (*Network, error)
	GetNetwork(networkId string) (*Network, error)
	DeleteNetwork(networkId string) error

	CreateSubnet(subnet Subnet) (*Subnet, error)

	CreateSecurityGroup(group NewSecurityGroup) (*SecurityGroup, error)
	GetSecurityGroup(securityGroupId string) (*SecurityGroup, error)
	DeleteSecurityGroup(securityGroupId string) error

	CreateSecurityGroupRule(rule SecurityGroupRule) (*SecurityGroupRule, error)

	CreateRouter(router NewRouter) (*Router, error)
	AddRouterInterface(routerId string, subnetId string) (*CreatedPortId, error)
}
