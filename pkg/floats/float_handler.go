package floats

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type FloatHandler struct {
	number            decimal.Decimal
	charNumAfterPoint int
}

func (f *FloatHandler) LessThan(compared float64) bool {
	decimalCompared := decimal.NewFromFloat(compared)
	return f.number.LessThan(decimalCompared)
}

func (f *FloatHandler) Sub(deductible float64) *FloatHandler {
	decimalDeductible := decimal.NewFromFloat(deductible)
	f.number = f.number.Sub(decimalDeductible)
	return f
}

func (f *FloatHandler) Add(summand float64) *FloatHandler {
	decimalSummand := decimal.NewFromFloat(summand)
	f.number = f.number.Add(decimalSummand)
	return f
}

func (f *FloatHandler) ConvertToFloat64() (float64, error) {
	floatAsString := f.number.String()

	floatNum, err := strconv.ParseFloat(floatAsString, 64)
	if err != nil {
		return 0, err
	}
	return floatNum, nil
}

func numDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}

func NewFloatHandler(num float64) (floatNum *FloatHandler) {
	numberFromFloat := decimal.NewFromFloat(num)
	return &FloatHandler{number: numberFromFloat, charNumAfterPoint: numDecPlaces(num)}
}
