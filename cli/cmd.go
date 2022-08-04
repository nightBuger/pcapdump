package cli

import (
	"PcapLib/pcapimpl"
	"fmt"
	"github.com/chzyer/readline"
	"os"
)

var completer = readline.NewPrefixCompleter()
var method = make(map[string]func([]string))
var globalDumper pcapimpl.Dumper

func init() {
	child := make([]readline.PrefixCompleterInterface, 0)

	//interface命令
	child = append(child,
		readline.PcItem("interface",
			readline.PcItem("list"),
			readline.PcItem("set", readline.PcItemDynamic(globalDumper.GetDevNameSlice)),
		))
	method["interface"] = inter

	//quit和exit命令
	child = append(child, readline.PcItem("quit"))
	child = append(child, readline.PcItem("exit"))
	method["quit"] = quit
	method["exit"] = quit

	//help命令
	child = append(child, readline.PcItem("help"))
	method["help"] = help

	//dump命令
	child = append(child,
		readline.PcItem("dump",
			readline.PcItem("show")))
	method["dump"] = dump

	completer.SetChildren(child)
}

func inter(subCmd []string) {
	switch {
	case len(subCmd) == 0 || subCmd[0] == "list":
		ShowDevList()
	case subCmd[0] == "set":
		if len(subCmd) == 1 {
			fmt.Println("set 模式需指定网卡名字,使用Tab键自动补全")
			break
		}
		err := globalDumper.SetInterface(subCmd[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("interface 没有该模式", subCmd[0])
	}
}
func quit([]string) {
	fmt.Println("退出入网包工具")
	os.Exit(0)
}

func help([]string) {
	fmt.Println(completer.Tree(""))
}

func dump(subCmd []string) {
	switch {
	case len(subCmd) == 0 || subCmd[0] == "run":
		fmt.Println(globalDumper.ToString())
	case subCmd[0] == "show":
		fmt.Println(globalDumper.ToString())
	default:
		fmt.Println("dump 没有该模式", subCmd[0])
	}
}
