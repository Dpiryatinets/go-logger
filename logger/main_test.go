package logger

import (
	"bytes"
	"testing"
	"time"
)

func TestDebugOutput(t *testing.T) {
	// TODO: move mocks to separate function/package
	out := logger2.stdOut
	logger2.stdOut = &bytes.Buffer{}
	actualNow := logger2.now
	logger2.now = fakeNow
	defer func() { logger2.stdOut = out;  logger2.now = actualNow }()

	logger := logger2.GetLogger("debug")
	logger.Debug([]int{1, 2, 3})

	testExpectation := "2019-10-04T17:24:12 - DEBUG - [1 2 3]"
	testResult := logger2.stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("DEBUG log incorrect value", testResult)
	}
}

func TestDebugLevel(t *testing.T) {
	out := logger2.stdOut
	logger2.stdOut = &bytes.Buffer{}
	defer func() { logger2.stdOut = out }()

	testExpectation := ""
	logger := logger2.GetLogger("error")
	logger.Debug([]int{1, 2, 3})

	testResult := logger2.stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("DEBUG log incorrect value", testResult)
	}
}

func TestInfoOutput(t *testing.T) {
	out := logger2.stdOut
	logger2.stdOut = &bytes.Buffer{}
	actualNow := logger2.now
	logger2.now = fakeNow
	defer func() { logger2.stdOut = out; logger2.now = actualNow }()

	logger := logger2.GetLogger("info")
	logger.Info([]int{1, 2, 3})

	testExpectation := "2019-10-04T17:24:12 - INFO - [1 2 3]"
	testResult := logger2.stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("INFO log incorrect value", testResult)
	}
}

func TestInfoLevel(t *testing.T) {
	out := logger2.stdOut
	logger2.stdOut = &bytes.Buffer{}
	defer func() { logger2.stdOut = out }()

	testExpectation := ""
	logger := logger2.GetLogger("error")
	logger.Info([]int{1, 2, 3})

	testResult := logger2.stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("INFO log incorrect value", testResult)
	}
}

func TestErrorOutput(t *testing.T) {
	out := logger2.stdOut
	logger2.stdOut = &bytes.Buffer{}
	err := logger2.stdError
	logger2.stdError = &bytes.Buffer{}
	actualNow := logger2.now
	logger2.now = fakeNow
	defer func() {
		logger2.stdOut = out
		logger2.stdError = err
		logger2.now = actualNow
	}()

	logger := logger2.GetLogger("info")
	logger.Error([]int{1, 2, 3})

	stdResult := logger2.stdOut.(*bytes.Buffer).String()
	if stdResult != "" {
		t.Fatal("ERROR log incorrect process out")
	}

	testExpectation := "2019-10-04T17:24:12 - ERROR - [1 2 3]"
	errResult := logger2.stdError.(*bytes.Buffer).String()
	if errResult != testExpectation {
		t.Fatal("ERROR log incorrect value", errResult)
	}
}


func fakeNow() time.Time {
	return time.Date(2019, 10, 04, 17, 24, 12, 0, time.UTC)
}