package network

type NetworkProvider interface {
	CreateNetwork(nn NewNetwork) (*Network, error)
	GetNetwork(networkId string) (*Network, error)
	DeleteNetwork(networkId string) error
}
