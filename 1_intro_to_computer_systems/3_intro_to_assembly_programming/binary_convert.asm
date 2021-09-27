section .text
global binary_convert
binary_convert:
  xor eax, eax ; set return register to 0
.loop:
  movzx ebx, byte [rdi] ; read byte of string
  shl eax, 1 ; shift accumulator left
  and ebx, 1 ; only consider low order bit
  add eax, ebx ; add value
  add rdi, 1 ; increment pointer
  cmp byte [rdi], 0 ; check for null terminator (end of string)
  jne .loop
	ret
