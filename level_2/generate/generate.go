package generate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BuildNIP interface {
	textNIP() string
}

type FirstAttr struct {
	gender      string
	year, month int
}

const (
	awalNIP   string = "AR"
	semester1        = "1"
	semester2        = "2"
)

func (attr FirstAttr) textNIP() string {
	yearNIP := strconv.Itoa(attr.year % 1e2)
	genderNIP := strings.ToUpper(string(attr.gender[5]))
	semesterNIP := semester1
	if attr.month > 6 && attr.month <= 12 {
		semesterNIP = semester2
	}
	return awalNIP + genderNIP + yearNIP + semesterNIP
}

func GenerateNIP(gender string, year, month, loop int) ([]string, error) {
	if gender == "" {
		return []string{}, errors.New("empty gender")
	}

	if loop == 0 {
		return []string{}, errors.New("empty loop")
	}

	var b BuildNIP
	b = FirstAttr{gender, year, month}
	f := b.textNIP()
	generateArr := []string{}
	j := 1
	for i := 0; i < loop; i++ {
		s := fmt.Sprintf("%05d", j)
		generateArr = append(generateArr, f+"-"+s)
		j++
	}
	return generateArr, nil
}

func GenerateNextNIP(nip string, loop int) ([]string, error) {
	if nip == "" {
		return []string{}, errors.New("empty NIP")
	}

	if loop == 0 {
		return []string{}, errors.New("empty loop")
	}

	rangeNIP := nip[0:6]
	rangeSeq, _ := strconv.Atoi(nip[8:12])
	generateArr := []string{}
	j := rangeSeq
	for i := 0; i < loop; i++ {
		s := fmt.Sprintf("%05d", j)
		generateArr = append(generateArr, rangeNIP+"-"+s)
		j++
	}
	return generateArr, nil
}
