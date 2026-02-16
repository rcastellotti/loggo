# loggo

`loggo` is a super simple logging library for go.
to use it, add `logger.go` to your project and import it in your main package.
see `main.go` for an example of how to use it.

NOTE: please use [rs/zerolog](https://github.com/rs/zerolog) for serious stuff.

```sh
rc@bearbook ~/loggo> go run main.go
[10:28:08] ERROR This is an error message (always shown)

rc@bearbook ~/loggo> go run main.go -v
[10:29:57] ERROR This is an error message (always shown)
[10:29:57] WARN  This is a warning message (shown with -v)
rc@bearbook ~/loggo> go run main.go -v -v
[10:29:59] ERROR This is an error message (always shown)
[10:29:59] WARN  This is a warning message (shown with -v)
[10:29:59] INFO  This is an info message (shown with -v -v)

rc@bearbook ~/loggo> go run main.go -v -v -v
[10:30:00] ERROR This is an error message (always shown)
[10:30:00] WARN  This is a warning message (shown with -v)
[10:30:00] INFO  This is an info message (shown with -v -v)
[10:30:00] DEBUG This is a debug message (shown with -v -v -v)

rc@bearbook ~/loggo> go run main.go -v -v -v -v
[10:30:02] ERROR This is an error message (always shown)
[10:30:02] WARN  This is a warning message (shown with -v)
[10:30:02] INFO  This is an info message (shown with -v -v)
[10:30:02] DEBUG This is a debug message (shown with -v -v -v)
[10:30:02] TRACE This is a trace message (shown with -v -v -v -v)
```
