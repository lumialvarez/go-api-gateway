package dao

import "gorm.io/gorm"

type Route struct {
	Id           int64  `gorm:"primaryKey,autoIncrement"`
	RelativePath string `gorm:"unique"`
	UrlTarget    string
	TypeTarget   string
	Secure       bool
	Enable       bool
	Methods      []*Method `gorm:"many2many:route_methods;"`
}

type Method struct {
	gorm.Model
	Name   string   `gorm:"unique"`
	Routes []*Route `gorm:"many2many:route_methods;"`
}
