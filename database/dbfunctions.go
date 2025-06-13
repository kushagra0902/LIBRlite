package database

import (
	"context"
	"github/Kushagra0902/LIBRlite/models"
)

func InsertMessage(msg *models.Message) bool{
	content:= msg.Content
	timestamp:= msg.Timestamp
	uuid := msg.ID
	status := msg.Status

	QueryString:= `INSERT INTO messages(uuid content timestamp status) VALUES($1 $2 $3 $4)`
	_,err:= Pool.Exec(context.Background(), QueryString, uuid, content, timestamp, status)

return err == nil //compares both and then if equal, returns true
}

func GetMessage(ts int64) (*[]models.Message,string){
QueryString:=`Select * FROM messages WHERE timestamp=($1) AND status=Verified`
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








