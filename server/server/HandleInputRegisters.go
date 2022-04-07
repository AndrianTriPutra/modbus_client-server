package server

import (
	"modbus_server/config"
	"modbus_server/counter"

	"github.com/simonvetter/modbus"
)

func (eh *exampleHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {

	if req.UnitId != uint8(config.ServerSetting.SlaveID) {
		err = modbus.ErrIllegalFunction
		return
	}

	// loop through all register addresses from req.addr to req.addr + req.Quantity - 1
	for regAddr := req.Addr; regAddr < req.Addr+req.Quantity; regAddr++ {
		switch regAddr {
		case 5:
			res = append(res, counter.ModbusData.Addr_5)
			// log.Printf(" [HandleInputRegisters] 5 :%x", res)
			// log.Printf(" [HandleInputRegisters] 5 :%v", res)

		case 6:
			res = append(res, uint16(counter.ModbusData.Addr_6))
			// log.Printf(" [HandleInputRegisters] 6 :%v", counter.ModbusData.Addr_6)
			// log.Printf(" [HandleInputRegisters] 6 :%x", res)
			// log.Printf(" [HandleInputRegisters] 6 :%v", res)

		case 7:
			res = append(res, counter.ModbusData.Addr_7)
			// log.Printf(" [HandleInputRegisters] 7 :%v", counter.ModbusData.UnixTs_s)
			// log.Printf(" [HandleInputRegisters] 7 :%x", counter.ModbusData.UnixTs_s)
			// log.Printf(" [HandleInputRegisters] 7 :%v", counter.ModbusData.Addr_7)
			// log.Printf(" [HandleInputRegisters] 7':%x", res)
			// log.Printf(" [HandleInputRegisters] 7':%v", res)

		case 8:
			res = append(res, counter.ModbusData.Addr_8)
			// log.Printf(" [HandleInputRegisters] 8 :%v", counter.ModbusData.UnixTs_s)
			// log.Printf(" [HandleInputRegisters] 8 :%x", counter.ModbusData.UnixTs_s)
			// log.Printf(" [HandleInputRegisters] 8 :%v", counter.ModbusData.Addr_8)
			// log.Printf(" [HandleInputRegisters] 8':%x", res)
			// log.Printf(" [HandleInputRegisters] 8':%v", res)

		case 9:
			res = append(res, counter.ModbusData.Addr_9)
			// log.Printf(" [HandleInputRegisters] 9 :%x", counter.ModbusData.Floting)
			// log.Printf(" [HandleInputRegisters] 9':%x", res)
			// log.Printf(" [HandleInputRegisters] 9':%v", res)

		case 10:
			res = append(res, counter.ModbusData.Addr_10)
			// log.Printf(" [HandleInputRegisters] 10':%x", counter.ModbusData.Floting)
			// log.Printf(" [HandleInputRegisters] 10':%x", res)
			// log.Printf(" [HandleInputRegisters] 10':%v", res)

		// attempting to access any input register address other than
		// those defined above will result in an illegal data address
		// exception client-side.
		default:
			err = modbus.ErrIllegalDataAddress
			return
		}
	}

	return
}
