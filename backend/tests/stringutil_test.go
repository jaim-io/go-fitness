package tests

import (
	"testing"

	"github.com/Jaim010/jaim-io/backend/pkg/utils/utils"
)

type strToUint32Test struct {
	GivenId      string
	ReturnsError bool
}

var strToUint32Tests = []strToUint32Test{
	{GivenId: "1", ReturnsError: false},
	{GivenId: "-1", ReturnsError: true},
	{GivenId: "a", ReturnsError: true},
	// Max uint32 + 1
	{GivenId: "4294967296", ReturnsError: true},
	// Max uint64 + 1
	{GivenId: "18446744073709551616", ReturnsError: true},
}

func TestStrToUint32(t *testing.T) {
	for _, test := range strToUint32Tests {
		_, err := utils.StrToUint32(test.GivenId)

		if (err != nil && test.ReturnsError == false) || (err == nil && test.ReturnsError == true) {
			t.Errorf("Expected %v. Got %v\n", !test.ReturnsError, test.ReturnsError)
		}
	}
}
