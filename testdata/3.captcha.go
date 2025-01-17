package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func main() {
	var driver = base64Captcha.DriverString{
		Width:           200,
		Height:          40,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          4,
		Source:          "0123456789",
	}
	c := base64Captcha.NewCaptcha(&driver, store)
	id, b64s, answer, err := c.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	fmt.Println(b64s)
	fmt.Println(answer)

}

func check() {
	if store.Verify("", "", true) {

	}
}
