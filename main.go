package main

import (
	"cube_cli/client"
	"cube_cli/utils"
	"fmt"
	"os"
)

const SvcId = 0x00000002

func main() {
	args, err := utils.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		fmt.Printf(utils.HelpMsgFmt, os.Args[0])
		return
	}

	cli, err := client.NewConnection(SvcId, args.Host, args.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.CloseConnection()

	resp, err := cli.VerifyToken(args.Token, args.Scope)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.ToString())
}
