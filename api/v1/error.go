package log_v1

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrOffsetOutOfRange struct {
	Offset uint64
}

func (e ErrOffsetOutOfRange) GRPCStatus() *status.Status {
	st := status.New(codes.Code(code.Code_OUT_OF_RANGE), fmt.Sprintf("offset out of range: %d", e.Offset))
	msg := fmt.Sprintf("The requested offset is outside the log's range: %d", e.Offset)
	// errdetails allows for attaching of additional metadata to the status.
	d := &errdetails.LocalizedMessage{
		Locale:  "en-US",
		Message: msg,
	}
	stWithDetails, err := st.WithDetails(d)
	if err != nil {
		return st
	}
	return stWithDetails
}

func (e ErrOffsetOutOfRange) Error() string {
	return e.GRPCStatus().Err().Error()
}
