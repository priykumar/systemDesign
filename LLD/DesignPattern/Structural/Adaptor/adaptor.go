package adaptor

import (
	"fmt"

	hdfc "github.com/priykumar/LLD/DesignPattern/Structural/Adaptor/hdfc"
	sbi "github.com/priykumar/LLD/DesignPattern/Structural/Adaptor/sbi"
)

// Code with phonePe
type BankAPI interface {
	SendMoney(from, to string, amount float64) string
}

// SBI Adapter
type SBIAdapter struct {
	sbi *sbi.SBI
}

// HDFC Adapter
type HDFCAdapter struct {
	hdfc *hdfc.HDFC
}

func (a *SBIAdapter) SendMoney(from, to string, amount float64) string {
	return a.sbi.MakeTransaction(from, to, amount)
}

func (a *HDFCAdapter) SendMoney(from, to string, amount float64) string {
	details := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	return a.hdfc.TransferMoney(details)
}

func MakePayment(api BankAPI) {
	response := api.SendMoney("Alice", "Bob", 1000)
	fmt.Println(response)
}

func Adaptor_DP() {
	sbiAPI := &SBIAdapter{sbi: &sbi.SBI{}}
	hdfcAPI := &HDFCAdapter{hdfc: &hdfc.HDFC{}}

	MakePayment(sbiAPI)  // SBI: Transferred ₹1000.00 from Alice to Bob
	MakePayment(hdfcAPI) // HDFC: Sent ₹1000.00 from Alice to Bob
}
