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

	fmt.Println("Here are your links and topics for today!")
	fmt.Println()

	randomOSINTLink := r.Intn(len(OSINT_tools_and_Links))
	fmt.Println(" THM provided study link:")
	fmt.Println("→ ", OSINT_tools_and_Links[randomOSINTLink])
	fmt.Println()

	randomCTIFeed := r.Intn(len(CTI_Feeds))
	fmt.Println(" CTI Feed to check out:")
	fmt.Println("→  ", CTI_Feeds[randomCTIFeed])
	fmt.Println()

	randomTHM_Module := r.Intn(len(THM_modules))
	fmt.Println(" THM Module to check out:")
	fmt.Println("→  ", THM_modules[randomTHM_Module])
	fmt.Println()

}
