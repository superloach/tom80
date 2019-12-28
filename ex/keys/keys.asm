.org	0xC00

loop:
	IN	A,	0x01
	AND	0b10000000
	JR	NZ,	noup
	up:
		v_poke	2,	1,	c_white
		JP loop
	noup:
		v_poke	2,	1,	c_black
		JP loop

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

k_up:		DB	0b10000000
k_down:		DB	0b01000000
k_left:		DB	0b00100000
k_right:	DB	0b00010000
k_a:		DB	0b00001000
k_b:		DB	0b00000100
k_c:		DB	0b00000010
k_menu:		DB	0b00000001
