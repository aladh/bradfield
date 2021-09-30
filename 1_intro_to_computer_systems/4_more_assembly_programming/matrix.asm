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
  imul edx, 4 ; * 4
  add rdi, rdx ; base + offset
  mov eax, dword [rdi]
	ret
