package requestgateway

import (
	"context"
	"log"
	"os"
	"strconv"

	lbcf "github.com/lidstromberg/config"
)

var (
	//EnvDebugOn controls verbose logging
	EnvDebugOn bool
	//EnvClientPool is the size of the client pool
	EnvClientPool int
)

//preflight config checks
func preflight(ctx context.Context, bc lbcf.ConfigSetting) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("Started Gateway preflight..")

	//get the session config and apply it to the config
	bc.LoadConfigMap(ctx, preflightConfigLoader())

	//then check that we have everything we need
	if bc.GetConfigValue(ctx, "EnvDebugOn") == "" {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	if bc.GetConfigValue(ctx, "EnvGtwayGcpProject") == "" {
		log.Fatal("Could not parse environment variable EnvGtwayGcpProject")
	}

	if bc.GetConfigValue(ctx, "EnvGtwayDsKind") == "" {
		log.Fatal("Could not parse environment variable EnvGtwayDsKind")
	}

	if bc.GetConfigValue(ctx, "EnvGtwayDsNamespace") == "" {
		log.Fatal("Could not parse environment variable EnvGtwayDsNamespace")
	}

	if bc.GetConfigValue(ctx, "EnvGtwayClientPool") == "" {
		log.Fatal("Could not parse environment variable EnvGtwayClientPool")
	}

	//set the debug value
	constlog, err := strconv.ParseBool(bc.GetConfigValue(ctx, "EnvDebugOn"))

	if err != nil {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	EnvDebugOn = constlog

	//set the poolsize
	pl, err := strconv.ParseInt(bc.GetConfigValue(ctx, "EnvGtwayClientPool"), 10, 64)

	if err != nil {
		log.Fatal("Could not parse environment variable EnvGtwayClientPool")
	}

	EnvClientPool = int(pl)

	log.Println("..Finished Gateway preflight.")
}

//preflightConfigLoader loads the config vars
func preflightConfigLoader() map[string]string {
	cfm := make(map[string]string)

	//EnvDebugOn controls verbose logging
	cfm["EnvDebugOn"] = os.Getenv("GTWAY_DEBUGON")
	//EnvGtwayGcpProject is the cloud project to target
	cfm["EnvGtwayGcpProject"] = os.Getenv("GTWAY_GCP_PROJECT")
	//EnvGtwayDsNamespace is the datastore namespace
	cfm["EnvGtwayDsNamespace"] = os.Getenv("GTWAY_NAMESP")
	//EnvGtwayDsKind is the Gateway entity
	cfm["EnvGtwayDsKind"] = os.Getenv("GTWAY_KD")
	//EnvGtwayClientPool is the client poolsize
	cfm["EnvGtwayClientPool"] = os.Getenv("GTWAY_CLIPOOL")

	if cfm["EnvDebugOn"] == "" {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	if cfm["EnvGtwayGcpProject"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwayGcpProject")
	}

	if cfm["EnvGtwayDsNamespace"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwayDsNamespace")
	}

	if cfm["EnvGtwayDsKind"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwayDsKind")
	}

	if cfm["EnvGtwayClientPool"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwayClientPool")
	}

	return cfm
}
