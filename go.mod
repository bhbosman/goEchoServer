module github.com/bhbosman/goEchoServer

go 1.15

require (
	github.com/bhbosman/gocommon v0.0.0-20201004145117-eae02ab42c9a
	github.com/bhbosman/gocomms v0.0.0-20210108094235-212b4e8c628c
	github.com/bhbosman/gologging v0.0.0-20200921180328-d29fc55c00bc
	github.com/bhbosman/gomessageblock v0.0.0-20200921180725-7cd29a998aa3
	github.com/reactivex/rxgo/v2 v2.1.0
	go.uber.org/fx v1.13.1
	github.com/cskr/pubsub v1.0.2
	github.com/bhbosman/goprotoextra v0.0.2-0.20210414124526-a342e2a9e82f
)

replace (
	github.com/bhbosman/goMessages => ../goMessages
	github.com/bhbosman/gocommon => ../gocommon
	github.com/bhbosman/gocomms => ../gocomms
	github.com/bhbosman/gologging => ../gologging
	github.com/bhbosman/gomessageblock => ../gomessageblock
	github.com/reactivex/rxgo/v2 => ../../reactivex/rxgo
)
