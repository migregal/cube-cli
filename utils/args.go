package utils

import "errors"

const HelpMsgFmt = "\nUsage: %s <host> <port> <token> <scope>\n\n"

type Args struct {
	Host  string
	Port  string
	Token string
	Scope string
}

func ParseArgs(args []string) (*Args, error) {
	if len(args) < 5 {
		return nil, errors.New("missing args detected")
	}

	if args[1] == "" {
		return nil, errors.New("missing host value")
	}

	if args[2] == "" {
		return nil, errors.New("missing port value")
	}

	if args[3] == "" {
		return nil, errors.New("missing token value")
	}

	res := Args{
		Host:  args[1],
		Port:  args[2],
		Token: args[3],
		Scope: args[4],
	}

	return &res, nil
}
