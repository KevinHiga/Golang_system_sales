package usecase

import (
	"context"
	"golang-project/config/dbiface"
	"golang-project/models"
	productRepo "golang-project/repository/mongodb"
	"io"
)

func FindProductAllData(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Product, error) {
	return productRepo.FindProducts(ctx, collection)
}

func FindProductOneData(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Product, error) {
	return productRepo.FindProduct(ctx, id, collection)
}

func UpdateProductData(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.Product, error) {
	return productRepo.ModifyProduct(ctx, id, reqBody, collection)
}

func CreateProductData(ctx context.Context, products []models.Product, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return productRepo.InsertProduct(ctx, products, collection)
}

func DeleteProductData(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	return productRepo.DeleteProduct(ctx, id, collection)
}
