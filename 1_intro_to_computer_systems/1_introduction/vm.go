package __introduction

import "fmt"

const (
	Load  = 0x01 // load    r1  addr    # Load value at given address into given register
	Store = 0x02 // store   r2  addr    # Store the value in register at the given memory address
	Add   = 0x03 // add     r1  r2      # Set r1 = r1 + r2
	Sub   = 0x04 // sub     r1  r2      # Set r1 = r1 - r2
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Address of program counter
const PcAddr = 0x00

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {
	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
	for {
		pc := registers[PcAddr]
		op := memory[pc]

		// decode and execute
		switch op {
		case Load:
			r1, addr := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			registers[r1] = memory[addr]
		case Store:
			r1, addr := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			// Only allow writing to data, prevent writing instructions
			if addr < 8 {
				memory[addr] = registers[r1]
			}
		case Add:
			r1, r2 := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			registers[r1] += registers[r2]
		case Sub:
			r1, r2 := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			registers[r1] -= registers[r2]
		case Addi:
			r1, val := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			registers[r1] += val
		case Subi:
			r1, val := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			registers[r1] -= val
		case Jump:
			addr := memory[pc+1]

			registers[PcAddr] = addr
		case Beqz:
			r1, offset := memory[pc+1], memory[pc+2]

			registers[PcAddr] += 3

			if registers[r1] == 0 {
				registers[PcAddr] += offset
			}
		case Halt:
			registers[PcAddr] += 1
			return
		default:
			panic(fmt.Sprintf("unknown instruction %x", op))
		}
	}
}
