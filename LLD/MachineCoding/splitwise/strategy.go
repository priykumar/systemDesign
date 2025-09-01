package main

type SplitStrategy interface {
	Split(Expense, *Splitwise)
}

type splitStrategy struct {
	strategy SplitStrategy
}

func (s splitStrategy) SplitTrans(e Expense, sp *Splitwise) {
	s.strategy.Split(e, sp)
}

type EqualSplit struct{}
type PercentSplit struct{}
type AmtSplit struct{}

func (s *Splitwise) MakeTrans(src, dst int, amt float32) {
	var srcOwesDst, dstOwesSrc float32 = 0.0, 0.0
	if _, exist := s.trans[dst]; exist {
		srcOwesDst = s.trans[dst][src]
	}
	if _, exist := s.trans[src]; exist {
		dstOwesSrc = s.trans[src][dst]
	}

	if dstOwesSrc > 0 {
		s.trans[src][dst] += amt
	} else if srcOwesDst > 0 {
		if srcOwesDst > amt {
			s.trans[dst][src] -= amt
		} else if srcOwesDst < amt {
			s.trans[dst][src] = 0
			delete(s.users[src].sendMoneyto, dst)

			if _, exist := s.trans[src]; !exist {
				s.trans[src] = make(map[int]float32)
			}
			s.trans[src][dst] = amt - srcOwesDst
			s.users[dst].sendMoneyto[src] = true
		} else {
			s.trans[dst][src] = 0
			delete(s.users[src].sendMoneyto, dst)
		}
	} else {
		if _, exist := s.trans[src]; !exist {
			s.trans[src] = make(map[int]float32)
		}
		s.trans[src][dst] += amt
		s.users[dst].sendMoneyto[src] = true
	}
}

func (e *EqualSplit) Split(expense Expense, s *Splitwise) {
	paidBy := expense.paidBy
	var lendAmt float32 = 0.0
	for o := range expense.owes {
		lendAmt = expense.totalAmt / float32(len(expense.owes)+1)
		s.MakeTrans(paidBy, o, lendAmt)
	}

	s.PrintLedger()
}

func (e *PercentSplit) Split(expense Expense, s *Splitwise) {
	paidBy := expense.paidBy
	var lendAmt float32 = 0.0
	for o, amt := range expense.owes {
		lendAmt = expense.totalAmt * (amt / 100.0)
		s.MakeTrans(paidBy, o, lendAmt)
	}

	s.PrintLedger()
}

func (e *AmtSplit) Split(expense Expense, s *Splitwise) {
	paidBy := expense.paidBy
	var lendAmt float32 = 0.0
	for o, amt := range expense.owes {
		lendAmt = amt
		s.MakeTrans(paidBy, o, lendAmt)
	}

	s.PrintLedger()
}
