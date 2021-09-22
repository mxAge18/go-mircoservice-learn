package product

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type prodService struct {
	UnimplementedProductServiceServer
}
func NewProductServiceServer() ProductServiceServer {
	return &prodService{}
}
func (s *prodService) GetStock(ctx context.Context,req *Request) (*Response, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "GetStock Request canceled")
	}
	return &Response{ProductId: 100, ProductName: "name", Stock: 1000}, nil
}

func (s *prodService) GetProductList(ctx context.Context, size *QuerySize) (*ProductResponseList, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "GetProductList Request canceled")
	}
	productListRes := &ProductResponseList{
		ProdRes: []*Response{
			&Response{ProductId: 11, ProductName: "name11", Stock: 1001},
			&Response{ProductId: 12, ProductName: "name12", Stock: 1002},
			&Response{ProductId: 13, ProductName: "name13", Stock: 1003},
			&Response{ProductId: 14, ProductName: "name14", Stock: 1004},
		},
	}
	if len(productListRes.ProdRes) > int(size.Size) {
		productListRes.ProdRes = productListRes.ProdRes[:int(size.Size)]
	}
	return productListRes, nil
}