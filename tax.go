package entities

//Tax represent tax
type Tax struct {
	Percentage float64 `json:"percentage"`
	Amount     float64 `json:"amount"`
}

//CalculateAmountAfterTax calculate amount after tax addition
func (me *Tax) CalculateAmountAfterTax(subTotalOrAmount float64) float64 {
	taxAmount := subTotalOrAmount * (me.Percentage / 100)
	return subTotalOrAmount + taxAmount
}

//CalculateTaxAmountFrom calculate the tax amount
func (me *Tax) CalculateTaxAmountFrom(subTotalOrAmount float64) float64 {
	taxAmount := subTotalOrAmount * (me.Percentage / 100)
	return taxAmount
}

//SetTaxAmount set the tax amount
func (me *Tax) SetTaxAmount(subTotalOrAmount float64) {
	me.Amount = me.CalculateTaxAmountFrom(subTotalOrAmount)
}
