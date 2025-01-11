package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	validateBalance := func(t *testing.T, output, expectation Bitcoin) {
		t.Helper()
		if output != expectation {
			t.Errorf("output %s, expected %s", output.String(), expectation.String())
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.0))

		fmt.Printf("O endereço do saldo no Teste é %v \n", &wallet.balance)
		output := wallet.balance
		expectation := Bitcoin(10.0)

		validateBalance(t, output, expectation)
	})

	t.Run("Withdraw", func(t *testing.T) {

		t.Run("Withdraw if has funds", func(t *testing.T) {
			wallet := Wallet{Bitcoin(20.0)}

			wallet.Withdraw(Bitcoin(10.0))

			fmt.Printf("O endereço do saldo no Teste é %v \n", &wallet.balance)
			output := wallet.balance
			expectation := Bitcoin(10.0)

			validateBalance(t, output, expectation)
		})

		t.Run("Return insufficient funds if hasnt funds", func(t *testing.T) {
			wallet := Wallet{Bitcoin(20.0)}

			error := wallet.Withdraw(Bitcoin(100.0))

			if error == nil {
				t.Errorf("Was expected insufficient funds error")
			}
		})
	})
}
