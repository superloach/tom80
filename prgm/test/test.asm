.cstr	"name:Test"

.org	0xC00

loop:
	IN	A,	(1)
	AND	10000000B
	JR	NZ,	set
	JP	unset

set:
	v_poke	0,	0,	c_red
	JP	loop

unset:
	v_poke	0,	0,	c_black
	JP	loop

v_poke:	.macro x, y, c
	LD	DE,	64 * (y) + (x)
	LD	A,	(c)
	LD	(DE),	A
.endm

c_red:		DB	0x30
c_yellow:	DB	0x3C
c_green:	DB	0x0C
c_forest:	DB	0x08
c_cyan:		DB	0x0F
c_blue:		DB	0x03
c_magenta:	DB	0x33
c_black:	DB	0x00
c_white:	DB	0x3F
