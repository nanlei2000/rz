package rz

import "testing"

func TestUnzip(t *testing.T) {
	rz := New(true)

	err := rz.Unzip("/Users/lielienan/Project/rz/testdata/npkill.zip", "/Users/lielienan/Project/rz/testdata/npkill")

	if err != nil {
		t.Log(err)
	}
}
