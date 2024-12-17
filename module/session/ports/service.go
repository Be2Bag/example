package ports

import "github.com/Be2Bag/example/module/session/dto"

type SessionService interface {
	Login(sessionRequest dto.SessionRequest) (string, error)
}
