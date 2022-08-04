package pcapimpl

import (
	"fmt"
)

func printByteToHex(arr []byte) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%02x ", arr[i])
	}
	fmt.Println()
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
