       CC       = gcc
       CXX      = gcc
       CFLAGS   += -I/usr/local/include
       LDFLAGS  += -L/usr/local/lib -lwkhtmltox -Wall -ansi -pedantic -ggdb
       CDEBUG   =
       OBJ_TST  = main.o

harness: $(OBJ_TST)
	$(CC) -o c_harness $(LDFLAGS) $(OBJ_TST) $(LIBS)

clean:
	rm -f *.o *.exe

fresh: clean harness

all: fresh

snurd:
	@echo "Fribble wibbly wib..."
