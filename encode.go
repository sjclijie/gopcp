package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell int64 `json:"cell"`
	} `json:"contact"`
}

func main() {

	var JSON = `{
		"name":"Gopher",
		"title":"Programmer",
		"contact":{
			"home":"aaa",
			"cell":18611990354
		}
	}`

	var c Contact

	err := json.Unmarshal([]byte(JSON), &c)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.Name, c.Title, c.Contact.Cell, c.Contact.Home)

	var b map[string]interface{}

	err = json.Unmarshal([]byte( JSON ), &b)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b["name"], b["title"], b["contact"].(map[string]interface{})["home"].(string))

	data, err := json.MarshalIndent(b, "", "    ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}
