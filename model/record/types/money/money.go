package money

type Amount struct {
	Value     int64
	Precision uint8
	Positive  bool
	Currency  *Currency
}

func (a *Amount) Subtract(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		if a.Value < amount.Value && a.Positive {
			a.Value = 0
		} else {
			a.Value = a.Value - amount.Value
		}
	}
	return a
}

func (a *Amount) Add(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		a.Value = a.Value + amount.Value
	}
	return a
}

// NOTE: These will be required to deal with integers representing decimal
//
//	places.
func (a *Amount) Multiply(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		a.Value = a.Value * amount.Value
	}
	return a
}

func (a *Amount) Divide(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		switch {
		case amount.Value != 0:
			if a.Value < amount.Value {
				a.Value = amount.Value / a.Value
			} else if amount.Value <= a.Value {
				a.Value = a.Value / amount.Value
			}
		default:
			continue
		}
	}
	return a
}
