package bluetoothemulation

import (
	"github.com/ecwid/control/protocol"
)

/*
Enable the BluetoothEmulation domain.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("BluetoothEmulation.enable", args, nil)
}

/*
Set the state of the simulated central.
*/
func SetSimulatedCentralState(c protocol.Caller, args SetSimulatedCentralStateArgs) error {
	return c.Call("BluetoothEmulation.setSimulatedCentralState", args, nil)
}

/*
Disable the BluetoothEmulation domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("BluetoothEmulation.disable", nil, nil)
}

/*
	Simulates a peripheral with |address|, |name| and |knownServiceUuids|

that has already been connected to the system.
*/
func SimulatePreconnectedPeripheral(c protocol.Caller, args SimulatePreconnectedPeripheralArgs) error {
	return c.Call("BluetoothEmulation.simulatePreconnectedPeripheral", args, nil)
}

/*
	Simulates an advertisement packet described in |entry| being received by

the central.
*/
func SimulateAdvertisement(c protocol.Caller, args SimulateAdvertisementArgs) error {
	return c.Call("BluetoothEmulation.simulateAdvertisement", args, nil)
}

/*
	Simulates the response code from the peripheral with |address| for a

GATT operation of |type|. The |code| value follows the HCI Error Codes from
Bluetooth Core Specification Vol 2 Part D 1.3 List Of Error Codes.
*/
func SimulateGATTOperationResponse(c protocol.Caller, args SimulateGATTOperationResponseArgs) error {
	return c.Call("BluetoothEmulation.simulateGATTOperationResponse", args, nil)
}

/*
	Simulates the response from the characteristic with |characteristicId| for a

characteristic operation of |type|. The |code| value follows the Error
Codes from Bluetooth Core Specification Vol 3 Part F 3.4.1.1 Error Response.
The |data| is expected to exist when simulating a successful read operation
response.
*/
func SimulateCharacteristicOperationResponse(c protocol.Caller, args SimulateCharacteristicOperationResponseArgs) error {
	return c.Call("BluetoothEmulation.simulateCharacteristicOperationResponse", args, nil)
}

/*
	Simulates the response from the descriptor with |descriptorId| for a

descriptor operation of |type|. The |code| value follows the Error
Codes from Bluetooth Core Specification Vol 3 Part F 3.4.1.1 Error Response.
The |data| is expected to exist when simulating a successful read operation
response.
*/
func SimulateDescriptorOperationResponse(c protocol.Caller, args SimulateDescriptorOperationResponseArgs) error {
	return c.Call("BluetoothEmulation.simulateDescriptorOperationResponse", args, nil)
}

/*
Adds a service with |serviceUuid| to the peripheral with |address|.
*/
func AddService(c protocol.Caller, args AddServiceArgs) (*AddServiceVal, error) {
	var val = &AddServiceVal{}
	return val, c.Call("BluetoothEmulation.addService", args, val)
}

/*
Removes the service respresented by |serviceId| from the simulated central.
*/
func RemoveService(c protocol.Caller, args RemoveServiceArgs) error {
	return c.Call("BluetoothEmulation.removeService", args, nil)
}

/*
	Adds a characteristic with |characteristicUuid| and |properties| to the

service represented by |serviceId|.
*/
func AddCharacteristic(c protocol.Caller, args AddCharacteristicArgs) (*AddCharacteristicVal, error) {
	var val = &AddCharacteristicVal{}
	return val, c.Call("BluetoothEmulation.addCharacteristic", args, val)
}

/*
	Removes the characteristic respresented by |characteristicId| from the

simulated central.
*/
func RemoveCharacteristic(c protocol.Caller, args RemoveCharacteristicArgs) error {
	return c.Call("BluetoothEmulation.removeCharacteristic", args, nil)
}

/*
	Adds a descriptor with |descriptorUuid| to the characteristic respresented

by |characteristicId|.
*/
func AddDescriptor(c protocol.Caller, args AddDescriptorArgs) (*AddDescriptorVal, error) {
	var val = &AddDescriptorVal{}
	return val, c.Call("BluetoothEmulation.addDescriptor", args, val)
}

/*
Removes the descriptor with |descriptorId| from the simulated central.
*/
func RemoveDescriptor(c protocol.Caller, args RemoveDescriptorArgs) error {
	return c.Call("BluetoothEmulation.removeDescriptor", args, nil)
}

/*
Simulates a GATT disconnection from the peripheral with |address|.
*/
func SimulateGATTDisconnection(c protocol.Caller, args SimulateGATTDisconnectionArgs) error {
	return c.Call("BluetoothEmulation.simulateGATTDisconnection", args, nil)
}
