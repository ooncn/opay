module github.com/ooncn/opay

go 1.15

replace github.com/ooncn/common => ..\common

require (
	github.com/kataras/iris/v12 v12.1.8
	github.com/ooncn/common v0.0.0-00010101000000-000000000000
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)
