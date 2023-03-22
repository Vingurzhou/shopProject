package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" form:"ID" column:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" form:"ProductName" gorm:"column:product_name"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" form:"ProductNum" gorm:"column:product_num"`
	ProductImage string `json:"ProductImage" sql:"productImage" form:"ProductImage" gorm:"column:product_image"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" form:"ProductUrl" gorm:"column:product_url"`
}
