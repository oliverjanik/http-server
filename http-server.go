package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app = kingpin.New("http-server", "Simple utility HTTP server")

	port = app.Flag("port", "Port to serve on.").Short('p').Default("8000").Int()

	serve    = app.Command("serve", "Serves files from a directory")
	serveDir = serve.Arg("dir", "Directory to serve").Default(".").String()

	proxy       = app.Command("proxy", "Proxies another HTTP server")
	proxyServer = proxy.Arg("server", "Server to proxy").Required().String()
)

func main() {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cmd {
	case serve.FullCommand():
		runServe(*serveDir, *port)
	case proxy.FullCommand():
		runProxy(*proxyServer, *port)
	}
}
