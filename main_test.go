package alfred

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetArg(t *testing.T) {
	var cases = []struct {
		name string
		arg  string
	}{
		{
			name: "simple arg",
			arg:  "foobar",
		},
	}

	for _, test := range cases {
		var alf ScriptFilter
		item := alf.Add("title", "title")

		item.SetArg(test.arg)

		assert.Equal(t, test.arg, alf.Items[0].Arg, test.name)

	}
}
