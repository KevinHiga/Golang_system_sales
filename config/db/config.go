package config

type Properties struct{
	Port string `env:"MY_APP_PORT" env-default:"8080"`
	Host string `env:"HOST" env-default:"localhost"`
	DBMongo string `env:"HOST" env-default:"mongodb+srv://platzi-admin:eRTM4Ly38IzPMmCi@curso-platzi.js3sy.mongodb.net/echo-mongo-project?retryWrites=true&w=majority"`
	DBHost string `env:"DB_HOST" env-default:"localhost"`
	DBPort string `env:"DB_PORT" env-default:"27017"`
	DBName string `env:"DB_NAME" env-default:"sales-system"`
	UsersCollection string `env:"COLLECTION_NAME" env-default:"users"`
	ProductCollection string `env:"COLLECTION_NAME" env-default:"product"`
	ProductTypeCollection string `env:"COLLECTION_NAME" env-default:"product-type"`
	ProductPTypeCollection string `env:"COLLECTION_NAME" env-default:"product.p-type"`
	ProductPriceTierCollection string `env:"COLLECTION_NAME" env-default:"product.pricetier"`
	SessionCollection string `env:"COLLECTION_NAME" env-default:"sessions"`
	ChartsRenderCollection string `env:"COLLECTION_NAME" env-default:"chartsrender"`
}