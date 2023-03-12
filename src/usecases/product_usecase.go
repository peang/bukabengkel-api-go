package usecase

import (
	"context"

	"github.com/peang/bukabengkel-api-go/src/domain/entity"
	"github.com/peang/bukabengkel-api-go/src/domain/repositories"
	"github.com/peang/bukabengkel-api-go/src/utils"
)

// AuthUsecase represent the todos usecase contract
type ProductUsecase interface {
	List(ctx context.Context, page int, perPage int, sort string, filter entity.ProductEntityRepositoryFilter) ([]*entity.ProductEntity, int, error)
}

type productUsecase struct {
	productRepository repositories.ProductRepositoryInterface
}

// NewAuthUsecase will create new an authUsecase object representation of AuthUsecase interface
func NewProductUsecase(
	productRepository repositories.ProductRepositoryInterface,
) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) List(ctx context.Context, page int, perPage int, sort string, filter entity.ProductEntityRepositoryFilter) (buildings []*entity.ProductEntity, count int, err error) {
	buildings, count, err = u.productRepository.List(ctx, page, perPage, sort, filter)
	if err != nil {
		err = utils.NewInternalServerError(err)
	}

	return
}
