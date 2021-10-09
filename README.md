# <a href="https://telegram.me/SpamWatch"><img src="https://avatars.githubusercontent.com/u/37397813?s=200&v=4" width="35px" align="left"></img></a> SpamWatch API Go Wrapper 

[![Go Reference](https://pkg.go.dev/badge/github.com/SpamWatch/spamwatch-go.svg)](https://pkg.go.dev/github.com/SpamWatch/spamwatch-go) [![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

spamwatch-go is official Go wrapper for [SpamWatch API](https://api.spamwat.ch), which is fast, secure and requires no additional packages to be installed.

<hr/>

## Features

- Can use custom SpamWatch API endpoint with the help of ClientOpts.
- It's in pure go, no need to install any kind of plugin or include any kind of additional files.
- No third party library bloat; only uses standard library.
- Type safe; no weird `interface{}` logic.

<hr/>

## Getting started

You can easily download the library with the standard `go get` command:

```bash
go get github.com/SpamWatch/spamwatch-go
```

Full documentation of this API, can be found [here](https://docs.spamwat.ch/).

<hr/>

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/SpamWatch/spamwatch-go"
)

var client = spamwatch.Client("API_KEY", nil)

func main() {
	ban, err := client.GetBan(777000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ban)
}
```

Still need more examples? Take a look at the [examples directory](examples).

Ask your doubts at the [support group](https://telegram.me/SpamWatchSupport).

<hr/>

## License

[![GNU General Public License v3.0](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html#header)

The spamwatch-go project is under the [GPL-3.0](https://opensource.org/licenses/GPL-3.0) license. You can find the license file [here](LICENSE).

