.org 0x0C00

tree 32, 16

loop1:
	happy_holidays	4,	3,	c_red
	.rept 10
		spin
	.endm

	happy_holidays	3,	4,	c_green
	.rept 10
		spin
	.endm

	happy_holidays	3,	4,	c_black

	JP	loop1

tree: .macro x, y
	v_poke	x-1,	y,	c_forest
	v_poke	x,	y,	c_forest

	v_poke	x-1,	y+1,	c_forest
	v_poke	x,	y+1,	c_forest

	v_poke	x-2,	y+2,	c_forest
	v_poke	x-1,	y+2,	c_forest
	v_poke	x,	y+2,	c_forest
	v_poke	x+1,	y+2,	c_forest

	v_poke	x-2,	y+3,	c_forest
	v_poke	x-1,	y+3,	c_forest
	v_poke	x,	y+3,	c_forest
	v_poke	x+1,	y+3,	c_forest

happy_holidays: .macro x_, y_, c_
	letter_h	x_,	y_,	c_
	letter_a	x_+4,	y_,	c_
	letter_p	x_+8,	y_,	c_
	letter_p	x_+12,	y_,	c_
	letter_y	x_+16,	y_,	c_

	letter_h	x_+24,	y_,	c_
	letter_o	x_+28,	y_,	c_
	letter_l	x_+32,	y_,	c_
	letter_i	x_+36,	y_,	c_
	letter_d	x_+40,	y_,	c_
	letter_a	x_+44,	y_,	c_
	letter_y	x_+48,	y_,	c_
	letter_s	x_+52,	y_,	c_
	letter_exc	x_+56,	y_,	c_
.endm

letter_h: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y+2,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
	v_poke	x+2,	y+3,	c
	v_poke	x+2,	y+4,	c
.endm

letter_a: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+2,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
	v_poke	x+2,	y+3,	c
	v_poke	x+2,	y+4,	c
.endm

letter_p: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+2,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
.endm

letter_y: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c

	v_poke	x+1,	y+2,	c
	v_poke	x+1,	y+3,	c
	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
.endm

letter_o: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
	v_poke	x+2,	y+3,	c
	v_poke	x+2,	y+4,	c
.endm

letter_l: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y+4,	c
.endm

letter_i: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+1,	c
	v_poke	x+1,	y+2,	c
	v_poke	x+1,	y+3,	c
	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+4,	c
.endm

letter_d: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+3,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y+1,	c
	v_poke	x+2,	y+2,	c
	v_poke	x+2,	y+3,	c
	v_poke	x+2,	y+4,	c
.endm

letter_s: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+4,	c

	v_poke	x+1,	y,	c
	v_poke	x+1,	y+2,	c
	v_poke	x+1,	y+4,	c

	v_poke	x+2,	y,	c
	v_poke	x+2,	y+2,	c
	v_poke	x+2,	y+3,	c
	v_poke	x+2,	y+4,	c
.endm

letter_exc: .macro x, y, c
	v_poke	x,	y,	c
	v_poke	x,	y+1,	c
	v_poke	x,	y+2,	c
	v_poke	x,	y+4,	c
.endm

spin:	.macro
	LD	B,	255
	spin%%M:
		DJNZ	spin%%M
.endm

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
