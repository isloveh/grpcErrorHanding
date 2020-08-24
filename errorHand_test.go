package grpcErrorHanding

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"testing"
	"time"
)

func TestError_ErrorHand(t *testing.T) {
	NewClient("./log", time.Duration(7)*time.Hour*24, time.Duration(1)*time.Hour*24, true, logrus.InfoLevel).RecodeError(codes.Internal, "服务器错误", codes.Internal.String(), xerrors.New("服务器错误"))
}
