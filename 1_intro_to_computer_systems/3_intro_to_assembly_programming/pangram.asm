section .text
global pangram
pangram:
  xor edx, edx
.loop:
  movzx ecx, byte [rdi] ; read byte of string
  cmp ecx, 0 ; check for null terminator (end of string)
  je .end
  or ecx, 32 ; force lowercase (by setting the third most significant bit)
  sub ecx, 'a' ; difference between letter and lowercase a
  bts edx, ecx ; flip the bit for the character in edx (bit set)
  inc rdi
  jmp .loop
.end:
  xor eax, eax
  and edx, 0x03ffffff ; ignore the 6 higher order bits (punctuation, etc)
  cmp edx, 0x03ffffff ; check if lower 26 bits are set
  sete al ; set return value to 1 (true) if both values are equal
  ret
