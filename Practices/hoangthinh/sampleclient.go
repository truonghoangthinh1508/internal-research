package main

import (
    "internal-research/Practices/hoangthinh/connection"
)

func main() {
    // connection.Listen ("localhost:8000")
    c := connection.Connect ("localhost:8")
    connection.SendMsg(c, "hello")
}
