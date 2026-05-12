package main

import (
	"encoding/json"
	"fmt"
)

// начало решения

// Genre описывает жанр фильма
type Genre string

func (g *Genre) UnmarshalJSON(data []byte) error {
	var tag struct {
		Name string
	}
	err := json.Unmarshal(data, &tag)
	if err != nil {
		return err
	}
	*g = Genre(tag.Name)
	return nil
}

// Movie описывает фильм
type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

// конец решения

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
