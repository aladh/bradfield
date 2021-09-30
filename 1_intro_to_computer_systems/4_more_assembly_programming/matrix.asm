section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex

  ; offset = (cols * rindex + cindex) * 4
  imul edx, ecx ; cols * rindex
  add edx, r8d ; + cindex
  mov eax, dword [rdi + rdx * 4] ; base + offset * 4
	ret
