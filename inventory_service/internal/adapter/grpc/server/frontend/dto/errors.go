package dto

import (
	"database/sql"
	"errors"
	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/internal/adapter/postgres/dao"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCError struct {
	Code    codes.Code
	Message string
}

var (
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrEditConflict        = errors.New("unable to update the record due to an edit conflict")

	ErrGRPCNotFoundResponse = &GRPCError{
		Code:    codes.NotFound,
		Message: "the requested resource could not be found",
	}
	ErrGRPCUnprocessableEntityResponse = &GRPCError{
		Code:    codes.InvalidArgument,
		Message: "unprocessable entity",
	}
	ErrGRPCEditConflictResponse = &GRPCError{
		Code:    codes.Aborted,
		Message: "unable to update the record due to an edit conflict",
	}
	ErrGRPCInternalResponse = &GRPCError{
		Code:    codes.Internal,
		Message: "something went wrong",
	}
)

// GRPCFromError maps internal errors to gRPC errors with custom struct
func GRPCFromError(err error) *GRPCError {
	switch {
	case errors.Is(err, sql.ErrNoRows),
		errors.Is(err, pgx.ErrNoRows),
		errors.Is(err, dao.ErrRecordNotFound):
		return ErrGRPCNotFoundResponse

	case errors.Is(err, ErrUnprocessableEntity):
		return ErrGRPCUnprocessableEntityResponse

	case errors.Is(err, ErrEditConflict):
		return ErrGRPCEditConflictResponse

	default:
		return ErrGRPCInternalResponse
	}
}

func StatusFromDomainError(err error) error {
	switch {
	case errors.Is(err, dto.ErrUnprocessableEntity):
		return status.Errorf(codes.InvalidArgument, "unprocessable entity")
	case errors.Is(err, dto.ErrEditConflict):
		return status.Errorf(codes.Aborted, "edit conflict")
	case errors.Is(err, dto.ErrInvalidFilters):
		return status.Errorf(codes.InvalidArgument, "invalid filters")
	case errors.Is(err, sql.ErrNoRows), errors.Is(err, dao.ErrRecordNotFound), errors.Is(err, pgx.ErrNoRows):
		return status.Errorf(codes.NotFound, "resource not found")
	default:
		return status.Errorf(codes.Internal, "internal error: %v", err)
	}
}
