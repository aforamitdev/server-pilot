package main

import (
	"context"
	"fmt"
	"net"
	"os"
)

func main() {

	if err := serve(context.Background(), ":5000"); err != nil {
		os.Exit(1)
	}
}

// serve is capable of answering to a single client at a time
func serve(ctx context.Context, address string) error {

	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		fmt.Printf("failed to UDP listen on '%s' with '%v'", address, err)
		return err
	}
	defer func() {
		if err := pc.Close(); err != nil {
			fmt.Printf("failed to close packet connection with '%v'", err)
		}
	}()

	doneChan := make(chan error, 1)
	// maxBufferSize specifies the size of the buffers that
	// are used to temporarily hold data from the UDP packets
	// that we receive.

	buffer := make([]byte, 1024)

	go func() {
		for {
			n, addr, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				// return err
			}
			fmt.Printf("address: '%+v' bytes: '%d' request: '%s'", addr, n, string(buffer))
		}
	}()

	var ret error
	select {
	case <-ctx.Done():
		ret = ctx.Err()
		fmt.Printf("cancelled with '%v'", err)
	case ret = <-doneChan:
	}

	return ret
}
