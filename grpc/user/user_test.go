package user

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
)

var user = &User{
	Id:       1,
	Name:     "miguel",
	Email:    "miguel@nakva.mx",
	Twitter:  "miguellgt",
	Street:   "matamoros",
	City:     "SLP",
	State:    "SLP",
	Zip:      "78100",
	Phone:    "934893843",
	Metadata: map[string]string{"groups": "foo"},
}

func BenchmarkUserProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(user)
		if err != nil {
			b.Fatalf("marshaling err: %v", err)
		}
	}
}

func BenchmarkUserJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(user)
		if err != nil {
			b.Fatalf("marshaling err: %v", err)
		}
	}
}

func BenchmarkUserProto3Unmarshal(b *testing.B) {
	data, err := proto.Marshal(user)
	if err != nil {
		b.Fatalf("unmarshaling err: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var user User
		err := proto.Unmarshal(data, &user)
		if err != nil {
			b.Fatalf("unmarshaling err: %v", err)
		}
	}
}

func BenchmarkUserJSONUnmarshal(b *testing.B) {
	data, err := json.Marshal(user)
	if err != nil {
		b.Fatalf("unmarshaling err: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var user User
		err := json.Unmarshal(data, &user)
		if err != nil {
			b.Fatalf("unmarshaling err: %v", err)
		}
	}
}
