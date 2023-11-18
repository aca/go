package decimalx

import "github.com/cockroachdb/apd/v3"

type Decimal struct {
	apd.Decimal
}

var Context = apd.BaseContext

var Zero = Decimal{}

func (x Decimal) Add(y Decimal) Decimal {
	_, err := Context.Add(&x.Decimal, &x.Decimal, &y.Decimal)
	if err != nil {
		panic(err)
	}
	return x
}

func (x Decimal) Minus(y Decimal) Decimal {
	v := Decimal{}
	v.Decimal.Set(&y.Decimal)
	v.Decimal.Neg(&v.Decimal)
	return x.Add(v)
}

func (x Decimal) IsSmaller(y Decimal) bool {
	v := x.Cmp(&y.Decimal)
	if v == -1 {
		return true
	}
	return false
}

func (x Decimal) IsBigger(y Decimal) bool {
	v := x.Cmp(&y.Decimal)
	if v == 1 {
		return true
	}
	return false
}

func New(coeff int64, exponent int32) Decimal {
	d := apd.New(coeff, exponent)
	v := Decimal{}
	v.Decimal = *d
	return v
}

func (x Decimal) String() string {
	return x.Decimal.Text('f')
}

func FromString(input string) Decimal {
	v := Decimal{}
	_, _, err := v.Decimal.SetString(input)
	if err != nil {
		panic(err)
	}
	return v
}
