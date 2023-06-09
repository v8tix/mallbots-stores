package grpc

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/v8tix/mallbots-stores-proto/pb"

	"github.com/v8tix/mallbots-stores/internal/application"
	"github.com/v8tix/mallbots-stores/internal/application/commands"
	"github.com/v8tix/mallbots-stores/internal/application/queries"
	"github.com/v8tix/mallbots-stores/internal/domain"
)

type server struct {
	app application.App
	pb.UnimplementedStoresServiceServer
}

var _ pb.StoresServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	pb.RegisterStoresServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateStore(ctx context.Context, request *pb.CreateStoreRequest) (*pb.CreateStoreResponse, error) {
	storeID := uuid.New().String()

	err := s.app.CreateStore(ctx, commands.CreateStore{
		ID:       storeID,
		Name:     request.GetName(),
		Location: request.GetLocation(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateStoreResponse{
		Id: storeID,
	}, nil
}

func (s server) EnableParticipation(ctx context.Context, request *pb.EnableParticipationRequest) (*pb.EnableParticipationResponse, error) {
	err := s.app.EnableParticipation(ctx, commands.EnableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.EnableParticipationResponse{}, nil
}

func (s server) DisableParticipation(ctx context.Context, request *pb.DisableParticipationRequest) (*pb.DisableParticipationResponse, error) {
	err := s.app.DisableParticipation(ctx, commands.DisableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DisableParticipationResponse{}, nil
}

func (s server) RebrandStore(ctx context.Context, request *pb.RebrandStoreRequest) (*pb.RebrandStoreResponse, error) {
	err := s.app.RebrandStore(ctx, commands.RebrandStore{
		ID:   request.GetId(),
		Name: request.GetName(),
	})

	return &pb.RebrandStoreResponse{}, err
}

func (s server) GetStore(ctx context.Context, request *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	store, err := s.app.GetStore(ctx, queries.GetStore{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetStoreResponse{Store: s.storeFromDomain(store)}, nil
}

func (s server) GetStores(ctx context.Context, request *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	stores, err := s.app.GetStores(ctx, queries.GetStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*pb.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &pb.GetStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s server) GetParticipatingStores(ctx context.Context, request *pb.GetParticipatingStoresRequest) (*pb.GetParticipatingStoresResponse, error) {
	stores, err := s.app.GetParticipatingStores(ctx, queries.GetParticipatingStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*pb.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &pb.GetParticipatingStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s server) AddProduct(ctx context.Context, request *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	id := uuid.New().String()
	err := s.app.AddProduct(ctx, commands.AddProduct{
		ID:          id,
		StoreID:     request.GetStoreId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		SKU:         request.GetSku(),
		Price:       request.GetPrice(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{Id: id}, nil
}

func (s server) RebrandProduct(ctx context.Context, request *pb.RebrandProductRequest) (*pb.RebrandProductResponse, error) {
	err := s.app.RebrandProduct(ctx, commands.RebrandProduct{
		ID:          request.GetId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
	})
	return &pb.RebrandProductResponse{}, err
}

func (s server) IncreaseProductPrice(ctx context.Context, request *pb.IncreaseProductPriceRequest) (*pb.IncreaseProductPriceResponse, error) {
	err := s.app.IncreaseProductPrice(ctx, commands.IncreaseProductPrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &pb.IncreaseProductPriceResponse{}, err
}

func (s server) DecreaseProductPrice(ctx context.Context, request *pb.DecreaseProductPriceRequest) (*pb.DecreaseProductPriceResponse, error) {
	err := s.app.DecreaseProductPrice(ctx, commands.DecreaseProductPrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &pb.DecreaseProductPriceResponse{}, err
}

func (s server) RemoveProduct(ctx context.Context, request *pb.RemoveProductRequest) (*pb.RemoveProductResponse, error) {
	err := s.app.RemoveProduct(ctx, commands.RemoveProduct{
		ID: request.GetId(),
	})

	return &pb.RemoveProductResponse{}, err
}

func (s server) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, queries.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{Product: s.productFromDomain(product)}, nil
}

func (s server) GetCatalog(ctx context.Context, request *pb.GetCatalogRequest) (*pb.GetCatalogResponse, error) {
	products, err := s.app.GetCatalog(ctx, queries.GetCatalog{StoreID: request.GetStoreId()})
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pb.Product, len(products))
	for i, product := range products {
		protoProducts[i] = s.productFromDomain(product)
	}

	return &pb.GetCatalogResponse{
		Products: protoProducts,
	}, nil
}

func (s server) storeFromDomain(store *domain.MallStore) *pb.Store {
	return &pb.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (s server) productFromDomain(product *domain.CatalogProduct) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.SKU,
		Price:       product.Price,
	}
}
