package requestgateway

//Gateway is the list of addresses authorised to use a given service
type Gateway struct {
	RemoteAddress string `json:"remoteaddress" datastore:"remoteaddress"`
}
