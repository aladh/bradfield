default rel
section .rodata
  pi: dd 3.141592654
  three: dd 3.0
section .text
global volume
volume:
  mulss xmm0, xmm0 ; r^2
  mulss xmm0, [pi] ; * pi
  mulss xmm0, xmm1 ; * h
  divss xmm0, [three] ; /3
 	ret
