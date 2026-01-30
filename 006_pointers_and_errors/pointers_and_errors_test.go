package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Initializing wallet with balance less than 10.0 would raise error", func(t *testing.T) {
		_, err := NewWallet(8.5)
		assertError(t, err, ErrInitialBalanceTooSmall)
	})

	t.Run("Wallet should return correct Balance", func(t *testing.T) {
		wallet, _ := NewWallet(15.0)

		wallet.Deposit(50.0)
		wallet.Deposit(150.0)
		wallet.Deposit(275.0)

		expected := 490.0
		assertBalance(t, *wallet, expected)
	})

	t.Run("Should be able to withdraw from wallet if balance is sufficient", func(t *testing.T) {
		wallet, _ := NewWallet(175.0)
		wallet.Deposit(275.0)

		withdrawAmount := 250.0
		wallet.Withdraw(withdrawAmount)

		expected := 200.0
		assertBalance(t, *wallet, expected)
	})

	t.Run("Should NOT be able to withdraw from wallet if balance is insufficient", func(t *testing.T) {
		wallet, _ := NewWallet(175.0)
		wallet.Deposit(275.0)

		withdrawAmount := 1850.0
		err := wallet.Withdraw(withdrawAmount)

		assertError(t, err, ErrInsufficientBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected float64) {
	t.Helper()
	balance := wallet.Balance()

	if balance != expected {
		t.Errorf("Balance error\nGot: %.2f, Expect: %.2f", balance, expected)
	}
}

func assertError(t testing.TB, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected an error to be raised")
	}

	if err != expected {
		t.Errorf("got %s, want %s", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Expected no error to be raised")
	}
}
