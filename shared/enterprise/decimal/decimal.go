package decimal

import (
	"fmt"
	"strconv"
	"strings"
)

type Decimal struct {
	Amount int64
	Scale  int32
}

func StringToDecimal(str string) (Decimal, error) {
	str = strings.Trim(strings.TrimRight(str, "."), " ")
	pieces := strings.Split(str, ".")

	if len(pieces) == 0 || len(pieces) > 2 {
		/*
			Input: "", "  ", "11.12.367"
			Expected Output: error
		*/

		return Decimal{}, fmt.Errorf("input string '%s' format is incorrect", str)
	}

	if str[0] == '.' {
		/*
			Input: ".1234567"
			Expected Output: {
			amount: 1234567
			scale: 7
			}
		*/

		str = str[1:]
		amount, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return Decimal{}, err
		}

		return Decimal{
			Amount: amount,
			Scale:  int32(len(str)),
		}, nil
	} else if len(pieces) == 1 {
		/*
			Input: "1234567"
			Expected Output: {
			amount: 1234567
			scale: 0
			}
		*/

		amount, err := strconv.ParseInt(pieces[0], 10, 64)
		if err != nil {
			return Decimal{}, err
		}

		return Decimal{
			Amount: amount,
			Scale:  0,
		}, nil
	} else {
		/*
			Input: "123.4567"
			Expected Output: {
			amount: 1234567
			scale: 4
			}
		*/

		amount, err := strconv.ParseInt(pieces[0]+pieces[1], 10, 64)
		if err != nil {
			return Decimal{}, err
		}

		return Decimal{
			Amount: amount,
			Scale:  int32(len(pieces[1])),
		}, nil
	}
}
