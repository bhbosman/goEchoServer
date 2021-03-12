package Server

import (
	"github.com/bhbosman/gocomms/impl"
	"github.com/bhbosman/gocomms/netListener"
	"go.uber.org/fx"
)

func RegisterEchoServiceListener() fx.Option {
	cfr := &connectionReactorFactory{
		name: "EchoServerConnectionReactorFactory",
	}
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(Empty)",
					"tcp4://127.0.0.1:3000",
					impl.CreateEmptyStack,
					cfr),
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(Compressed)",
					"tcp4://127.0.0.1:3001",
					impl.CreateCompressedStack,
					cfr),
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netListener.NewNetListenApp(
					"EchoServerConnectionManager(UnCompressed)",
					"tcp4://127.0.0.1:3002",
					impl.CreateUnCompressedStack,
					cfr),
			}),
	)
}
