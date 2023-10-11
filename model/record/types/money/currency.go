package money

type Country struct {
	Name string
}

type Currency struct {
	Name          string
	Countries     Country
	Symbol        CurrencySymbol
	ExchangeRates []*ExchangeRate
}

type SymbolLocation bool

const (
	Before SymbolLocation = iota
	After
)

type CurrencySymbol struct {
	Symbol   rune
	Location SymbolLocation
}
