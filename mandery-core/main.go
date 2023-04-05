package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xomander/mandery/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
