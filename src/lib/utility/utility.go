package utility

import (
	"encoding/binary"
	"math/rand"
	"time"
)

//IntToBytes :Turn int type to bytes
func IntToBytes(value int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(value))
	return bs
}

//BytesToInt :Turn bytes to int type
func BytesToInt(bs []byte) int {
	value := int(binary.LittleEndian.Uint32(bs))
	return value
}

//UIntToBytes :Turn uint type to bytes
func UIntToBytes(value uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, value)
	return bs
}

//BytesToUInt :Turn bytes to uint type
func BytesToUInt(bs []byte) uint32 {
	value := binary.LittleEndian.Uint32(bs)
	return value
}

//Int16ToBytes :Turn int16 type to bytes
func Int16ToBytes(value int16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(value))
	return bs
}

//BytesToInt :Turn bytes to int type
func BytesToInt16(bs []byte) int {
	value := int(binary.LittleEndian.Uint16(bs))
	return value
}

//StringToBytes :Turn string type to bytes
func StringToBytes(str string) []byte {
	bs := []byte(str)
	return bs
}

//BytesToString :Turn bytes to string type
func BytesToString(bs []byte) string {
	str := string(bs)
	return str
}

//XORKey :Cipher Key
const XORKey = "yourgolangxorkey"

//EncryptDecrypt :With embedded key
func EncryptDecrypt(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ XORKey[i%len(XORKey)]
	}
	return output
}

//RandInteger :Get a random number less equal to maximum
func RandInteger(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// writer := util.NewBinaryWriter()
// writer.WriteInt(5)
// writer.WriteString("Hello World!")
// writer.WriteInt(10)
// writer.WriteString("BIG5 雙字元")
// writer.WriteString("JIS  アメジストワーム")
// writer.WriteInt(20)

// reader := util.NewBinaryReader(writer.ToBytes())
// fmt.Println(reader.ReadInt())
// fmt.Println(reader.ReadString())
// fmt.Println(reader.ReadInt())
// fmt.Println(reader.ReadString())
// fmt.Println(reader.ReadString())
// fmt.Println(reader.ReadInt())
