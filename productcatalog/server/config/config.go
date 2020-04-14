/**** Amit Chatter (amitsosimple@gmail.com) ****/

package config

import (
	"../../../productutil/config"
	"../../../productutil/sdk"
)

type ProductInfo struct {
	Product_id 	string `json:"product_id"`
	SellerId 	string `json:"sellerId"`
	Title		string `json:"title"`
	Manufacturer string `json:"manufacturer"`
	IsLowQuantity bool	`json:"isLowQuantity"`
	IsSoldOut	bool `json:"isSoldOut"`
	IsBackorder	bool `json:"isBackorder"`
	Metafields	[]config.MetaFields `json:"metafields,omitempty"`
	RequiresShipping bool `json:"requiresShipping"`
	IsVisible	bool `json:"isVisible"`
	PublishedAt config.DateTime `json:"publishedAt,omitempty"`
	CreatedAt   config.DateTime `json:"createdAt,omitempty"`
	UpdatedAt	config.DateTime	`json:"updatedAt,omitempty"`
	Workflow	*config.Workflow  `json:"workflow,omitempty"`
}

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

func GetProductFileName() string {
	return sdkClient.GetProductFile()
}

func TranslateToProductInfo(singleProduct config.Product) ProductInfo{
	var productInfo ProductInfo

	productInfo.Product_id = singleProduct.Product_id
	productInfo.SellerId = singleProduct.SellerId
	productInfo.Title = singleProduct.Title
	productInfo.Manufacturer = singleProduct.Manufacturer
	productInfo.IsLowQuantity = singleProduct.IsLowQuantity
	productInfo.IsSoldOut = singleProduct.IsSoldOut
	productInfo.IsBackorder = singleProduct.IsBackorder
	productInfo.Metafields = singleProduct.Metafields
	productInfo.RequiresShipping = singleProduct.RequiresShipping
	productInfo.IsVisible = singleProduct.IsVisible
	productInfo.PublishedAt = singleProduct.PublishedAt
	productInfo.CreatedAt = singleProduct.CreatedAt
	productInfo.UpdatedAt = singleProduct.UpdatedAt
	productInfo.Workflow = singleProduct.Workflow

	return productInfo
}

func TranslateToProduct(singleProduct ProductInfo) config.Product{
	var productInfo config.Product

	productInfo.Product_id = singleProduct.Product_id
	productInfo.SellerId = singleProduct.SellerId
	productInfo.Title = singleProduct.Title
	productInfo.Manufacturer = singleProduct.Manufacturer
	productInfo.IsLowQuantity = singleProduct.IsLowQuantity
	productInfo.IsSoldOut = singleProduct.IsSoldOut
	productInfo.IsBackorder = singleProduct.IsBackorder
	productInfo.Metafields = singleProduct.Metafields
	productInfo.RequiresShipping = singleProduct.RequiresShipping
	productInfo.IsVisible = singleProduct.IsVisible
	productInfo.PublishedAt = singleProduct.PublishedAt
	productInfo.CreatedAt = singleProduct.CreatedAt
	productInfo.UpdatedAt = singleProduct.UpdatedAt
	productInfo.Workflow = singleProduct.Workflow

	return productInfo
}