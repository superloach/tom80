v_init	31,	14,	0x30
v_init	31,	15,	0x30
v_init	31,	16,	0x30
v_init	31,	17,	0x30
v_init	31,	18,	0x30
v_init	31,	19,	0x30
v_init	31,	20,	0x30
v_init	31,	21,	0x30
v_init	31,	22,	0x30
v_init	31,	23,	0x30
v_init	31,	24,	0x30
v_init	31,	25,	0x30
v_init	31,	26,	0x30
v_init	31,	27,	0x30
v_init	31,	28,	0x30
v_init	31,	29,	0x30
v_init	31,	30,	0x30
v_init	31,	31,	0x30
v_init	31,	32,	0x30
v_init	31,	33,	0x30

v_init	32,	22,	0x30
v_init	32,	23,	0x30
v_init	32,	24,	0x30
v_init	32,	25,	0x30

v_init:	.macro	$x,	$y,	$c
	.org	$y * 64 + $x
	DB	$c
.endm
