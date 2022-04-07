package simmon

import (
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

func Write(host string, slave_id uint8, timeout time.Duration) error {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     host,
		Timeout: timeout,
	})

	err = client.Open()
	if err != nil {
		fmt.Println("error client open")
		return err
	}
	defer client.Close()

	client.SetUnitId(slave_id)

	var val []uint16
	val = append(val, 800) //reg 2
	val = append(val, 1)   //reg 3
	val = append(val, 70)  //reg 4
	err = client.WriteRegisters(2, val)
	if err != nil {
		log.Println("errror write registers")
		return err
	} else {
		log.Println("succes write registers")
	}

	return nil
}
