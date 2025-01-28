package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/internal/modal"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hellow from server");
}

func (app *application) Home(w http.ResponseWriter, r *http.Request){
	msg:="We are Home"
	
	fmt.Fprintf(w,"Welcome to %v", msg);
}


func (app *application) Endpt(w http.ResponseWriter, r *http.Request){
	var payLoad = struct {
		Status string `json:"status"`;
		Message string `json:"message"`;
		

	}{
		Status:  "Success",
        Message: "API endpoint accessed",
	}


	//marshal it and Store in out
	out, err :=json.Marshal(payLoad);
	if err!=nil {
		fmt.Println(err);

	}

	//Headers
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);
	w.Write(out);
	
	


}


func (app *application) GetNames(w http.ResponseWriter , r *http.Request){

	//json banao
	var pload = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name : "BhosadPappu",
		Age : 23,
	}

	out,err := json.Marshal(pload);

	if err!= nil {
        fmt.Println(err);

    }


	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);
	w.Write(out);




}

//Get Movies

func (app *application) GetMovies(w http.ResponseWriter, R *http.Request){

	var movies [] modal.Movies
	rd,_ := time.Parse("2006-01-02", "1986-03-07")

	highlander:= modal.Movies{
	ID          :1,
	Title      :"Highlander",
	ReleaseDate :rd,
	RunTime :100,

	MPAARating :"R",
	Description :"Good Movie",
	Image :"https://www.google.com",
	CreatedAt : time.Now(),
	UpdatedAt : time.Now(),

	}
	movies = append(movies,highlander )


	rd,_= time.Parse("2006-01-02", "1999-10-15")
	fightClub:= modal.Movies{
		ID          :2,
		Title      :"Fight Club",
		ReleaseDate :rd,
		RunTime :100,
	
		MPAARating :"R",
		Description :"Goat Movie!!!!",
		Image :"https://www.google.com",
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	
		}
		movies = append(movies,fightClub);
		out,err:= json.Marshal(movies)
		if  err!=nil {
			fmt.Println(err);
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.WriteHeader(http.StatusOK);
		w.Write(out);
		
	
}