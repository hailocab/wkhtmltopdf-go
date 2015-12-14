package main

import (
	"fmt"
	tos "os"
	"github.com/hailocab/wkhtmltopdf-go/wkhtmltopdf"
)

func main() {
	// global settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfGlobal
	gs := wkhtmltopdf.NewGolbalSettings()
	gs.Set("outputFormat", "pdf")
	//gs.Set("out", "test.pdf")
	gs.Set("out", "")
	gs.Set("orientation", "Portrait")
	gs.Set("colorMode", "Color")
	gs.Set("size.paperSize", "A4")
	//gs.Set("load.cookieJar", "myjar.jar")
	// object settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfObject
	os := wkhtmltopdf.NewObjectSettings()
	//os.Set("page", "http://www.slashdot.org")
	//os.Set("page", "")
	os.Set("load.debugJavascript", "false")
	os.Set("load.loadErrorHandling", "ignore")
	//os.Set("load.jsdelay", "1000") // wait max 1s
	os.Set("web.enableJavascript", "false")
	os.Set("web.enablePlugins", "false")
	os.Set("web.loadImages", "true")
	os.Set("web.background", "true")

	c := gs.NewConverter()
	//c.Add(os)
	c.AddHtml(os, "<html><body><h3>HELLO</h3><p>Hailo's World of Cruft</p></body></html>")

	c.ProgressChanged = func(c *wkhtmltopdf.Converter, b int) {
		fmt.Printf("Progress: %d\n", b)
	}
	c.Error = func(c *wkhtmltopdf.Converter, msg string) {
		fmt.Printf("error: %s\n", msg)
	}
	c.Warning = func(c *wkhtmltopdf.Converter, msg string) {
		fmt.Printf("error: %s\n", msg)
	}
	c.Phase = func(c *wkhtmltopdf.Converter) {
		fmt.Printf("Phase\n")
	}
	c.Convert()

	fmt.Printf("Got error code: %d\n", c.ErrorCode())

	lout, outp := c.Output()

	fmt.Printf("Output %d char.s from conversion\n", lout)

        f, err := tos.OpenFile("direct_test.pdf", tos.O_WRONLY|tos.O_CREATE, tos.ModePerm)
        if err != nil {
                fmt.Printf("Failed to open file: %s\n", err)
        }
        defer func() { f.Close() }()
        f.Truncate(0)
        f.Write([]byte(outp))
}
