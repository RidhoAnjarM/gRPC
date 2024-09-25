package services

import (
	"context"
	productPB "go-grpc/pb/product"
	pagingPB "go-grpc/pb/pagination"
)

type ProductService struct{
	productPB.UnimplementedProductServiceServer
}

func (p *ProductService) GetProducts(context.Context, *productPB.Empty) (*productPB.Products, error){
	Products := &productPB.Products{
		Pagination: &pagingPB.Pagination{
			Total: 10,
			PerPage: 10,
			CurrentPage: 11,
			LastPage: 5,
		},
		Data: []*productPB.Product{
			{
				Id: 1,
				Name: "Jersey Ronaldo",
				Price: 500000.00,
				Stock: 2,
				Category: &productPB.Category{
					Id: 1,
					Name: "T-Shirt",		
				},
			},
			{
				Id: 1,
				Name: "Jersey Ronaldo",
				Price: 500000.00,
				Stock: 2,
				Category: &productPB.Category{
					Id: 1,
					Name: "T-Shirt",		
				},
			},
		},
	}

	return Products, nil
}