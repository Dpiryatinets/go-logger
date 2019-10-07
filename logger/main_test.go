package logger

import (
	"bytes"
	"testing"
	"time"
)

func TestDebugOutput(t *testing.T) {
	// TODO: move mocks to separate function/package
	out := stdOut
	stdOut = &bytes.Buffer{}
	actualNow := now
	now = fakeNow
	defer func() { stdOut = out;  now = actualNow }()

	logger := GetLogger("debug")
	logger.Debug([]int{1, 2, 3})

	testExpectation := "2019-10-04T17:24:12 - DEBUG - [1 2 3]"
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
	actualNow := now
	now = fakeNow
	defer func() { stdOut = out; now = actualNow }()

	logger := GetLogger("info")
	logger.Info([]int{1, 2, 3})

	testExpectation := "2019-10-04T17:24:12 - INFO - [1 2 3]"
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
	actualNow := now
	now = fakeNow
	defer func() {
		stdOut = out
		stdError = err
		now = actualNow
	}()

	logger := GetLogger("info")
	logger.Error([]int{1, 2, 3})

	stdResult := stdOut.(*bytes.Buffer).String()
	if stdResult != "" {
		t.Fatal("ERROR log incorrect process out")
	}

	testExpectation := "2019-10-04T17:24:12 - ERROR - [1 2 3]"
	errResult := stdError.(*bytes.Buffer).String()
	if errResult != testExpectation {
		t.Fatal("ERROR log incorrect value", errResult)
	}
}


func fakeNow() time.Time {
	return time.Date(2019, 10, 04, 17, 24, 12, 0, time.UTC)
}