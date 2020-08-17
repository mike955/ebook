package ebook

type AddRequest struct {
	Username     string `form:"ebookName" binding:"required"`
	EnglishName string `form:"englishName" binding:"required"`
	AliasName string `form:"aliasName" binding:"required"`
	Category string `form:"category" binding:"required"`
	PublishDate string `form:"publishDate" binding:"required"`
	KeyWords string `form:"keyWords" binding:"required"`
	Pic []byte `form:"pic" binding:"required"`
	File []byte `form:"file" binding:"required"`
}


type GetRequest struct {
	Id     uint64 `json:"id" binding:"required"`
}

type GetListRequest struct {
	Ids     []uint64 `json:"ids" binding:"required"`
}

type DownloadRequest struct {
	Id     uint64 `json:"id" binding:"required"`
}

type ViewRequest struct {
	Id     uint64 `json:"id" binding:"required"`
	Type     string `json:"type" binding:"required"`
}
