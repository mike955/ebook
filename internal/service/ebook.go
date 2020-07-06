package service

import (
	"context"
	pb "ebook/api/ebook"
	"ebook/internal/data"
	"fmt"
	"io"
)

type ebookService struct {
	EbookData data.EbookData
}

var EbookService = &ebookService{}

// func (serrvice *ebookService) Add(ctx context.Context, req *pb.AddRequest) (response *pb.AddResponse, err error) {
// 	fmt.Println("=============")
// 	fmt.Println(req)
// 	return
// }

func (serrvice *ebookService) Add(stream pb.Ebook_AddServer) (err error) {
	var lastEbook *pb.AddRequest
	//startTime := time.Now()
	for  {
		input, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println(err)
			return nil
			//fmt.Println("=err: ", err)
		}
		lastEbook = input
		fmt.Println("=============")
		fmt.Println(input)
	}
	
	fmt.Println("=============")
	fmt.Println(lastEbook)
	return
}

func (serrvice *ebookService) Delete(ctx context.Context, req *pb.DeleteRequest) (response *pb.DeleteResponse, err error){
	return
}

func (serrvice *ebookService) Update(ctx context.Context, req *pb.UpdateRequest) (response *pb.UpdateResponse, err error){
	return
}

func (serrvice *ebookService) GetEbook(ctx context.Context, req *pb.GetEbookRequest) (response *pb.GetEbookResponse, err error){
	return
}

func (serrvice *ebookService) GetEbooks(ctx context.Context, req *pb.GetEbooksRequest) (response *pb.GetEbooksResponse, err error){
	return
}