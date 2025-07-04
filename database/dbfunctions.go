package database

import (
	"context"
	"fmt"
	"github/Kushagra0902/LIBRlite/models"
)

func InsertMessage(msg *models.Message) bool{
	fmt.Println("Inserting Data")
	content:= msg.Content
	timestamp:= msg.Timestamp
	uuid := msg.ID
	status := msg.Status

	QueryString:= `INSERT INTO messages(id ,content, "timestamp" , status) VALUES($1, $2, $3, $4)`
	_,err:= Pool.Exec(context.Background(), QueryString, uuid, content, timestamp, status)
	if err!=nil{
		panic(err)
	}
return err == nil //compares both and then if equal, returns true
}

func GetMessage(ts int64) (*[]models.Message,string){
QueryString:=`Select * FROM messages WHERE timestamp=($1) AND status='Approved'`
rows, err := Pool.Query(context.Background(), QueryString, ts)

if err!=nil{
	return nil, "Error"
}

var messages []models.Message

for rows.Next() {
	var message models.Message
	err := rows.Scan(&message.ID, &message.Content, &message.Timestamp, &message.Status)
	if err != nil {
		return nil, "Error scanning row"
	}
	messages = append(messages, message)
}

return &messages, "Success"
}

func GetAllMessages() *[]models.Message{
QueryString:=`Select * FROM messages WHERE status='Approved'`
rows,err:= Pool.Query(context.Background(),QueryString)

if err!=nil{
	fmt.Println("Error getting all msgs")
	return nil
}
var allMessages []models.Message
var message models.Message


for rows.Next(){
	err2:=rows.Scan(&message.ID, &message.Content, &message.Timestamp,&message.Status)
	if err2!=nil{
		fmt.Println("Error getting data from rows")
	}

	allMessages = append(allMessages, message)
	

}
return &allMessages
}





