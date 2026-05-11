package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	dur := time.Duration(d)
	h := int(dur.Hours())
	dur -= time.Hour * time.Duration(h)
	m := int(dur.Minutes())

	var res string
	if h > 0 {
		res += fmt.Sprintf("%dh", h)
	}
	if m > 0 {
		res += fmt.Sprintf("%dm", m)
	}

	return []byte(`"` + res + `"`), nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	var res string
	for range r {
		res += "★"
	}
	for range 5 - r {
		res += "☆"
	}
	return []byte(`"` + res + `"`), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	if indent > 0 {
		return marshalIndent(indent, movies...)
	}

	res, err := json.Marshal(movies)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func marshalIndent(indent int, movies ...Movie) (string, error) {
	var s string
	for range indent {
		s += " "
	}
	res, err := json.MarshalIndent(movies, "", s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(0*time.Hour + 49*time.Minute),
		Rating:   1,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 0*time.Minute),
		Rating:   0,
	}

	s, err := MarshalMovies(8, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
