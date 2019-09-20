package client

import (
    "log"
    "strconv"
    "strings"
)

type Client interface {
    Command(cmd Request) (*Response, error)
    Commands(cmds []Request) ([]Response, error)
}

type Request struct {
    Command string
    Args []string
}

type Response struct {
    Command string
    Status int
    StatusTag string
    ErrorTag string
    Details string
}

func HasError(rs []Response) bool {
    for _, r := range rs {
        if r.Status != 0 {
            return true
        }
    }
    return false
}

func PrintResponse(rs []Response) string {
    out := []string{}
    for _, r := range rs {
        out = append(out, r.String())
    }
    return strings.Join(out, "\n")
}

func NewResponse(s string) *Response {
    lines := strings.Split(strings.Trim(s, "\r\n"), "\r\n")
    code, err := strconv.Atoi(strings.Split(lines[1], "=")[1])
    if err != nil {
        log.Printf("invalid status code in response: %s", lines[1])
        code = -1
    }
    r := Response{
        Command: lines[0],
        Status: code, 
        StatusTag: strings.Split(lines[2], "=")[1],
    }
    if len(lines) > 3 {
        if strings.HasPrefix(lines[3], "error_tag") {
            r.ErrorTag = strings.Split(lines[3], "=")[1]
            if len(lines) > 4 {
                r.Details = strings.Join(lines[4:], "\n")
            }
        } else {
            r.Details = strings.Join(lines[3:], "\n")
        }
    }
    return &r
}

func (r *Response) String() string {
    s := []string {
        r.Command,
        "status=" + strconv.Itoa(r.Status),
        "status_tag=" + r.StatusTag,
    }
    if r.ErrorTag != "" {
        s = append(s, "error_tag=" + r.ErrorTag)
    }
    s = append(s, r.Details)
    return strings.Join(s, "\n")
}
