package ieproxy

import "os"

type ProxyConf struct {
	Static    StaticProxyConf
	Automatic ProxyScriptConf
}

type StaticProxyConf struct {
	Active    bool              // Is Proxy Active
	Protocols map[string]string // proxy addresses for each scheme, "" is fallback
	NoProxy   string            // comma separated address not to use Proxy
}

type ProxyScriptConf struct {
	Active           bool
	PreConfiguredURL string // pre-configured .pac file
}

// GetConf retrieves the proxy configuration from the Windows Regedit
func GetConf() ProxyConf {
	return getConf()
}

// ReloadConf reloads the proxy configuration
func ReloadConf() ProxyConf {
	return reloadConf()
}

// OverrideEnvWithStaticProxy writes new values to the
// `http_proxy`, `https_proxy` and `no_proxy` environment variables.
// The values are taken from the Windows Regedit (should be called in `init()` function - see example)
func OverrideEnvWithStaticProxy() {
	overrideEnvWithStaticProxy(GetConf(), os.Setenv)
}

// FindProxyForURL computes the proxy for a given URL according to the pac file
func (psc *ProxyScriptConf) FindProxyForURL(URL string) string {
	return psc.findProxyForURL(URL)
}

type envSetter func(string, string) error
