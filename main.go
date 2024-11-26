package main

import (
	"errors"
	"fmt"
	"strconv"

	"example.com/emi-app/emi"
)

func main() {
	fmt.Println("Loan EMI calculator by Utkarsh Gaikwad")
	loan, err := loanData()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------")
		return
	}

	emi, amt, interest := loan.GetEMICalcs()

	fmt.Println("\n======================================")
	loan.OutputLoanDetails()
	fmt.Printf("EMI calculated : %.0f INR\n", emi)
	fmt.Printf("Total Amount to pay : %.0f INR\n", amt)
	fmt.Printf("Total Intrest to pay : %.0f INR\n", interest)
}

func getInputData[T float64 | uint32](promptText string) (T, error) {
	var input string
	fmt.Print(promptText)
	fmt.Scanln(&input)

	var value T
	switch any(value).(type) {
	case float64:
		parsed, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return value, fmt.Errorf("invalid float64 input: %v", err)
		}
		value = any(parsed).(T)
	case uint32:
		parsed, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return value, fmt.Errorf("invalid uint32 input: %v", err)
		}
		value = any(uint32(parsed)).(T)
	}
	return value, nil
}

func loanData() (*emi.LoanDetails, error) {
	p, err := getInputData[float64]("Please enter principal in INR : ")
	if err != nil {
		return nil, errors.New("unable to parse principal")
	}

	n, err := getInputData[uint32]("Please enter period in years : ")
	if err != nil {
		return nil, errors.New("unable to parse years")
	}

	r, err := getInputData[float64]("Please enter rate of interest in percent p.a. : ")
	if err != nil {
		return nil, errors.New("unable to parse rate of intrest")
	}

	var loan *emi.LoanDetails
	loan, err = emi.New(p, n, r)
	return loan, err
}
