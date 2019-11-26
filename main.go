package main

import (
	"Yoga-Go/internal/yogaconfig"
	"Yoga-Go/yogaliner"
	"Yoga-Go/yogaliner/args"
	"Yoga-Go/yogautil"
	"Yoga-Go/yogautil/escaper"
	"Yoga-Go/yogaverbose"
	"fmt"
	"github.com/peterh/liner"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
)

const (
	NameShortDisplayNum = 16
)

var (
	Version = "1.0.0"

	isCli bool

	historyFilePath = filepath.Join(yogaconfig.GetConfigDir(), "yoga_command_history.txt")
)

func main() {
	app := cli.NewApp()
	app.Name = "Yoga-Go"
	app.Version = Version
	app.Author = "yoga Huang"
	app.Usage = "Yoga-Pan for " + runtime.GOOS + "/" + runtime.GOARCH
	app.Description = `Yoga-Go 使用Go语言编写的Yoga网盘命令行客户端, 为操作Yoga网盘, 提供实用功能.
	具体功能, 参见 COMMANDS 列表

	特色:
		网盘内列出文件和目录, 支持通配符匹配路径;
		下载网盘内文件, 支持网盘内目录 (文件夹) 下载, 支持多个文件或目录下载, 支持断点续传和高并发高速下载.`

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "启动调试",
			EnvVar:      yogaverbose.EnvVerbose,
			Destination: &yogaverbose.IsVerbose,
		},
	}

	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}

		isCli = true
		yogaverbose.Verbosef("VERBOSE: 这是一条调试信息\n\n")

		var (
			line = yogaliner.NewLiner()
			err  error
		)
		line.History, err = yogaliner.NewLineHistory(historyFilePath)
		if err != nil {
			fmt.Printf("警告: 读取历史命令文件错误, %s\n", err)
		}

		line.ReadHistory()
		defer func() {
			line.DoWriteHistory()
			line.Close()
		}()

		line.State.SetCompleter(func(line string) (s []string) {
			var (
				lineArgs                   = args.Parse(line)
				numArgs                    = len(lineArgs)
				acceptCompleteFileCommands = []string{
					"cd", "ls", "mkdir", "rm", "share", "tree", "upload",
				}

				closed = strings.LastIndex(line, " ") == len(line)-1
			)

			for _, cmd := range app.Commands {
				for _, name := range cmd.Names() {
					if !strings.HasPrefix(name, line) {
						continue
					}

					s = append(s, name+" ")
				}
			}

			switch numArgs {
			case 0:
				return
			case 1:
				if !closed {
					return
				}
			}

			thisCmd := app.Command(lineArgs[0])
			if thisCmd == nil {
				return
			}

			if !yogautil.ContainsString(acceptCompleteFileCommands, thisCmd.FullName()) {
				return
			}

			var (
				runeFunc = unicode.IsSpace
				//pcsRuneFunc = func(r rune) bool {
				//	switch r {
				//	case '\'', '"':
				//		return true
				//	}
				//	return unicode.IsSpace(r)
				//}
				targetPath string
			)

			if !closed {
				targetPath = lineArgs[numArgs-1]
				escaper.EscapeStringsByRuneFunc(lineArgs[:numArgs-1], runeFunc) // 转义
			} else {
				escaper.EscapeStringsByRuneFunc(lineArgs, runeFunc)
			}

			switch {
			case targetPath == "." || strings.HasSuffix(targetPath, "/."):
				s = append(s, line+"/")
				return
			case targetPath == ".." || strings.HasSuffix(targetPath, "/.."):
				s = append(s, line+"/")
				return
			}

			return

		})

		for {
			prompt := app.Name + " > "
			commandLine, err := line.State.Prompt(prompt)
			switch err {
			case liner.ErrPromptAborted:
				return
			case nil:
				// continue
			default:
				fmt.Println(err)
				return
			}

			line.State.AppendHistory(commandLine)

			cmdArgs := args.Parse(commandLine)
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{os.Args[0]}
			s = append(s, cmdArgs...)

			// 恢复原始终端状态
			// 防止运行命令时程序被结束, 终端出现异常
			line.Pause()
			c.App.Run(s)
			line.Resume()
		}

	}

	app.Commands = []cli.Command{
		{
			Name:  "env",
			Usage: "显示程序环境变量",
			Description: `
	YOGA_GO_CONFIG_DIR: 配置文件路径,
	YOGA_GO_VERBOSE: 是否启用调试.`,
			Category: "其他",
			Action: func(c *cli.Context) error {
				envStr := "%s=\"%s\"\n"
				envVar, ok := os.LookupEnv(yogaverbose.EnvVerbose)
				if ok {
					fmt.Printf(envStr, yogaverbose.EnvVerbose, envVar)
				} else {
					fmt.Printf(envStr, yogaverbose.EnvVerbose, "0")
				}

				envVar, ok = os.LookupEnv(yogaconfig.EnvConfigDir)
				if ok {
					fmt.Printf(envStr, yogaconfig.EnvConfigDir, envVar)
				} else {
					fmt.Printf(envStr, yogaconfig.EnvConfigDir, yogaconfig.GetConfigDir())
				}

				return nil
			},
		},
		{
			Name:    "quit",
			Aliases: []string{"exit"},
			Usage:   "退出程序",
			Action: func(c *cli.Context) error {
				return cli.NewExitError("", 0)
			},
			Hidden:   true,
			HideHelp: true,
		},
	}


	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
