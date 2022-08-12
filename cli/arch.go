package cli

import (
	"os"
	"os/exec"
	"runtime"
)

var (
	Clear func() = nil
)

func init() {
	switch runtime.GOOS {
	case "linux":
		initLinux()
	case "windows":
		initWin()
	default:
		panic("本程序不支持你的平台:" + runtime.GOOS)
	}
}

//初始化 把对应平台的函数赋值给函数指针
func initLinux() {
	Clear = clearLinux
}
func initWin() {
	Clear = clearWin
}

//arch func
func clearWin() {
	runCommand("cmd", "/c", "cls")
}
func clearLinux() {
	runCommand("clear")
}

//执行shell命令
func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
