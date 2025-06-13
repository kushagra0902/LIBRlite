package models



type Message struct {
	ID        string   `json:"id"`
	Content   string  `json:"content"`
	Timestamp int64  `json:"timestamp"`
	Status   string  `json:"status"`
}


type ModReturn struct {
	ModID       string   `json:"modid"`
	MessageID  string  `json:"messageid"`
	
	Status   string  `json:"status"`
}




