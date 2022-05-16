package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"main.go/structs"
	"main.go/views"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		
		d := getResponse(r.URL.Path)

		
		// w.WriteHeader(http.StatusNotFound)
		 json.NewEncoder(w).Encode(d)
		

	})
	http.ListenAndServe("localhost:3000",mux)
}


func getResponse(p string)  views.Response {
	fmt.Print(p)
	id := strings.TrimPrefix(p,"/")
	_, err := strconv.Atoi(id)
	var d views.Response

	if err != nil {
		d = views.Response{
			Status : http.StatusNotFound,
			Name: "Enter a valid rank",
		}
	}else{
		url := "https://pokeapi.co/api/v2/pokemon/" + id
		response , err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		resData, err := ioutil.ReadAll(response.Body)

		if err != nil{
			log.Fatal(err)
		}

		var responseObject structs.Response
   		json.Unmarshal(resData, &responseObject)
		retName := responseObject.Name

		if  retName == ""{
			d = views.Response{
				Status : http.StatusNotFound,
				Name: "Enter a valid rank",
			}
				
		}else{
			d = views.Response{
				Status : http.StatusOK,
				Name : retName,
			}
		}
	}
			
	return d
				
}