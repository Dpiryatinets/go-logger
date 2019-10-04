package main

import (
	"bytes"
	"testing"
)

func TestDebugOutput(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	defer func() { stdOut = out }()

	logger := GetLogger("debug")
	logger.Debug([]int{1, 2, 3})

	testExpectation := "DEBUG - [1 2 3]"
	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("DEBUG log incorrect value", testResult)
	}
}

func TestDebugLevel(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	defer func() { stdOut = out }()

	testExpectation := ""
	logger := GetLogger("error")
	logger.Debug([]int{1, 2, 3})

	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("DEBUG log incorrect value", testResult)
	}
}

func TestInfoOutput(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	defer func() { stdOut = out }()

	logger := GetLogger("info")
	logger.Info([]int{1, 2, 3})

	testExpectation := "INFO - [1 2 3]"
	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("INFO log incorrect value", testResult)
	}
}

func TestInfoLevel(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	defer func() { stdOut = out }()

	testExpectation := ""
	logger := GetLogger("error")
	logger.Info([]int{1, 2, 3})

	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("INFO log incorrect value", testResult)
	}
}

func TestErrorOutput(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	err := stdError
	stdError = &bytes.Buffer{}
	defer func() {
		stdOut = out
		stdError = err
	}()

	logger := GetLogger("info")
	logger.Error([]int{1, 2, 3})

	stdResult := stdOut.(*bytes.Buffer).String()
	if stdResult != "" {
		t.Fatal("ERROR log incorrect process out")
	}

	testExpectation := "ERROR - [1 2 3]"
	errResult := stdError.(*bytes.Buffer).String()
	if errResult != testExpectation {
		t.Fatal("ERROR log incorrect value", errResult)
	}
}
