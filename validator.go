//
// Copyright 2016 Dairon Medina <http://github.com/codeadict>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	LEN_INDIVIDUAL_ID = 10
	PROVINCES = 24 // Number of provinces, actually 24 in Ecuador
)

// Error messages
var (
	ErrOnlyDigits = errors.New("Identification can only contain digits")
	ErrInvalidFormat = errors.New("Identification format is incorrect")
	ErrInvalidProvince = errors.New("Province is invalid")
	ErrInvalidThirdDigit = errors.New("Third digit is invalid")
	ErrInvalidCheckdigit = errors.New("Check digit is incorrect")
)

type IdNo struct {
	Valid  bool              `json:"valid"`
	Source string            `json:"source"`
	Data   map[string]string `json:"data"`
}

/**
 * 00 0 000000 0  (10 digits)
 * |  |        |
 * |  |        `- Check digit
 * |  |
 * |  |
 * |  `- Third Digit (Must be < 6)
 * `- Province (Must be between 1 and PROVINCES)
 */
func Validate(id string) (*IdNo, error){
	n := &IdNo{Source: id}
	n.Data = make(map[string]string)

	if _, e := strconv.Atoi(n.Source); e != nil || len(n.Source) != LEN_INDIVIDUAL_ID {
		return n, ErrInvalidFormat
	}

	tt := strings.Split(n.Source, "")
	source := make([]int, len(tt))
	for i := range tt {
		source[i], _ = strconv.Atoi(tt[i])
	}

	// Extract province, third digit and check digit
	province := source[0]*10 + source[1]
	third_digit := source[2]
	check_digit := source[9]
	var coeficients = [][]int{
		{2, 1, 2, 1, 2, 1, 2, 1, 2},
	}

	if (province > 0 && province <= PROVINCES){
		if third_digit > 6 {
			return n, ErrInvalidThirdDigit
		} else {
			// Control checkdigits
			for _, coeficient := range coeficients {
				sum := 0
				for i, c := range coeficient {
					sum += source[i] * c
				}
				verify := LEN_INDIVIDUAL_ID - (sum/LEN_INDIVIDUAL_ID)
				if verify == LEN_INDIVIDUAL_ID {
					verify = 0
				}
				if check_digit != verify {
					return n, ErrInvalidCheckdigit
				}
		}
	}
	} else {
		return n, ErrInvalidProvince
	}

	// Everything looks OK, go ahead and return value
	n.Valid = true
	n.Data["province"] = fmt.Sprintf("%02s", strconv.Itoa(province))
	return n, nil
}
