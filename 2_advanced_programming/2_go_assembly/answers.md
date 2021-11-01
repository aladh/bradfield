1. In what ways is Go assembly a “pseudo-assembly”? How is it different from something like x86-64?
   1. It is not a direct representation of the underlying machine. Some instructions map directly to the machine, and others don't.
2. Where does Go’s assembler fit into the build pipeline?
   1. The compiler outputs pseudo-assembly, and the assembler+linker output the final machine code.
3. What details are specified by an “Application Binary Interface”, and what specific decisions does the Go ABI make?
   1. ABI specifies calling convention and memory layout of a program.
   2. There is a stable ABI (ABI0) and another ABI (ABIInternal) which are able to call each other.
4. How do the two (stack-based or register-based) calling conventions handle Go’s multiple return values?
   1. In the stack-based calling convention, space for return values is allocated on the caller's stack.
   2. In register-based callign convention, mulitple registers are used.
5. Can you experimentally check (e.g. by looking at assembler output) whether a locally-compiled program uses a stack-based or register-based calling convention?
