package qrcode

import (
	"testing"
)

func TestQRCode(t *testing.T) {
	qr := NewQRCode("https://github.com/pibigstar/qrcode", true)
	qr.Debug()
	qr.Output()
}
