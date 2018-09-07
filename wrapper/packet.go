package wrapper

type Packet struct {
	Freq       uint32  // central frequency of the IF chain (in Hz)
	IFChain    uint8   // by which IF chain was packet received
	Status     uint8   // status of the received Packet
	CountUS    uint32  // internal concentrator counter for timestamping, 1 microsecond resolution
	RFChain    uint8   // by which RF chain was packet received
	Modulation uint8   // modulation used by the packet
	Bandwidth  uint8   // modulation bandwidth (LoRa only)
	Datarate   uint32  // RX datarate of the packet (SF for LoRa)
	Coderate   uint8   // error-correcting code of the packet (LoRa only)
	RSSI       float32 // average packet RSSI in dB
	SNR        float32 // average packet SNR, in dB (LoRa only)
	MinSNR     float32 // minimum packet SNR, in dB (LoRa only)
	MaxSNR     float32 // maximum packet SNR, in dB (LoRa only)
	CRC        uint16  // CRC that was received in the payload
	Size       uint32  // Payload size in bytes
	Payload    []byte  // Buffer containing the payload, not yet base64-encoded
}
