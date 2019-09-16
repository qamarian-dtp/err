package err

import (
	"errors"
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

// TestError () tests the data type "Error".
func TestError (t *testing.T) {
	str.PrintEtr ("Testing data type 'Error' ...", "std", "TestError ()")

	errCause := errors.New ("Some cause error.")
	errDescp := "Some error."
	someErr := New (errDescp, 9, 15, errCause)

	if someErr.Error () != errDescp {
		str.PrintEtr ("Test has failed. Ref: 1", "err", "TestError ()")
		t.Fail ()
	}
	if someErr.Class () != 9 {
		str.PrintEtr ("Test has failed. Ref: 2", "err", "TestError ()")
		t.Fail ()
	}
	if someErr.Type () != 15 {
		str.PrintEtr ("Test has failed. Ref: 3", "err", "TestError ()")
		t.Fail ()
	}
	if someErr.Unwrap () != errCause {
		str.PrintEtr ("Test has failed. Ref: 5", "err", "TestError ()")
		t.Fail ()
	}

	err2 := New (errDescp, 1, 2)
	if err2.Unwrap () != nil {
		str.PrintEtr ("Test has failed. Ref: 6", "err", "TestError ()")
		t.Fail ()
	}

	str.PrintEtr ("Data type 'Error' passed test!", "std", "TestError ()")
}