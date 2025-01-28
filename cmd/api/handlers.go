package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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