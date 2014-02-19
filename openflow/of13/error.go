package of13

import . "github.com/oshothebig/goflow/openflow"

type ErrorType uint16
type ErrorCode uint16

type Error struct {
	Header
	Type ErrorType
	Code ErrorCode
	Data []uint8
}

func NewError(typ ErrorType, code ErrorCode, data []uint8) *Error {
	m := &Error{*NewHeader(MessageTypes.Error), typ, code, newData(data)}
	m.Length = m.Length + 4 + uint16(len(data))

	return m
}

const (
	OFPET_HELLO_FAILED ErrorType = iota
	OFPET_BAD_REQUEST
	OFPET_BAD_ACTION
	OFPET_BAD_INSTRUCTION
	OFPET_BAD_MATCH
	OFPET_FLOW_MOD_FAILED
	OFPET_GROUP_MOD_FAILED
	OFPET_PORT_MOD_FAILED
	OFPET_TABLE_MOD_FAILED
	OFPET_QUEUE_OP_FAILED
	OFPET_SWITCH_CONFIG_FAILED
	OFPET_ROLE_REQUEST_FAILED
	OFPET_METER_MOD_FAILED
	OFPET_TABLE_FEATURES_FAILED
	OFPET_EXPERIMENTER ErrorType = 0xffff
)

var ErrorTypes = struct {
	HelloFailed         ErrorType
	BadRequest          ErrorType
	BadAction           ErrorType
	BadInstruction      ErrorType
	BadMatch            ErrorType
	FlowModFailed       ErrorType
	GroupModFailed      ErrorType
	PortModFailed       ErrorType
	TableModFailed      ErrorType
	QueueOpFailed       ErrorType
	SwitchConfigFailed  ErrorType
	RoleRequestFailed   ErrorType
	MeterModFailed      ErrorType
	TableFeaturesFailed ErrorType
	Experimenter        ErrorType
}{
	OFPET_HELLO_FAILED,
	OFPET_BAD_REQUEST,
	OFPET_BAD_ACTION,
	OFPET_BAD_INSTRUCTION,
	OFPET_BAD_MATCH,
	OFPET_FLOW_MOD_FAILED,
	OFPET_GROUP_MOD_FAILED,
	OFPET_PORT_MOD_FAILED,
	OFPET_TABLE_MOD_FAILED,
	OFPET_QUEUE_OP_FAILED,
	OFPET_SWITCH_CONFIG_FAILED,
	OFPET_ROLE_REQUEST_FAILED,
	OFPET_METER_MOD_FAILED,
	OFPET_TABLE_FEATURES_FAILED,
	OFPET_EXPERIMENTER,
}

const (
	OFPBRC_BAD_VERSION ErrorCode = iota
	OFPBRC_BAD_TYPE
	OFPBRC_BAD_MULTIPART
	OFPBRC_BAD_EXPERIMENTER
	OFPBRC_BAD_EXP_TYPE
	OFPBRC_EPERM
	OFPBRC_BAD_LEN
	OFPBRC_BUFFER_EMPTY
	OFPBRC_BUFFER_UNKNOWN
	OFPBRC_BAD_TABLE_ID
	OFPBRC_IS_SLAVE
	OFPBRC_BAD_PORT
	OFPBRC_BAD_PACKET
	OFPBRC_MULTIPART_BUFFER_OVERFLOW
)

var BadRequestCodes = struct {
	BadVersion              ErrorCode
	BadType                 ErrorCode
	BadMultipart            ErrorCode
	BadExperimenter         ErrorCode
	BadExperimenterType     ErrorCode
	PermissionError         ErrorCode
	BadLength               ErrorCode
	BufferEmpty             ErrorCode
	BufferUnknown           ErrorCode
	BadTableId              ErrorCode
	IsSlave                 ErrorCode
	BadPort                 ErrorCode
	BadPacket               ErrorCode
	MultipartBufferOverflow ErrorCode
}{
	OFPBRC_BAD_VERSION,
	OFPBRC_BAD_TYPE,
	OFPBRC_BAD_MULTIPART,
	OFPBRC_BAD_EXPERIMENTER,
	OFPBRC_BAD_EXP_TYPE,
	OFPBRC_EPERM,
	OFPBRC_BAD_LEN,
	OFPBRC_BUFFER_EMPTY,
	OFPBRC_BUFFER_UNKNOWN,
	OFPBRC_BAD_TABLE_ID,
	OFPBRC_IS_SLAVE,
	OFPBRC_BAD_PORT,
	OFPBRC_BAD_PACKET,
	OFPBRC_MULTIPART_BUFFER_OVERFLOW,
}

const (
	OFPBAC_BAD_TYPE ErrorCode = iota
	OFPBAC_BAD_LEN
	OFPBAC_BAD_EXPERIMENTER
	OFPBAC_BAD_EXP_TYPE
	OFPBAC_BAD_OUT_PORT
	OFPBAC_BAD_ARGUMENT
	OFPBAC_EPERM
	OFPBAC_TOO_MANY
	OFPBAC_BAD_QUEUE
	OFPBAC_BAD_OUT_GROUP
	OFPBAC_MATCH_INCONSISTENT
	OFPBAC_UNSUPPORTED_ORDER
	OFPBAC_BAD_TAG
	OFPBAC_BAD_SET_TYPE
	OFPBAC_BAD_SET_LEN
	OFPBAC_BAD_SET_ARGUMENT
)

var BadActionCodes = struct {
	BadType             ErrorCode
	BadLength           ErrorCode
	BadExperimenter     ErrorCode
	BadExperimenterType ErrorCode
	BadOutPort          ErrorCode
	BadArgument         ErrorCode
	PermissionError     ErrorCode
	TooMany             ErrorCode
	BadQueue            ErrorCode
	BadOutGroup         ErrorCode
	MatchInconsistent   ErrorCode
	UnsupportedOrder    ErrorCode
	BadTag              ErrorCode
	BadSetType          ErrorCode
	BadSetLength        ErrorCode
	BadSetArgument      ErrorCode
}{
	OFPBAC_BAD_TYPE,
	OFPBAC_BAD_LEN,
	OFPBAC_BAD_EXPERIMENTER,
	OFPBAC_BAD_EXP_TYPE,
	OFPBAC_BAD_OUT_PORT,
	OFPBAC_BAD_ARGUMENT,
	OFPBAC_EPERM,
	OFPBAC_TOO_MANY,
	OFPBAC_BAD_QUEUE,
	OFPBAC_BAD_OUT_GROUP,
	OFPBAC_MATCH_INCONSISTENT,
	OFPBAC_UNSUPPORTED_ORDER,
	OFPBAC_BAD_TAG,
	OFPBAC_BAD_SET_TYPE,
	OFPBAC_BAD_SET_LEN,
	OFPBAC_BAD_SET_ARGUMENT,
}

const (
	OFPBIC_UNKNOWN_INST ErrorCode = iota
	OFPBIC_UNSUP_INST
	OFPBIC_BAD_TABLE_ID
	OFPBIC_UNSUP_METADATA
	OFPBIC_UNSUP_METADATA_MASK
	OFPBIC_BAD_EXPERIMENTER
	OFPBIC_BAD_EXP_TYPE
	OFPBIC_BAD_LEN
	OFPBIC_EPERM
)

var BadInstructionCodes = struct {
	UnknownInstruction      ErrorCode
	UnsupportedInstruction  ErrorCode
	BadTableId              ErrorCode
	UnsupportedMetadata     ErrorCode
	UnsupportedMetadataMask ErrorCode
	BadExperimenter         ErrorCode
	BadExperimenterType     ErrorCode
	BadLength               ErrorCode
	PermissionError         ErrorCode
}{
	OFPBIC_UNKNOWN_INST,
	OFPBIC_UNSUP_INST,
	OFPBIC_BAD_TABLE_ID,
	OFPBIC_UNSUP_METADATA,
	OFPBIC_UNSUP_METADATA_MASK,
	OFPBIC_BAD_EXPERIMENTER,
	OFPBIC_BAD_EXP_TYPE,
	OFPBIC_BAD_LEN,
	OFPBIC_EPERM,
}

const (
	OFPBMC_BAD_TYPE ErrorCode = iota
	OFPBMC_BAD_LEN
	OFPBMC_BAD_TAG
	OFPBMC_BAD_DL_ADDR_MASK
	OFPBMC_BAD_NW_ADDR_MASK
	OFPBMC_BAD_WILDCARD
	OFPBMC_BAD_FIELD
	OFPBMC_BAD_VALUE
	OFPBMC_BAD_MASK
	OFPBMC_BAD_PREREQ
	OFPBMC_DUP_FIELD
	OFPBMC_EPERM
)

var BadMatchCodes = struct {
	BadType         ErrorCode
	BadLength       ErrorCode
	BadTag          ErrorCode
	BadMacAddrMask  ErrorCode
	BadIpAddrMask   ErrorCode
	BadWildcard     ErrorCode
	BadField        ErrorCode
	BadValue        ErrorCode
	BasMask         ErrorCode
	BadPrerequiste  ErrorCode
	DuplicatedField ErrorCode
	PermissionError ErrorCode
}{
	OFPBMC_BAD_TYPE,
	OFPBMC_BAD_LEN,
	OFPBMC_BAD_TAG,
	OFPBMC_BAD_DL_ADDR_MASK,
	OFPBMC_BAD_NW_ADDR_MASK,
	OFPBMC_BAD_WILDCARD,
	OFPBMC_BAD_FIELD,
	OFPBMC_BAD_VALUE,
	OFPBMC_BAD_MASK,
	OFPBMC_BAD_PREREQ,
	OFPBMC_DUP_FIELD,
	OFPBMC_EPERM,
}

const (
	OFPFMFC_UNKNOWN ErrorCode = iota
	OFPFMFC_TABLE_FULL
	OFPFMFC_BAD_TABLE_ID
	OFPFMFC_OVERLAP
	OFPFMFC_EPERM
	OFPFMFC_BAD_TIMEOUT
	OFPFMFC_BAD_COMMAND
	OFPFMFC_BAD_FLAGS
)

var FlowModFailedCodes = struct {
	Unknown         ErrorCode
	TableFull       ErrorCode
	BadTableId      ErrorCode
	Overlap         ErrorCode
	PermissionError ErrorCode
	BadTimeout      ErrorCode
	BadCommand      ErrorCode
	BadFlags        ErrorCode
}{
	OFPFMFC_UNKNOWN,
	OFPFMFC_TABLE_FULL,
	OFPFMFC_BAD_TABLE_ID,
	OFPFMFC_OVERLAP,
	OFPFMFC_EPERM,
	OFPFMFC_BAD_TIMEOUT,
	OFPFMFC_BAD_COMMAND,
	OFPFMFC_BAD_FLAGS,
}

const (
	OFPGMFC_GROUP_EXISTS ErrorCode = iota
	OFPGMFC_INVALID_GROUP
	OFPGMFC_WEIGHT_UNSUPPORTED
	OFPGMFC_OUT_OF_GROUPS
	OFPGMFC_OUT_OF_BUCKETS
	OFPGMFC_CHAINING_UNSUPPORTED
	OFPGMFC_WATCH_UNSUPPORTED
	OFPGMFC_LOOP
	OFPGMFC_UNKNOWN_GROUP
	OFPGMFC_CHAINED_GROUP
	OFPGMFC_BAD_TYPE
	OFPGMFC_BAD_COMMAND
	OFPGMFC_BAD_BUCKET
	OFPGMFC_BAD_WATCH
	OFPGMFC_EPERM
)

var GroupModFailedCodes = struct {
	GroupExists         ErrorCode
	InvalidGroup        ErrorCode
	WeightUnsupported   ErrorCode
	OutOfGroups         ErrorCode
	OutOfBuckets        ErrorCode
	ChainingUnsupported ErrorCode
	WatchUnsupported    ErrorCode
	Loop                ErrorCode
	UnknownGroup        ErrorCode
	ChainedGroup        ErrorCode
	BadType             ErrorCode
	BadCommand          ErrorCode
	BadBucket           ErrorCode
	BadWatch            ErrorCode
	PermissionError     ErrorCode
}{
	OFPGMFC_GROUP_EXISTS,
	OFPGMFC_INVALID_GROUP,
	OFPGMFC_WEIGHT_UNSUPPORTED,
	OFPGMFC_OUT_OF_GROUPS,
	OFPGMFC_OUT_OF_BUCKETS,
	OFPGMFC_CHAINING_UNSUPPORTED,
	OFPGMFC_WATCH_UNSUPPORTED,
	OFPGMFC_LOOP,
	OFPGMFC_UNKNOWN_GROUP,
	OFPGMFC_CHAINED_GROUP,
	OFPGMFC_BAD_TYPE,
	OFPGMFC_BAD_COMMAND,
	OFPGMFC_BAD_BUCKET,
	OFPGMFC_BAD_WATCH,
	OFPGMFC_EPERM,
}

const (
	OFPPMFC_BAD_PORT ErrorCode = iota
	OFPPMFC_BAD_HW_ADDR
	OFPPMFC_BAD_CONFIG
	OFPPMFC_BAD_ADVERTISE
	OFPPMFC_EPERM
)

var PortModFailedCodes = struct {
	BadPort         ErrorCode
	BadMacAddr      ErrorCode
	BadConfig       ErrorCode
	BadAdvertise    ErrorCode
	PermissionError ErrorCode
}{
	OFPPMFC_BAD_PORT,
	OFPPMFC_BAD_HW_ADDR,
	OFPPMFC_BAD_CONFIG,
	OFPPMFC_BAD_ADVERTISE,
	OFPPMFC_EPERM,
}

const (
	OFPTMFC_BAD_TABLE ErrorCode = iota
	OFPTMFC_BAD_CONFIG
	OFPTMFC_EPERM
)

var TableModFailedCodes = struct {
	BadTable        ErrorCode
	BadConfig       ErrorCode
	PermissionError ErrorCode
}{
	OFPTMFC_BAD_TABLE,
	OFPTMFC_BAD_CONFIG,
	OFPTMFC_EPERM,
}

const (
	OFPQOFC_BAD_PORT ErrorCode = iota
	OFPQOFC_BAD_QUEUE
	OFPQOFC_EPERM
)

var QueueOpFailedCodes = struct {
	BadPort         ErrorCode
	BadQueue        ErrorCode
	PermissionError ErrorCode
}{
	OFPQOFC_BAD_PORT,
	OFPQOFC_BAD_QUEUE,
	OFPQOFC_EPERM,
}

const (
	OFPSCFC_BAD_FLAGS ErrorCode = iota
	OFPSCFC_BAD_LEN
	OFPSCFC_EPERM
)

var SwitchConfigFailedCodes = struct {
	BadFlags        ErrorCode
	BadLength       ErrorCode
	PermissionError ErrorCode
}{
	OFPSCFC_BAD_FLAGS,
	OFPSCFC_BAD_LEN,
	OFPSCFC_EPERM,
}

const (
	OFPRRFC_STALE ErrorCode = iota
	OFPRRFC_UNSUP
	OFPRRFC_BAD_ROLE
)

var RoleRequestFailedCodes = struct {
	Stale       ErrorCode
	Unsupported ErrorCode
	BadRole     ErrorCode
}{
	OFPRRFC_STALE,
	OFPRRFC_UNSUP,
	OFPRRFC_BAD_ROLE,
}

const (
	OFPMMFC_UNKNOWN ErrorCode = iota
	OFPMMFC_METER_EXISTS
	OFPMMFC_INVALID_METER
	OFPMMFC_UNKNOWN_METER
	OFPMMFC_BAD_COMMAND
	OFPMMFC_BAD_FLAGS
	OFPMMFC_BAD_RATE
	OFPMMFC_BAD_BURST
	OFPMMFC_BAD_BAND
	OFPMMFC_BAD_BAND_VALUE
	OFPMMFC_OUT_OF_METERS
	OFPMMFC_OUT_OF_BANDS
)

var MeterModFailedCodes = struct {
	Unknown      ErrorCode
	MeterExists  ErrorCode
	InvalidMeter ErrorCode
	UnknownMeter ErrorCode
	BadCommand   ErrorCode
	BadFlags     ErrorCode
	BadRate      ErrorCode
	BadBurst     ErrorCode
	BadBand      ErrorCode
	BadBandValue ErrorCode
	OutOfMeters  ErrorCode
	OutOfBands   ErrorCode
}{
	OFPMMFC_UNKNOWN,
	OFPMMFC_METER_EXISTS,
	OFPMMFC_INVALID_METER,
	OFPMMFC_UNKNOWN_METER,
	OFPMMFC_BAD_COMMAND,
	OFPMMFC_BAD_FLAGS,
	OFPMMFC_BAD_RATE,
	OFPMMFC_BAD_BURST,
	OFPMMFC_BAD_BAND,
	OFPMMFC_BAD_BAND_VALUE,
	OFPMMFC_OUT_OF_METERS,
	OFPMMFC_OUT_OF_BANDS,
}

const (
	OFPTFFC_BAD_TABLE ErrorCode = iota
	OFPTFFC_BAD_METADATA
	OFPTFFC_BAD_TYPE
	OFPTFFC_BAD_LEN
	OFPTFFC_BAD_ARGUMENT
	OFPTFFC_EPERM
)

var TableFeaturesFailedCode = struct {
	BadTable        ErrorCode
	BadMetadata     ErrorCode
	BadType         ErrorCode
	BadLength       ErrorCode
	BadArgument     ErrorCode
	PermissionError ErrorCode
}{
	OFPTFFC_BAD_TABLE,
	OFPTFFC_BAD_METADATA,
	OFPTFFC_BAD_TYPE,
	OFPTFFC_BAD_LEN,
	OFPTFFC_BAD_ARGUMENT,
	OFPTFFC_EPERM,
}

type ErrorExperimenter struct {
	Header
	Type            ErrorType
	ExperimenerType uint16
	Experimenter    uint32
	Data            []uint8
}
