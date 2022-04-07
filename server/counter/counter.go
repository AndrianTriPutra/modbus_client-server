package counter

import (
	"log"
	"math"
	"os"
	"time"
)

var ModbusData struct {
	//readholding
	Addr_1 uint16
	Addr_2 uint16
	Addr_3 uint16
	Addr_4 int16

	//readinput
	Addr_5  uint16
	Addr_6  int16
	Addr_7  uint16
	Addr_8  uint16
	Addr_9  uint16
	Addr_10 uint16

	UnixTs_s uint32
	Floting  uint32
}

func Counter(ticker time.Ticker, stop chan bool) {
	ModbusData.Addr_1 = 1000
	ModbusData.Addr_2 = 500
	ModbusData.Addr_3 = 1
	ModbusData.Addr_4 = -1000

	ModbusData.Addr_5 = 10000
	ModbusData.Addr_6 = -10000
	for {
		select {
		case <-ticker.C:
			ModbusData.Addr_1++
			ModbusData.Addr_2++
			ModbusData.Addr_4++

			if ModbusData.Addr_1 >= 65535 {
				ModbusData.Addr_1 = 100
			} else if ModbusData.Addr_2 >= 65535 {
				ModbusData.Addr_2 = 500
			} else if ModbusData.Addr_4 >= 1000 {
				ModbusData.Addr_4 = -1000
			}

			ModbusData.Addr_5++
			ModbusData.Addr_6++
			//current unix time
			ModbusData.UnixTs_s = uint32(time.Now().Unix() & 0xffffffff)
			//the 16 most significant bits of the current unix time
			ModbusData.Addr_7 = uint16((ModbusData.UnixTs_s >> 16) & 0xffff)
			//the 16 least significant bits of the current unix time
			ModbusData.Addr_8 = uint16(ModbusData.UnixTs_s & 0xffff)

			if ModbusData.Addr_5 > 65535 {
				ModbusData.Addr_5 = 10000
			} else if ModbusData.Addr_6 > 10000 {
				ModbusData.Addr_6 = -10000
			}

			ModbusData.Floting = math.Float32bits(3.1415)
			//return 3.1415, encoded as a 32-bit floating point number in input
			ModbusData.Addr_9 = uint16((ModbusData.Floting >> 16) & 0xffff)
			// returh the 16 least significant bits of the number
			ModbusData.Addr_10 = uint16(ModbusData.Floting & 0xffff)

		case <-stop:
			log.Println(" stop Counter ")
			os.Exit(1)
		}
	}
}
