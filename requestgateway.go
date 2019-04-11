package requestgateway

import (
	"strconv"

	lbcf "github.com/lidstromberg/config"
	lblog "github.com/lidstromberg/log"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

//GtwyMgr handles interactions with the datastore
type GtwyMgr struct {
	ds *datastore.Client
	bc lbcf.ConfigSetting
}

//NewGtwyMgr creates a new gateway manager
func NewGtwyMgr(ctx context.Context, bc lbcf.ConfigSetting) (*GtwyMgr, error) {
	preflight(ctx, bc)

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "NewGtwyMgr", "info", "start")
	}

	datastoreClient, err := datastore.NewClient(ctx, bc.GetConfigValue(ctx, "EnvGtwayGcpProject"), option.WithGRPCConnectionPool(EnvClientPool))
	if err != nil {
		return nil, err
	}

	cm1 := &GtwyMgr{
		ds: datastoreClient,
		bc: bc,
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "NewGtwyMgr", "info", "end")
	}

	return cm1, nil
}

//Set sets a gateway address
func (gt GtwyMgr) Set(ctx context.Context, appcontext, remoteAddress string) error {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Set", "info", "start")
	}

	glst := &Gateway{AppContext: appcontext, RemoteAddress: remoteAddress}

	ky, err := gt.newKey(ctx, gt.bc.GetConfigValue(ctx, "EnvGtwayDsNamespace"), gt.bc.GetConfigValue(ctx, "EnvGtwayDsKind"))
	if err != nil {
		return err
	}

	tx, err := gt.ds.NewTransaction(ctx)
	if err != nil {
		return err
	}

	if _, err := tx.Put(ky, glst); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.Commit(); err != nil {
		return err
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Set", "info", "end")
	}
	return nil
}

//IsPermitted indicates if the address is approved
func (gt GtwyMgr) IsPermitted(ctx context.Context, appcontext, remoteAddress string) (bool, error) {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", "start")
	}

	//check the approval list
	q := datastore.NewQuery(gt.bc.GetConfigValue(ctx, "EnvGtwayDsKind")).
		Namespace(gt.bc.GetConfigValue(ctx, "EnvGtwayDsNamespace")).
		Filter("appcontext =", appcontext).
		Filter("remoteaddress =", remoteAddress).
		KeysOnly()

	//get the count
	n, err := gt.ds.Count(ctx, q)
	//if there was an error return it and false
	if err != nil {
		if err != datastore.ErrNoSuchEntity {
			return false, err
		}
		return false, nil
	}

	//return false if the count was zero
	if n == 0 {
		return false, nil
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", strconv.Itoa(n))
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", "end")
	}

	//otherwise the address is valid
	return true, nil
}

//newKey is datastore specific and returns a key using datastore.AllocateIDs
func (gt GtwyMgr) newKey(ctx context.Context, dsNS, dsKind string) (*datastore.Key, error) {
	var keys []*datastore.Key

	//create an incomplete key of the type and namespace
	newKey := datastore.IncompleteKey(dsKind, nil)
	newKey.Namespace = dsNS

	//append it to the slice
	keys = append(keys, newKey)

	//allocate the ID from datastore
	keys, err := gt.ds.AllocateIDs(ctx, keys)
	if err != nil {
		return nil, err
	}

	//return only the first key
	return keys[0], nil
}

//Delete removes a gateway address
func (gt GtwyMgr) Delete(ctx context.Context, appcontext, remoteAddress string) error {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Delete", "info", "start")
	}

	//check the approval list
	q := datastore.NewQuery(gt.bc.GetConfigValue(ctx, "EnvGtwayDsKind")).
		Namespace(gt.bc.GetConfigValue(ctx, "EnvGtwayDsNamespace")).
		Filter("appcontext =", appcontext).
		Filter("remoteaddress =", remoteAddress).
		KeysOnly()

	var arr []Gateway
	keys, err := gt.ds.GetAll(ctx, q, &arr)
	if err != nil {
		return err
	}

	tx, err := gt.ds.NewTransaction(ctx)
	if err != nil {
		return err
	}

	if err := tx.DeleteMulti(keys); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.Commit(); err != nil {
		return err
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Delete", "info", "end")
	}
	return nil
}
