package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	toBytes := IntToBytes(100)
	fmt.Println(toBytes)

	bytesToInt := BytesToInt([]byte("0000000000000000000000000000000000000000000000000000000000000461"))
	fmt.Println("bytesToInt:", bytesToInt)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
