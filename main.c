/*
 * main.c
 *
 *  Created on: 14 Jan 2016
 *      Author: john
 */

#include <stdio.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <wkhtmltox/pdf.h>

#define NB 256
#define NH 1024*1024

void error_callback(wkhtmltopdf_converter * converter, const char * str) {
	printf("ERROR reported: %s\n", str);
}

void warning_callback(wkhtmltopdf_converter * converter, const char * str) {
	printf("WARNING reported: %s\n", str);
}

void phase_changed_callback(wkhtmltopdf_converter * converter) {
	printf("Phase changed to: %s\n", wkhtmltopdf_phase_description(converter, wkhtmltopdf_current_phase(converter)));
}

void progress_changed_callback(wkhtmltopdf_converter * converter, const int progess) {
	printf("Progress counter changed to: %d\n", progess);
}

int main() {
	wkhtmltopdf_init(0);

	FILE* sample = fopen("sample.html", "r");
	char* buffer = malloc(sizeof(char) * NB);
	char* html = malloc(sizeof(char) * NH);
	char* current = "";

	if(sample == NULL) {
		printf("Failed to find the sample file");
		exit(1);
	}

	printf("Successfully opened file...\n");

	while((current = fgets(buffer, NB, sample)) != NULL) {
		html = strcat(html, buffer);
//		strncat(html, buffer, NB);
	}

	printf("Len. HTML: %zd\n", strlen(html));

	{
		/* global settings */
		wkhtmltopdf_global_settings* gsettings = wkhtmltopdf_create_global_settings();

		/* object settings */
		wkhtmltopdf_object_settings* osettings = wkhtmltopdf_create_object_settings();

		/* global settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfGlobal
		gs := wkhtmltopdf.NewGolbalSettings()
		gs.Set("outputFormat", "pdf")
		// Output will be to an internal buffer
		gs.Set("out", "")
		gs.Set("orientation", "Portrait")
		gs.Set("colorMode", "Color")
		gs.Set("size.paperSize", "A4") */
		/* object settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfObject
		os := wkhtmltopdf.NewObjectSettings()
		os.Set("load.debugJavascript", "false")
		os.Set("load.loadErrorHandling", "ignore")
		os.Set("load.jsdelay", "1000") // wait max 1s
		os.Set("web.enableJavascript", "false")
		os.Set("web.enablePlugins", "false")
		os.Set("web.loadImages", "true")
		os.Set("web.background", "true") */

		wkhtmltopdf_set_global_setting(gsettings, "outputFormat", "pdf");
//		wkhtmltopdf_set_global_setting(gsettings, "out", "");
		wkhtmltopdf_set_global_setting(gsettings, "out", "sample-out.pdf");
		wkhtmltopdf_set_global_setting(gsettings, "orientation", "portrait");
		wkhtmltopdf_set_global_setting(gsettings, "colorMode", "Color");
		wkhtmltopdf_set_global_setting(gsettings, "size.paperSize", "A4");

		wkhtmltopdf_set_object_setting(osettings, "load.debugJavascript", "false");
		wkhtmltopdf_set_object_setting(osettings, "load.loadErrorHandling", "ignore");
		wkhtmltopdf_set_object_setting(osettings, "load.jsdelay", "1000");
		wkhtmltopdf_set_object_setting(osettings, "web.enableJavascript", "false");
		wkhtmltopdf_set_object_setting(osettings, "web.enablePlugins", "false");
		wkhtmltopdf_set_object_setting(osettings, "web.loadImages", "true");
		wkhtmltopdf_set_object_setting(osettings, "web.background", "true");

		{
			int count = 0;
			int len_out = 0;
			int len_buf = 0;
			const unsigned char** out;
			wkhtmltopdf_converter* converter = wkhtmltopdf_create_converter(gsettings);
			wkhtmltopdf_add_object(converter, osettings, html);
			wkhtmltopdf_set_error_callback(converter, error_callback);
			wkhtmltopdf_set_warning_callback(converter, warning_callback);
			wkhtmltopdf_set_phase_changed_callback(converter, phase_changed_callback);
			wkhtmltopdf_set_progress_changed_callback(converter, progress_changed_callback);

			count = wkhtmltopdf_convert(converter);

			printf("\nCompleted conversion\n\n");

			printf("Any error? Code: %d\n\n", wkhtmltopdf_http_error_code(converter));

//			len_out = wkhtmltopdf_get_output(converter, out);
//			len_buf = strlen((const char*)*out);

			printf("Len. converted: %d Len. output: %d\n", count, len_out);

			wkhtmltopdf_destroy_converter(converter);
			wkhtmltopdf_destroy_object_settings(osettings);
			wkhtmltopdf_destroy_global_settings(gsettings);
		}
	}

	fclose(sample);

	return 0;
}
