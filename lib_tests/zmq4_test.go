package lib_tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-zeromq/zmq4"
)

func zmq4Req() {
	logger := log.New(log.Writer(), "rrclient: ", log.LstdFlags)

	req := zmq4.NewReq(context.Background())
	defer req.Close()

	err := req.Dial("tcp://localhost:5559")
	if err != nil {
		logger.Fatalf("could not dial: %v", err)
	}

	for i := 0; i < 10; i++ {
		err := req.Send(zmq4.NewMsgString("Hello"))
		if err != nil {
			logger.Fatalf("could not send greeting: %v", err)
		}

		msg, err := req.Recv()
		if err != nil {
			logger.Fatalf("could not recv greeting: %v", err)
		}
		logger.Printf("received reply %d [%s]\n", i, msg.Frames[0])
	}
}

func TestZMQ4Req(t *testing.T) {
	zmq4Req()
}

func zmq4Rep() {
	logger := log.New(log.Writer(), "rrworker: ", log.LstdFlags)

	//  Socket to talk to clients
	rep := zmq4.NewRep(context.Background())
	defer rep.Close()

	err := rep.Listen("tcp://*:5559")
	if err != nil {
		logger.Fatalf("could not dial: %v", err)
	}

	for {
		//  Wait for next request from client
		msg, err := rep.Recv()
		if err != nil {
			logger.Fatalf("could not recv request: %v", err)
		}

		logger.Printf("received request: [%s]\n", msg.Frames[0])

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		err = rep.Send(zmq4.NewMsgString("World"))
		if err != nil {
			logger.Fatalf("could not send reply: %v", err)
		}
	}
}

func TestZMQ4Rep(t *testing.T) {
	zmq4Rep()
}

func TestZMQ4ReqRep(t *testing.T){
	go zmq4Rep()
	go zmq4Req()

	time.Sleep(15 * time.Second)
}
