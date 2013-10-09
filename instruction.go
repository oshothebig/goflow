package goflow

type InstructionType uint16

// corresponds to ofp_instruction_type
const (
	OFPIT_GOTO_TABLE InstructionType = iota
	OFPIT_WRITE_METADATA
	OFPIT_WRITE_ACTIONS
	OFPIT_APPLY_ACTIONS
	OFPIT_CLEAR_ACTIONS
	OFPIT_METER
	OFPIT_EXPERIMENTER = 0xffff
)

type Instruction interface {
	Packetizable
	Header() *InstructionHeader
}

type InstructionHeader struct {
	Type   InstructionType
	Length uint16
}

func (h *InstructionHeader) Header() *InstructionHeader {
	return h
}

type GoToTable struct {
	InstructionHeader
	TableId uint8
	pad     [3]uint8
}

type WriteMetadata struct {
	InstructionHeader
	pad      [4]uint8
	Metadata Metadata
	Mask     MetadataMask
}

type InstructionActions struct {
	InstructionHeader
	pad     [4]uint8
	Actions []Action
}

type WriteActions struct {
	InstructionActions
}

type ApplyActions struct {
	InstructionActions
}

type ClearActions struct {
	InstructionActions
}

type Meter struct {
	InstructionHeader
	MeterId uint32
}

// TODO: implement ofp_instruction_experimenter
