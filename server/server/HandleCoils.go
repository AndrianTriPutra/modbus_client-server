package server

import (
	"modbus_server/config"

	"github.com/simonvetter/modbus"
)

func (eh *exampleHandler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	if req.UnitId != uint8(config.ServerSetting.SlaveID) {

		err = modbus.ErrIllegalFunction
		return
	}

	return
}
