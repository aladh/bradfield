### 1.1 

9 = 0x9

136 = 0x88

247 = 0xf7

### 1.2

Hex form: 6 digits = 3 bytes = 24 bits = 2^24 = 16,777,216

RGB form: 256^3 = 16,777,216

### hellohex

17 bytes = 34 hexadecimal chars

First five bytes: 0b01101000 0b01100101 0b01101100 0b01101100 0b01101111

### 2.1

4 = 0b100

65 = 0b1000001

105 = 0b1101001

255 = 11111111

0b10 = 2

0b11 = 3

0b1101100 = 108  

0b1010101 = 85 

### 2.2

0b11111111 + 0b00001101 = 0b100001100

With 8-bit registers, the value returned is 0b00001100. Referred to as overflow.

### 2.3

127 = 0b01111111

-128 = 0b10000000

-1 = 0b11111111

1 = 0b00000001

-14 = 0b11110010

0b10000011 = -128 + 3 = -125
0b11000100 = -128 + 64 + 4 = -60

### 2.4

0b01111111 + 0b10000000 = 0b11111111 = -1

To negate a number in two's complement, flip all the bits and add one to it.

To compute subtraction of two's complement numbers, negate the value being subtracted and add the numbers together.

The value of the most significant bit in 8-bit two's complement is -128. In 32-bit, it's -2^31 (-2,147,483,648).

### 2.5

An overflow has occurred if, for the most significant bit, the carry-in is different from the carry-out.

### 3.1

9001 = 0b00100011 00101001 in big endian.

### 3.2

Sequence: 0x441e 7368 = 1,142,846,312

Acknowledgement: 0xeff2 a002 = 4,025,655,298

Source port: 0xaf00 = 44,800

Destination port: 0xbc06 = 48,134

Header size: 0x8 = 8 32-bit words = 256 bits = 32 bytes

Optional fields start at the twentieth byte, so there are 12 bytes of optional data.

### 3.3

Variant: BM - Windows 3.1x, 95, NT, ... etc.

`xxd -c4 -e -s 14 -l 100 image1.bmp`

image1: 24 (0x18) x 48 (0x30)

image2: 32 (0x20) x 64 (0x40)

`xxd -c4 -e -s 28 -l 2 image2.bmp`

0x18 = 24 bits per pixel

`xxd -c4 -e -s 10 -l 4 image1.bmp`

The image data starts at 0x8a = 138 bytes

`xxd -c1 -e -s 138 -l 100 image1.bmp`

### 4.1

0 10000100 01010100000000000000000 = + 2^(132-127) * (1 + 2^-2 + 2^-4 + 2^-6) = 2^5 * 1.328125 = 42.5

The smallest change on the largest exponent is (1 + 2^-23) * 2^127 = 2^104

The smallest change that can be made is (1 + 2^-23) * 2^-126 = 2^-149

### 5.1

It takes 3 bytes to encode a snowman (U+2603).

0b1110 0010 1001 1000 1000 0011

Hex: 0xe2 98 83 

### 5.2

hello 

0b11110000 10011111 10011000 10000000 -> 0b00000001 11110110 00000000 = 0x1F600 = 😀

### 5.3

`echo "\a"`
