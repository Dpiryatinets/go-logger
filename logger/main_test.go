package logger

import (
	"bytes"
	"testing"
	"time"
)

const fakePid = 112
const fakeServiceName = "fake_api"

func TestLogFormat(t *testing.T) {
	// TODO: move mocks to separate function/package
	out := stdOut
	stdOut = &bytes.Buffer{}
	actualNow := now
	now = fakeNow
	actualGetPid := getPid
	getPid = getFakePid
	defer func() { stdOut = out; now = actualNow; getPid = actualGetPid }()

	logger := &Logger{LoggingOptions{
		LogLevel:    debug,
		ServiceName: fakeServiceName,
	}}
	logger.Debug([]int{1, 2, 3})

	testExpectation := "{\"message\":[1,2,3],\"type\":\"debug\",\"processId\":112,\"date\":" +
		"\"2019-10-04T17:24:12Z\",\"serviceName\":\"fake_api\"}\n"
	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("log format is wrong", testResult, testExpectation)
	}
}

func TestLogLevel(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	defer func() { stdOut = out }()

	testExpectation := ""
	logger := &Logger{LoggingOptions{
		LogLevel:    errorLevel,
		ServiceName: fakeServiceName,
	}}
	logger.Debug([]int{1, 2, 3})

	testResult := stdOut.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("output is not empty", testResult, testExpectation)
	}
}

func TestErrorOutput(t *testing.T) {
	out := stdOut
	stdOut = &bytes.Buffer{}
	err := stdError
	stdError = &bytes.Buffer{}
	actualNow := now
	now = fakeNow
	actualGetPid := getPid
	getPid = getFakePid
	defer func() {
		stdOut = out
		stdError = err
		now = actualNow
		getPid = actualGetPid
	}()

	logger := GetLogger(LoggingOptions{
		LogLevel:    info,
		ServiceName: fakeServiceName,
	})
	logger.Error([]int{1, 2, 3})

	testResult := stdOut.(*bytes.Buffer).String()
	testExpectation := ""
	if testResult != testExpectation {
		t.Fatal("wrong process out", testResult, testExpectation)
	}

	testExpectation = "{\"message\":[1,2,3],\"type\":\"error\",\"processId\":112," +
		"\"date\":\"2019-10-04T17:24:12Z\",\"serviceName\":\"fake_api\"}\n"
	testResult = stdError.(*bytes.Buffer).String()
	if testResult != testExpectation {
		t.Fatal("wrong message using right out", testResult, testExpectation)
	}
}

func fakeNow() time.Time {
	return time.Date(2019, 10, 04, 17, 24, 12, 0, time.UTC)
}

func getFakePid() int {
	return fakePid
}
