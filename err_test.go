package err

import (
	"errors"
	"gopkg.in/qamarian-lib/str.v2"
	"math/big"
	"testing"
)

// TestError () tests the data type "Error".
func TestError (t *testing.T) {
	str.PrintEtr ("Testing has started!' ...", "std", "TestError ()")

	// Testing method Error () ... {
	desc := "Some error description."
	errC := New (desc, big.NewInt (0), big.NewInt (0))

	if errC.Error () != desc {
		str.PrintEtr ("Test failed. Ref: 0", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Class () ... {
	errD := New ("0", big.NewInt (1), big.NewInt (2))

	if errD.Class ().Cmp (big.NewInt (1)) != 0 {
		str.PrintEtr ("Test failed. Ref: 1", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Type () ... {
	errE := New ("0", big.NewInt (1), big.NewInt (2))

	if errE.Type ().Cmp (big.NewInt (2)) != 0 {
		str.PrintEtr ("Test failed. Ref: 2", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Unwrap () ... {
	secX := errors.New ("Some other secondary.")
	errX := New ("0", big.NewInt (0), big.NewInt (0), secX)
	errY := New ("0", big.NewInt (0), big.NewInt (0))

	if errX.Unwrap () != secX {
		str.PrintEtr ("Test failed. Ref: 3", "err", "TestError ()")
		t.FailNow ()
	}
	if errY.Unwrap () != nil {
		str.PrintEtr ("Test failed. Ref: 4", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	str.PrintEtr ("Test passed!", "std", "TestError ()")
}
