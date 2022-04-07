package goburrow

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"

	"github.com/goburrow/modbus"
)

func Write(host string, slave_id byte, timeout time.Duration) error {
	handler := modbus.NewTCPClientHandler(host)
	handler.Timeout = timeout
	handler.SlaveId = slave_id

	var dw = [3]int{700, 0, 50}
	var output []byte
	outputFinal := []byte{}
	for i := 0; i < 3; i++ {
		output = intToHex(int64(dw[i]))
		output = append(output[:0], output[6:]...)
		outputFinal = append(outputFinal, output...)
	}
	// log.Printf("outputFinal :%v", outputFinal)
	// log.Printf("outputFinal :%x", outputFinal)

	err := handler.Connect()
	if err != nil {
		log.Println("error connect")
		return err
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(2, 3, outputFinal) //suc
	if err != nil {
		log.Println("errror write registers")
	} else {
		log.Println("succes write registers")
	}

	//handler.Close()
	return nil
}

func intToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
