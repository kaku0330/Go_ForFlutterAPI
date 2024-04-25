package models

type Servey struct {
	Food        string `json:"food" gorm:"food"`
	Employee    string `json:"employee" gorm:"employee"`
	Environment string `json:"environment" gorm:"environment"`
	Price       string `json:"price" gorm:"price"`
}
