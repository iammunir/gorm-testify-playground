package models

type Account struct {
	NIK         string `json:"nik" gorm:"column:nik"`
	AccountNum  string `json:"account_num" gorm:"column:account_num"`
	AccountName string `json:"account_name" gorm:"column:account_name"`
	Balance     string `json:"balance" gorm:"column:account_balance"`
}

func (Account) TableName() string {
	return "customer_accounts"
}
