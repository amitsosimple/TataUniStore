/**** Amit Chatter (amitsosimple@gmail.com) ****/

package config

import (
	"../../../productutil/config"
	"../../../productutil/sdk"
)

var Products []config.Product
var sdkClient *sdk.Sdk

func InitCatalogClient() {
	sdkClient,_ = sdk.NewProductClient()
}

func GetProductList() []config.Product {
	return sdkClient.GetProducts()
}

func UpdateCatalogInfo(products []config.Product) error {
	return sdkClient.UpdateCatalog(products)
}