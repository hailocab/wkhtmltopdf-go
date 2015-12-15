package demonstrate

import (
	"testing"

	//	"fmt"
	"github.com/hailocab/wkhtmltopdf-go/wkhtmltopdf"
	tos "os"
)

func TestPdfFromStream(t *testing.T) {
	// global settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfGlobal
	gs := wkhtmltopdf.NewGolbalSettings()
	gs.Set("outputFormat", "pdf")
	// Output will be to an internal buffer
	gs.Set("out", "")
	gs.Set("orientation", "Portrait")
	gs.Set("colorMode", "Color")
	gs.Set("size.paperSize", "A4")
	// object settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfObject
	os := wkhtmltopdf.NewObjectSettings()
	os.Set("load.debugJavascript", "false")
	os.Set("load.loadErrorHandling", "ignore")
	//os.Set("load.jsdelay", "1000") // wait max 1s
	os.Set("web.enableJavascript", "false")
	os.Set("web.enablePlugins", "false")
	os.Set("web.loadImages", "true")
	os.Set("web.background", "true")

	c := gs.NewConverter()
	// Some sample text
	c.AddHtml(os, "<html><body><h3>HELLO</h3><p>Hailo's World of Cruft</p></body></html>")

	c.ProgressChanged = func(c *wkhtmltopdf.Converter, b int) {
		t.Logf("Progress: %d\n", b)
	}
	c.Error = func(c *wkhtmltopdf.Converter, msg string) {
		t.Logf("error: %s\n", msg)
	}
	c.Warning = func(c *wkhtmltopdf.Converter, msg string) {
		t.Logf("error: %s\n", msg)
	}
	c.Phase = func(c *wkhtmltopdf.Converter) {
		t.Logf("Phase\n")
	}
	c.Convert()

	t.Logf("Got error code: %d\n", c.ErrorCode())

	if c.ErrorCode() != 0 {
		t.Errorf("Conversion to PDF failed: incomprehensible error-code: %v", c.ErrorCode())
	}

	lout, outp := c.Output()
	lo := int(lout)

	t.Logf("Output %d char.s from conversion\n", lout)
	if lo != 10406 || lo != len(outp) {
		t.Errorf("Conversion to PDF incorrect: lengths out of kilter: expected: %d lout: %d len text: %d", 10406, lout, len(outp))
	}

	t.Logf("Open file for writing... direct_test.pdf")
	f, err := tos.OpenFile("direct_test.pdf", tos.O_WRONLY|tos.O_CREATE, tos.ModePerm)
	if err != nil {
		t.Errorf("Failed to open file: %s\n", err)
	}
	defer func() { f.Close(); t.Logf("Closed PDF file") }()
	f.Truncate(0)
	f.Write([]byte(outp))
}
