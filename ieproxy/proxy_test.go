package ieproxy

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func init() {
	fmt.Println("Initialize Static Proxy in proxy_test")
	OverrideEnvWithStaticProxy()
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment
}

func TestExample(t *testing.T) {
	fmt.Println("== Proxy Configuration ==")
	for _, name := range []string{"http_proxy", "https_proxy", "no_proxy"} {
		fmt.Println(name + ": " + os.Getenv(name))
	}

	fmt.Println("== Proxy test ==")

	req, err := http.NewRequest("GET", "https://golang.org/", nil)
	if err != nil {
		panic(err)
	}
	url, err := http.DefaultTransport.(*http.Transport).Proxy(req)
	if err != nil {
		panic(err)
	}
	if url != nil {
		fmt.Println("PROXY " + url.String())
	} else {
		fmt.Println("DIRECT")
	}

}
