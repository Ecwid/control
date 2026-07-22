package bluetoothemulation

/*
	Event for when a GATT operation of |type| to the peripheral with |address|

happened.
*/
type GattOperationReceived struct {
	Address string            `json:"address"`
	Type    GATTOperationType `json:"type"`
}

/*
	Event for when a characteristic operation of |type| to the characteristic

respresented by |characteristicId| happened. |data| and |writeType| is
expected to exist when |type| is write.
*/
type CharacteristicOperationReceived struct {
	CharacteristicId string                      `json:"characteristicId"`
	Type             CharacteristicOperationType `json:"type"`
	Data             []byte                      `json:"data,omitempty"`
	WriteType        CharacteristicWriteType     `json:"writeType,omitempty"`
}

/*
	Event for when a descriptor operation of |type| to the descriptor

respresented by |descriptorId| happened. |data| is expected to exist when
|type| is write.
*/
type DescriptorOperationReceived struct {
	DescriptorId string                  `json:"descriptorId"`
	Type         DescriptorOperationType `json:"type"`
	Data         []byte                  `json:"data,omitempty"`
}
