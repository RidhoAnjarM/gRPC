package services

import (
    "context"
    productPB "go-grpc/pb/product"
    pagingPB "go-grpc/pb/pagination"
    "sync"
)

type ProductService struct {
    productPB.UnimplementedProductServiceServer
    products map[uint64]*productPB.Product 
    mu       sync.Mutex                    
}

func NewProductService() *ProductService {
    return &ProductService{
        products: make(map[uint64]*productPB.Product),
    }
}

func (p *ProductService) GetProducts(ctx context.Context, req *productPB.Empty) (*productPB.Products, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    var productList []*productPB.Product
    for _, product := range p.products {
        productList = append(productList, product)
    }

    return &productPB.Products{
        Pagination: &pagingPB.Pagination{
            Total:        uint64(len(productList)),
            PerPage:      10,
            CurrentPage:  1,
            LastPage:     1,
        },
        Data: productList,
    }, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *productPB.Id) (*productPB.Product, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    product, exists := p.products[id.Id]
    if !exists {
        return nil, nil
    }

    return product, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, product *productPB.Product) (*productPB.Id, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    product.Id = uint64(len(p.products) + 1)
    p.products[product.Id] = product

    return &productPB.Id{Id: product.Id}, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, product *productPB.Product) (*productPB.Status, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    if _, exists := p.products[product.Id]; exists {
        p.products[product.Id] = product
        return &productPB.Status{Id: uint32(product.Id)}, nil
    }
    return nil, nil 
}

func (p *ProductService) DeleteProduct(ctx context.Context, id *productPB.Id) (*productPB.Status, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    if _, exists := p.products[id.Id]; exists {
        delete(p.products, id.Id)
        return &productPB.Status{Id: uint32(id.Id)}, nil 
    }
    return nil, nil 
}
