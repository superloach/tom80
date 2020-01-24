.org 0x0C00

int_setup:
	; set interrupt mode to 2
	IM	2

	; fill the interrupt table
	LD	HL,	0xFA00
	LD	DE,	0xFA01
	LD	(HL),	0xF8	; routine at 0xF8F8
	LD	BC,	0x100
	LDIR

	; copy the frame routine to 0xF8F8
	LD	HL,	frame
	LD	DE,	0xF8F8
	LD	BC,	frame_end - frame
	LDIR

	; set interrupt table location to 0xFA00
	LD	A,	0xFA
	LD	I,	A

frame:
	; set debug mode to text
	LD	A,	1
	OUT	(0),	A

	; send period to debug
	LD	A,	"."
	OUT	(0),	A

	; exit routine
	RET

frame_end:
