package Server

import (
	"github.com/bhbosman/gocomms/impl"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/netListener"
	"go.uber.org/fx"
)

func RegisterEchoServiceListener() fx.Option {
	const CreateClientHandlerFactoryName = "EchoServerConnectionReactorFactory"
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Group: impl.ConnectionReactorFactoryConst,
				Target: func() (intf.IConnectionReactorFactory, error) {
					return &connectionReactorFactory{
						name: CreateClientHandlerFactoryName,
					}, nil
				},
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(Empty)",
					"tcp4://127.0.0.1:3000",
					impl.TransportFactoryEmptyName,
					impl.CreateEmptyStack,
					CreateClientHandlerFactoryName),
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(Compressed)",
					"tcp4://127.0.0.1:3001",
					impl.TransportFactoryCompressedName,
					impl.CreateCompressedStack,
					CreateClientHandlerFactoryName),
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(UnCompressed)",
					"tcp4://127.0.0.1:3002",
					impl.TransportFactoryUnCompressedName,
					impl.CreateUnCompressedStack,
					CreateClientHandlerFactoryName),
			}),
	)
}
