package service

import (
	"context"
	pb "ebook/ebook-ebook/api/ebook"
	"ebook/ebook-ebook/internel/dao"
	"ebook/ebook-ebook/pkg/err_code"
)

type ebookService struct {
	ebookDao dao.EbookDao
}

var EbookService = &ebookService{
}


func (service *ebookService) Add(ctx context.Context, req *pb.AddRequest) (response *pb.AddResponse, err error) {
	response = new(pb.AddResponse)
	newEbook := map[string]interface{}{
		"userId": req.UserId,
		"ebookName": req.EbookName,
		"englishName": req.EnglishName,
		"aliasName": req.AliasName,
		"category": req.Category,
		"type": req.Type,
		"publishDate": req.PublishDate,
		"keyWords": req.KeyWords,
		"hashValue": req.HashValue,
		"ebookDir": req.EbookDir,
		"previewDir": req.PreviewDir,
	}
	// TODO  判断文件是否存在
	err = service.ebookDao.Add(newEbook)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_EBOOK_ERROR")
		return
	}
	ebooks, err := service.ebookDao.FindByFields(map[string]interface{}{"userId": req.UserId, "ebookName": req.EbookName})
	if err != nil || len(ebooks) == 0{
		response.Errno, response.Errmsg = err_code.Code("GET_EBOOK_ERROR")
		return
	}
	ebook := ebooks[0]
	response.Data = &pb.EbookResponseInfo{
		Id:                   ebook.ID,
		UserId:               ebook.UserId,
		EbookName:             ebook.EbookName,
		EnglishName:                ebook.EbookName,
		AliasName:               ebook.AliasName,
		Category:               ebook.Category,
		Type:               ebook.Type,
		PublishDate:               ebook.PublishDate,
		KeyWords:               ebook.KeyWords,
		HashValue:               ebook.HashValue,
		EbookDir:               ebook.EbookDir,
		PreviewDir:               ebook.EbookDir,
		IsDelete:             ebook.IsDelete,
		CreateTime:           ebook.CreateTime,
		UpdateTime:           ebook.UpdateTime,
	}
	return
}

func (service *ebookService) Delete(ctx context.Context, req *pb.DeleteRequest) (response *pb.DeleteResponse, err error){
	response = new(pb.DeleteResponse)
	response = new(pb.DeleteResponse)
	users, err := service.ebookDao.FindByFields(map[string]interface{}{"is": req.Id})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_EBOOK_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("EBOOK_IS_NOT_EXIST_ERROR")
		return
	}
	if err = service.ebookDao.DeleteById(req.Id); err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_EBOOK_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *ebookService) Update(ctx context.Context, req *pb.UpdateRequest) (response *pb.UpdateResponse, err error){
	users, err := service.ebookDao.FindByFields(map[string]interface{}{"is": req.Id, "user_id": req.UserId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_EBOOK_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("EBOOK_IS_NOT_EXIST_ERROR")
		return
	}
	where := map[string]interface{}{"is": req.Id}
	updateFields := map[string]interface{}{}
	if ebookName := req.EbookName; ebookName != "" {
		updateFields["ebook_name"] = ebookName
	}
	if englishName := req.EnglishName; englishName != "" {
		updateFields["english_name"] = englishName
	}
	if aliasName := req.AliasName; aliasName != "" {
		updateFields["alias_name"] = aliasName
	}
	if ebookType := req.Type; ebookType != "" {
		updateFields["type"] = ebookType
	}
	if category := req.Category; category != 0 {
		updateFields["status"] = category
	}
	if publishDate := req.PublishDate; publishDate != "" {
		updateFields["publish_date"] = publishDate
	}
	if keyWords := req.KeyWords; keyWords != "" {
		updateFields["key_words"] = keyWords
	}
	if hashValue := req.HashValue; hashValue != "" {
		updateFields["hash_value"] = hashValue
	}
	if ebookDir := req.EbookDir; ebookDir != "" {
		updateFields["ebook_dir"] = ebookDir
	}
	if previewDir := req.PreviewDir; previewDir != "" {
		updateFields["preview_dir"] = previewDir
	}
	if isDelete := req.IsDelete; isDelete != 0 {
		updateFields["is_delete"] = isDelete
	}
	err = service.ebookDao.UpdateFields(where, updateFields)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("UPDATE_EBOOK_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *ebookService) Get(ctx context.Context, req *pb.GetRequest) (response *pb.GetResponse, err error){
	response = new(pb.GetResponse)
	condition := map[string]interface{}{
		"id": req.Id,
		"user_id": req.UserId,
	}
	ebooks, err := service.ebookDao.FindByFields(condition)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_EBOOK_ERROR")
		return
	}
	if len(ebooks) == 0 {
		response.Errno, response.Errmsg = err_code.Code("EBOOK_IS_NOT_EXIST_ERROR")
		return
	}
	ebook := ebooks[0]
	response.Data = &pb.EbookResponseInfo{
		Id:                   ebook.ID,
		UserId:               ebook.UserId,
		EbookName:             ebook.EbookName,
		EnglishName:                ebook.EbookName,
		AliasName:               ebook.AliasName,
		Category:               ebook.Category,
		Type:               ebook.Type,
		PublishDate:               ebook.PublishDate,
		KeyWords:               ebook.KeyWords,
		HashValue:               ebook.HashValue,
		EbookDir:               ebook.EbookDir,
		PreviewDir:               ebook.EbookDir,
		IsDelete:             ebook.IsDelete,
		CreateTime:           ebook.CreateTime,
		UpdateTime:           ebook.UpdateTime,
	}
	return
}

func (service *ebookService) Gets(ctx context.Context, req *pb.GetsRequest) (response *pb.GetsResponse, err error){
	response = new(pb.GetsResponse)
	condition := make(map[string]interface{})
	if ids := req.Ids; len(ids) > 0  {
		condition["id"] = ids
	}
	if userId := req.UserId; userId != "" {		// todo convert like
		condition["user_id"] = "%" + userId + "%"
	}
	if ebookName := req.EbookName; ebookName != "" {
		condition["ebook_name"] = "%" + ebookName + "%"		// todo convert like
	}
	if englishName := req.EnglishName; englishName != "" {
		condition["role_id"] = englishName
	}
	if aliasName := req.AliasName; aliasName != "" {
		condition["status"] = aliasName
	}
	if category := req.Category; category != 0 {
		condition["is_delete"] = category
	}
	if ebookType := req.Type; ebookType != "" {
		condition["type"] = ebookType
	}
	if category := req.Category; category != 0 {
		condition["category"] = category
	}
	if publishDate := req.PublishDate; publishDate != "" {
		condition["publish_date"] = publishDate
	}
	if keyWords := req.KeyWords; keyWords != "" {
		condition["key_words"] = keyWords
	}
	if hashValue := req.HashValue; hashValue != "" {
		condition["hash_value"] = hashValue
	}
	if ebookDir := req.EbookDir; ebookDir != "" {
		condition["ebook_dir"] = ebookDir
	}
	if previewDir := req.PreviewDir; previewDir != "" {
		condition["preview_dir"] = previewDir
	}
	if isDelete := req.IsDelete; isDelete != 0 {
		condition["is_delete"] = isDelete
	}
	ebooks, err := service.ebookDao.FindByFields(condition)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_EBOOK_ERROR")
		return
	}
	for _, ebook := range ebooks {
		userInfo := &pb.EbookResponseInfo{
			Id:                   ebook.ID,
			UserId:               ebook.UserId,
			EbookName:             ebook.EbookName,
			EnglishName:                ebook.EbookName,
			AliasName:               ebook.AliasName,
			Category:               ebook.Category,
			Type:               ebook.Type,
			PublishDate:               ebook.PublishDate,
			KeyWords:               ebook.KeyWords,
			HashValue:               ebook.HashValue,
			EbookDir:               ebook.EbookDir,
			PreviewDir:               ebook.EbookDir,
			IsDelete:             ebook.IsDelete,
			CreateTime:           ebook.CreateTime,
			UpdateTime:           ebook.UpdateTime,
		}
		response.Data = append(response.Data, userInfo)
	}
	return
}
