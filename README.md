## wkhtmltopdf Go Shim

### Overview

This project contains a fork from https://github.com/jimmyw/wkhtmltopdf-go

What we have added is the ability to:

* pass an HTML page as a string (Converter.AddHtml(...))
* pull back the output as a string (Converter.Output() (int64, string))

Depends on wkhtmltopdf library
http://wkhtmltopdf.org/downloads.html

