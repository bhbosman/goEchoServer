package Server

import (
	"context"
	"github.com/bhbosman/gocomms/intf"
	log "github.com/bhbosman/gologging"
)

type connectionReactorFactory struct {
	name string
}

func (self *connectionReactorFactory) Create(
	name string,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	logger *log.SubSystemLogger,
	userContext interface{}) intf.IConnectionReactor {
	result := newConnectionReactor(logger, cancelCtx, cancelFunc, name, userContext)
	return result
}

func (self *connectionReactorFactory) Name() string {
	return self.name
}
