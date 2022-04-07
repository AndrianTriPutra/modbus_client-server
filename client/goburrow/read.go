package goburrow

import (
	"fmt"
	"log"
	"mod_client/util"
	"strconv"
	"time"

	"github.com/goburrow/modbus"
)

func Read(host string, slave_id byte, timeout time.Duration) ([]string, []string, error) {
	handler := modbus.NewTCPClientHandler(host)
	handler.Timeout = timeout
	handler.SlaveId = slave_id

	var (
		hex           []string
		data          []string
		start, length uint16
	)

	err := handler.Connect()
	if err != nil {
		log.Println("error connect")
		return hex, data, err
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	// ================ ReadHoldingRegisters ================
	//register 1-4
	start = 1
	length = 4
	databyte, err := client.ReadHoldingRegisters(start, length)
	if err != nil {
		log.Println("error cause RH timeout register 1-4")
		return hex, data, err
	}
	str := fmt.Sprintf("%x", databyte)

	//this optional, if you want return hex
	//partisi 4 character if uint16/int16
	var trimhex [4]string
	j := 0
	for i := start; i <= length; i++ {
		trimhex[j] = str[0:4]
		hex = append(hex, trimhex[j])
		th := len(str)
		if th > 4 {
			str = str[4:th]
		}
		j++
	}

	//parsing
	RHu16s := util.BytesToUint16s(util.BIG_ENDIAN, databyte)
	i := 0
	var ui string
	for _, val := range RHu16s {
		switch i {
		case 0, 1, 2:
			ui = strconv.Itoa(int(val))
		case 3:
			ui = strconv.Itoa(int(int16(val)))
		}

		data = append(data, ui)
		i++
	}

	// ================ ReadHoldingRegisters ================

	// ================ ReadInputRegisters ================
	//register 5-6
	start = 5
	length = 2
	databyte, err = client.ReadInputRegisters(start, length)
	if err != nil {
		log.Println("error cause RI timeout register 5-6")
		return hex, data, err
	}
	str = fmt.Sprintf("%x", databyte)

	j = 0
	for i := start; i < start+length; i++ {
		trimhex[j] = str[0:4]
		hex = append(hex, trimhex[j])
		th := len(str)
		if th > 4 {
			str = str[4:th]
		}
		j++
	}

	RIu16s := util.BytesToUint16s(util.BIG_ENDIAN, databyte)
	i = 0
	for _, val := range RIu16s {
		switch i {
		case 0:
			ui = strconv.Itoa(int(val))
		case 1:
			ui = strconv.Itoa(int(int16(val)))
		}

		data = append(data, ui)
		i++
	}

	//register 7-8
	start = 7
	length = 2
	databyte, err = client.ReadInputRegisters(start, length)
	if err != nil {
		log.Println("error cause RI timeout register 7-8")
		return hex, data, err
	}
	str = fmt.Sprintf("%x", databyte)
	hex = append(hex, str)

	//marge register 7-8
	RIu32s := util.BytesToUint32s(util.BIG_ENDIAN, util.HIGH_WORD_FIRST, databyte)
	for _, val := range RIu32s {
		ui = strconv.Itoa(int(val))
		data = append(data, ui)
	}

	//register 9-10
	start = 9
	length = 2
	databyte, err = client.ReadInputRegisters(start, length)
	if err != nil {
		log.Println("error cause RI timeout register 9-10")
		return hex, data, err
	}
	str = fmt.Sprintf("%x", databyte)
	hex = append(hex, str)

	//marge register 7-8
	RIf32s := util.BytesToFloat32s(util.BIG_ENDIAN, util.HIGH_WORD_FIRST, databyte)
	for _, val := range RIf32s {
		fui := float32(val)
		ui = fmt.Sprintf("%.2f", fui)
		data = append(data, ui)
	}
	// ================ ReadInputRegisters ================

	//handler.Close()
	return hex, data, nil
}
