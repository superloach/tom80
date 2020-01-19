# tom80
a Z80-based fantasy console.

![pixel pattern](https://media.discordapp.net/attachments/314487938949971980/663244487438630922/2020-01-04-235619_514x386_scrot.png)

## running emulator
currently, the only distribution is ebiten, which should run fine on desktops.
```bash
go run -game name
```
web and mobile support is untested and thus likely incomplete.

## making programs
### compiling
make sure you have [asm80](https://github.com/maly/asm80-node) installed
```bash
asm80 -m Z80 -t bin -o <name>.bin <path>/<name>.asm
# <name>.bin will be output in <path>
```
or use [asm80.com](https://www.asm80.com/) ([asm80 docs](https://maly.gitbooks.io/asm80/))
### writing
programs should start like:
```as
.cstr	"name:program name"
.cstr	"author:your name"
.cstr	"version:0.0"
.org	0x0C00
```
if you know what you're doing, you may add `.cstr "clear:false"`, and VRAM will not be cleared after loading the ROM.

## basic specs
- Screen
  - 64 x 48 pixels
  - 6-bit colour (`0b[__RRGGBB]`)
  - VRAM 0x0000 - 0x0BFF
- Controls
  - ports 0x01 - 0x08
  - Up, Down, Left, Right, A, B, C, Menu (1<<7 - 1<<0)

## todo
order of unfinished items is mostly arbitrary
- [x] rom loading
  - [x] rom info in vram area on load
- [x] screen working
  - [x] use screen in asm
- [x] controls working
  - [x] use controls in asm
- [x] COMPLETE DOCS
- [x] use flag
- [x] add simple way to pause cpu ([hajimehoshi/ebiten#1037](https://github.com/hajimehoshi/ebiten/issues/1037))
- [x] folder layout
  - [x] move programs from /ex to /prgm
  - [x] move system from / to /lib
  - [x] move frontend from /dist/ebiten to /
  - [x] use name in /prgm rather than path
- [x] use port 0x00 for system commands
  - [x] mode setting
  - [ ] debug print mode
  - [ ] file list mode
  - [ ] file load mode
- [ ] "loader" program
  - loaded into rom by default
  - uses file listing/loading
- [ ] audio working
  - [ ] gather 16 sounds
  - [x] use ports 0x09 - 0x18
  - [ ] define event type (`0bVVVPPPPP`)
  - [ ] use audio in asm
- [ ] standard library for asm
  - [ ] simplify control keys
  - [ ] shape drawing macros
  - [ ] fonts?
- [ ] port asm80 to go-asm80
- [ ] integrate go-asm80
  - [ ] compile from asm if bin is missing
  - [ ] attempt recompile if asm is newer
