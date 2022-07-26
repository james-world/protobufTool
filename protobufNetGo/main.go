package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	pb "protobufNetGo/imburse_protobuf"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func convertDateToCSDateTime(date time.Time) *pb.DateTime {
	datetime := &pb.DateTime{
		Value: date.Unix() / 60 / 60 / 24,
		Scale: 0,
		Kind:  0,
	}
	return datetime
}

func convertUintToDecimal(value uint64) *pb.Decimal {
	decimal := &pb.Decimal{
		Lo:        value,
		Hi:        0,
		SignScale: 0,
	}
	return decimal
}

func convertStringToGuid(guidStr string) *pb.Guid {
	s := strings.Replace(guidStr, "-", "", -1)
	lo := s[12:16] + s[8:12] + s[0:8]
	hi := s[30:] + s[28:30] + s[26:28] + s[24:26] + s[22:24] + s[20:22] + s[18:20] + s[16:18]
	a, _ := strconv.ParseUint(lo, 16, 64)
	b, _ := strconv.ParseUint(hi, 16, 64)
	guid := &pb.Guid{
		Lo: a,
		Hi: b,
	}
	return guid
}

func main() {

	msg := &pb.ScheduleInstructionStep3V2{
		CorrelationId:               "123456",
		RequestId:                   convertStringToGuid("30f6432c-ae25-4fbf-bd80-3b537e8ef0f2"),
		TenantId:                    convertStringToGuid("a70d4280-7117-4fcc-a7a1-a185e465fddc"),
		CustomerRef:                 "MyCustomerRef",
		FinancialInstrumentId:       convertStringToGuid("a8f9f2f8-36f3-40b1-871a-70f92feb4c44"),
		WorkflowId:                  "MyWorkflowId",
		OrderRef:                    "MyOrderRef",
		InstructionRef:              "MyInstructionRef",
		Country:                     "UK",
		Currency:                    "GBP",
		Amount:                      convertUintToDecimal(10),
		RequestedSettlementDate:     convertDateToCSDateTime(time.Date(2023, time.August, 1, 0, 0, 0, 0, time.UTC)),
		ForecastedSettlementDate:    convertDateToCSDateTime(time.Date(2023, time.August, 3, 0, 0, 0, 0, time.UTC)),
		ScheduledExecutionTimestamp: convertDateToCSDateTime(time.Date(2023, time.July, 28, 0, 0, 0, 0, time.UTC)),
		AppId:                       "MyAppId",
		AppQueue:                    "Braintree",
		PaymentMethod:               "CreditCard",
	}

	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Failed to encode msg:", err)
	}
	if err := ioutil.WriteFile("message_go.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write file:", err)
	}
	b64 := base64.StdEncoding.EncodeToString(out)
	if err := ioutil.WriteFile("message_go.txt", []byte(b64), 0644); err != nil {
		log.Fatalln("Failed to write base64 file:", err)
	}

}
