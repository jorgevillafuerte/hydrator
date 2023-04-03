package main

import (
	"fmt"

	"github.com/yalochat/go-commerce-components/integrations/data"

	"github.com/yalochat/hydrator/nesting"
)

func main() {
	o := new(nesting.Order)

	hydrator := data.NewProtoHydrator()
	hydrator.RegisterMessage(o.ProtoReflect().Type())

	data := data.Record{
		"channel":     "some_channel",
		"storeId":     "some_store",
		"phoneNumber": "some_phone",
		"currency":    "some_currency",
		"from":        "some_from",
		"order": data.Record{
			"id":           "some_order_id",
			"total":        123.123,
			"sequence":     1,
			"status":       "some_status",
			"customerCode": "some_customer_code",
			"checkoutRules": []data.Record{
				{
					"type": "some_type_1",
				},
				{
					"type": "some_type_2",
				},
			},
			"items": []data.Record{
				{
					"id":        "some_item_id",
					"updatedAt": "1972-01-01T10:00:20.021Z",
					"createdAt": "1971-01-01T10:00:20.021Z",
					"price":     123.123,
					"imageUrls": []string{"some_image_1", "some_image_2"},
				},
			},
		},
	}

	msg, err := hydrator.Hydrate("tmp.hydrator.nesting.Order", data)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}
