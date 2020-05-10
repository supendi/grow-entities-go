package entities

import "testing"

func TestDiscountCalculateAmountAfterDiscount(t *testing.T) {
	discount := &Discount{
		DiscountType: DiscountTypePercent,
		Percentage:   10,
		Amount:       0,
	}

	subTotal := 1000.0
	expected := 900.0
	amountAfterDiscount := discount.CalculateAmountAfterDiscount(subTotal)
	if amountAfterDiscount != expected {
		t.Fatalf("amountAfterDiscount value expected to be %v but got %v", expected, amountAfterDiscount)
	}
}

func TestDiscountCalculateDiscountAmountFrom(t *testing.T) {
	discount := &Discount{
		DiscountType: DiscountTypePercent,
		Percentage:   10,
		Amount:       0,
	}

	subTotal := 1000.0
	expected := 100.0 // 1000 * (10/100)
	discountAmount := discount.CalculateDiscountAmountFrom(subTotal)
	if discountAmount != expected {
		t.Fatalf("discount amount value expected to be %v but got %v", expected, discountAmount)
	}
}

func TestDiscountCalculateDiscountPercentageFrom(t *testing.T) {
	discount := &Discount{
		DiscountType: DiscountTypeAmount,
		Percentage:   0,
		Amount:       100,
	}

	subTotal := 1000.0
	expected := 10.0
	discountPercentage := discount.CalculateDiscountPercentageFrom(subTotal)
	if discountPercentage != expected {
		t.Fatalf("Discount percentage value expected to be %v but got %v", expected, discountPercentage)
	}
}

func TestDiscountSetAmountOrPercentage_SetPercentage(t *testing.T) {
	discount := &Discount{
		DiscountType: DiscountTypeAmount,
		Percentage:   0,
		Amount:       100,
	}

	subTotal := 1000.0
	expected := 10.0
	discount.SetAmountOrPercentage(subTotal)
	if discount.Percentage != expected {
		t.Fatalf("discount.Percentage value expected to be %v but got %v", expected, discount.Percentage)
	}
}

func TestDiscountSetAmountOrPercentage_SetAmount(t *testing.T) {
	discount := &Discount{
		DiscountType: DiscountTypePercent,
		Percentage:   10,
		Amount:       1000000, //sengaja salah
	}

	subTotal := 1000.0
	expected := 100.0
	discount.SetAmountOrPercentage(subTotal)
	if discount.Amount != expected {
		t.Fatalf("discount.Amount value expected to be %v but got %v", expected, discount.Amount)
	}
}
