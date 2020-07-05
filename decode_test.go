package gcast

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type T struct {
	A string `json:"a"`
	B string `json:"b"`
	C struct {
		D []string `json:"d"`
	} `json:"c"`
}

func Test_Decode(t *testing.T) {
	var src = map[string]interface{}{
		"a": "1",
		"b": 2,
		"c": map[string]interface{}{
			"d": []string{"1", "2", "3"},
		},
	}

	var p T

	err := Decode(src, &p)
	assert.Nil(t, err)
	fmt.Printf("%+v\n", p)
}

func BenchmarkDecode(b *testing.B) {
	var src = map[string]interface{}{
		"a": "1",
		"b": 2,
		"c": map[string]interface{}{
			"d": []string{"1", "2", "3"},
		},
	}
	for i := 0; i < b.N; i++ {
		var p T
		Decode(src, &p)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	var src = map[string]interface{}{
		"a": "1",
		"b": 2,
		"c": map[string]interface{}{
			"d": []string{"1", "2", "3"},
		},
	}
	var str, _ = json.Marshal(src)
	for i := 0; i < b.N; i++ {
		var p T
		json.Unmarshal(str, &p)
	}
}

func BenchmarkAssign(b *testing.B) {
	var src = map[string]interface{}{
		"a": "1",
		"b": 2,
		"c": map[string]interface{}{
			"d": []string{"1", "2", "3"},
		},
	}
	for i := 0; i < b.N; i++ {
		var p T
		p.A = src["a"].(string)
		p.B = fmt.Sprintf("%v", src["b"])
		v := src["c"].(map[string]interface{})
		p.C.D = v["d"].([]string)
	}
}

func TestAssign(t *testing.T) {
	var src = map[string]interface{}{
		"a": "1",
		"b": 2,
		"c": map[string]interface{}{
			"d": []string{"1", "2", "3"},
		},
	}

	var p T
	//p.A = src["a"].(string)
	//p.B = fmt.Sprintf("%v", src["b"])
	//v := src["c"].(map[string]interface{})
	//p.C.D = v["d"].([]string)
	Decode(src, &p)

	t.Logf("%+v\n", p)
}
