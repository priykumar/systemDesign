package sbi

import "fmt"

// This code is with SBI bank
type SBI struct{}

func (s *SBI) MakeTransaction(sourceAcc, destAcc string, amt float64) string {
	return fmt.Sprintf("SBI: Transferred ₹%.2f from %s to %s", amt, sourceAcc, destAcc)
}
