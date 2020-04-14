/**** Amit Chatter (amitsosimple@gmail.com) ****/

package productutil

import (
	"./config"
)

type AllProducts []config.Product

var Products = AllProducts{
	{
		Product_id:    		"1",
		SellerId:      		"1",
		Title:         		"1",
		Manufacturer:  		"1",
		IsLowQuantity: 		false,
		IsSoldOut:     		false,
		IsBackorder:   		false,
		Metafields:    		[]config.MetaFields{{Key:"Capacity", Value:""}, {Key:"Capacity1", Value:""}},
		RequiresShipping: 	true,
		IsVisible:        	true,
		PublishedAt:		config.DateTime{Date:"2020-02-12T08:05:39.743Z"},
		CreatedAt:			config.DateTime{"2010-08-23T05:53:16.134Z"},
		UpdatedAt:			config.DateTime{"2019-08-23T05:53:16.134Z"},
		Workflow:			&config.Workflow{Status:"new"},
		Price:				&config.Price{Range:"4.50-5.00",Min:4.50, Max:5.00},
	},
}
