package queenattack

import (
	"errors"
	"strconv"
	"strings"
)

// ValidLetterPos ...
var ValidLetterPos = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
}

// CanQueenAttack ...
func CanQueenAttack(whitePos string, blackPos string) (bool, error) {

	//check if pos format is invalid
	if len(whitePos) != 2 || len(blackPos) != 2 {
		return false, errors.New("error not valid pos")
	}

	//check if white column is in range a-g
	w := ValidLetterPos[whitePos[0:1]]
	if w == 0 {
		return false, errors.New("error white out of board")
	}

	//check if black column is in range a-g
	b := ValidLetterPos[blackPos[0:1]]
	if b == 0 {
		return false, errors.New("error black out of valid pos")
	}

	wp, err := strconv.Atoi(whitePos[1:2])
	if err != nil {
		return false, err
	}

	bp, err := strconv.Atoi(blackPos[1:2])
	if err != nil {
		return false, err
	}

	//check if white out of row
	if wp < 1 || wp > 8 {
		return false, errors.New("error white out of board")
	}

	//check if black out of row
	if bp < 1 || bp > 8 {
		return false, errors.New("error black out of board")
	}

	//check if in the same position
	if strings.EqualFold(whitePos, blackPos) {
		return false, errors.New("error same position")
	}

	//if in the same column but different row
	if ValidLetterPos[whitePos[:1]] == ValidLetterPos[blackPos[:1]] {
		return true, nil
	}

	//if in the same row
	if wp == bp {
		return true, nil
	}

	//if diagonal
	if (w-b)*(w-b) == (wp-bp)*(wp-bp) {
		return true, nil
	}

	//otherwise
	return false, nil
}
