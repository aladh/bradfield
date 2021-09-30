section .text
global fib
; int fib(int n) {
;    if (n <= 1) return n;
;
;   return fib(n-1) + fib(n-2);
;  }

fib:
  cmp rdi, 1 ; check for n <= 1
  jle .base_case ; skip recursive calls and return n

  push rdi ; preserve existing value of n
  dec rdi ; prepare arg n-1
  call fib ; recursive call
  push rbx ; preserve existing value of rbx (used below)
  mov rbx, rax ; move return value to a non-volatile register

  dec rdi ; prepare arg n-2
  call fib ; recursive call
  add rax, rbx ; add the two return values from recursive calls

  pop rbx ; restore rbx
  pop rdi ; restore n
  ret
.base_case:
  mov rax, rdi ; return n
  ret
