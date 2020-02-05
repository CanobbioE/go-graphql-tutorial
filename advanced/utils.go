package main

func populate() []Tutorial {
	author := &Author{Name: "Edoardo Canobbio", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "Frist Comment"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	return tutorials
}
