package main

import (
    "internal-research/Practices/phuocthanh/connection"
)

func main() {
    // connection.Listen ("localhost:8686")
    c := connection.Connect ("localhost:8686")
    connection.SendMsg(c, "hello")
}
