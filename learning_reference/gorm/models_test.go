package main

import (
	"github.com/stretchr/testify/assert"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
	"testing"
)

func TestProduct(t *testing.T) {
	myProduct := Product{Code: "string", Price: 10}
	assert := assert.New(t)
	assert.Equal(10, myProduct.Price)
}
func TestProductNew(t *testing.T) {
	product, err := New(&stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String("Test Name"),
		Description: stripe.String("This is a description"),
		Caption:     stripe.String("This is a caption"),
		Attributes: []*string{
			stripe.String("Attr1"),
			stripe.String("Attr2"),
		},
		URL:       stripe.String("http://example.com"),
		Shippable: stripe.Bool(true),
		PackageDimensions: &stripe.PackageDimensionsParams{
			Height: stripe.Float64(2.234),
			Length: stripe.Float64(5.10),
			Width:  stripe.Float64(6.50),
			Weight: stripe.Float64(10),
		},
		Type: stripe.String(string(stripe.ProductTypeGood)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, product)
}
