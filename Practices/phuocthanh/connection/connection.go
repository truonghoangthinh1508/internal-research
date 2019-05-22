package connection 

import ("net"; "fmt"; "encoding/gob")

func Connect(addr string) net.Conn {
    c,e := net.Dial ("tcp", addr)
    fmt.Println("trying to connect to remote server at " + addr )
    if e != nil {fmt.Println(e)}
    return c
}

func SendMsg(conn net.Conn, msg string) {
    gob.NewEncoder(conn).Encode(msg)
}

func Listen(sock string) {
    bound_sock, err :=  net.Listen ("tcp", sock) 
    if err != nil {fmt.Println(err)}
    defer bound_sock.Close()

    fmt.Println ("server is listening at " + sock )
    for {
        conn, err := bound_sock.Accept()
        if err != nil {fmt.Print(err)}
        go HandleConn(conn)
        }
}

func HandleConn(conn net.Conn) {
    var msg string
    gob.NewDecoder(conn).Decode(&msg)
    fmt.Println("Received: " + msg)
}


