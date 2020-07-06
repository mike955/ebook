package dao

type Ebook struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	EbookName string	`json:"ebook_name"`
	EnglishName string	`json:"english_name"`
	AliasName string	`json:"alias_name"`
	Category uint	`json:"category"`
	PublishDate string	`json:"publish_date"`
	KeyWords string	`json:"key_words"`
}

type EbookDao struct {
}

func (dao EbookDao) Add(data map[string]interface{}) (err error) {
	ebook := Ebook{
		EbookName:     data["accountName"].(string),
		EnglishName:    data["accountEmail"].(string),
		AliasName: data["accountPassword"].(string),
		Category:            data["salt"].(uint),
		PublishDate:     data["accountRole"].(string),
		KeyWords:          data["status"].(string),
	}
	if err := DB.Create(&ebook).Error; err !=nil {
		return err
	}
	return nil
}