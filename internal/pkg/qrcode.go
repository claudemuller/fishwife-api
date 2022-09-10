package pkg

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func GenQrCode(url string) {
	const tmpLoc = "/tmp/fishwife-qr.png"

	err := qrcode.WriteFile(url, qrcode.Medium, 256, tmpLoc)
	if err != nil {
		log.Fatalf("error generating QR code: %v", err)
	}
}
