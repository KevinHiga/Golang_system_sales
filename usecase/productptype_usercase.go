package usecase

import (
	"context"
	"golang-project/config/dbiface"
	"golang-project/models"
	productTypeRepo "golang-project/repository/mongodb"
)

func FindOrderbyProductsType(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Productptype, error) {
	return productTypeRepo.FindOrderProductsPType(ctx, collection)
}

func FindOrderbyPriceTier(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Productptype, error) {
	return productTypeRepo.FindOrderPriceTier(ctx, collection)
}
