package models


type Header struct {
	x_askrindo_api_gateway   string    `header:"x_askrindo_api_gateway" binding:"required"`
}