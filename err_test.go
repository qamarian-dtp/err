package err

import (
	"errors"
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

// TestError () tests the data type "Error".
func TestError (t *testing.T) {
	str.PrintEtr ("Testing data type 'Error' ...", "std", "TestError ()")

	// Testing function New () ... {
	sec_ := errors.New ("0")
	err_ := New ("0", 1, 2, sec_)

	if err_.errString != "0" || err_.errClass != 1 || err_.errType != 2 ||
		err_.secondary != sec_ {

		str.PrintEtr ("Test has failed. Ref: 0", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method SetSecondary () ... {
	errA := New ("0", 0, 0, errors.New ("Some secondary."))
	errB := New ("0", 0, 0)
	secA := errors.New ("Some other secondary.")

	if errA.SetSecondary (secA) == true {
		str.PrintEtr ("Test has failed. Ref: 1", "err", "TestError ()")
		t.FailNow ()
	}
	if errB.SetSecondary (secA) == false {
		str.PrintEtr ("Test has failed. Ref: 2", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Error () ... {
	desc := "Some error description."
	errC := New (desc, 0, 0)

	if errC.Error () != desc {
		str.PrintEtr ("Test has failed. Ref: 3", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Class () ... {
	errD := New ("0", 1, 2)

	if errD.Class () != 1 {
		str.PrintEtr ("Test has failed. Ref: 4", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Type () ... {
	errE := New ("0", 1, 2)

	if errE.Type () != 2 {
		str.PrintEtr ("Test has failed. Ref: 5", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Unwrap () ... {
	secX := errors.New ("Some other secondary.")
	errX := New ("0", 0, 0, secX)
	errY := New ("0", 0, 0)

	if errX.Unwrap () != secX {
		str.PrintEtr ("Test has failed. Ref: 6", "err", "TestError ()")
		t.FailNow ()
	}
	if errY.Unwrap () != nil {
		str.PrintEtr ("Test has failed. Ref: 7", "err", "TestError ()")
		t.FailNow ()
	}
	// ... }

	str.PrintEtr ("Data type 'Error' passed test!", "std", "TestError ()")
}
