package main

import (
    "fmt"
    "net"
)

const (
    host = ""
    port = "8080"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    fmt.Println("New connection from", conn.RemoteAddr())

    // Echo back whatever is sent
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            if err != io.EOF {
                fmt.Println("Error reading:", err)
            }
            return
        }
        
        fmt.Println("Received:", string(buf[:n]))
        conn.Write(buf[:n])
    }
}

func main() {
    addr, err := net.ResolveTCPAddr("tcp", host+":"+port)
    if err != nil {
        fmt.Println("Error resolving address:", err)
        return
    }

    listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on", host, ":", port)

    for {
        conn, err := listener.AcceptTCP()
        if err != nil {
            fmt.Println("Error accepting:", err)
            continue
        }

        go handleConnection(conn)
    }
}