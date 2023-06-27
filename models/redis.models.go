package models

type ExchangeSymbols struct {
	Success bool              `json:"success"`
	Symbols map[string]string `json:"symbols"`
}

type SymbolsCache struct {
	Symbols map[string]string `json:"symbols"`
}

type ExchangeRates struct {
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
	Success bool               `json:"success"`
}

type RatesCache struct {
	Rates map[string]float64 `json:"rates"`
}