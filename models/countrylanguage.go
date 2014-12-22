package models

type Countrylanguage struct {
	CountryCode *Country `orm:"column(CountryCode);rel(fk)"`
	Language    string   `orm:"column(Language);size(30)"`
	IsOfficial  string   `orm:"column(IsOfficial)"`
	Percentage  float32  `orm:"column(Percentage)"`
}
