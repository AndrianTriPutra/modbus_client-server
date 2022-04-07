package main

import (
	"fmt"
	"log"
	"mod_client/goburrow"
	"strconv"
	"time"
)

func goburrowRead() {
	fmt.Println()
	log.Println(" =============== goburrow.Read =============== ")
	timeout := 300 * time.Millisecond
	hex, data, err := goburrow.Read("localhost:5502", 1, timeout)
	if err != nil {
		log.Println("[goburrow.Polling] error get data")
	} else {
		// log.Printf("hex :%s", hex)
		// log.Printf("data:%s", data)

		var (
			resultHex  [8]string
			resultData [8]string
		)
		i := 0
		for _, val := range hex {
			resultHex[i] = val
			i++
		}
		i = 0
		for _, val := range data {
			resultData[i] = val
			i++
		}

		log.Printf("[reg]  [ hex ]      [data]")
		for i = 0; i < 8; i++ {
			log.Printf("[%v]     [%s]       [%s]", i+1, resultHex[i], resultData[i])
		}

		intTime, _ := strconv.ParseInt(resultData[6], 10, 64)
		//ts := time.Unix(intTime, 0).UTC()
		ts := time.Unix(intTime, 0)
		timestamp := ts.Format(time.RFC3339)

		fmt.Println()
		log.Printf("[reg]   [ hex ]           [data]              [time]")
		log.Printf("[%v]     [%s]       [%s]      [%s]", 6, resultHex[6], resultData[6], timestamp)
	}
	fmt.Println()
}

func goburrowWrite() {
	fmt.Println()
	log.Println(" =============== goburrow.Write =============== ")
	timeout := 300 * time.Millisecond
	goburrow.Write("localhost:5502", 1, timeout)
	fmt.Println()
}
