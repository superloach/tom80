.org	0xC00

LD	A,	1
OUT	(0),	A
LD	A,	"h"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"e"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"l"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"l"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"o"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	","
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	" "
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"w"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"o"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"r"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"l"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"d"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"!"
OUT	(0),	A
LD	A,	1
OUT	(0),	A
LD	A,	"\n"
OUT	(0),	A
HALT

loop:
	JP	loop
