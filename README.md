# QRCode

fork: https://github.com/qianlnk/qrcode
add use mod

a tool to generate qrcode and print it on console.

## Install

```shell
go get -u github.com/pibigstar/qrcode
```

## Usage

* [ ] cmd

```shell
qrcode 'https://github.com/pibigstar/qrcode'
```

* [ ] package

```golang
package main

import (
	"github.com/pibigstar/qrcode"
)

func main() {
	qr := qrcode.NewQRCode("https://github.com/pibigstar/qrcode", false)
	qr.Output()
}

```

* [ ] result

![](qrcode.gif)

## Note

max context length is 60.

