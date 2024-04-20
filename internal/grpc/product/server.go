package product

import (
	"context"
	"errors"

	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/services"
	productv1 "github.com/GosMachine/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Product interface {
	GetCategory(name string) (category *models.Category, err error)
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Product
}

func RegisterAuthServer(gRPC *grpc.Server, product Product) {
	productv1.RegisterProductServer(gRPC, &serverAPI{product: product})
}

func (s *serverAPI) GetCategory(ctx context.Context, req *productv1.GetCategoryRequest) (*productv1.GetGategoryResponse, error) {
	category, err := s.product.GetCategory(req.Name)
	if err != nil {
		if errors.Is(err, services.ErrCategoryNotFound) {
			return nil, status.Error(codes.InvalidArgument, "category not found")
		}
		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
	}
	var items []*productv1.Item
	for _, v := range category.Products {
		items = append(items, &productv1.Item{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			Image:       v.ImageURL,
		})
	}
	return &productv1.GetGategoryResponse{Description: category.Description, Items: items}, nil
}
