package main

import (
	"context"
	echoServer "github.com/bhbosman/goEchoServer/internal/components"
	app2 "github.com/bhbosman/gocommon/app"
	"github.com/bhbosman/gocomms/connectionManager"
	"github.com/bhbosman/gocomms/connectionManager/endpoints"
	"github.com/bhbosman/gocomms/connectionManager/view"
	commsImpl "github.com/bhbosman/gocomms/impl"
	//"github.com/bhbosman/gocommon/comms/http"
	log2 "github.com/bhbosman/gologging"
	"go.uber.org/fx"
	"log"
	"os"
)

func main() {
	app := fx.New(
		log2.ProvideLogFactory(log.New(os.Stderr, "EchoServer: ", log.LstdFlags), nil),
		connectionManager.RegisterDefaultConnectionManager(),
		commsImpl.RegisterAllConnectionRelatedServices(),
		endpoints.RegisterConnectionManagerEndpoint(),
		view.RegisterConnectionsHtmlTemplate(),
		app2.RegisterRootContext(),
		echoServer.RegisterEchoServiceListener(),
		fx.Provide(
			func(params struct {
				fx.In
				Factory *log2.Factory
			}) *log2.SubSystemLogger {
				return params.Factory.Create("Main")
			}),

		fx.Invoke(
			func(params struct {
				fx.In
				Lifecycle fx.Lifecycle
				Apps      []*fx.App `group:"Apps"`
				Logger    *log2.SubSystemLogger
				//RunTimeManager *app2.RunTimeManager
			}) {

				for _, item := range params.Apps {
					localApp := item
					params.Lifecycle.Append(fx.Hook{
						OnStart: func(ctx context.Context) error {
							return localApp.Start(ctx)
						},
						OnStop: func(ctx context.Context) error {
							return localApp.Stop(ctx)
						},
					})
				}
			}),
	)
	if app.Err() != nil {
		return
	}
	app.Run()
}
