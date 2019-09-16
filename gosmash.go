package main

import (
    "flag"
    "fmt"
    "golang.org/x/crypto/ssh"
    "gosmash/client"
)

//An example CLI program demonstrating the usage of GoSMASH client.

func main() {
    host := flag.String("h", "", "host string in the form of user@host:port")
    key := flag.String("i", "", "path to the private key")
    password := flag.String("p", "", "password")
    boot := flag.String("boot", "once", "boot options, either always or once")
    flag.Parse()
    fmt.Println("password:", *password, "\nkey:", *key)    

    endpoint := client.NewEndpoint(host)

    var auth ssh.AuthMethod
    if *password != "" {
        auth = client.PasswordAuth(password)
    } else if *key != "" {
        auth = client.KeyAuth(key)
    } else {
        panic("missing SSH key or password in the arguemnts")
    }
    c, ce := client.NewClient(*endpoint, auth)
    if ce != nil {
        panic(ce)
    }

    args := flag.Args()
    fmt.Println(args)

    if len(args) == 0 {
        fmt.Println("Missing command name")
        return
    }

    var (
        res string
        err error
    )

    switch cmd := args[0]; cmd {
    case "insert_cdrom":
        if len(args) < 2 {
            fmt.Println("Usage: -b=[once|always] insert_cdrom ISO_url")
            return
        }
        url := args[1]
        if *boot == "once" {
            res, err = c.InsertCDRomImageSingleBoot(url)
        } else {             
            res, err = c.InsertCDRomImage(url)
        }
    case "eject_cdrom":
        res, err = c.EjectCDRomImage()
    case "insert_usb":
        if len(args) < 2 {
            fmt.Println("Usage: -b=[once|always] insert_usb ISO_url")
            return
        }
        url := args[2]
        if *boot == "once" {
            res, err = c.InsertUSBImageSingleBoot(url)
        } else {
            res, err = c.InsertUSBImage(url)
        }
    case "eject_usb":
        res, err = c.EjectUSBImage()
    case "start":
        res, err = c.StartServer()
    case "stop":
        res, err = c.StopServer()
    case "reset_hard":
        res, err = c.ResetServerHard()
    case "reset_soft":
        res, err = c.ResetServerSoft()
    default:
        fmt.Println("Unrecognized command:", cmd)
        return
    }
    if err != nil {
        fmt.Println("Command execution error: " + err.Error())
    }
    fmt.Println("Command completed:\n", res)
}

