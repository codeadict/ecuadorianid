package validator

import(
  "testing"
)


func TestValid(t *testing.T) {
	tests := []string{"1104680135"}

	for _, tt := range tests {
		if p, err := Validate(tt); err != nil {
			t.Errorf("Fail at '%v', out '%+v', err: %+v", tt, p, err)
		}
	}
}

func TestInvalid(t *testing.T) {
	tests := []string{"1105680134"}

	for _, tt := range tests {
		if p, err := Validate(tt); err == nil {
			t.Errorf("Fail at '%v', out '%+v', should have produced error", tt, p)
		}
	}
}
