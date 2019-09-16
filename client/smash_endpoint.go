package client

import (
    "fmt"
    "strconv"
    "strings"
)

type endpoint struct {
    Host string
    Port int
    User string
}

func NewEndpoint(s *string) *endpoint {
    e := endpoint{
        Host: *s,
    }

    if parts := strings.Split(e.Host, "@"); len(parts) > 1 {
        e.User = parts[0]
        e.Host = parts[1]
    }

    if parts := strings.Split(e.Host, ":"); len(parts) > 1 {
        e.Host = parts[0]
        e.Port, _ = strconv.Atoi(parts[1])
    }

    return &e
}

func (e *endpoint) HostString() *string {
     s := e.Host
     if e.Port != 0 {
         s = fmt.Sprintf("%s:%d", s, e.Port)
     }
     return &s
}
