package emi

import (
	"errors"
	"fmt"
	"math"
)

type LoanDetails struct {
	principal      float64
	period         uint32
	rateOfInterest float64
}

func New(principal float64, period uint32, rateOfInterest float64) (*LoanDetails, error) {
	if principal <= 0 {
		return nil, errors.New("principal amount must be greater than zero")
	}
	if period == 0 {
		return nil, errors.New("loan period cannot be zero")
	}
	if rateOfInterest <= 0 || rateOfInterest > 100 {
		return nil, errors.New("rate of interest must be between 0 and 100")
	}
	return &LoanDetails{
		principal:      principal,
		period:         period,
		rateOfInterest: rateOfInterest,
	}, nil
}

func (l *LoanDetails) OutputLoanDetails() {
	fmt.Printf("Principal : %.2f INR\n", l.principal)
	fmt.Printf("Period : %d years\n", l.period)
	fmt.Printf("Rate of Intrest : %.2f percent per annum\n", l.rateOfInterest)
}

func (l *LoanDetails) GetEMICalcs() (float64, float64, float64) {
	n := float64(l.period) * 12.0
	r := l.rateOfInterest / 1200.0
	x := math.Pow(1+r, n)
	emi := l.principal * r * x / (x - 1)
	amt := emi * n
	interest := amt - l.principal
	return emi, amt, interest
}
