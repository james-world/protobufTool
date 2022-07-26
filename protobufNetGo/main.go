package main

import (
	"io/ioutil"
	"log"
	pb "protobufNetGo/imburse_protobuf"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func convertStringToGuid(s string) *pb.Guid {
	a, _ := strconv.ParseUint(s[:16], 16, 64)
	b, _ := strconv.ParseUint(s[16:], 16, 64)
	guid := &pb.Guid{
		Lo: a,
		Hi: b,
	}
	return guid
}

func main() {

	guid := convertStringToGuid("a8f9f2f836f340b1871a70f92feb4c44")

	msg := &pb.ScheduleInstructionStep3V2{
		CorrelationId: "123456",
		RequestId:     guid,
	}

	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Failed to encode msg:", err)
	}
	if err := ioutil.WriteFile("message_go.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write file:", err)
	}

}
