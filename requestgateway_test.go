package requestgateway

import (
	"testing"

	lbcf "github.com/lidstromberg/config"

	context "golang.org/x/net/context"
)

func Test_DataRepoConnect(t *testing.T) {
	ctx := context.Background()

	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)

	if err != nil {
		t.Fatal(err)
	}

	tx, err := gt.dsclient.NewTransaction(ctx)

	if err != nil {
		t.Fatal(err)
	}

	tx.Rollback()
}
func Test_SetGateway(t *testing.T) {
	ctx := context.Background()

	bc := lbcf.NewConfig(ctx)

	gt, err := NewGtwyMgr(ctx, bc)

	if err != nil {
		t.Fatal(err)
	}

	err = gt.Set(ctx, "0.0.0.0")

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

	chk, err := gt.IsPermitted(ctx, "0.0.0.0")

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

	chk, err := gt.IsPermitted(ctx, "0.0.0.1")

	if err != nil {
		t.Fatal(err)
	}

	if chk {
		t.Fatal("Gateway address should not be permitted")
	}
}
