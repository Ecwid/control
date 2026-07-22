package bluetoothemulation

/*
Indicates the various states of Central.
*/
type CentralState string

/*
Indicates the various types of GATT event.
*/
type GATTOperationType string

/*
Indicates the various types of characteristic write.
*/
type CharacteristicWriteType string

/*
Indicates the various types of characteristic operation.
*/
type CharacteristicOperationType string

/*
Indicates the various types of descriptor operation.
*/
type DescriptorOperationType string

/*
Stores the manufacturer data
*/
type ManufacturerData struct {
	Key  int    `json:"key"`
	Data []byte `json:"data"`
}

/*
Stores the byte data of the advertisement packet sent by a Bluetooth device.
*/
type ScanRecord struct {
	Name             string              `json:"name,omitempty"`
	Uuids            []string            `json:"uuids,omitempty"`
	Appearance       int                 `json:"appearance,omitempty"`
	TxPower          int                 `json:"txPower,omitempty"`
	ManufacturerData []*ManufacturerData `json:"manufacturerData,omitempty"`
}

/*
Stores the advertisement packet information that is sent by a Bluetooth device.
*/
type ScanEntry struct {
	DeviceAddress string      `json:"deviceAddress"`
	Rssi          int         `json:"rssi"`
	ScanRecord    *ScanRecord `json:"scanRecord"`
}

/*
	Describes the properties of a characteristic. This follows Bluetooth Core

Specification BT 4.2 Vol 3 Part G 3.3.1. Characteristic Properties.
*/
type CharacteristicProperties struct {
	Broadcast                 bool `json:"broadcast,omitempty"`
	Read                      bool `json:"read,omitempty"`
	WriteWithoutResponse      bool `json:"writeWithoutResponse,omitempty"`
	Write                     bool `json:"write,omitempty"`
	Notify                    bool `json:"notify,omitempty"`
	Indicate                  bool `json:"indicate,omitempty"`
	AuthenticatedSignedWrites bool `json:"authenticatedSignedWrites,omitempty"`
	ExtendedProperties        bool `json:"extendedProperties,omitempty"`
}

type EnableArgs struct {
	State       CentralState `json:"state"`
	LeSupported bool         `json:"leSupported"`
}

type SetSimulatedCentralStateArgs struct {
	State CentralState `json:"state"`
}

type SimulatePreconnectedPeripheralArgs struct {
	Address           string              `json:"address"`
	Name              string              `json:"name"`
	ManufacturerData  []*ManufacturerData `json:"manufacturerData"`
	KnownServiceUuids []string            `json:"knownServiceUuids"`
}

type SimulateAdvertisementArgs struct {
	Entry *ScanEntry `json:"entry"`
}

type SimulateGATTOperationResponseArgs struct {
	Address string            `json:"address"`
	Type    GATTOperationType `json:"type"`
	Code    int               `json:"code"`
}

type SimulateCharacteristicOperationResponseArgs struct {
	CharacteristicId string                      `json:"characteristicId"`
	Type             CharacteristicOperationType `json:"type"`
	Code             int                         `json:"code"`
	Data             []byte                      `json:"data,omitempty"`
}

type SimulateDescriptorOperationResponseArgs struct {
	DescriptorId string                  `json:"descriptorId"`
	Type         DescriptorOperationType `json:"type"`
	Code         int                     `json:"code"`
	Data         []byte                  `json:"data,omitempty"`
}

type AddServiceArgs struct {
	Address     string `json:"address"`
	ServiceUuid string `json:"serviceUuid"`
}

type AddServiceVal struct {
	ServiceId string `json:"serviceId"`
}

type RemoveServiceArgs struct {
	ServiceId string `json:"serviceId"`
}

type AddCharacteristicArgs struct {
	ServiceId          string                    `json:"serviceId"`
	CharacteristicUuid string                    `json:"characteristicUuid"`
	Properties         *CharacteristicProperties `json:"properties"`
}

type AddCharacteristicVal struct {
	CharacteristicId string `json:"characteristicId"`
}

type RemoveCharacteristicArgs struct {
	CharacteristicId string `json:"characteristicId"`
}

type AddDescriptorArgs struct {
	CharacteristicId string `json:"characteristicId"`
	DescriptorUuid   string `json:"descriptorUuid"`
}

type AddDescriptorVal struct {
	DescriptorId string `json:"descriptorId"`
}

type RemoveDescriptorArgs struct {
	DescriptorId string `json:"descriptorId"`
}

type SimulateGATTDisconnectionArgs struct {
	Address string `json:"address"`
}
