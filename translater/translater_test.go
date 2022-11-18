package translater

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func setKey() {

	buf, err := ioutil.ReadFile("../app.yaml")
	if err != nil {
		panic(fmt.Sprintf("cannot read config file %v", err))
	}

	scanner := bufio.NewScanner(strings.NewReader(string(buf)))
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		if strings.HasPrefix(line, "naverClientId") {
			buf := strings.Split(line, ":")
			NaverClientId = strings.Trim(buf[1], " ")
		} else if strings.HasPrefix(line, "naverClientSecret") {
			buf := strings.Split(line, ":")
			NaverClientSecret = strings.Trim(buf[1], " ")
		}
	}
}

func TestKo2En(t *testing.T) {

	setKey()

	r, err := Ko2En("안녕하세요")
	if r != "Hello" {
		if err != nil {
			t.Errorf("Wrong result "+r+" %v", err)
		} else {
			t.Error("Wrong result " + r)
		}
	}
}

func TestEn2Ko(t *testing.T) {

	setKey()

	r, err := En2Ko("Hello")
	if r != "안녕하세요." {
		if err != nil {
			t.Errorf("Wrong result "+r+" %v", err)
		} else {
			t.Error("Wrong result " + r)
		}
	}
}
