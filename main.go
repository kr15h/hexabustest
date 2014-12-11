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
		// Create write packet to switch on and off
		var wPack hexabus.WritePacket = hexabus.WritePacket{hexabus.FLAG_NONE,
			1, hexabus.DTYPE_BOOL, switchState}

		// Debug
		var switchStateStr string = "Off"
		if switchState {
			switchStateStr = "On"
		}
		fmt.Printf("Switching %s\n", switchStateStr)

		// Send packet to switch
		err = wPack.Send(switchAddress)

		if err != nil {
			fmt.Println(err)
		} else {
			switchState = !switchState
		}

		time.Sleep(2000 * time.Millisecond)
	}

}
