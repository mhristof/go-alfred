package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	Items []ScriptFilterItem `json:"items"`
}

type ScriptFilterItem struct {
	UID      string `json:"uid,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

// Arg Set the argument for the item.
func (i *ScriptFilterItem) Arg(arg string) {
	i.Arg = arg
}

// Add Adds an item to the list. It returns the item in case it requires modifications.
func (sf *ScriptFilter) Add(title, subtitle string) *ScriptFilterItem {
	var item = ScriptFilterItem{
		UID:      title,
		Arg:      title,
		Title:    title,
		Subtitle: subtitle,
	}
	sf.Items = append(sf.Items, item)
	return &item
}

// Print Prints the items in JSON format
func (sf *ScriptFilter) Print() {
	sfJSON, err := json.MarshalIndent(sf, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sfJSON))

}
