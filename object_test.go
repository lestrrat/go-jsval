package jsval

import (
	"strings"
	"testing"

	"github.com/lestrrat/go-jsschema"
	"github.com/stretchr/testify/assert"
)

func TestObject(t *testing.T) {
	const src = `{
  "type": "object",
  "additionalProperties": false,
  "properties": {
    "name": {
      "type": "string",
      "maxLength": 20,
      "pattern": "^[a-z ]+$"
    },
	  "age": {
		  "type": "integer",
	    "minimum": 0
	  },
	  "tags": {
      "type": "array",
	    "items": {
        "type": "string"
      }
    }
  }
}`

	s, err := schema.Read(strings.NewReader(src))
	if !assert.NoError(t, err, "reading schema should succeed") {
		return
	}

	v := New()
	if !assert.NoError(t, v.Build(s), "Validator.Build should succeed") {
		return
	}

	data := []interface{}{
		map[string]interface{}{"Name": "World"},
		map[string]interface{}{"name": "World"},
		map[string]interface{}{"name": "wooooooooooooooooooooooooooooooorld"},
		map[string]interface{}{
			"tags": []interface{}{ 1, "foo", false },
		},
	}
	for _, input := range data {
		t.Logf("Testing %#v (should FAIL)", input)
		if !assert.Error(t, v.Validate(input), "validation fails") {
			return
		}
	}

	data = []interface{}{
		map[string]interface{}{"name": "world"},
		map[string]interface{}{"tags": []interface{}{"foo", "bar", "baz"}},
	}
	for _, input := range data {
		t.Logf("Testing %#v (should PASS)", input)
		if !assert.NoError(t, v.Validate(input), "validation passes") {
			return
		}
	}
}
