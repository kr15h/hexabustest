package main

import (
	"fmt"
	"github.com/morriswinkler/hexabus"
	"time"
)

var switchAddress string = "[fafa::50:c4ff:fe04:8390]"

func main() {
	fmt.Println("Hello Hexabus")
	fmt.Printf("Version: %s\n", hexabus.VERSION)

	// Get endpoint IDs
	//func QueryEids(address string, eid_qty uint16) ([]EID, error)
	eids, err := hexabus.QueryEids(switchAddress, 32)
	if err != nil {
		panic(err)
	}

	// Print eids
	for i := range eids {
		fmt.Printf("%v\n", eids[i])
	}

	// Create a bool for storing the on/off state of the remote switch
	var switchState bool = false

	// Enter infinite loop of switching on and off
	for {
		timeLoopStart := time.Now()

		// Launch hexabus switch function as a concurent routine
		// as it takes more than 3 seconds to exec
		go Switch(switchState)
		switchState = !switchState

		// Sleep for some time
		time.Sleep(500 * time.Millisecond)

		// Measure how much time the loop did take
		timeLoopElapsed := time.Since(timeLoopStart)
		fmt.Printf("Loop took %f seconds\n", timeLoopElapsed.Seconds())
	}

}

func Switch(switchState bool) {

	// Create write packet to switch on and off
	var wPack hexabus.WritePacket = hexabus.WritePacket{hexabus.FLAG_NONE,
		1, hexabus.DTYPE_BOOL, switchState}

	// Debug
	var switchStateStr string = "Off"
	if switchState {
		switchStateStr = "On"
	}

	// Send packet to switch
	timeStart := time.Now()
	fmt.Printf("Sending packet with state %s", switchStateStr)
	err := wPack.Send(switchAddress)
	timeElapsed := time.Since(timeStart)
	fmt.Printf(" took %f seconds\n", timeElapsed.Seconds())
	if err != nil {
		fmt.Println(err)
	}
}
