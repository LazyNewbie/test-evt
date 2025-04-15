package adevent

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
)

type AdEventEncoder struct {
}

func (e AdEventEncoder) UnmarshalJson(data []byte, event *AdEvent) error {
	return protojson.Unmarshal(data, event)
}

func (e AdEventEncoder) MarshalJson(event *AdEvent) ([]byte, error) {
	return protojson.Marshal(event)
}

func (e AdEventEncoder) UnmarshalBin(data []byte, event *AdEvent) error {
	return prototext.Unmarshal(data, event)
}

func (e AdEventEncoder) MarshalBin(event *AdEvent) ([]byte, error) {
	return prototext.Marshal(event)
}
