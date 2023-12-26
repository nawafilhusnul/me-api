package constants

import (
	"fmt"
)

type APIError int

const (
	ValidationException APIError = iota + 1
	InvalidParameterException
	EntityNotFoundException
	UnprocessableEntityException
	UnauthorizedException
	ForbiddenException
	InternalServerErrorException
)

func (e APIError) Error() error {
	switch e {
	case ValidationException:
		return fmt.Errorf("some request failed the validation. please refer to API Docs")
	case InvalidParameterException:
		return fmt.Errorf("request parameter is not accepted in this endpoint. please refer to API Docs")
	case EntityNotFoundException:
		return fmt.Errorf("entity not found in our database")
	case UnprocessableEntityException:
		return fmt.Errorf("can not perform the action for requested entity")
	case UnauthorizedException:
		return fmt.Errorf("seem you do not have a valid session. please relogin")
	case ForbiddenException:
		return fmt.Errorf("this endpoint is not for you. try other")
	case InternalServerErrorException:
		return fmt.Errorf("oops. We have internal problem. we are about to fix it")
	default:
		return fmt.Errorf("please contact our developer! the system does not know this error")
	}
}

func (e APIError) Code() int {
	switch e {
	case ValidationException:
		return int(ValidationException)
	case InvalidParameterException:
		return int(InvalidParameterException)
	case EntityNotFoundException:
		return int(EntityNotFoundException)
	case UnprocessableEntityException:
		return int(UnprocessableEntityException)
	case UnauthorizedException:
		return int(UnauthorizedException)
	case ForbiddenException:
		return int(ForbiddenException)
	case InternalServerErrorException:
		return int(InternalServerErrorException)
	default:
		return int(0)
	}

}
