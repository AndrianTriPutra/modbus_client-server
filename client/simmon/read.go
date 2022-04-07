package simmon

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/simonvetter/modbus"
)

func Read(host string, slave_id uint8, timeout time.Duration) ([]string, []string, error) {

	var (
		hex           []string
		data          []string
		start, length uint16
	)

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     host,
		Timeout: timeout,
	})

	err = client.Open()
	if err != nil {
		fmt.Println("error client open")
		return hex, data, err
	}
	defer client.Close()

	client.SetUnitId(slave_id)

	// ================ ReadHoldingRegisters ================
	start = 1
	length = 4
	RHu16s, err := client.ReadRegisters(start, length, modbus.HOLDING_REGISTER)
	if err != nil {
		fmt.Println("error client read reg 1-4")
	}

	i := 0
	var ui, str string
	for _, val := range RHu16s {
		switch i {
		case 0, 1, 2:
			ui = strconv.Itoa(int(val))
		case 3:
			ui = strconv.Itoa(int(int16(val)))
		}

		str = fmt.Sprintf("%x", val)
		hex = append(hex, str)
		data = append(data, ui)
		i++
	}
	// ================ ReadHoldingRegisters ================

	// ================ ReadInputRegisters ================
	start = 5
	length = 2
	RIu16s, err := client.ReadRegisters(start, length, modbus.INPUT_REGISTER)
	if err != nil {
		fmt.Println("error client read reg 5-6")
	}
	i = 0
	for _, val := range RIu16s {
		switch i {
		case 0:
			ui = strconv.Itoa(int(val))
		case 1:
			ui = strconv.Itoa(int(int16(val)))
		}

		str = fmt.Sprintf("%x", val)
		hex = append(hex, str)
		data = append(data, ui)
		i++
	}

	start = 7
	RIu32, err := client.ReadUint32(start, modbus.INPUT_REGISTER)
	if err != nil {
		fmt.Println("error client read reg 7")
	}
	ui = strconv.Itoa(int(RIu32))
	str = fmt.Sprintf("%x", RIu32)
	hex = append(hex, str)
	data = append(data, ui)

	start = 9
	RIf32s, err := client.ReadFloat32(start, modbus.INPUT_REGISTER)
	if err != nil {
		fmt.Println("error client read reg 9")
	}
	ui = fmt.Sprintf("%.2f", RIf32s)
	str = fmt.Sprintf("%x", math.Float32bits(RIf32s))
	hex = append(hex, str)
	data = append(data, ui)

	// ================ ReadInputRegisters ================
	//client.Close()
	return hex, data, nil
}
