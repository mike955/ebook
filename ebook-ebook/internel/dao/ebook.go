package dao

import (
	"github.com/jinzhu/gorm"
)

type Ebook struct {
	CommonModel
	ID          uint64 `json:"id" gorm"index"`
	UserId      string `json:"user_id"`
	EbookName   string `json:"ebook_name"`
	EnglishName string `json:"english_name"`
	AliasName   string `json:"alias_name"`
	Category    string `json:"category"`
	PublishDate string `json:"publish_date"`
	KeyWords    string `json:"key_words"`

	PreviewType       string `json:"preview_type"`
	PreviewSize       int64  `json:"preview_size"`
	PreviewDir        string `json:"preview_dir"`
	PreviewUploadName string `json:"preview_upload_name"`
	PreviewHashValue  string `json:"preview_hash_value"`

	EbookType       string `json:"ebook_type"`
	EbookSize       int64  `json:"ebook_size"`
	EbookDir        string `json:"ebook_dir"`
	EbookUploadName string `json:"ebook_upload_name"`
	EbookHashValue  string `json:"ebook_hash_value"`
	IsDelete        uint64 `gorm:"default:0" json:"is_delete"`
}

type EbookDao struct {
}

func (dao EbookDao) Add(data map[string]interface{}) (err error) {
	account := Ebook{
		UserId:      data["userId"].(string),
		EbookName:   data["ebookName"].(string),
		EnglishName: data["englishName"].(string),
		AliasName:   data["aliasName"].(string),
		Category:    data["category"].(string),
		PublishDate: data["publishDate"].(string),
		KeyWords:    data["keyWords"].(string),

		PreviewType:       data["previewType"].(string),
		PreviewSize:       data["previewSize"].(int64),
		PreviewDir:        data["previewDir"].(string),
		PreviewUploadName: data["previewUploadName"].(string),
		PreviewHashValue:  data["previewHashValue"].(string),

		EbookType:       data["ebookType"].(string),
		EbookSize:       data["ebookSize"].(int64),
		EbookDir:        data["ebookDir"].(string),
		EbookUploadName: data["ebookUploadName"].(string),
		EbookHashValue:  data["ebookHashValue"].(string),
	}
	if err := DB.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

func (dao EbookDao) FindByID(id uint64) (*Ebook, error) {
	var ebook = new(Ebook)
	err := DB.Where(&Ebook{ID: id}).First(ebook).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ebook, nil
}

func (dao EbookDao) FindByFields(fields map[string]interface{}) ([]*Ebook, error) {
	var ebooks []*Ebook
	err := DB.Where(fields).Find(&ebooks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ebooks, nil
}

func (dao EbookDao) DeleteById(id uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": id}).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao EbookDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFileds).Error
	return
}
