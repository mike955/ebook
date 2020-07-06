package data

import "ebook/internal/dao"
import pb "ebook/api/ebook"

type EbookData struct {
	EbookDao dao.EbookDao
}

func (ebookData EbookData) Add (params *pb.AddRequest) (response *pb.AddResponse, err error){
	return
}