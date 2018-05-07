package validate_test

import (
	"testing"
	"github.com/knocus/validate"
)

var (
	testcases = []struct{
		email	string
		err     error	
	}{
		{email:"sahalsajjad@gmail.com", err:nil,},
		{email:"sahalsajjad@gmaial.com", err:validate.ErrInvalidEmail,},
		{email:"sahalsajjad@gmail", err:validate.ErrInvalidEmail,},
		{email:"", err:validate.ErrInvalidEmail,},

	}
)

func TestEmail(t *testing.T) {
	for _, c := range testcases {
		

		err := validate.Email(c.email)
		if err != c.err {
			t.Errorf(`"Testcase: %s failed. Expected: %s, Got: %s`, c.email, c.err, err)
		}
	}
}

