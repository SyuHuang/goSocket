package proto

import (
	"reflect"

	util "../lib/utility"
)

//IProtocol :Interface of protocol, all protocols must implement below functions.
type IProtocol interface {
	FromBytes(*util.BinaryReader)
	ToBytes(*util.BinaryWriter)
	Todo()
}

//Collection of all dummys for dup by name.
var protoFactory = make(map[string]IProtocol)

//InitProtocol :Call for regist all the protocols in this game.
func InitProtocol() {
	RegistProtocol(new(HeartBeatProto))
	RegistProtocol(new(PingProto))
}

//RegistProtocol :Put a dummy on map for dynamic creating.
func RegistProtocol(dummy IProtocol) {
	dummyType := reflect.TypeOf(dummy).Elem()
	protoFactory[dummyType.Name()] = dummy
}

//MakeInstance :Create a new instance by it's dummy.
func MakeInstance(name string) IProtocol {
	if dummy, ok := protoFactory[name]; ok {
		t := reflect.New(reflect.TypeOf(dummy).Elem()).Interface()
		if instance, ok := t.(IProtocol); ok {
			return instance
		}
	}
	return nil
}

//TransferRawDataToProtocol :
func TransferRawDataToProtocol(rawData []byte) []IProtocol {

	decrypted := util.EncryptDecrypt(rawData)

	reader := util.NewBinaryReader(decrypted)
	counts := reader.ReadInt()
	protos := make([]IProtocol, counts)

	for i := 0; i < counts; i++ {
		name := reader.ReadString()
		clone := MakeInstance(name)
		clone.FromBytes(reader)

		protos[i] = clone
	}

	return protos
}

//TransferProtocolToRawData :Response to transfer game protocols to rawdata
func TransferProtocolToRawData(protos []IProtocol) []byte {

	writer := util.NewBinaryWriter()
	writer.WriteInt(len(protos))

	for _, proto := range protos {
		name := reflect.TypeOf(proto).Elem().Name()
		writer.WriteString(name)
		proto.ToBytes(writer)
	}

	rawData := writer.ToBytes()
	encryped := util.EncryptDecrypt(rawData)

	return encryped
}
