package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RunHTTPServerOnAddr(httpAddr string, methodNames []string) {
	webserver, err := newWebserver(*host, methodNames)
	if err != nil {
		log.Fatalf("failed to init webserver: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", webserver.indexPage)
	mux.HandleFunc("/favicon.ico", webserver.favicon)

	log.Printf("listening http on %s\n", httpAddr)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpAddr), mux); err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}
}

type webserver struct {
	t       *template.Template
	content *webpageContent
}

type webpageContent struct {
	Host    string
	Methods []string
}

func newWebserver(host string, methods []string) (*webserver, error) {
	t := template.New("")
	t, err := t.Parse(index)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %w", err)
	}
	return &webserver{t: t, content: &webpageContent{Host: host, Methods: methods}}, nil
}

func (s *webserver) indexPage(w http.ResponseWriter, r *http.Request) {
	if err2 := s.t.Execute(w, s.content); err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}

func (s *webserver) favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	if _, err := fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII="); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// By default proxy serves 80 and 443 for web page, 9000 and 9001 (with TLS) for grpc
var index = `<!DOCTYPE html>https://github.com/grafana/k6-infrastructure/scripts/grpcbin
<html>
  <body>
    <h1>grpcbin: gRPC Request & Response Service</h1>
    <h2>Endpoints</h2>
	<ul>
    <li>grpc://{{.Host}}:9000 (without TLS)</li>
	  <li>grpc://{{.Host}}:9001 (with TLS)</li>
	  <li><a href=https://{{.Host}}>https://{{.Host}}</a>(this web page)</li>
    </ul>
    <h2>Methods</h2>
    <ul>
      <li>
        <a href="https://github.com/moul/pb/blob/master/grpcbin/grpcbin.proto">grpcbin.proto</a>
        <ul>
          {{- range .Methods}}
          <li>{{.}}</li>
          {{- end}}
        </ul>
      </li>
      <li>
        <a href="https://github.com/moul/pb/blob/master/hello/hello.proto">hello.proto</a>
        <ul>
          <li>SayHello</li>
          <li>LotsOfReplies</li>
          <li>LotsOfGreetings</li>
          <li>BidiHello</li>
        </ul>
      </li>
      <li>
        <a href="https://github.com/moul/pb/blob/master/addsvc/addsvc.proto">addsvc.proto</a>
        <ul>
          <li>Sum</li>
          <li>Concat</li>
        </ul>
      </li>
      <li>
        <a href="https://github.com/moul/pb/blob/master/a_bit_of_everything/lib/examples/examplepb/a_bit_of_everything.proto">a_bit_of_everything.proto</a>
        <ul>
          <li>Create</li>
          <li>CreateBody</li>
          <li>Lookup</li>
          <li>Update</li>
          <li>Delete</li>
          <li>GetQuery</li>
          <li>Echo</li>
          <li>DeepPathEcho</li>
          <li>NoBindings</li>
          <li>Timeout</li>
          <li>ErrorWithDetails</li>
          <li>GetMessageWithBody</li>
          <li>PostWithEmptyBody</li>
        </ul>
      </li>
    </ul>
    <h2>Examples</h2>
    <ul>
      <li><a href="https://k6.io/docs/javascript-api/k6-net-grpc/">https://k6.io/docs/javascript-api/k6-net-grpc/</a></li>
      <li><a href="https://k6.io/blog/performance-testing-grpc-services/">https://k6.io/blog/performance-testing-grpc-services/</a></li>
    </ul>
    <h2>About</h2>
	<a href="https://github.com/moul/grpcbin">Developed</a> by <a href="https://manfred.life">Manfred Touron</a>, inspired by <a href="https://httpbin.org/">https://httpbin.org/</a>
	and slightly tuned by <a href="https://k6.io">k6.io</a>
  </body>
</html>
`
