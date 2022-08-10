package pcapimpl

import (
	"fmt"
)

func PrintByteToHex(arr []byte) string {
	var ret string
	for i := 0; i < len(arr); i++ {
		ret += fmt.Sprintf("%02x ", arr[i])
	}
	return ret
}

func getDevNameSlice() (devnames []string, err error) {
	devs, err := GetDeviceList()
	for _, i2 := range devs {
		devnames = append(devnames, i2.Name)
	}
	return
}

func IfEmptyToNil(input string) (output string) {
	if len(input) == 0 {
		return "<nil>"
	}
	return input
}
