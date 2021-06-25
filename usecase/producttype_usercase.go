package usecase

import (
	"context"
	"golang-project/config/dbiface"
	"golang-project/models"
	productTypeRepo "golang-project/repository/mongodb"
	"io"
)


func FindProductTypeAllData(ctx context.Context, collection dbiface.CollectionAPI) ([]models.ProductType, error) {
	return productTypeRepo.FindProductsType(ctx, collection)
}

func FindProductTypeOneData(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.ProductType, error) {
	return productTypeRepo.FindProductType(ctx, id, collection)
}

func UpdateProductTypeData(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.ProductType, error) {
	return productTypeRepo.ModifyProductType(ctx, id, reqBody, collection)
}

func CreateProductTypeData(ctx context.Context, products []models.Product, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return productTypeRepo.InsertProductType(ctx, products, collection)
}

func DeleteProductTypeData(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	return productTypeRepo.DeleteProductType(ctx, id, collection)
}