package currencies

import "testing"

func TestFormat_DefaultOptions(t *testing.T) {
	result := Format(1000, CurrencyOptions{})

	expected := "$1.000"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestFormat_WithCurrencyCode(t *testing.T) {
	result := Format(1000, CurrencyOptions{
		Code: "COP",
	})

	expected := "$1.000 COP"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestFormat_CustomSeparators(t *testing.T) {
	result := Format(1234567, CurrencyOptions{
		ThousandsSeparator: ",",
		DecimalSeparator:   ".",
		Code:               "USD",
	})

	expected := "$1,234,567 USD"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestFormat_SmallValue(t *testing.T) {
	result := Format(50, CurrencyOptions{
		Code: "COP",
	})

	expected := "$50 COP"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
