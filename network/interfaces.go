package network

type NetworkProvider interface {
	CreateNetwork(nn NewNetwork) (*Network, error)
	GetNetwork(networkId string) (*Network, error)
	GetNetworks() ([]Network, error)
	UpdateNetwork(networkId string, updatedNetwork UpdatedNetwork) (*Network, error)
	DeleteNetwork(networkId string) error

	CreateSubnet(subnet NewSubnet) (*Subnet, error)
	GetSubnet(subnetId string) (*Subnet, error)
	UpdateSubnet(subnetId string, updatedSubnet UpdatedSubnet) (*Subnet, error)
	DeleteSubnet(subnetId string) error

	CreateSecurityGroup(group NewSecurityGroup) (*SecurityGroup, error)
	GetSecurityGroup(securityGroupId string) (*SecurityGroup, error)
	DeleteSecurityGroup(securityGroupId string) error

	CreateSecurityGroupRule(rule SecurityGroupRule) (*SecurityGroupRule, error)

	CreateRouter(router NewRouter) (*Router, error)
	GetRouter(routerId string) (*Router, error)
	UpdateRouter(router Router) (*Router, error)
	AddRouterInterface(routerId string, subnetId string) (*Port, error)
	RemoveRouterInterface(routerId string, portId string) error
	DeleteRouter(routerId string) error
	GetPorts() ([]*Port, error)

	CreatePool(newPool NewPool) (*Pool, error)
	GetPool(poolId string) (*Pool, error)
	UpdatePool(pool Pool) (*Pool, error)
	DeletePool(poolId string) error

	CreateMember(newMember NewMember) (*Member, error)
	GetMember(memberId string) (*Member, error)
	DeleteMember(memberId string) error

	CreateMonitor(newMonitor NewMonitor) (*Monitor, error)
	GetMonitor(monitorId string) (*Monitor, error)
	DeleteMonitor(monitorId string) error

	AssociateMonitor(monitorId string, poolId string) error
	UnassociateMonitor(monitorId string, poolId string) error

	CreateVip(newVio NewVip) (*Vip, error)
	GetVip(vipId string) (*Vip, error)
	DeleteVip(vipId string) error

	ListFloatingIps() ([]FloatingIp, error)
	AssociateFloatingIp(portId string, floatingId string) error
	UnassociateFloatingIp(floatingIpId string) error

	CreateFirewall(newFirewall NewFirewall) (*Firewall, error)
	GetFirewall(firewallId string) (*Firewall, error)
	DeleteFirewall(firewallId string) error
}
