package ieproxy

// In Go, the `init()` function is a special function used for package initialization tasks. Here are some key points about it:
// 1. **Automatic Execution**: The `init()` function is automatically executed when a package is initialized. You don't need to call it explicitly.
// 2. **Multiple `init()` Functions**: A single package can have multiple `init()` functions. They are executed in the order they are defined.
// 3. **Purpose**: It is typically used to set up initial values, configurations, or environment settings required by the package before the `main()` function runs.
// Common uses include initializing database connections, setting log levels, or loading configuration files.
// Although init() is a useful tool, it can sometimes make code difficult to read, since a hard-to-find init() instance will greatly affect the order
// in which the code is run. Because of this, it is important for developers who are new to Go to understand the facets of this function, so that they
// can make sure to use init() in a legible manner when writing code.
// https://www.gyata.ai/golang/golang-init-function.
// https://www.golinuxcloud.com/golang-init-function/.
// https://www.educative.io/answers/what-are-init-functions-in-go.
// https://www.digitalocean.com/community/tutorials/understanding-init-in-go.
// https://golangdocs.com/init-function-in-golang.

func init() {
	OverrideEnvWithStaticProxy()
}
