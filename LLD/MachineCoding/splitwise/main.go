package main

import "fmt"

func main() {
	swise := NewSplitwise()
	u1 := swise.AddUser("U1", "1234", "u1@abc.com")
	u2 := swise.AddUser("U2", "5678", "u2@abc.com")
	u3 := swise.AddUser("U3", "1357", "u3@abc.com")

	fmt.Println("1 --------------------------------")
	e1 := Expense{
		paidBy:   u1,
		totalAmt: 100,
		owes: map[int]float32{
			u2: 10,
		},
		exType: PERCENT,
	}
	swise.AddExpense(e1) // u2 owes u1 10

	fmt.Println("2 --------------------------------")
	e2 := Expense{
		paidBy:   u2,
		totalAmt: 1000,
		owes: map[int]float32{
			u1: 0,
			// u3: 0,
		},
		exType: EQUAL,
	}
	swise.AddExpense(e2) // u1 owes u2 490

	fmt.Println("3 --------------------------------")
	e3 := Expense{
		paidBy:   u1,
		totalAmt: 100,
		owes: map[int]float32{
			u2: 500,
			u3: 300,
		},
		exType: PRICE,
	}
	swise.AddExpense(e3) // u2 owes u1 10 // u3 owes u1 300

	fmt.Println("4 --------------------------------")
	swise.Settle(u3, u1, 10)

	fmt.Println("5 --------------------------------")
	swise.Settle(u3, u1, 290)

}
