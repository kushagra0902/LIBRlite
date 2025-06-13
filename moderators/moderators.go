package moderators

import (

	"fmt"
	"github/Kushagra0902/LIBRlite/models"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/google/uuid"
)
var wg *sync.WaitGroup
func Validate(msg *models.Message) *[]models.ModReturn{
	var ModResponses = new([]models.ModReturn)
	distri_chan := make(chan *models.Message)
	responseChan:= make(chan *models.ModReturn)
	// timeout,cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()


	for i := 0; i<3;i++{
		distri_chan <- msg
		wg.Add(1)

		go func(msg *models.Message){
	time.Sleep(time.Duration(rand.IntN(2)+1)*time.Second)
	var response = new(models.ModReturn) // have to use new as if not, them initially nil fields, which gives error on reassigning the values
	validate_no :=rand.IntN(2)
	if validate_no == 1{
		response.Status = "Approved"
	}else{
		response.Status = "Rejected"
	}
	response.ModID = uuid.New().String()
	response.MessageID = msg.ID
	responseChan <-response
	 defer wg.Done()
	}(<-distri_chan)


	wg.Wait()
	}
	for modResponse := range responseChan{
		err:=append(*ModResponses, *modResponse)
		if err!=nil{
			fmt.Println("Error Taking Mod response")
		}
	}

	return ModResponses
}

