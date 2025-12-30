package mongoutil

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Aggregate []bson.M

func Float64ToDecimal128(value float64) primitive.Decimal128 {
	decimal128, err := primitive.ParseDecimal128(fmt.Sprintf("%0.2f", value))
	if err != nil {
		panic(err)
	}
	return decimal128
}

func Decimal128ToFloat64(decimal128 primitive.Decimal128) float64 {

	value, err := strconv.ParseFloat(decimal128.String(), 64)
	if err != nil {
		panic(err)
	}
	return value
}
