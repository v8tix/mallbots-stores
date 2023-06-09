package grpc

import (
	"context"
	"database/sql"

	"google.golang.org/grpc"

	"github.com/v8tix/eda/di"
	"github.com/v8tix/mallbots-stores-proto/pb"
	"github.com/v8tix/mallbots-stores/internal/application"
)

type serverTx struct {
	c di.Container
	pb.UnimplementedStoresServiceServer
}

var _ pb.StoresServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	pb.RegisterStoresServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) CreateStore(ctx context.Context, request *pb.CreateStoreRequest) (resp *pb.CreateStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.CreateStore(ctx, request)
}

func (s serverTx) EnableParticipation(ctx context.Context, request *pb.EnableParticipationRequest) (resp *pb.EnableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.EnableParticipation(ctx, request)
}

func (s serverTx) DisableParticipation(ctx context.Context, request *pb.DisableParticipationRequest) (resp *pb.DisableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.DisableParticipation(ctx, request)
}

func (s serverTx) RebrandStore(ctx context.Context, request *pb.RebrandStoreRequest) (resp *pb.RebrandStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.RebrandStore(ctx, request)
}

func (s serverTx) GetStore(ctx context.Context, request *pb.GetStoreRequest) (resp *pb.GetStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.GetStore(ctx, request)
}

func (s serverTx) GetStores(ctx context.Context, request *pb.GetStoresRequest) (resp *pb.GetStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.GetStores(ctx, request)
}

func (s serverTx) GetParticipatingStores(ctx context.Context, request *pb.GetParticipatingStoresRequest) (resp *pb.GetParticipatingStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.GetParticipatingStores(ctx, request)
}

func (s serverTx) AddProduct(ctx context.Context, request *pb.AddProductRequest) (resp *pb.AddProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.AddProduct(ctx, request)
}

func (s serverTx) RebrandProduct(ctx context.Context, request *pb.RebrandProductRequest) (resp *pb.RebrandProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.RebrandProduct(ctx, request)
}

func (s serverTx) IncreaseProductPrice(ctx context.Context, request *pb.IncreaseProductPriceRequest) (resp *pb.IncreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.IncreaseProductPrice(ctx, request)
}

func (s serverTx) DecreaseProductPrice(ctx context.Context, request *pb.DecreaseProductPriceRequest) (resp *pb.DecreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.DecreaseProductPrice(ctx, request)
}

func (s serverTx) RemoveProduct(ctx context.Context, request *pb.RemoveProductRequest) (resp *pb.RemoveProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.RemoveProduct(ctx, request)
}

func (s serverTx) GetProduct(ctx context.Context, request *pb.GetProductRequest) (resp *pb.GetProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.GetProduct(ctx, request)
}

func (s serverTx) GetCatalog(ctx context.Context, request *pb.GetCatalogRequest) (resp *pb.GetCatalogResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.App)}

	return next.GetCatalog(ctx, request)
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}
