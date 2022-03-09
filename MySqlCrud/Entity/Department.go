package Entity

import (
	_"gorm.io/driver/mysql"
	_"gorm.io/gorm"
	_"github.com/jinzhu/gorm"
)

type Department struct{
	DeptId  string  `json:"DeptId"` //`gorm:"primaryKey"`		// `json:"DeptId"`
	DeptName string  `json:"DeptName"`
	DeptLocation string `json:"DeptLocation"`
}