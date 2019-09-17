package client

type Client interface {
    Command(cmd string) (string, error)
    Commands(cmds []string) (string, error)
}
