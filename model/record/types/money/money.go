package money

type Amount struct {
	Value     int64
	Precision uint8
	Positive  bool
	Currency  *Currency
}

func (self *Amount) Subtract(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		if self.Value < amount.Value && self.Positive {
			self.Value = 0
		} else {
			self.Value = self.Value - amount.Value
		}
	}
	return self
}

func (self *Amount) Add(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		self.Value = self.Value + amount.Value
	}
	return self
}

// NOTE: These will be required to deal with integers representing decimal
//       places.
func (self *Amount) Multiply(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		self.Value = self.Value * amount.Value
	}
	return self
}

func (self *Amount) Divide(amounts ...[]*Amount) *Amount {
	for _, amount := range amounts {
		switch {
		case amount.Value != 0:
			if self.Value < amount.Value {
				self.Value = amount.Value / self.Value
			} else if amount.Value <= self.Value {
				self.Value = self.Value / amount.Value
			}
		default:
			continue
		}
	}
	return self
}
