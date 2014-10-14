package network

type Networks struct {
	Networks []Network `json:"networks,omitempty"`
}

type Network struct {
	Status       string   `json:"status,omitempty"`
	Name         string   `json:"name,omitempty"`
	AdminStateUp bool     `json:"admin_state_up,omitempty"`
	TenantId     string   `json:"tenant_id,omitempty"`
	Subnets      []string `json:"subnets,omitempty"`
	Id           string   `json:"id,omitempty"`
	IsExternal   bool     `json:"router:external,omitempty"`
}

type NewNetwork struct {
	Name         string `json:"name,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	TenantId     string `json:"tenant_id,omitempty"`
}

type UpdatedNetwork struct {
	Name         string `json:"name"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
}

type Subnet struct {
	NetworkId  string `json:"network_id,omitempty"`
	Name       string `json:"name,omitempty"`
	TenantId   string `json:"tenant_id,omitempty"`
	Cidr       string `json:"cidr,omitempty"`
	IPVersion  int    `json:"ip_version,omitempty"`
	EnableDhcp bool   `json:"enable_dhcp,omitempty"`
	Id         string `json:"id,omitempty"`
	GatewayIp  string `json:"gateway_ip,omitempty"`
}

type NewSubnet struct {
	NetworkId  string `json:"network_id,omitempty"`
	Name       string `json:"name,omitempty"`
	TenantId   string `json:"tenant_id,omitempty"`
	Cidr       string `json:"cidr,omitempty"`
	IPVersion  int    `json:"ip_version,omitempty"`
	EnableDhcp bool   `json:"enable_dhcp,omitempty"`
}

type UpdatedSubnet struct {
	Name       string `json:"name"`
	EnableDhcp bool   `json:"enable_dhcp,omitempty"`
}

type SecurityGroup struct {
	Id          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	TenantId    string              `json:"tenant_id,omitempty"`
	Description string              `json:"description,omitempty"`
	Rules       []SecurityGroupRule `json:"security_group_rules,omitempty"`
}

type NewSecurityGroup struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	TenantId    string `json:"tenant_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type SecurityGroupRule struct {
	Id              string `json:"id,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
	Direction       string `json:"direction,omitempty"`
	PortRangeMin    int    `json:"port_range_min,omitempty"`
	PortRangeMax    int    `json:"port_range_max,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	RemoteIpPrefix  string `json:"remote_ip_prefix,omitempty"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
}

type NewRouter struct {
	Name            string          `json:"name,omitempty"`
	ExternalGateway ExternalGateway `json:"external_gateway_info,omitempty"`
	AdminStateUp    bool            `json:"admin_state_up,omitempty"`
}

type Router struct {
	Id              string          `json:"id,omitempty"`
	TenantId        string          `json:"tenant_id,omitempty"`
	Name            string          `json:"name,omitempty"`
	ExternalGateway ExternalGateway `json:"external_gateway_info,omitempty"`
	AdminStateUp    bool            `json:"admin_state_up,omitempty"`
	Status          string          `json:"status,omitempty"`
}

type ExternalGateway struct {
	NetworkId string `json:"network_id,omitempty"`
}

type SubnetId struct {
	SubnetId string `json:"subnet_id"`
}

type Port struct {
	NetworkId    string    `json:"network_id"`
	PortId       string    `json:"id"`
	DeviceId     string    `json:"device_id"`
	name         string    `json:"name"`
	Id           string    `json:"id"`
	FixedIps     []FixedIp `json:"fixed_ips"`
	AdminStateUp bool      `json:"admin_state_up"`
}

type FixedIp struct {
	SubnetId  string `json:"subnet_id"`
	IpAddress string `json:"ip_address"`
}

type NewPool struct {
	SubnetId     string `json:"subnet_id,omitempty"`
	LoadMethod   string `json:"lb_method,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Provider     string `json:"provider,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
}

type Pool struct {
	SubnetId       string   `json:"subnet_id,omitempty"`
	LoadMethod     string   `json:"lb_method,omitempty"`
	Protocol       string   `json:"protocol,omitempty"`
	Name           string   `json:"name,omitempty"`
	AdminStateUp   bool     `json:"admin_state_up,omitempty"`
	Status         string   `json:"status,omitempty"`
	Description    string   `json:"description,omitempty"`
	Id             string   `json:"id,omitempty"`
	VipId          string   `json:"vip_id,omitempty"`
	Members        []string `json:"members,omitempty"`
	HealthMonitors []string `json:"health_monitors,omitempty"`
}

type NewMember struct {
	ProtocolPort int    `json:"protocol_port,omitempty"`
	Address      string `json:"address,omitempty"`
	PoolId       string `json:"pool_id,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
}

type Member struct {
	Id           string `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	ProtocolPort int    `json:"protocol_port,omitempty"`
	Address      string `json:"address,omitempty"`
	PoolId       string `json:"pool_id,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
}

type NewMonitor struct {
	Type          string `json:"type,omitempty"`
	Delay         int    `json:"delay,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	MaxRetries    int    `json:"max_retries,omitempty"`
	UrlPath       string `json:"url_path,omitempty"`
	ExpectedCodes string `json:"expected_codes ,omitempty"`
	HttpMethod    string `json:"http_method ,omitempty"`
}

type Monitor struct {
	Id            string `json:"id,omitempty"`
	Type          string `json:"type,omitempty"`
	Delay         int    `json:"delay,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	MaxRetry      int    `json:"max_retry,omitempty"`
	UrlPath       string `json:"url_path,omitempty"`
	ExpectedCodes string `json:"expected_codes ,omitempty"`
	HttpMethod    string `json:"http_method ,omitempty"`
	Status        string `json:"status ,omitempty"`
	AdminStateUp  bool   `json:"admin_state_up  ,omitempty"`
}

type NewVip struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	SubnetId     string `json:"subnet_id,omitempty"`
	ProtocolPort int    `json:"protocol_port,omitempty"`
	PoolId       string `json:"pool_id,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
}

type Vip struct {
	Id              string `json:"id,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	SubnetId        string `json:"subnet_id,omitempty"`
	Status          string `json:"status ,omitempty"`
	ProtocolPort    int    `json:"protocol_port,omitempty"`
	PoolId          string `json:"pool_id,omitempty"`
	PortId          string `json:"port_id,omitempty"`
	Address         string `json:"address,omitempty"`
	ConnectionLimit int    `json:"connection_limit,omitempty"`
	AdminStateUp    bool   `json:"admin_state_up  ,omitempty"`
}
