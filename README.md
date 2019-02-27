# httplog for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/lassik/go-httplog)](https://goreportcard.com/report/github.com/lassik/go-httplog)

This is a composable HTTP logging library. It doesn't dictate how you
should log your stuff. Instead, it provides a toolkit to help roll
your own logger. Use the parts you like and ignore the rest.

* It's based around a `LogRequest` struct that contains parsed info
  about a HTTP request.
* The `LogHandler` middleware lets you log any HTTP request. You give
  it a function that gets a `LogRequest`. The function can format and
  write log entries any way it likes.
* To format log entries, you can use the `CommonLogLine` or
  `CombinedLogLine` function. It's also easy to roll your own
  formatting: `LogRequest` should have all the info you need.

The hard parts of this library were extracted from the [Gorilla
Handlers](https://github.com/gorilla/handlers) package. Credit for
them goes to Mahmud Ridwan. I just added the composable framework
around them. The Gorilla libraries are great, but at the time of
writing they impose some restrictions that make it hard to do fancy
stuff like log into multiple places and use custom formats. This
logging library is also framework-agnostic and doesn't depend on
Gorilla or anything else outside the Go standard libraries.

[Documentation](https://godoc.org/github.com/lassik/go-httplog)
