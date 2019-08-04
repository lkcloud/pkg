package errcode

import (
	"github.com/lkcloud/errors"
)

// 错误代码说明：20502
// 2: 服务级错误（1为系统级错误）
// 05: 服务模块代码 (从00开始)
// 02: 具体错误代码 (从00开始)

// error codes below 1000 are reserved future use by the
// "github.com/lkcloud/errors" package.

// system level: common errors
const (
	// ErrBind - error occurred while binding the request body to the struct
	ErrBind errors.Code = iota + 10000

	// ErrValidation - validation failed
	ErrValidation

	// ErrPermissionDenied - permission denied
	ErrPermissionDenied

	// ErrPageNotFound - page not found
	ErrPageNotFound
)

// system level:  database errors
const (
	// ErrDatabase - database error
	ErrDatabase errors.Code = iota + 10101
)

// service level: user module errors
const (
	// ErrEncrypt - error occurred while encrypting the user password
	ErrEncrypt errors.Code = iota + 20000

	//  ErrUserNotFound - user was not found
	ErrUserNotFound

	// ErrTokenInvalid - token was invalid
	ErrTokenInvalid

	// ErrPasswordIncorrect - password was incorrect
	ErrPasswordIncorrect
)

// service level: secret module errors
const (
	// ErrEncrypt - error occurred while encrypting the user password
	ErrReachMaxCount errors.Code = iota + 20100

	//  ErrSecretNotFound - user was not found
	ErrSecretNotFound
)

func init() {
	errors.Codes[ErrBind] = errors.ErrCode{
		Ext: "error occurred while binding the request body to the struct",
	}
	errors.Codes[ErrValidation] = errors.ErrCode{
		Ext: "validation failed",
	}
	errors.Codes[ErrPermissionDenied] = errors.ErrCode{
		Ext: "permission denied",
	}
	errors.Codes[ErrPageNotFound] = errors.ErrCode{
		Ext: "page not found",
	}
	errors.Codes[ErrDatabase] = errors.ErrCode{
		Ext: "database error",
	}
	errors.Codes[ErrEncrypt] = errors.ErrCode{
		Ext: "error occurred while encrypting the user password",
	}
	errors.Codes[ErrUserNotFound] = errors.ErrCode{
		Ext: "user was not found",
	}
	errors.Codes[ErrTokenInvalid] = errors.ErrCode{
		Ext: "token was invalid",
	}
	errors.Codes[ErrPasswordIncorrect] = errors.ErrCode{
		Ext: "password was incorrect",
	}
	errors.Codes[ErrReachMaxCount] = errors.ErrCode{
		Ext: "secret reach to max count",
	}
	errors.Codes[ErrSecretNotFound] = errors.ErrCode{
		Ext: "secret was not found",
	}
}
