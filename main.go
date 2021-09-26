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

	cli, err := client.NewConnection(SvcId, args.Host, args.Port)
	if err != nil {
		panic(err)
	}
	defer cli.CloseConnection()

	resp, err := cli.VerifyToken(args.Token, args.Scope)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ToString())
}
