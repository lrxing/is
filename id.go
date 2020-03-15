package is

import (
	"errors"
	"math"
	"strconv"
)

//ID 判断身份证是否合法
func ID(ID string) (bool, error) {
    var (
        sum      int
        checkVal int
    )

    if len(ID) != 18 {
        return false, errors.New("ID illegal")
    }
    if ID[17] == 'X' || ID[17] == 'x' {
        checkVal = 10
    } else {
        if !isNumber(ID[17]) {
            return false, errors.New("check value illegal")
        }
        t, err := strconv.Atoi(string(ID[17]))
        if err != nil {
            return false, err
        }
        checkVal = t
    }
    for i := 0; i < 17; i++ {
        if !isNumber(ID[i]) {
            return false, errors.New("ID illegal")
        }
        if id, err := strconv.Atoi(string(ID[i])); err == nil {
            sum += id * (int(math.Pow(2, float64(18-i-1))) % 11)
        }
    }
    checkVal1 := (12 - sum%11) % 11
    if checkVal == checkVal1 {
        return true, nil
    }
    return false, errors.New("Error")
}

func isNumber(b byte) bool {
    if b > '9' || b < '0' {
        return false
    }
    return true
}
