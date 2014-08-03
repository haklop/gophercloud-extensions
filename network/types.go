package network

type Networks struct {
	Networks []Network `json:"networks,omitempty"`
}

type Network struct {
	Status       string `json:"status,omitempty"`
	Name         string `json:"name,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
	TenantId     string `json:"tenant_id,omitempty"`
	Id           string `json:"id,omitempty"`
}

type NewNetwork struct {
	Name         string `json:"name,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	TenantId     string `json:"tenant_id,omitempty"`
}
