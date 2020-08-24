package grpcErrorHanding

import (
	"encoding/json"
	log "github.com/isloveh/grpclog"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Error struct {
	Client *log.Logs `json:"client"`
}

func NewClient(path string, withMaxAge, withRotationTime time.Duration, isWrite bool, level logrus.Level) *Error {
	client := &log.Logs{
		Path:             path,
		WithMaxAge:       withMaxAge,
		WithRotationTime: withRotationTime,
		IsWrite:          isWrite,
		Level:            level,
	}
	e := new(Error)
	e.Client = client
	return e
}

func (e *Error) RecordInfo(message string) {
	e.Client.Info(message)
}

func (e *Error) RecordWarn(warn string) {
	e.Client.Warn(warn)
}

func (e *Error) RecordDebug(debug string) {
	e.Client.Debug(debug)
}

func (e *Error) RecodeError(code codes.Code, message string, errorType string, err error) error {
	httpCode := HTTPStatusFromCode(code)
	e.Client.Error(httpCode, tracerr.Wrap(err), errorType)
	response := struct {
		Message   string
		ErrorType string
		Uuid      string
		Error     int
	}{
		Message:   message,
		ErrorType: errorType,
		Uuid:      uuid.NewV1().String(),
		Error:     httpCode,
	}
	b, _ := json.Marshal(response)
	return status.Errorf(codes.Internal, "%s", b)
}

func (e *Error) RecodeFatal(err error) {
	e.Client.Fatal(tracerr.Wrap(err))
}

func (e *Error) RecordPanic(err error) {
	e.Client.Panic(err)
}
