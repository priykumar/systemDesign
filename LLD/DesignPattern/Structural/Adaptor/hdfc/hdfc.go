package hdfc

import "fmt"

type HDFC struct{}

func (h *HDFC) TransferMoney(details map[string]interface{}) string {
	return fmt.Sprintf("HDFC: Sent ₹%.2f from %s to %s",
		details["amount"], details["from"], details["to"])
}
