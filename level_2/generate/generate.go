package generate
import (
	"strconv"
	"errors"
	"strings"
)

type BuildNIP interface {
	textNIP() string
}

type FirstAttr struct {
    gender string
	year, month int
}

type SecondAttr struct {
    firstStr string
	seq int
}

const (
	digit1 string = "0"
	digit2        = "00"
	digit3        = "000"
	digit4        = "0000"
	digit5        = "00000"
)

func (attr FirstAttr) textNIP() string {
	const awalNIP = "AR"
	yearNIP := strconv.Itoa(attr.year % 1e2)
	genderNIP := strings.ToUpper(string(attr.gender[5]))
	semesterNIP := "1"
	if attr.month > 6 && attr.month <= 12 {
		semesterNIP = "2"
	}
    return awalNIP + genderNIP + yearNIP + semesterNIP
}

func (attr SecondAttr) textNIP() string {
	seqText := digit5
	if attr.seq > 0 && attr.seq <= 9 {
		seqText = digit4 + strconv.Itoa(attr.seq)
	} else if attr.seq >= 10 && attr.seq <= 99 {
		seqText = digit3 + strconv.Itoa(attr.seq)
	} else if attr.seq >= 100 && attr.seq <= 999 {
		seqText = digit2 + strconv.Itoa(attr.seq)
	} else if attr.seq >= 1000 && attr.seq <= 9999 {
		seqText = digit1 + strconv.Itoa(attr.seq)
	} else if attr.seq >= 10000 && attr.seq <= 99999 {
		seqText = strconv.Itoa(attr.seq)
	}
	return seqText
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
		b = SecondAttr{f, j}
		s := b.textNIP()
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

	var b BuildNIP
	rangeNIP := nip[0:6]
	rangeSeq,_ := strconv.Atoi(nip[8:12])
	generateArr := []string{}
	j := rangeSeq
	for i := 0; i < loop; i++ {
		b = SecondAttr{rangeNIP, j}
		s := b.textNIP()
		generateArr = append(generateArr, rangeNIP+"-"+s)
		j++
	}
    return generateArr, nil
}