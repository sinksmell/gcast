package gcast

import (
	"encoding/json"
	"testing"
)

type T struct {
	A string `json:"a"`
	B string `json:"b"`
	C struct {
		D []string `json:"d"`
	} `json:"c"`
}

func TestDecodeS2S(t *testing.T) {
	var src, dst T
	src.A = "test"
	src.B = "test"
	src.C = struct {
		D []string `json:"d"`
	}{D: []string{"test"}}

	Decode(src, &dst)
	t.Logf("%+v\n", dst)
}

func BenchmarkDecode(b *testing.B) {
	var src T
	src.A = "test"
	src.B = "test"
	src.C = struct {
		D []string `json:"d"`
	}{D: []string{"test"}}
	for i := 0; i < b.N; i++ {
		var p T
		Decode(src, &p)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	var src T
	src.A = "test"
	src.B = "test"
	src.C = struct {
		D []string `json:"d"`
	}{D: []string{"test"}}
	var str, _ = json.Marshal(src)
	for i := 0; i < b.N; i++ {
		var p T
		json.Unmarshal(str, &p)
	}
}

func BenchmarkAssign(b *testing.B) {
	var src T
	src.A = "test"
	src.B = "test"
	src.C = struct {
		D []string `json:"d"`
	}{D: []string{"test"}}
	for i := 0; i < b.N; i++ {
		var p T
		p.A = src.A
		p.B = src.B
		p.C = struct {
			D []string `json:"d"`
		}{}
		p.C.D = make([]string, 0)
		copy(p.C.D, src.C.D)
	}
}
