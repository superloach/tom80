.cstr	"name:Test"

.org	0xC00

loop:
	IN	A,	(1)
	AND	10000000B
	JR	NZ,	upset
	JP	upunset
	updone:
	IN	A,	(1)
	AND	01000000B
	JR	NZ,	downset
	JP	downunset
	downdone:
	IN	A,	(1)
	AND	00100000B
	JR	NZ,	leftset
	JP	leftunset
	leftdone:
	IN	A,	(1)
	AND	00010000B
	JR	NZ,	rightset
	JP	rightunset
	rightdone:
	JP	loop

upset:
	v_poke	2,	1,	c_red
	JP	updone

upunset:
	v_poke	2,	1,	c_black
	JP	updone

downset:
	v_poke	2,	3,	c_red
	JP	downdone

downunset:
	v_poke	2,	3,	c_black
	JP	downdone

leftset:
	v_poke	1,	2,	c_red
	JP	leftdone

leftunset:
	v_poke	1,	2,	c_black
	JP	leftdone

rightset:
	v_poke	3,	2,	c_red
	JP	rightdone

rightunset:
	v_poke	3,	2,	c_black
	JP	rightdone

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
