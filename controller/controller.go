package controller

import (
	"encoding/json"
	"fmt"
	"github/Kushagra0902/LIBRlite/database"
	"github/Kushagra0902/LIBRlite/models"
	"github/Kushagra0902/LIBRlite/moderators"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)


type Request struct{
	Content string `json:"content"`
}

func AddMessage(w http.ResponseWriter, r *http.Request){
decoder := json.NewDecoder(r.Body)
var message_sent Request
err:=decoder.Decode(&message_sent)
fmt.Println(message_sent.Content)
if err!=nil{
	fmt.Fprintln(w,"Error getting the messgae")
}
response2 := new(models.ResponseClient)

success,response2:=addMessageHelper(message_sent.Content)
if success{
	err:= json.NewEncoder(w).Encode(response2)
	if err!=nil{
		panic(err)
	}
}else{
	err2:= json.NewEncoder(w).Encode(response2)
	if err2!=nil{
		panic(err)
	}
}
}

func addMessageHelper(content string) (bool, *models.ResponseClient) {
	var message = new(models.Message)
	//fmt.Println("Hello from helper")
	message.Content = content
	message.ID = uuid.New().String()
	message.Timestamp = time.Now().Unix()
	_,status := moderators.Validate(message)

	message.Status = status


	database.InsertMessage(message)

	Response := new(models.ResponseClient)
	Response.ID = message.ID
	Response.Status = message.Status
	Response.Timestamp = message.Timestamp
	return status == "Approved", Response
	
}


func GetMessages(w http.ResponseWriter, r *http.Request){

	queryParams := mux.Vars(r)
	timestamp:=queryParams["ts"]
	ts,_ := strconv.ParseInt(timestamp, 10, 64)
	fmt.Println(ts)
	messageList,_ := database.GetMessage(ts)
	if len(*messageList)!=0 {
		json.NewEncoder(w).Encode(messageList)
		//fmt.Println("test print")
	}else{
	json.NewEncoder(w).Encode("Message is rejected by mods. Please stick to communtiy guidelines")
	
	}

}

func GetAllMessages(w http.ResponseWriter, r *http.Request){

	
	messageList := database.GetAllMessages()
	if len(*messageList)!=0 {
		json.NewEncoder(w).Encode(messageList)
		//fmt.Println("test print")
	}else{
	json.NewEncoder(w).Encode("Message is rejected by mods. Please stick to communtiy guidelines")
	
	}

}





