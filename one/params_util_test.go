package params_test

import (
	"bytes"
	"testing"

	params "github.com/aimzeter/wuts/one"
)

func TestGetStudentID(t *testing.T) {
	tests := []struct {
		name string
		body string

		isError bool
		wantID  uint64
	}{

		{
			name: "valid body",
			body: `
				{
					"student_id": 1
				}`,
			isError: false,
			wantID:  1,
		},
		{
			name: "field not found",
			body: `
				{
					"random_id": 1
				}`,
			isError: true,
			wantID:  0,
		},
		{
			name:    "empty body",
			body:    `null`,
			isError: true,
			wantID:  0,
		},
		{
			name: "invalid body",
			body: `
				{
					student_id: 1
				}`,
			isError: true,
			wantID:  0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			input := bytes.NewBufferString(tc.body)

			id, err := params.GetStudentID(input)
			assertError(t, tc.isError, err)
			assertID(t, tc.wantID, id)
		})
	}
}

func assertError(t *testing.T, want bool, err error) {
	t.Helper()
	got := err != nil
	if want != got {
		if want {
			t.Fatalf("❌ FAIL ❌: should return error\n")
		}
		t.Fatalf("❌ FAIL ❌: should not return error, got error '%s'\n", err.Error())
	}
}

func assertID(t *testing.T, want, got uint64) {
	t.Helper()
	if want != got {
		t.Errorf("❌ FAIL ❌: did not return correct id.\n"+
			"\twant:\t%d\n"+
			"\tgot:\t%d\n", want, got)
	}
}