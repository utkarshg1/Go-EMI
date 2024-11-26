# Loan EMI Calculator

A simple command-line Loan EMI (Equal Monthly Installment) calculator built in Go by **Utkarsh Gaikwad**.

---

## Features

- Calculates EMI for loans based on:
  - **Principal Amount** (in INR)
  - **Loan Period** (in years)
  - **Annual Interest Rate** (percentage)
- Provides:
  - Monthly EMI
  - Total Amount Payable
  - Total Interest Payable

---

## Requirements

- Go 1.18 or later (supports generics).

---

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/utkarshg1/Go-EMI.git
   cd Go-EMI
   ```

2. Run the program:

   ```bash
   go run main.go
   ```

3. Follow the on-screen prompts:

   - Enter the **principal amount** in INR.
   - Enter the **loan period** in years.
   - Enter the **rate of interest** in percentage.

4. View the calculated results:
   - Monthly EMI
   - Total amount payable
   - Total interest payable

---

## Example Output

```plaintext
Loan EMI calculator by Utkarsh Gaikwad
Please enter principal in INR : 1000000
Please enter period in years : 10
Please enter rate of interest in percent p.a. : 7.5

======================================
Loan Details:
Principal Amount : 1000000 INR
Loan Period      : 10 years
Rate of Interest : 7.5% p.a.

EMI calculated        : 11870 INR
Total Amount to pay   : 1424421 INR
Total Interest to pay : 424421 INR
```

---

## Author

Developed by **Utkarsh Gaikwad**.
