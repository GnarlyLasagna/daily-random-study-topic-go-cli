package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if len(IT_Topics) == 0 {
		fmt.Println("No study topics configured.")
		return
	}

	randomITTopic := r.Intn(len(IT_Topics))
	fmt.Println("📘 Today's IT study topic:")
	fmt.Println("→ ", IT_Topics[randomITTopic])

	randomTHMLink := r.Intn(len(THM_Links))
	fmt.Println("📘 Today's THM provided study link:")
	fmt.Println("→ ", THM_Links[randomTHMLink])

	randomCTIFeed := r.Intn(len(CTI_Feeds))
	fmt.Println("📘 Today's CTI Feed:")
	fmt.Println("→ ", CTI_Feeds[randomCTIFeed])

}
