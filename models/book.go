package models

import (
  "gorm.io/gorm"
)


type Book struct {
  gorm.Model
  Title  string `json:"title"`
  Author string `json:"author"`
  Code  string  `json:"code"`
  Price uint    `json:"price"`
}

type CreateBookInput struct {
	Title  string `json:"title" xml:"title" form:"title" binding:"required"`
	Author string `json:"author" xml:"author" form:"author" binding:"required"`
	Code string `json:"code" xml:"code" form:"code"  binding:"required"`
	Price uint `json:"price" xml:"price" form:"price"  binding:"required"`
}

type CreateBookInputForm struct {
	Title  string `form:"title" binding:"required"`
	Author string `form:"author" binding:"required"`
	Code string `form:"code" binding:"required"`
	Price uint `form:"price" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Code string `json:"code"`
	Price uint `json:"price"`
}