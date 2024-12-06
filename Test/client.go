package main

import (
    "fmt"
    "net"
    "bufio"
    "log"
    "os"
}

const (
    host = ""
    port = "8080"
)

func handleConnectionToServer(conn net.Conn) {
    defer conn.Close()

    // Echo back whatever is sent
    reader := bufio.NewReader(os.Stdin)
    buf := make([]byte, 1024)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from termina!")
            return
        }
                buf = []byte(str)
        conn.Write(buf[:n])

        n, err := conn.Read(buf)
        if err != nil {
            if err != io.EOF {
                fmt.Println("Error reading:", err)
            }
            return
        }
        fmt.Println("Received:", string(buf[:n]))
    }
}

func main() {
    conn, err := net.Dial("tcp", "localhost:80")    

    if err != nil: {
        fmt.Println("Error connected to server: ", err)
        return
    }

    fmt.Println("Successful connected to server", host, ":", port)

    go handleConnectionToServer(conn)
}