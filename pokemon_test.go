package main

import (
	"encoding/json"
	"net/http"
	"testing"
	"main.go/views"
)

type addTest struct{
	path string
	response views.Response
}

var addTests = [] addTest {
	{"/abc",views.Response{Status : http.StatusNotFound,Name: "Enter a valid rank",}},
	{"/",views.Response{Status : http.StatusNotFound,Name: "Enter a valid rank",}},
	{"/23423524526",views.Response{Status : http.StatusNotFound,Name: "Enter a valid rank",}},
	{"/2",views.Response{Status : http.StatusOK,Name: "ivysaur",}}}



func TestAdd(t *testing.T){

    for _, test := range addTests{
        if output := getResponse(test.path); output != test.response {
			val,_ := json.MarshalIndent(output, "", "    ")
            t.Errorf("Output %q not equal to expected %q", val, test.response)
        }
    }
}
