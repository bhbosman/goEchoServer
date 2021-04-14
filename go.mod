module github.com/bhbosman/goEchoServer


go 1.15

require (
	github.com/bhbosman/gocomms v0.0.0-20210414144344-fb75f75793be
	github.com/bhbosman/gocommon v0.0.0-20210414135919-fd7afceec0b0
	github.com/bhbosman/gologging v0.0.0-20200921180328-d29fc55c00bc
	github.com/bhbosman/gomessageblock v0.0.0-20210414135653-cd754835d03b
	github.com/bhbosman/goprotoextra v0.0.2-0.20210414124526-a342e2a9e82f
	github.com/cskr/pubsub v1.0.2
	github.com/reactivex/rxgo/v2 v2.1.0
	go.uber.org/fx v1.13.1

)

replace (
	//github.com/bhbosman/goMessages => ../goMessages
	//github.com/bhbosman/gocommon => ../gocommon
	//github.com/bhbosman/gocomms => ../gocomms
	//github.com/bhbosman/gologging => ../gologging
	github.com/bhbosman/gomessageblock => ../gomessageblock
	github.com/reactivex/rxgo/v2 => ../../reactivex/rxgo
)

