package test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	formatter := new(prefixed.TextFormatter)
	logrus.SetFormatter(formatter)
}

func TestURI(t *testing.T) {
	str := "tcp://0.0.0.0:1234"
	uri, err := url.Parse(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uri.Scheme)
	fmt.Println(uri.Host)
	fmt.Println(uri.Hostname())
	fmt.Println(uri.Port())
}
