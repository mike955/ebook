package dao

import (
	"github.com/jinzhu/gorm"
)

type Ebook struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	UserId string	`json:"user_id"`
	EbookName string	`json:"ebook_name"`
	EnglishName string	`json:"english_name"`
	AliasName string	`json:"alias_name"`
	Category uint64	`json:"category"`
	Type string	`json:"type"`
	PublishDate string	`json:"publish_date"`
	KeyWords string	`json:"key_words"`
	HashValue string	`json:"hash_value"`
	EbookDir string	`json:"ebook_dir"`
	PreviewDir string	`json:"preview_dir"`
	IsDelete  uint64 `gorm:"default:0" json:"is_delete"`
}

type EbookDao struct {
}

func (dao EbookDao) Add(data map[string]interface{}) (err error) {
	account := Ebook{
		UserId:       data["userId"].(string),
		EbookName:       data["ebookName"].(string),
		EnglishName:     data["englishName"].(string),
		AliasName:    data["aliasName"].(string),
		Category: data["category"].(uint64),
		Type:            data["type"].(string),
		PublishDate:     data["publishDate"].(string),
		KeyWords:          data["keyWords"].(string),
		HashValue:          data["hashValue"].(string),
		EbookDir:          data["ebookDir"].(string),
		PreviewDir:          data["previewDir"].(string),
	}
	if err := DB.Create(&account).Error; err !=nil {
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

func  (dao EbookDao) FindByFields (fields map[string]interface{}) ([]*Ebook, error)  {
	var ebooks []*Ebook
	err := DB.Where(fields).Find(&ebooks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ebooks, nil
}

func  (dao EbookDao) DeleteById (id uint64) (err error)  {
	if err := DB.Where(&map[string]interface{}{"id": id}).Update("is_delete", 1).Error; err !=nil {
		return err
	}
	return nil
}

func (dao EbookDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFileds).Error
	return
}