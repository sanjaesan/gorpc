package main

import (
	"context"
	"flag"
	"fmt"
	pb "gorpc/api"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listento")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}
	srv := grpc.NewServer()
	pb.RegisterTextToSpeechServer(srv, server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("Could not serve:  %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("Could not create temp file")
	}
	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("Could not close %s: %v", file.Name(), err)
	}
	cmd := exec.Command("flite", "-t", text.Text, "o", "output.wav")
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("Could not read temp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}