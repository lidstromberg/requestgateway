package requestgateway

//Gateway is the list of addresses authorised to use a given service
type Gateway struct {
	AppContext    string `json:"appcontext" datastore:"appcontext"`
	RemoteAddress string `json:"remoteaddress" datastore:"remoteaddress"`
}
