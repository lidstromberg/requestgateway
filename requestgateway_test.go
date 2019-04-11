package requestgateway

import (
	"testing"

	lbcf "github.com/lidstromberg/config"

	context "golang.org/x/net/context"
)

var (
	appname = "gatewaytester"
	iptest  = "0.0.0.10"
)

func Test_DataRepoConnect(t *testing.T) {
	ctx := context.Background()

	bc := lbcf.NewConfig(ctx)

	_, err := NewGtwyMgr(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}
}
func Test_SetGateway(t *testing.T) {
	ctx := context.Background()
	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	err = gt.Set(ctx, appname, iptest)
	if err != nil {
		t.Fatal(err)
	}
}
func Test_ShouldBeValid(t *testing.T) {
	ctx := context.Background()
	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	chk, err := gt.IsPermitted(ctx, appname, iptest)
	if err != nil {
		t.Fatal(err)
	}

	if !chk {
		t.Fatal("Gateway address should be permitted")
	}
}
func Test_ShouldNotBeValid(t *testing.T) {
	ctx := context.Background()

	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	chk, err := gt.IsPermitted(ctx, appname, "0.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	if chk {
		t.Fatal("Gateway address should not be permitted")
	}
}
func Test_DeleteGateway(t *testing.T) {
	ctx := context.Background()
	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	err = gt.Delete(ctx, appname, iptest)
	if err != nil {
		t.Fatal(err)
	}
}
