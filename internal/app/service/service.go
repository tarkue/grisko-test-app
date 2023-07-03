package service

import (
	"grisko-test-app/internal/app/endpoint"
	"grisko-test-app/internal/app/models"

	"github.com/labstack/echo/v4"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateBsonProduct(ctx echo.Context) (bp *models.BsonProduct) {
	return &models.BsonProduct{
		Name:  ctx.QueryParam(endpoint.NameParam),
		Info:  ctx.QueryParam(endpoint.InfoParam),
		Img:   ctx.QueryParam(endpoint.ImgParam),
		Price: ctx.QueryParam(endpoint.PriceParam),
	}

}
