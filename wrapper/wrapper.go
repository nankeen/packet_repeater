package wrapper

// #cgo CFLAGS: -I${SRCDIR}/../../lora_gateway/libloragw/inc
// #cgo LDFLAGS: -lm ${SRCDIR}/../../lora_gateway/libloragw/libloragw.a -lrt
// #include "config.h"
// #include "loragw_hal.h"
// #include "loragw_gps.h"
// void setType(struct lgw_conf_rxrf_s *rxrfConf, enum lgw_radio_type_e val) {
// 	rxrfConf->type = val;
// }
import "C"
import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

const NbMaxPackets = 8

// Lock to prevent concentrator conflict
var concentratorMutex = &sync.Mutex{}

var loraChannelBandwidths = map[uint32]C.uint8_t{
	7800:   C.BW_7K8HZ,
	15600:  C.BW_15K6HZ,
	31200:  C.BW_31K2HZ,
	62500:  C.BW_62K5HZ,
	125000: C.BW_125KHZ,
	250000: C.BW_250KHZ,
	500000: C.BW_500KHZ,
}

var loraChannelSpreadingFactors = map[uint32]C.uint32_t{
	7:  C.DR_LORA_SF7,
	8:  C.DR_LORA_SF8,
	9:  C.DR_LORA_SF9,
	10: C.DR_LORA_SF10,
	11: C.DR_LORA_SF11,
	12: C.DR_LORA_SF12,
}

/*
============================
|                          |
|      INITIALIZATION      |
|                          |
============================
*/

func SetBoardConf(clksrc uint, lorawan_public bool) error {
	var boardConf = C.struct_lgw_conf_board_s{
		clksrc:         C.uint8_t(clksrc),
		lorawan_public: C.bool(lorawan_public),
	}

	if C.lgw_board_setconf(boardConf) != C.LGW_HAL_SUCCESS {
		return errors.New("Board configuration failed")
	}
	return nil
}

// StartLoRaGateway wraps the HAL function to start the concentrator once configured
func StartLoRaGateway() error {
	state := C.lgw_start()

	if state != C.LGW_HAL_SUCCESS {
		return errors.New("Failed to start concentrator")
	}
	return nil
}

// StopLoRaGateway wraps the HAL function to stop the concentrator once started
func StopLoRaGateway() error {
	state := C.lgw_stop()

	if state != C.LGW_HAL_SUCCESS {
		return errors.New("Failed to stop concentrator gracefully")
	}
	return nil
}

/*
============================
|                          |
|       TX GAIN LUT        |
|                          |
============================
*/

func prepareTXLut(txLut *C.struct_lgw_tx_gain_s, txConf GainTableConf) {
	if txConf.DacGain != nil {
		txLut.dac_gain = C.uint8_t(*txConf.DacGain)
	} else {
		txLut.dac_gain = 3
	}
	txLut.pa_gain = C.uint8_t(txConf.PaGain)
	txLut.mix_gain = C.uint8_t(txConf.MixGain)
	txLut.rf_power = C.int8_t(txConf.RfPower)
	txLut.dig_gain = C.uint8_t(txConf.DigGain)
}

// SetTXGainConf prepares, and then sends the configuration of the TX Gain LUT to the concentrator
func SetTXGainConf() error {
	txLuts := GetLuts()
	var gainLut = C.struct_lgw_tx_gain_lut_s{
		size: 0,
		lut:  [C.TX_GAIN_LUT_SIZE_MAX]C.struct_lgw_tx_gain_s{},
	}
	for i, txLut := range txLuts {
		prepareTXLut(&gainLut.lut[i], txLut)
	}
	gainLut.size = C.uint8_t(len(txLuts))

	if C.lgw_txgain_setconf(&gainLut) != C.LGW_HAL_SUCCESS {
		return errors.New("Failed to configure concentrator TX Gain LUT")
	}
	return nil
}

/*
============================
|                          |
| RF CHANNEL CONFIGURATION |
|                          |
============================
*/

// initRadio initiates a radio configuration in the C.struct_lgw_conf_rxrf_s format, given
// the configuration for that radio.
func initRadio(radio RadioConf) (C.struct_lgw_conf_rxrf_s, error) {
	var cRadio = C.struct_lgw_conf_rxrf_s{
		enable:      C.bool(radio.Enabled),
		freq_hz:     C.uint32_t(radio.Freq),
		rssi_offset: C.float(radio.RssiOffset),
		tx_enable:   C.bool(radio.TxEnabled),
	}

	// Checking the radio is of a pre-defined type
	switch radio.RadioType {
	case "SX1257":
		C.setType(&cRadio, C.LGW_RADIO_TYPE_SX1257)
	case "SX1255":
		C.setType(&cRadio, C.LGW_RADIO_TYPE_SX1255)
	default:
		return cRadio, errors.New("Invalid radio type (should be SX1255 or SX1257)")
	}
	return cRadio, nil
}

// enableRadio is enabling the radio
func enableRadio(radio RadioConf, nb uint8) error {
	// Checking if radio is enabled and thus needs to be activated
	if !radio.Enabled {
		return nil
	}

	cRadio, err := initRadio(radio)
	if err != nil {
		return err
	}

	if C.lgw_rxrf_setconf(C.uint8_t(nb), cRadio) != C.LGW_HAL_SUCCESS {
		return errors.New("Radio configuration failed")
	}

	return nil
}

// SetRFChannels send the configuration of the radios to the concentrator
func SetRFChannels() error {
	for i, radio := range GetRFConfs() {
		err := enableRadio(radio, uint8(i))
		if err != nil {
			return err
		}
	}
	return nil
}

/*
============================
|                          |
| SF CHANNEL CONFIGURATION |
|                          |
============================
*/

func enableSFChannel(channelConf ChannelConf, nb uint8) error {
	if !channelConf.Enabled {
		return nil
	}

	var cChannel = C.struct_lgw_conf_rxif_s{
		enable:   C.bool(channelConf.Enabled),
		rf_chain: C.uint8_t(channelConf.Radio),
		freq_hz:  C.int32_t(channelConf.IfValue),
	}

	if C.lgw_rxif_setconf(C.uint8_t(nb), cChannel) != C.LGW_HAL_SUCCESS {
		return errors.New(fmt.Sprintf("Missing configuration for SF channel %d", nb))
	}
	return nil
}

// SetSFChannels enables the different SF channels
func SetSFChannels() error {
	for i, sfChannel := range GetMultiSFChannels() {
		err := enableSFChannel(sfChannel, uint8(i))
		if err != nil {
			return err
		}
	}
	return nil
}

/*
============================
|                          |
|       STD AND FSK        |
|   CHANNEL CONFIGURATION  |
|                          |
============================
*/

// initLoRaStdChannel initiates a C.struct_lgw_conf_rxif_s from a LoRaChannelConf
func initLoRaStdChannel(stdChan ChannelConf) C.struct_lgw_conf_rxif_s {
	var cChannel = C.struct_lgw_conf_rxif_s{
		enable:   C.bool(stdChan.Enabled),
		rf_chain: C.uint8_t(stdChan.Radio),
		freq_hz:  C.int32_t(stdChan.IfValue),
	}

	switch *stdChan.Bandwidth {
	case 125000, 250000, 500000:
		cChannel.bandwidth = loraChannelBandwidths[*stdChan.Bandwidth]
	default:
		cChannel.bandwidth = C.BW_UNDEFINED
	}

	if stdChan.Datarate != nil && *stdChan.Datarate >= 7 && *stdChan.Datarate <= 12 {
		cChannel.datarate = loraChannelSpreadingFactors[*stdChan.Datarate]
	} else {
		cChannel.datarate = C.DR_UNDEFINED
	}

	return cChannel
}

// SetStandardChannel enables the LoRa standard channel from the configuration
func SetStandardChannel(stdChan ChannelConf) error {
	if !stdChan.Enabled {
		return nil
	}

	var cChannel = initLoRaStdChannel(stdChan)

	if C.lgw_rxif_setconf(8, cChannel) != C.LGW_HAL_SUCCESS {
		return errors.New("Configuration for LoRa standard channel failed")
	}
	return nil
}

// SetFSKChannel sets the FSK Channel configuration on the concentrator
func SetFSKChannel(fskChan ChannelConf) error {
	if !fskChan.Enabled {
		return nil
	}

	var cFSKChan = C.struct_lgw_conf_rxif_s{
		enable:    C.bool(fskChan.Enabled),
		rf_chain:  C.uint8_t(fskChan.Radio),
		freq_hz:   C.int32_t(fskChan.IfValue),
		bandwidth: C.BW_UNDEFINED,
	}
	if fskChan.Datarate != nil {
		cFSKChan.datarate = C.uint32_t(*fskChan.Datarate)
	}
	if fskChan.Bandwidth == nil {
		return errors.New("No bandwidth information in the configuration for the FSK channel - cannot retransmit the FSK packet")
	}

	val := *fskChan.Bandwidth
	switch {
	case val > 0 && val <= 7800:
		cFSKChan.bandwidth = loraChannelBandwidths[7800]
	case val > 7800 && val <= 15600:
		cFSKChan.bandwidth = loraChannelBandwidths[15600]
	case val > 15600 && val <= 31200:
		cFSKChan.bandwidth = loraChannelBandwidths[31200]
	case val > 31200 && val <= 62500:
		cFSKChan.bandwidth = loraChannelBandwidths[62500]
	case val > 62500 && val <= 125000:
		cFSKChan.bandwidth = loraChannelBandwidths[125000]
	case val > 125000 && val <= 250000:
		cFSKChan.bandwidth = loraChannelBandwidths[250000]
	case val > 250000 && val <= 500000:
		cFSKChan.bandwidth = loraChannelBandwidths[500000]
	}

	if C.lgw_rxif_setconf(9, cFSKChan) != C.LGW_HAL_SUCCESS {
		return errors.New("Configuration for FSK channel failed")
	}
	return nil
}

/*
============================
|                          |
|     Packet Handling      |
|                          |
============================
*/

func packetsFromCPackets(cPackets [8]C.struct_lgw_pkt_rx_s, nbPackets int) []Packet {
	var packets = make([]Packet, nbPackets)
	for i := 0; i < nbPackets && i < 8; i++ {
		packets[i] = packetFromCPacket(cPackets[i])
	}
	return packets
}

func packetFromCPacket(cPacket C.struct_lgw_pkt_rx_s) Packet {
	// When using packetFromCPacket, it is assumed that accessing gpsTimeReferenceMutex
	// is safe => Use gpsTimeReferenceMutex before calling packetFromCPacket /before/
	// using this function
	var p = Packet{
		Freq:       uint32(cPacket.freq_hz),
		IFChain:    uint8(cPacket.if_chain),
		Status:     uint8(cPacket.status),
		CountUS:    uint32(cPacket.count_us),
		RFChain:    uint8(cPacket.rf_chain),
		Modulation: uint8(cPacket.modulation),
		Bandwidth:  uint8(cPacket.bandwidth),
		Datarate:   uint32(cPacket.datarate),
		Coderate:   uint8(cPacket.coderate),
		RSSI:       float32(cPacket.rssi),
		SNR:        float32(cPacket.snr),
		MinSNR:     float32(cPacket.snr_min),
		MaxSNR:     float32(cPacket.snr_max),
		CRC:        uint16(cPacket.crc),
		Size:       uint32(cPacket.size),
	}

	p.Payload = make([]byte, p.Size)
	var i uint32
	for i = 0; i < p.Size; i++ {
		p.Payload[i] = byte(cPacket.payload[i])
	}

	return p
}

func Receive() ([]Packet, error) {
	var packets [NbMaxPackets]C.struct_lgw_pkt_rx_s
	concentratorMutex.Lock()
	nbPackets := C.lgw_receive(NbMaxPackets, &packets[0])
	concentratorMutex.Unlock()
	if nbPackets == C.LGW_HAL_ERROR {
		return nil, errors.New("Failed packet fetch from the concentrator")
	}
	return packetsFromCPackets(packets, int(nbPackets)), nil
}

/*
============================
|                          |
|       Broadcasting       |
|                          |
============================
*/

func insertPayload(pkt Packet, txPkt *C.struct_lgw_pkt_tx_s) error {
	payload := pkt.Payload
	if len(payload) > 256 {
		return errors.New("Payload too big to transmit")
	}
	txPkt.size = C.uint16_t(len(payload))
	for i := 0; i < len(payload); i++ {
		txPkt.payload[i] = C.uint8_t(payload[i])
	}
	return nil
}

func SendPacket(pkt Packet) error {
	var txPacket = C.struct_lgw_pkt_tx_s{
		freq_hz: C.uint32_t(pkt.Freq),
		// rf_chain:   C.uint8_t(pkt.RFChain),
		no_crc:    C.bool(false),
		no_header: C.bool(false),
		payload:   [256]C.uint8_t{},
		tx_mode:   C.IMMEDIATE,
		count_us:  C.uint32_t(pkt.CountUS),
		bandwidth: C.uint8_t(pkt.Bandwidth),
		datarate:  C.uint32_t(pkt.Datarate),
		// datarate:   C.DR_LORA_SF9,
		coderate: C.uint8_t(pkt.Coderate),
		// coderate:   C.CR_LORA_4_5,
		rf_power:   14,
		modulation: C.uint8_t(pkt.Modulation),
	}

	// Inserting payload
	if err := insertPayload(pkt, &txPacket); err != nil {
		return err
	}

	return sendPacketConcentrator(txPacket)
}

func sendPacketConcentrator(txPacket C.struct_lgw_pkt_tx_s) error {
	for {
		var txStatus C.uint8_t
		concentratorMutex.Lock()
		var result = C.lgw_status(C.TX_STATUS, &txStatus)
		concentratorMutex.Unlock()
		if result == C.LGW_HAL_ERROR {
			fmt.Fprintln(os.Stderr, "Couldn't get concentrator status")
		} else if txStatus == C.TX_EMITTING {
			// XX: Should we stop emission (like in the legacy packet forwarder) or retry?
			// If we retry, we might overwrite a normally scheduled downlink, that might
			// then not be relayed by the concentrator...
			return errors.New("Concentrator is already emitting")
		} else if txStatus == C.TX_SCHEDULED {
			fmt.Fprintln(os.Stderr, "A downlink was already scheduled, overwriting it")
		}
		break
	}

	concentratorMutex.Lock()
	result := C.lgw_send(txPacket)
	concentratorMutex.Unlock()

	if result == C.LGW_HAL_ERROR {
		return errors.New("Downlink transmission to the concentrator failed")
	}

	return nil
}

func WaitForConcentrator() error {
	for {
		var txStatus C.uint8_t
		concentratorMutex.Lock()
		var result = C.lgw_status(C.TX_STATUS, &txStatus)
		concentratorMutex.Unlock()
		if result == C.LGW_HAL_ERROR {
			fmt.Fprintln(os.Stderr, "Couldn't get concentrator status")
		} else if txStatus == C.TX_OFF {
			// XX: Should we stop emission (like in the legacy packet forwarder) or retry?
			// If we retry, we might overwrite a normally scheduled downlink, that might
			// then not be relayed by the concentrator...
			return errors.New("Concentrator is off")
		} else if txStatus == C.TX_STATUS_UNKNOWN {
			return errors.New("Concentrator status unknown")
		} else if txStatus == C.TX_FREE {
			break
		}
		time.Sleep(100 * time.Microsecond)
	}
	return nil
}
