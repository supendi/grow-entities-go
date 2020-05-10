package entities

import "testing"

func TestTax_CalculateAmountAfterTax(t *testing.T) {
	subTotal := 1000.0

	tax := &Tax{
		Amount:     0,
		Percentage: 10,
	}

	taxAmount := subTotal * (tax.Percentage / 100)
	expected := subTotal + taxAmount
	amountAfterTax := tax.CalculateAmountAfterTax(subTotal)

	if expected != amountAfterTax {
		t.Fatalf("The amount after tax expected to be %v but got %v", expected, amountAfterTax)
	}
}

func TestTax_CalculateTaxAmountFrom(t *testing.T) {
	subTotal := 1000.0

	tax := &Tax{
		Amount:     0,
		Percentage: 10,
	}

	expected := subTotal * (tax.Percentage / 100)
	taxAmount := tax.CalculateTaxAmountFrom(subTotal)

	if expected != taxAmount {
		t.Fatalf("The amount after tax expected to be %v but got %v", expected, taxAmount)
	}
}

func TestTax_SetTaxAmount(t *testing.T) {
	subTotal := 1000.0
	tax := &Tax{
		Amount:     0,
		Percentage: 10,
	}

	tax.SetTaxAmount(subTotal)
	expected := tax.CalculateTaxAmountFrom(subTotal)

	if expected != tax.Amount {
		t.Fatalf("tax.Amount expected to be %v but got %v", expected, tax.Amount)
	}
}
