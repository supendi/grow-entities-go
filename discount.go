package entities

const (
	//DiscountTypeNotSet is not applying discount
	DiscountTypeNotSet = "NOTSET"

	//DiscountTypePercent is discount type in percent
	DiscountTypePercent = "PERCENT"

	//DiscountTypeAmount is discount type in amount
	DiscountTypeAmount = "AMOUNT"
)

//Discount represent discount
type Discount struct {
	DiscountType string  `json:"discountType"`
	Percentage   float64 `json:"percentage"`
	Amount       float64 `json:"amount"` //store the discount
}

//CalculateAmountAfterDiscount calculate the amount after discount: amount - discount
func (me *Discount) CalculateAmountAfterDiscount(amount float64) float64 {

	if me.DiscountType == DiscountTypeNotSet {
		return amount
	}

	if (me.DiscountType == DiscountTypeAmount) || (me.Amount > 0 && me.Percentage == 0) {
		return amount - me.Amount
	}
	if (me.DiscountType == DiscountTypePercent) || (me.Percentage > 0 && me.Amount == 0) {
		discountAmount := amount * (me.Percentage / 100)
		return amount - discountAmount
	}

	return amount
}

//CalculateDiscountAmountFrom get the discount amount if discount info is in percentage
func (me *Discount) CalculateDiscountAmountFrom(subTotalOrAmount float64) float64 {
	discountAmount := subTotalOrAmount * (me.Percentage / 100)
	return discountAmount
}

//CalculateDiscountPercentageFrom get the discount percentage if discount info is in amount/nominal
func (me *Discount) CalculateDiscountPercentageFrom(subTotalOrAmount float64) float64 {
	return (me.Amount / subTotalOrAmount) * 100
}

//SetAmountOrPercentage recalculate the amount or percentage of discount
//if discount type is Amount, then percentage will be (re)calculated and set
//if discount type is Percentage, then the amount will be (re)calculated and set
func (me *Discount) SetAmountOrPercentage(subTotalOrAmount float64) {
	if (me.DiscountType == DiscountTypeAmount) || (me.Amount > 0 && me.Percentage == 0) {
		me.Percentage = me.CalculateDiscountPercentageFrom(subTotalOrAmount)
		return
	}
	me.Amount = me.CalculateDiscountAmountFrom(subTotalOrAmount)
}
