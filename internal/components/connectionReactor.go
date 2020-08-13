package Server

import (
	"context"
	"github.com/bhbosman/gocomms/connectionManager"
	commsImpl "github.com/bhbosman/gocomms/impl"
	log "github.com/bhbosman/gologging"
	multiBlock "github.com/bhbosman/gomessageblock"
	"github.com/bhbosman/goprotoextra"
	"github.com/reactivex/rxgo/v2"
	"net"
	"net/url"
)

type connectionReactor struct {
	commsImpl.BaseConnectionReactor
}

func (self *connectionReactor) Init(
	conn net.Conn,
	url *url.URL,
	connectionId string,
	connectionManager connectionManager.IConnectionManager,
	onSend goprotoextra.ToConnectionFunc,
	ddd goprotoextra.ToReactorFunc) (rxgo.NextExternalFunc, error) {
	_, err := self.BaseConnectionReactor.Init(conn, url, connectionId, connectionManager, onSend, ddd)
	if err != nil {
		return nil, err
	}
	self.Logger.NameChange(self.ConnectionId)
	return self.doNext, nil
}

func (self *connectionReactor) Close() error {
	return self.BaseConnectionReactor.Close()
}

func (self *connectionReactor) Open() error {
	return self.BaseConnectionReactor.Open()
}

func (self *connectionReactor) doNext(external bool, i interface{}) {
	if self.CancelCtx.Err() != nil {
		return
	}
	defer func() {
		recover()
	}()
	switch v := i.(type) {
	case *multiBlock.ReaderWriter:
		self.ToConnection(v)
	}
}

func newConnectionReactor(
	logger *log.SubSystemLogger,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	name string,
	userContext interface{}) *connectionReactor {
	result := &connectionReactor{
		BaseConnectionReactor: commsImpl.NewBaseConnectionReactor(
			logger,
			name,
			cancelCtx,
			cancelFunc,
			userContext),
	}
	return result
}
