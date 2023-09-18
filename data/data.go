package data

type Tutorial struct {
	ID       int
	Title    string
	Author   string
	Comments []Comment
}

type Comment struct {
	Body string
}

func Populate() []Tutorial {
	tutorial := Tutorial{
		ID:     1,
		Title:  "Jujutsu Kaijen",
		Author: "Ashutosh",
		Comments: []Comment{
			Comment{Body: "First Comment"},
			Comment{Body: "Second Comment"},
		},
	}
	tutorial2 := Tutorial{
		ID:     2,
		Title:  "Naruto",
		Author: "Ashutosh",
		Comments: []Comment{
			Comment{Body: "Second Comment"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	tutorials = append(tutorials, tutorial2)
	return tutorials
}
