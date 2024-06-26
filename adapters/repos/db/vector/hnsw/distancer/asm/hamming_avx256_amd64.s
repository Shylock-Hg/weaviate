//go:build !noasm && amd64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·hamming_256(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ res+16(FP), DX
	MOVQ len+24(FP), CX
	BYTE $0x55               // pushq	%rbp
	WORD $0x8948; BYTE $0xe5 // movq	%rsp, %rbp
	BYTE $0x53               // pushq	%rbx
	LONG $0xf8e48348         // andq	$-8, %rsp
	WORD $0x8b48; BYTE $0x19 // movq	(%rcx), %rbx
	WORD $0xfb83; BYTE $0x07 // cmpl	$7, %ebx
	JG   LBB0_8
	WORD $0x8941; BYTE $0xd9 // movl	%ebx, %r9d
	LONG $0xffc18341         // addl	$-1, %r9d
	WORD $0xc031             // xorl	%eax, %eax
	LONG $0x1ff98341         // cmpl	$31, %r9d
	JAE  LBB0_3
	WORD $0x8949; BYTE $0xf0 // movq	%rsi, %r8
	WORD $0x8949; BYTE $0xfb // movq	%rdi, %r11
	JMP  LBB0_6

LBB0_8:
	LONG $0xc0eff9c5         // vpxor	%xmm0, %xmm0, %xmm0
	WORD $0xfb83; BYTE $0x20 // cmpl	$32, %ebx
	JB   LBB0_9
	LONG $0xc0eff9c5         // vpxor	%xmm0, %xmm0, %xmm0
	LONG $0xc9eff1c5         // vpxor	%xmm1, %xmm1, %xmm1
	LONG $0xd2efe9c5         // vpxor	%xmm2, %xmm2, %xmm2
	LONG $0xdbefe1c5         // vpxor	%xmm3, %xmm3, %xmm3

LBB0_15:
	LONG $0x2610fcc5               // vmovups	(%rsi), %ymm4
	LONG $0x6e10fcc5; BYTE $0x20   // vmovups	32(%rsi), %ymm5
	LONG $0x7610fcc5; BYTE $0x40   // vmovups	64(%rsi), %ymm6
	LONG $0x7e10fcc5; BYTE $0x60   // vmovups	96(%rsi), %ymm7
	LONG $0x27c2dcc5; BYTE $0x0c   // vcmpneq_oqps	(%rdi), %ymm4, %ymm4
	LONG $0x6fc2d4c5; WORD $0x0c20 // vcmpneq_oqps	32(%rdi), %ymm5, %ymm5
	LONG $0xdcfae5c5               // vpsubd	%ymm4, %ymm3, %ymm3
	LONG $0xd5faedc5               // vpsubd	%ymm5, %ymm2, %ymm2
	LONG $0x67c2ccc5; WORD $0x0c40 // vcmpneq_oqps	64(%rdi), %ymm6, %ymm4
	LONG $0xccfaf5c5               // vpsubd	%ymm4, %ymm1, %ymm1
	LONG $0x67c2c4c5; WORD $0x0c60 // vcmpneq_oqps	96(%rdi), %ymm7, %ymm4
	LONG $0xc4fafdc5               // vpsubd	%ymm4, %ymm0, %ymm0
	WORD $0xc383; BYTE $0xe0       // addl	$-32, %ebx
	LONG $0x80ef8348               // subq	$-128, %rdi
	LONG $0x80ee8348               // subq	$-128, %rsi
	WORD $0xfb83; BYTE $0x1f       // cmpl	$31, %ebx
	JA   LBB0_15
	WORD $0xfb83; BYTE $0x08       // cmpl	$8, %ebx
	JAE  LBB0_10
	JMP  LBB0_12

LBB0_3:
	LONG $0x01c18349         // addq	$1, %r9
	WORD $0x894d; BYTE $0xca // movq	%r9, %r10
	LONG $0xe0e28349         // andq	$-32, %r10
	WORD $0x2944; BYTE $0xd3 // subl	%r10d, %ebx
	LONG $0x96048d4e         // leaq	(%rsi,%r10,4), %r8
	LONG $0x971c8d4e         // leaq	(%rdi,%r10,4), %r11
	QUAD $0x000000008d048d4a // leaq	(,%r9,4), %rax
	LONG $0x80e08348         // andq	$-128, %rax
	LONG $0xc0eff9c5         // vpxor	%xmm0, %xmm0, %xmm0
	WORD $0xc931             // xorl	%ecx, %ecx
	LONG $0xc9eff1c5         // vpxor	%xmm1, %xmm1, %xmm1
	LONG $0xd2efe9c5         // vpxor	%xmm2, %xmm2, %xmm2
	LONG $0xdbefe1c5         // vpxor	%xmm3, %xmm3, %xmm3

LBB0_4:
	LONG $0x2410fcc5; BYTE $0x0e               // vmovups	(%rsi,%rcx), %ymm4
	LONG $0x6c10fcc5; WORD $0x200e             // vmovups	32(%rsi,%rcx), %ymm5
	LONG $0x7410fcc5; WORD $0x400e             // vmovups	64(%rsi,%rcx), %ymm6
	LONG $0x7c10fcc5; WORD $0x600e             // vmovups	96(%rsi,%rcx), %ymm7
	LONG $0x24c2dcc5; WORD $0x040f             // vcmpneqps	(%rdi,%rcx), %ymm4, %ymm4
	LONG $0xc4fafdc5                           // vpsubd	%ymm4, %ymm0, %ymm0
	LONG $0x64c2d4c5; WORD $0x200f; BYTE $0x04 // vcmpneqps	32(%rdi,%rcx), %ymm5, %ymm4
	LONG $0xccfaf5c5                           // vpsubd	%ymm4, %ymm1, %ymm1
	LONG $0x64c2ccc5; WORD $0x400f; BYTE $0x04 // vcmpneqps	64(%rdi,%rcx), %ymm6, %ymm4
	LONG $0x6cc2c4c5; WORD $0x600f; BYTE $0x04 // vcmpneqps	96(%rdi,%rcx), %ymm7, %ymm5
	LONG $0xd4faedc5                           // vpsubd	%ymm4, %ymm2, %ymm2
	LONG $0xddfae5c5                           // vpsubd	%ymm5, %ymm3, %ymm3
	LONG $0x80e98348                           // subq	$-128, %rcx
	WORD $0x3948; BYTE $0xc8                   // cmpq	%rcx, %rax
	JNE  LBB0_4
	LONG $0xc0fef5c5                           // vpaddd	%ymm0, %ymm1, %ymm0
	LONG $0xc0feedc5                           // vpaddd	%ymm0, %ymm2, %ymm0
	LONG $0xc0fee5c5                           // vpaddd	%ymm0, %ymm3, %ymm0
	LONG $0x397de3c4; WORD $0x01c1             // vextracti128	$1, %ymm0, %xmm1
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0xc870f9c5; BYTE $0xee               // vpshufd	$238, %xmm0, %xmm1              # xmm1 = xmm0[2,3,2,3]
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0xc870f9c5; BYTE $0x55               // vpshufd	$85, %xmm0, %xmm1               # xmm1 = xmm0[1,1,1,1]
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0xc07ef9c5                           // vmovd	%xmm0, %eax
	WORD $0x394d; BYTE $0xd1                   // cmpq	%r10, %r9
	JE   LBB0_24

LBB0_6:
	WORD $0xd989 // movl	%ebx, %ecx
	WORD $0xf631 // xorl	%esi, %esi

LBB0_7:
	LONG $0x107ac1c4; WORD $0xb004             // vmovss	(%r8,%rsi,4), %xmm0             # xmm0 = mem[0],zero,zero,zero
	LONG $0xc27ac1c4; WORD $0xb304; BYTE $0x04 // vcmpneqss	(%r11,%rsi,4), %xmm0, %xmm0
	LONG $0xc77ef9c5                           // vmovd	%xmm0, %edi
	WORD $0xf829                               // subl	%edi, %eax
	LONG $0x01c68348                           // addq	$1, %rsi
	WORD $0xf139                               // cmpl	%esi, %ecx
	JNE  LBB0_7
	JMP  LBB0_24

LBB0_9:
	LONG $0xc9eff1c5 // vpxor	%xmm1, %xmm1, %xmm1
	LONG $0xd2efe9c5 // vpxor	%xmm2, %xmm2, %xmm2
	LONG $0xdbefe1c5 // vpxor	%xmm3, %xmm3, %xmm3

LBB0_10:
	LONG $0x2610fcc5             // vmovups	(%rsi), %ymm4
	LONG $0x27c2dcc5; BYTE $0x0c // vcmpneq_oqps	(%rdi), %ymm4, %ymm4
	LONG $0xdcfae5c5             // vpsubd	%ymm4, %ymm3, %ymm3
	WORD $0xc383; BYTE $0xf8     // addl	$-8, %ebx
	LONG $0x20c78348             // addq	$32, %rdi
	LONG $0x20c68348             // addq	$32, %rsi
	WORD $0xfb83; BYTE $0x07     // cmpl	$7, %ebx
	JA   LBB0_10

LBB0_12:
	WORD $0xdb85             // testl	%ebx, %ebx
	JE   LBB0_13
	LONG $0xff4b8d44         // leal	-1(%rbx), %r9d
	WORD $0xc931             // xorl	%ecx, %ecx
	LONG $0x1ff98341         // cmpl	$31, %r9d
	JAE  LBB0_18
	WORD $0x8949; BYTE $0xf8 // movq	%rdi, %r8
	WORD $0x8949; BYTE $0xf2 // movq	%rsi, %r10
	JMP  LBB0_21

LBB0_13:
	WORD $0xc931 // xorl	%ecx, %ecx
	JMP  LBB0_23

LBB0_18:
	LONG $0x01c18349         // addq	$1, %r9
	WORD $0x894d; BYTE $0xcb // movq	%r9, %r11
	LONG $0xe0e38349         // andq	$-32, %r11
	LONG $0x9f048d4e         // leaq	(%rdi,%r11,4), %r8
	LONG $0x9e148d4e         // leaq	(%rsi,%r11,4), %r10
	WORD $0x2944; BYTE $0xdb // subl	%r11d, %ebx
	LONG $0xe4efd9c5         // vpxor	%xmm4, %xmm4, %xmm4
	WORD $0xc031             // xorl	%eax, %eax
	LONG $0xedefd1c5         // vpxor	%xmm5, %xmm5, %xmm5
	LONG $0xf657c8c5         // vxorps	%xmm6, %xmm6, %xmm6
	LONG $0xff57c0c5         // vxorps	%xmm7, %xmm7, %xmm7

LBB0_19:
	LONG $0x04107cc5; BYTE $0x86               // vmovups	(%rsi,%rax,4), %ymm8
	LONG $0x4c107cc5; WORD $0x2086             // vmovups	32(%rsi,%rax,4), %ymm9
	LONG $0x54107cc5; WORD $0x4086             // vmovups	64(%rsi,%rax,4), %ymm10
	LONG $0x5c107cc5; WORD $0x6086             // vmovups	96(%rsi,%rax,4), %ymm11
	LONG $0x04c23cc5; WORD $0x0487             // vcmpneqps	(%rdi,%rax,4), %ymm8, %ymm8
	LONG $0xfa5dc1c4; BYTE $0xe0               // vpsubd	%ymm8, %ymm4, %ymm4
	LONG $0x44c234c5; WORD $0x2087; BYTE $0x04 // vcmpneqps	32(%rdi,%rax,4), %ymm9, %ymm8
	LONG $0xfa55c1c4; BYTE $0xe8               // vpsubd	%ymm8, %ymm5, %ymm5
	LONG $0x44c22cc5; WORD $0x4087; BYTE $0x04 // vcmpneqps	64(%rdi,%rax,4), %ymm10, %ymm8
	LONG $0x4cc224c5; WORD $0x6087; BYTE $0x04 // vcmpneqps	96(%rdi,%rax,4), %ymm11, %ymm9
	LONG $0xfa4dc1c4; BYTE $0xf0               // vpsubd	%ymm8, %ymm6, %ymm6
	LONG $0xfa45c1c4; BYTE $0xf9               // vpsubd	%ymm9, %ymm7, %ymm7
	LONG $0x20c08348                           // addq	$32, %rax
	WORD $0x3949; BYTE $0xc3                   // cmpq	%rax, %r11
	JNE  LBB0_19
	LONG $0xe4fed5c5                           // vpaddd	%ymm4, %ymm5, %ymm4
	LONG $0xe4fecdc5                           // vpaddd	%ymm4, %ymm6, %ymm4
	LONG $0xe4fec5c5                           // vpaddd	%ymm4, %ymm7, %ymm4
	LONG $0x397de3c4; WORD $0x01e5             // vextracti128	$1, %ymm4, %xmm5
	LONG $0xe5fed9c5                           // vpaddd	%xmm5, %xmm4, %xmm4
	LONG $0xec70f9c5; BYTE $0xee               // vpshufd	$238, %xmm4, %xmm5              # xmm5 = xmm4[2,3,2,3]
	LONG $0xe5fed9c5                           // vpaddd	%xmm5, %xmm4, %xmm4
	LONG $0xec70f9c5; BYTE $0x55               // vpshufd	$85, %xmm4, %xmm5               # xmm5 = xmm4[1,1,1,1]
	LONG $0xe5fed9c5                           // vpaddd	%xmm5, %xmm4, %xmm4
	LONG $0xe17ef9c5                           // vmovd	%xmm4, %ecx
	WORD $0x394d; BYTE $0xd9                   // cmpq	%r11, %r9
	JE   LBB0_23

LBB0_21:
	WORD $0xd889 // movl	%ebx, %eax
	WORD $0xf631 // xorl	%esi, %esi

LBB0_22:
	LONG $0x107ac1c4; WORD $0xb224             // vmovss	(%r10,%rsi,4), %xmm4            # xmm4 = mem[0],zero,zero,zero
	LONG $0xc25ac1c4; WORD $0xb024; BYTE $0x04 // vcmpneqss	(%r8,%rsi,4), %xmm4, %xmm4
	LONG $0xe77ef9c5                           // vmovd	%xmm4, %edi
	WORD $0xf929                               // subl	%edi, %ecx
	LONG $0x01c68348                           // addq	$1, %rsi
	WORD $0xf039                               // cmpl	%esi, %eax
	JNE  LBB0_22

LBB0_23:
	LONG $0xcafef5c5               // vpaddd	%ymm2, %ymm1, %ymm1
	LONG $0xc0fef5c5               // vpaddd	%ymm0, %ymm1, %ymm0
	LONG $0xc3fefdc5               // vpaddd	%ymm3, %ymm0, %ymm0
	LONG $0x027de2c4; BYTE $0xc0   // vphaddd	%ymm0, %ymm0, %ymm0
	LONG $0x027de2c4; BYTE $0xc0   // vphaddd	%ymm0, %ymm0, %ymm0
	LONG $0x397de3c4; WORD $0x01c1 // vextracti128	$1, %ymm0, %xmm1
	LONG $0xc0fef1c5               // vpaddd	%xmm0, %xmm1, %xmm0
	LONG $0xc07ef9c5               // vmovd	%xmm0, %eax
	WORD $0xc801                   // addl	%ecx, %eax

LBB0_24:
	LONG $0xc02a9ac5         // vcvtsi2ss	%eax, %xmm12, %xmm0
	LONG $0x0211fac5         // vmovss	%xmm0, (%rdx)
	LONG $0xf8658d48         // leaq	-8(%rbp), %rsp
	BYTE $0x5b               // popq	%rbx
	BYTE $0x5d               // popq	%rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // retq
