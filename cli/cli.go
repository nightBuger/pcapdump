package cli

import (
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"io/ioutil"
	"strings"
)

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree("    "))
}

// Function constructor - constructs new function for listing given directory
func listFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func StartCli() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m入网包工具命令行 #\033[0m ",
		HistoryFile:     "./readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		cmdSlice := strings.Fields(line)
		if len(cmdSlice) == 0 {
			continue
		}
		cmd, ok := method[cmdSlice[0]]
		if !ok {
			fmt.Println("无法识别的命令,使用<help>查看用法")
			continue
		}
		cmd(cmdSlice[1:])
	}

	quit(make([]string, 0))
}
