package controller

import (
	"encoding/json"
	"fmt"
	"github/Kushagra0902/LIBRlite/database"
	"github/Kushagra0902/LIBRlite/models"
	"github/Kushagra0902/LIBRlite/moderators"
	"net/http"
	"time"

	"github.com/google/uuid"
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
success:=addMessageHelper(message_sent.Content)
if success{
	fmt.Fprintln(w,"Message Stored")
}else{
	fmt.Fprintln(w,"Message Rejected")
}
}

func addMessageHelper(content string) bool {
	var message = new(models.Message)
	fmt.Println("Hello from helper")
	message.Content = content
	message.ID = uuid.New().String()
	message.Timestamp = time.Now().Unix()
	_,status := moderators.Validate(message)

	message.Status = status


	database.InsertMessage(message)
	return status == "Approved"
}


