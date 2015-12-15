package converter

import (
	"fmt"

	log "github.com/cihub/seelog"

	"github.com/hailocab/wkhtmltopdf-go/wkhtmltopdf"
)

func ConvertHtmlStringToPdf(html string) (string, error) {
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
	c.AddHtml(os, html)

	errors := []string{}	
	c.ProgressChanged = func(c *wkhtmltopdf.Converter, b int) {
		log.Debugf("Progress: %d", b)
	}
	c.Error = func(c *wkhtmltopdf.Converter, msg string) {
		log.Errorf("Error converting to PDF: %s", msg)
		errors = append(errors, msg)
	}
	c.Warning = func(c *wkhtmltopdf.Converter, msg string) {
		log.Errorf("Problem converting to PDF: %s", msg)
	}
	c.Phase = func(c *wkhtmltopdf.Converter) {
		log.Debugf("Phase\n")
	}
	c.Convert()
	
	if len(errors) > 0 {
		return "", fmt.Errorf("Error converting to PDF: %s", errors[0])		
	}
	if c.ErrorCode() != 0 {
		return "", fmt.Errorf("Error-code converting to PDF: %d", c.ErrorCode())
	}

	_, outp := c.Output()
	return outp, nil
}