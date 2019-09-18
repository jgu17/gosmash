package client

type Client interface {
    Command(cmd Request) (string, error)
    Commands(cmds []Request) (string, error)
}

type Request struct {
    Command string
    Args []string
}
