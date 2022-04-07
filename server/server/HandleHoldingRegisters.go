package server

import (
	"modbus_server/config"
	"modbus_server/counter"

	"github.com/simonvetter/modbus"
)

func (eh *exampleHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	var regAddr uint16

	if req.UnitId != uint8(config.ServerSetting.SlaveID) {
		// only accept unit ID #1
		err = modbus.ErrIllegalFunction
		return
	}

	// loop through `quantity` registers
	for i := 0; i < int(req.Quantity); i++ {
		// compute the target register address
		regAddr = req.Addr + uint16(i)

		switch regAddr {
		// expose the static, read-only
		case 1:
			res = append(res, counter.ModbusData.Addr_1)
			//log.Printf(" [HandleHoldingRegisters] 1 :%x", res)
			//log.Printf(" [HandleHoldingRegisters] 1 :%v", res)

		// expose holdingReg1 in register 2 (RW)
		case 2:
			if req.IsWrite {
				counter.ModbusData.Addr_2 = req.Args[i]
				eh.holdingReg1 = counter.ModbusData.Addr_2
			}
			eh.holdingReg1 = counter.ModbusData.Addr_2
			res = append(res, eh.holdingReg1)
			//log.Printf(" [HandleHoldingRegisters] 2 :%x", res)
			//log.Printf(" [HandleHoldingRegisters] 2 :%v", res)

		// expose holdingReg2 in register 3 (RW)
		case 3:
			if req.IsWrite {
				// only accept values 1 and 2
				switch req.Args[i] {
				case 0, 1:
					counter.ModbusData.Addr_3 = req.Args[i]
					eh.holdingReg2 = counter.ModbusData.Addr_2

					// make note of the change (e.g. for auditing purposes)
					// fmt.Printf("%s set reg#3 to %v\n", req.ClientAddr, eh.holdingReg2)
				default:
					// if the written value is neither 1 nor 2,
					// return a modbus "illegal data value" to
					// let the client know that the value is
					// not acceptable.
					err = modbus.ErrIllegalDataValue
					return
				}
			}
			eh.holdingReg2 = counter.ModbusData.Addr_3
			res = append(res, eh.holdingReg2)
			//log.Printf(" [HandleHoldingRegisters] 3 :%x", res)
			//log.Printf(" [HandleHoldingRegisters] 3 :%v", res)

		// expose eh.holdingReg3 in register 4 (RW)
		// note: eh.holdingReg3 is a signed 16-bit integer
		case 4:
			if req.IsWrite {
				// cast the 16-bit unsigned integer passed by the server
				// to a 16-bit signed integer when writing
				counter.ModbusData.Addr_4 = int16(req.Args[i])
				eh.holdingReg3 = counter.ModbusData.Addr_4
			}
			// cast the 16-bit signed integer from the handler to a 16-bit unsigned
			// integer so that we can append it to `res`.
			eh.holdingReg3 = counter.ModbusData.Addr_4
			res = append(res, uint16(eh.holdingReg3))
			//log.Printf(" [HandleHoldingRegisters] 4 :%v", counter.ModbusData.Addr_4)
			//log.Printf(" [HandleHoldingRegisters] 4 :%x", res)
			//log.Printf(" [HandleHoldingRegisters] 4 :%v", res)

		// any other address is unknown
		default:
			err = modbus.ErrIllegalDataAddress
			return
		}
	}

	return
}
