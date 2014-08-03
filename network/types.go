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
}

type NewNetwork struct {
	Name         string `json:"name,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	TenantId     string `json:"tenant_id,omitempty"`
}

type Subnet struct {
	NetworkId  string `json:"network_id,omitempty"`
	Name       string `json:"name,omitempty"`
	TenantId   string `json:"tenant_id,omitempty"`
	Cidr       string `json:"cidr,omitempty"`
	IPVersion  int    `json:"ip_version,omitempty"`
	EnableDhcp bool   `json:"enable_dhcp,omitempty"`
	Id         string `json:"id,omitempty"`
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
