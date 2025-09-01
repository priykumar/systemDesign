package main

import (
	"fmt"
)

type ExType int

var latestUserId int

const (
	PERCENT ExType = iota
	EQUAL
	PRICE
)

type Splitwise struct {
	users map[int]*User           // list of users onboareded in splitwise
	trans map[int]map[int]float32 // A -> B -> 10 == B owes A Rs 10
}

type User struct {
	name        string
	phone       string
	email       string
	sendMoneyto map[int]bool
	// takeMoneyFrom map[int]bool // we can get this from trans in Splitwise
}

type Expense struct {
	paidBy   int
	totalAmt float32
	owes     map[int]float32
	exType   ExType
	// metadata interface{}
}

type SplitwiseService interface {
	AddUser(string, string, string)
	AddExpense(Expense)
	Settle()
}

func NewSplitwise() *Splitwise {
	return &Splitwise{
		users: make(map[int]*User),
		trans: make(map[int]map[int]float32),
	}
}

func NewUser(name, ph, email string) *User {
	return &User{
		name:        name,
		phone:       ph,
		email:       email,
		sendMoneyto: make(map[int]bool),
	}
}

func (s *Splitwise) AddUser(name, ph, email string) int {
	latestUserId += 1
	u := NewUser(name, ph, email)
	s.users[latestUserId] = u
	s.trans[latestUserId] = make(map[int]float32)

	return latestUserId
}

func (s *Splitwise) AddExpense(expense Expense) {
	st := splitStrategy{}
	switch expense.exType {
	case EQUAL:
		st.strategy = &EqualSplit{}
	case PERCENT:
		st.strategy = &PercentSplit{}
	default:
		st.strategy = &AmtSplit{}
	}

	st.SplitTrans(expense, s)
}

func (s *Splitwise) PrintLedger() {
	fmt.Println("After adding expense", s.trans)
	for _, u := range s.users {
		fmt.Println(u)
	}
}

func (s *Splitwise) Settle(src, dst int, amt float32) {
	dstTrans := s.trans[dst]
	if val, exist := dstTrans[src]; exist && val >= amt {
		dstTrans[src] -= amt
		if dstTrans[src] == 0 {
			user := s.users[src]
			delete(user.sendMoneyto, dst)
		}
		fmt.Println(src, "paid", dst, "amount", amt)
	}

	s.PrintLedger()
}
