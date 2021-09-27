section .text
global sum_to_n
sum_to_n:
  xor eax, eax ; total = 0
.loop:
  add eax, edi ; total += n
  sub edi, 1 ; count down from n instead of up from 0
  jg .loop ; loop if n > 1
  ret ; return value (total) is already in eax
