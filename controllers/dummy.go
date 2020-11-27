package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple_restfull_api/models"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	//Contoh => //https://www.soberkoder.com/consume-rest-api-go/
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err == nil {
		fmt.Println(resp.Status)
		defer resp.Body.Close()
	} else {
		panic(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct models.Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)
	fmt.Println(todoStruct.Page)
}
