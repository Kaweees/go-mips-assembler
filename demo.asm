.data               # Data segment
hello_str: .asciiz "Hello, World!\n"

.text               # Code segment
.globl main         # Declare main function as global
main:
    li $v0, 4       # Load immediate value 4 into $v0 (print string syscall)
    la $a0, hello_str # Load address of hello_str into $a0
    syscall         # Make the syscall

    li $v0, 10      # Load immediate value 10 into $v0 (exit syscall)
    syscall         # Make the syscall to exit
