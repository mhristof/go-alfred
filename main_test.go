package alfred

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetArg(t *testing.T) {
	var cases = []struct {
		name string
		arg  string
		icon string
	}{
		{
			name: "simple arg",
			arg:  "foobar",
		},
		{
			name: "simple arg",
			icon: "/dev/null",
		},
	}

	for _, test := range cases {
		var alf ScriptFilter
		item := alf.Add("title", "title")

		if test.arg != "" {
			item.SetArg(test.arg)

			assert.Equal(t, test.arg, alf.Items[0].Arg, test.name)
		}

		if test.icon != "" {
			item.SetIcon(test.icon)

			assert.Equal(t, test.icon, alf.Items[0].Icon["path"], test.name)
		}

	}
}
