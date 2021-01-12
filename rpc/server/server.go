package main

import (
	"log"
	"net"
	"net/rpc"
)

type MathService struct {
}

type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	rpc.RegisterName("MathService", new(MathService))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)

}
