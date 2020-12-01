package main

import (
	"context"
	"flag"
	pb "gorpc/api"
	"io/ioutil"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("b", "localhost:5050", "address to the Server")
	output := flag.String("o", "output.wav", "wav file where output will be written")

	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to %s: %v", *address, err)

	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	request := &pb.Text{Text: "Hello, There .. it's me"}
	res, err := client.Say(context.Background(), request)

	if err != nil {
		log.Fatalf("Could not say %s: %v", request.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("Could not write to %s: %v", *output, err)
	}
}
