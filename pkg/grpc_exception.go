package pkg

import (
	"strconv"

	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/version"
)

// NewExceptionFromTrailer todo
func NewExceptionFromTrailer(md metadata.MD, err error) exception.APIException {
	ctx := newGrpcCtx(md)
	code, _ := strconv.Atoi(ctx.get(ResponseCodeHeader))
	reason := ctx.get(ResponseReasonHeader)
	message := ctx.get(ResponseDescHeader)
	ctx.get(ResponseMetaHeader)
	ctx.get(ResponseDataHeader)
	return exception.NewAPIException(version.ServiceName, code, reason, message)
}
