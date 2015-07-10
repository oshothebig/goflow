package of10

const WireProtocolVersion = 0x01

var generateXid func() uint32 = NewXidGenerator()
