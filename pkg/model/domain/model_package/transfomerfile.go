package model_package

type TransformerFile struct {
	ID   string `gorm:"primarykey"`
	Name string
	URL  string
}
