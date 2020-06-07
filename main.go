package main

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/business-owner/api"
	"google.golang.org/grpc"

)

func main() {
	cc, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer cc.Close()


	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.ResetBusinessOwnerPasswordRequest{
		BusinessOwnerEmail: "esbolatovakezhan@gmail.com",
		BusinessOwnerPassword: "aki",
	}

	a, err  := c.ResetBusinessOwnerPassword(context.Background(), &r)
	if err != nil {
		panic(err)
	}

	log.Println(a)
}
