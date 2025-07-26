package adaptor

import (
	"testing"

	hdfc "github.com/priykumar/LLD/DesignPattern/Structural/Adaptor/hdfc"
	sbi "github.com/priykumar/LLD/DesignPattern/Structural/Adaptor/sbi"
)

func TestSBIAdapter_SendMoney(t *testing.T) {
	sbiAdapter := &SBIAdapter{sbi: &sbi.SBI{}}

	expected := "SBI: Transferred ₹1000.00 from Alice to Bob"
	result := sbiAdapter.SendMoney("Alice", "Bob", 1000)
	if result != expected {
		t.Errorf("Expected: %s, got: %s", expected, result)
	}
}

func TestHDFCAdapter_SendMoney(t *testing.T) {
	hdfcAdapter := &HDFCAdapter{hdfc: &hdfc.HDFC{}}

	expected := "HDFC: Sent ₹1000.00 from Alice to Bob"
	result := hdfcAdapter.SendMoney("Alice", "Bob", 1000)

	if result != expected {
		t.Errorf("Expected: %s, got: %s", expected, result)
	}
}
