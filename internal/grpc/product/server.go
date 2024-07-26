package product

import (
	"context"
	"errors"

	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/services"
	"github.com/GosMachine/ProductService/internal/storage/database"
	productv1 "github.com/GosMachine/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Product interface {
	GetCategory(slug string) (category *models.Category, err error)
	GetProduct(categorySlug, productSlug string) (product *models.Product, err error)
	GetCategories() (categories []database.Category, err error)
	CreateTicket(name, email, message, ip string) error
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
		items = append(items, &productv1.Item{
			Name:  v.Name,
			Slug:  v.Slug,
			Price: v.Price,
			Image: v.ImageURL,
		})
	}
	return &productv1.GetGategoryResponse{Description: category.Description, Items: items, Name: category.Name}, nil
}

func (s *serverAPI) GetCategories(ctx context.Context, req *emptypb.Empty) (*productv1.GetCategoriesResponse, error) {
	categories, err := s.product.GetCategories()
	if err != nil {
		if errors.Is(err, services.ErrCategoryNotFound) {
			return nil, status.Error(codes.InvalidArgument, "category not found")
		}
		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
	}
	var resultCategories []*productv1.Category
	for _, v := range categories {
		resultCategories = append(resultCategories, &productv1.Category{Name: v.Name, Slug: v.Slug})
	}
	return &productv1.GetCategoriesResponse{Categories: resultCategories}, nil
}

func (s *serverAPI) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	product, err := s.product.GetProduct(req.CategorySlug, req.ProductSlug)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	productResponse := productv1.GetProductResponse{Description: product.Description, Stock: product.Stock}
	productResponse.Fields = append(productResponse.Fields, &productv1.InputFields{Label: "Quantity", Type: "quantity"})
	for _, value := range product.Fields {
		productResponse.Fields = append(productResponse.Fields, &productv1.InputFields{Label: value.Label, Type: value.Type})
	}
	productResponse.Item = &productv1.Item{Name: product.Name, Slug: product.Slug, Price: product.Price, Image: product.ImageURL}

	return &productResponse, nil
}

func (s *serverAPI) CreateTicket(ctx context.Context, req *productv1.CreateTicketRequest) (*emptypb.Empty, error) {
	err := s.product.CreateTicket(req.Name, req.Email, req.Message, req.IP)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// func (s *serverAPI) Checkout(ctx context.Context, req *productv1.CheckoutRequest) (*productv1.CheckoutResponse, error) {
// 	url, err := s.product.Checkout(req)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal server error. Please try again.")
// 	}
// 	return &productv1.CheckoutResponse{Url: url}, nil
// }
