package product

import (
	"context"
	"errors"

	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/services"
	productv1 "github.com/GosMachine/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Product interface {
	GetCategory(slug string) (category *models.Category, err error)
	GetCategories() (categories *postgres.Categories, err error)
	Checkout(req *productv1.CheckoutRequest) (url string, err error)
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Product
}

func RegisterAuthServer(gRPC *grpc.Server, product Product) {
	productv1.RegisterProductServer(gRPC, &serverAPI{product: product})
}

func (s *serverAPI) GetCategory(ctx context.Context, req *productv1.GetCategoryRequest) (*productv1.GetGategoryResponse, error) {
	category, err := s.product.GetCategory(req.Slug)
	if err != nil {
		if errors.Is(err, services.ErrCategoryNotFound) {
			return nil, status.Error(codes.InvalidArgument, "category not found")
		}
		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
	}
	var items []*productv1.Item
	for _, v := range category.Products {
		var fields []*productv1.InputFields
		v.Fields = append(v.Fields, models.InputField{Label: "Count", Type: "count"})
		for _, value := range v.Fields {
			fields = append(fields, &productv1.InputFields{Label: value.Label, Type: value.Type})
		}
		items = append(items, &productv1.Item{
			Name:        v.Name,
			Slug:        v.Slug,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			Image:       v.ImageURL,
			Fields:      fields,
		})
	}
	return &productv1.GetGategoryResponse{Description: category.Description, Items: items, Name: category.Name}, nil
}

func (s *serverAPI) GetCategories(ctx context.Context, req *emptypb.Empty) (*productv1.GetCategoriesResponse, error) {
	categories, err := s.product.GetCategories()
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
	}
	return &productv1.GetCategoriesResponse{Names: categories.Names, Slugs: categories.Slugs}, nil
}

func (s *serverAPI) Checkout(ctx context.Context, req *productv1.CheckoutRequest) (*productv1.CheckoutResponse, error) {
	url, err := s.product.Checkout(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
	}
	return &productv1.CheckoutResponse{Url: url}, nil
}
