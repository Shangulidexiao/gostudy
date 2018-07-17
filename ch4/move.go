package ch4

import 
(	
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey Bogart","Ingrid Bergman"}},
	 {Title:"Cool Hand Luke",Year:1967,Color:true,Actors:[]string{"Plau Newman"}}}
	 
func Exec(){
	dataJson, err := json.Marshal(movies)
	if err!=nil {
		fmt.Printf("Json Marshal faild %s", err)
	}

	var myMovices []Movie
	if err := json.Unmarshal(dataJson, &myMovices); err != nil {
		fmt.Printf("Json Marshal faild %s", err)
	}
	fmt.Printf("%v", myMovices );
}