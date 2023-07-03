package endpoint

import (
	"grisko-test-app/internal/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateBsonProduct(ctx echo.Context) (bp *models.BsonProduct)
}

type DataBase interface {
	GetAll(filter *models.BsonProduct) (*models.BsonProductList, error)
	InsertOne(product *models.BsonProduct) (*mongo.InsertOneResult, error)
	DeleteProduct(id string) *models.BsonProduct
	UpdateProduct(id string, updates string) *models.BsonProduct
}

type Endpoint struct {
	s  Service
	db DataBase
}

func New(s Service, db DataBase) *Endpoint {
	return &Endpoint{
		s:  s,
		db: db,
	}
}

func (e *Endpoint) ReadProduct(ctx echo.Context) error {
	filters := e.s.CreateBsonProduct(ctx)
	all, err := e.db.GetAll(filters)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, all)

	return nil
}

func (e *Endpoint) CreateProduct(ctx echo.Context) error {
	if ctx.QueryParams().Has(NameParam) {

		newProduct := e.s.CreateBsonProduct(ctx)

		doc, err := e.db.InsertOne(newProduct)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, &doc.InsertedID)
		return nil

	} else {
		er := &models.HandlerError{
			Error: NameParam + RequiredParamError,
		}

		ctx.JSON(http.StatusBadRequest, &er)
		return nil
	}

}

func (e *Endpoint) DeleteProduct(ctx echo.Context) error {
	if ctx.QueryParams().Has(IDParam) {
		model := e.db.DeleteProduct(ctx.QueryParam(IDParam))

		ctx.JSON(http.StatusOK, &model)
		return nil

	}
	er := &models.HandlerError{
		Error: IDParam + RequiredParamError,
	}

	ctx.JSON(http.StatusBadRequest, &er)
	return nil

}

func (e *Endpoint) UpdateProduct(ctx echo.Context) error {
	if ctx.QueryParams().Has(IDParam) && ctx.QueryParams().Has(UpdateParam) {
		model := e.db.UpdateProduct(
			ctx.QueryParam(IDParam), ctx.QueryParam(UpdateParam),
		)

		ctx.JSON(http.StatusOK, &model)
		return nil

	}
	er := &models.HandlerError{
		Error: IDParam + ", " + UpdateParam + RequiredParamError,
	}

	ctx.JSON(http.StatusBadRequest, &er)
	return nil
}
