package main

import (
	"go-grpc/cmd/services"
	"log"
	"net"

	"google.golang.org/grpc"

	productPB "go-grpc/pb/product"
)

const(
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("gagal listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	ProductService := services.ProductService{}
	productPB.RegisterProductServiceServer(grpcServer, &ProductService)


	log.Printf("server started %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil{
		log.Fatalf("gagal serve %v", err.Error())
	}
}