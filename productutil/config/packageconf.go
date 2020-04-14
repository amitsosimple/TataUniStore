/**** Amit Chatter (amitsosimple@gmail.com) ****/

package config

type Product struct {
	Product_id 	string `json:"product_id"`
	SellerId 	string `json:"sellerId"`
	Title		string `json:"title"`
	Manufacturer string `json:"manufacturer"`
	IsLowQuantity bool	`json:"isLowQuantity"`
	IsSoldOut	bool `json:"isSoldOut"`
	IsBackorder	bool `json:"isBackorder"`
	Metafields	[]MetaFields `json:"metafields,omitempty"`
	RequiresShipping bool `json:"requiresShipping"`
	IsVisible	bool `json:"isVisible"`
	PublishedAt DateTime `json:"publishedAt,omitempty"`
	CreatedAt   DateTime `json:"createdAt,omitempty"`
	UpdatedAt	DateTime	`json:"updatedAt,omitempty"`
	Workflow	*Workflow  `json:"workflow,omitempty"`
	Price		*Price `json:"price,omitempty"`
}

type MetaFields struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Workflow struct {
	Status string `json:"status"`
}

type DateTime struct {
	Date string `json:"$date"`
}

type Price struct {
	Range string `json:"range,omitempty"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

