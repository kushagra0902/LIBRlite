package moderators

import (
	
	"github/Kushagra0902/LIBRlite/models"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/google/uuid"
)

func Validate(msg *models.Message) (*[]models.ModReturn, string) {
	//fmt.Println("Hello from mod file start")

	var wg sync.WaitGroup
	var ModResponses = new([]models.ModReturn)

	distri_chan := make(chan *models.Message, 3)   // buffered so it doesn't block
	responseChan := make(chan *models.ModReturn, 3) // buffered to collect results

	// Start 3 worker goroutines that read from distri_chan
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			m := <-distri_chan // read from channel
			time.Sleep(2 * time.Second)

			var response = new(models.ModReturn)
			validate_no := rand.IntN(2)
			if validate_no == 1 {
				response.Status = "Approved"
			} else {
				response.Status = "Rejected"
			}
			response.ModID = uuid.New().String()
			response.MessageID = m.ID
			responseChan <- response

			//fmt.Println("Hello from mod")
		}()
	}

	// Send 3 copies of msg into the channel
	for i := 0; i < 3; i++ {
		distri_chan <- msg
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(responseChan)
	//fmt.Println("Hello from mod file end")

	// Collect responses
	var Approved_counter int
	for modResponse := range responseChan {
		*ModResponses = append(*ModResponses, *modResponse)
		if modResponse.Status == "Approved" {
			Approved_counter++
		}
	}

	if Approved_counter >= 2 {
		return ModResponses, "Approved"
	}
	return ModResponses, "Rejected"
}
 

// here earlier I was sending the msg to the districhain and passing the value from districhain in the same loop. What this was causing is data simultaneously flow nhi kar rha tha. Therefore we separated the process of sending itno distri chain in a diff for loop and since now my routines are working, They spontaneously listens to the channel thus ensuring no deadlock, which we witnessed earlier.