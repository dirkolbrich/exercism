package queenattack

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

var (
	// ErrInvalid is returned, if an invalid position is found
	ErrInvalid = errors.New("Position in invalid.")
	// ErrEqual is returned, if two figures share the same position
	ErrEqual = errors.New("Positions are equal.")
)

type position struct {
	file, rank int
	ok         bool
}

func isInvalid(p position) bool {
	return !(p.ok)
}

func isEqual(p1, p2 position) bool {
	return p1 == p2
}

func parse(s string) (p position) {
	// set regex
	var validAN = regexp.MustCompile("^[a-h][1-8]$")

	if validAN.MatchString(s) {
		rank, err := strconv.Atoi(s[1:])
		if err == nil && rank < 8 {
			p.file = int(s[0]) - int('a')
			p.rank = rank
			p.ok = true
		}
	}
	return p
}

func horizontal(w, b position) bool {
	return w.file == b.file
}

func vertical(w, b position) bool {
	return w.rank == b.rank
}

func diagonal(w, b position) bool {
	// if distance vertical equals distance horizontal
	return math.Abs(float64(w.file-b.file)) == math.Abs(float64(w.rank-b.rank))
}

func canAttack(w, b position) bool {
	return horizontal(w, b) || vertical(w, b) || diagonal(w, b)
}

// CanQueenAttack indicates whether or not two queens are positioned so that they can attack each other.
func CanQueenAttack(w, b string) (attack bool, error error) {
	white, black := parse(w), parse(b)
	//fmt.Println(white, black)
	if isInvalid(white) || isInvalid(black) {
		return false, ErrInvalid
	}
	if isEqual(white, black) {
		return false, ErrEqual
	}
	return canAttack(white, black), nil
}
