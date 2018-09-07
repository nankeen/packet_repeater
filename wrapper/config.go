package wrapper

// LBTConf wraps lbt configuration for SX1301
type LbtConf struct {
	Enabled        bool              `json:"enable"`
	RssiTarget     int               `json:"rssi_target"`
	RssiOffset     int               `json:"sx127x_rssi_offset"`
	ChannelsConfig []ChannelFreqConf `json:"chan_cfg"`
}

type ChannelFreqConf struct {
	Freq     int `json:"freq_hz"` // Frequency in hertz
	ScanTime int `json:"scan_time_us"`
}

type ChannelConf struct {
	Enabled      bool    `json:"enable"`
	Radio        uint8   `json:"radio"`
	IfValue      int32   `json:"if"`
	Bandwidth    *uint32 `json:"bandwidth,omitempty"`
	Datarate     *uint32 `json:"datarate,omitempty"`
	SpreadFactor *uint8  `json:"spread_factor,omitempty"`
	Description  *string `json:"desc,omitempty"`
}

type RadioConf struct {
	Enabled     bool    `json:"enable"`
	RadioType   string  `json:"type"`
	Freq        int     `json:"freq"`
	RssiOffset  float32 `json:"rssi_offset"`
	TxEnabled   bool    `json:"tx_enable"`
	TxMinFreq   *int    `json:"tx_freq_min,omitempty"`
	TxMaxFreq   *int    `json:"tx_freq_max,omitempty"`
	TxNotchFreq *int    `json:"tx_notch_freq,omitempty"`
}

type GainTableConf struct {
	PaGain      uint8   `json:"pa_gain"`
	MixGain     uint8   `json:"mix_gain"`
	RfPower     int8    `json:"rf_power"`
	DigGain     uint8   `json:"dig_gain"`
	DacGain     *uint8  `json:"dac_gain,omitempty"`
	Description *string `json:"desc,omitempty"`
}

type SX1301Conf struct {
	LorawanPublic          bool           `json:"lorawan_public"`
	Clksrc                 int            `json:"clksrc"`
	ClksrcDescription      *string        `json:"clksrc_desc,omitempty"`
	AntennaGain            *int           `json:"antenna_gain,omitempty"`
	AntennaGainDescription *string        `json:"antenna_gain_desc,omitempty"`
	LbtConfig              *LbtConf       `json:"lbt_cfg,omitempty"`
	Radio0                 *RadioConf     `json:"radio_0,omitempty"`
	Radio1                 *RadioConf     `json:"radio_1,omitempty"`
	MultiSFChan0           *ChannelConf   `json:"chan_multiSF_0,omitempty"`
	MultiSFChan1           *ChannelConf   `json:"chan_multiSF_1,omitempty"`
	MultiSFChan2           *ChannelConf   `json:"chan_multiSF_2,omitempty"`
	MultiSFChan3           *ChannelConf   `json:"chan_multiSF_3,omitempty"`
	MultiSFChan4           *ChannelConf   `json:"chan_multiSF_4,omitempty"`
	MultiSFChan5           *ChannelConf   `json:"chan_multiSF_5,omitempty"`
	MultiSFChan6           *ChannelConf   `json:"chan_multiSF_6,omitempty"`
	MultiSFChan7           *ChannelConf   `json:"chan_multiSF_7,omitempty"`
	MultiSFChan8           *ChannelConf   `json:"chan_multiSF_8,omitempty"`
	MultiSFChan9           *ChannelConf   `json:"chan_multiSF_9,omitempty"`
	MultiSFChan10          *ChannelConf   `json:"chan_multiSF_10,omitempty"`
	MultiSFChan11          *ChannelConf   `json:"chan_multiSF_11,omitempty"`
	MultiSFChan12          *ChannelConf   `json:"chan_multiSF_12,omitempty"`
	MultiSFChan13          *ChannelConf   `json:"chan_multiSF_13,omitempty"`
	MultiSFChan14          *ChannelConf   `json:"chan_multiSF_14,omitempty"`
	MultiSFChan15          *ChannelConf   `json:"chan_multiSF_15,omitempty"`
	MultiSFChan16          *ChannelConf   `json:"chan_multiSF_16,omitempty"`
	MultiSFChan17          *ChannelConf   `json:"chan_multiSF_17,omitempty"`
	MultiSFChan18          *ChannelConf   `json:"chan_multiSF_18,omitempty"`
	MultiSFChan19          *ChannelConf   `json:"chan_multiSF_19,omitempty"`
	MultiSFChan20          *ChannelConf   `json:"chan_multiSF_20,omitempty"`
	MultiSFChan21          *ChannelConf   `json:"chan_multiSF_21,omitempty"`
	MultiSFChan22          *ChannelConf   `json:"chan_multiSF_22,omitempty"`
	MultiSFChan23          *ChannelConf   `json:"chan_multiSF_23,omitempty"`
	MultiSFChan24          *ChannelConf   `json:"chan_multiSF_24,omitempty"`
	MultiSFChan25          *ChannelConf   `json:"chan_multiSF_25,omitempty"`
	MultiSFChan26          *ChannelConf   `json:"chan_multiSF_26,omitempty"`
	MultiSFChan27          *ChannelConf   `json:"chan_multiSF_27,omitempty"`
	MultiSFChan28          *ChannelConf   `json:"chan_multiSF_28,omitempty"`
	MultiSFChan29          *ChannelConf   `json:"chan_multiSF_29,omitempty"`
	MultiSFChan30          *ChannelConf   `json:"chan_multiSF_30,omitempty"`
	MultiSFChan31          *ChannelConf   `json:"chan_multiSF_31,omitempty"`
	MultiSFChan32          *ChannelConf   `json:"chan_multiSF_32,omitempty"`
	MultiSFChan33          *ChannelConf   `json:"chan_multiSF_33,omitempty"`
	MultiSFChan34          *ChannelConf   `json:"chan_multiSF_34,omitempty"`
	MultiSFChan35          *ChannelConf   `json:"chan_multiSF_35,omitempty"`
	MultiSFChan36          *ChannelConf   `json:"chan_multiSF_36,omitempty"`
	MultiSFChan37          *ChannelConf   `json:"chan_multiSF_37,omitempty"`
	MultiSFChan38          *ChannelConf   `json:"chan_multiSF_38,omitempty"`
	MultiSFChan39          *ChannelConf   `json:"chan_multiSF_39,omitempty"`
	MultiSFChan40          *ChannelConf   `json:"chan_multiSF_40,omitempty"`
	MultiSFChan41          *ChannelConf   `json:"chan_multiSF_41,omitempty"`
	MultiSFChan42          *ChannelConf   `json:"chan_multiSF_42,omitempty"`
	MultiSFChan43          *ChannelConf   `json:"chan_multiSF_43,omitempty"`
	MultiSFChan44          *ChannelConf   `json:"chan_multiSF_44,omitempty"`
	MultiSFChan45          *ChannelConf   `json:"chan_multiSF_45,omitempty"`
	MultiSFChan46          *ChannelConf   `json:"chan_multiSF_46,omitempty"`
	MultiSFChan47          *ChannelConf   `json:"chan_multiSF_47,omitempty"`
	MultiSFChan48          *ChannelConf   `json:"chan_multiSF_48,omitempty"`
	MultiSFChan49          *ChannelConf   `json:"chan_multiSF_49,omitempty"`
	MultiSFChan50          *ChannelConf   `json:"chan_multiSF_50,omitempty"`
	MultiSFChan51          *ChannelConf   `json:"chan_multiSF_51,omitempty"`
	MultiSFChan52          *ChannelConf   `json:"chan_multiSF_52,omitempty"`
	MultiSFChan53          *ChannelConf   `json:"chan_multiSF_53,omitempty"`
	MultiSFChan54          *ChannelConf   `json:"chan_multiSF_54,omitempty"`
	MultiSFChan55          *ChannelConf   `json:"chan_multiSF_55,omitempty"`
	MultiSFChan56          *ChannelConf   `json:"chan_multiSF_56,omitempty"`
	MultiSFChan57          *ChannelConf   `json:"chan_multiSF_57,omitempty"`
	MultiSFChan58          *ChannelConf   `json:"chan_multiSF_58,omitempty"`
	MultiSFChan59          *ChannelConf   `json:"chan_multiSF_59,omitempty"`
	MultiSFChan60          *ChannelConf   `json:"chan_multiSF_60,omitempty"`
	MultiSFChan61          *ChannelConf   `json:"chan_multiSF_61,omitempty"`
	MultiSFChan62          *ChannelConf   `json:"chan_multiSF_62,omitempty"`
	MultiSFChan63          *ChannelConf   `json:"chan_multiSF_63,omitempty"`
	LoraSTDChannel         *ChannelConf   `json:"chan_Lora_std,omitempty"`
	FSKChannel             *ChannelConf   `json:"chan_FSK,omitempty"`
	TxLut0                 *GainTableConf `json:"tx_lut_0,omitempty"`
	TxLut1                 *GainTableConf `json:"tx_lut_1,omitempty"`
	TxLut2                 *GainTableConf `json:"tx_lut_2,omitempty"`
	TxLut3                 *GainTableConf `json:"tx_lut_3,omitempty"`
	TxLut4                 *GainTableConf `json:"tx_lut_4,omitempty"`
	TxLut5                 *GainTableConf `json:"tx_lut_5,omitempty"`
	TxLut6                 *GainTableConf `json:"tx_lut_6,omitempty"`
	TxLut7                 *GainTableConf `json:"tx_lut_7,omitempty"`
	TxLut8                 *GainTableConf `json:"tx_lut_8,omitempty"`
	TxLut9                 *GainTableConf `json:"tx_lut_9,omitempty"`
	TxLut10                *GainTableConf `json:"tx_lut_10,omitempty"`
	TxLut11                *GainTableConf `json:"tx_lut_11,omitempty"`
	TxLut12                *GainTableConf `json:"tx_lut_12,omitempty"`
	TxLut13                *GainTableConf `json:"tx_lut_13,omitempty"`
	TxLut14                *GainTableConf `json:"tx_lut_14,omitempty"`
	TxLut15                *GainTableConf `json:"tx_lut_15,omitempty"`
}

func GetLuts() []GainTableConf {
	return []GainTableConf{
		{0, 15, 2, 0, nil, nil},
		{1, 8, 1, 0, nil, nil},
		{1, 10, 4, 0, nil, nil},
		{1, 12, 6, 0, nil, nil},
		{1, 13, 7, 0, nil, nil},
		{2, 8, 8, 0, nil, nil},
		{2, 9, 10, 0, nil, nil},
		{2, 10, 11, 0, nil, nil},
		{2, 11, 13, 0, nil, nil},
		{2, 12, 14, 0, nil, nil},
		{2, 15, 15, 0, nil, nil},
		{3, 8, 17, 0, nil, nil},
		{3, 9, 19, 0, nil, nil},
		{3, 10, 20, 0, nil, nil},
		{3, 12, 22, 0, nil, nil},
		{3, 14, 24, 0, nil, nil},
	}
}

func GetRFConfs() []RadioConf {
	radio0_min := 919000000
	radio0_max := 923000000
	return []RadioConf{
		{true, "SX1257", 922400000, -166.0, true, &radio0_min, &radio0_max, nil},
		{true, "SX1257", 919800000, -166.0, false, nil, nil, nil},
	}
}

func GetMultiSFChannels() []ChannelConf {
	return []ChannelConf{
		{true, 0, 300000, nil, nil, nil, nil},
		{true, 0, 100000, nil, nil, nil, nil},
		{true, 0, 100000, nil, nil, nil, nil},
		{true, 0, 300000, nil, nil, nil, nil},
		{true, 1, 300000, nil, nil, nil, nil},
		{true, 1, 100000, nil, nil, nil, nil},
		{true, 1, 100000, nil, nil, nil, nil},
		{true, 1, 300000, nil, nil, nil, nil},
	}
}

func GetLoraSTDChannel() *ChannelConf {
	var bandwidth uint32 = 0
	var spread_factor uint8 = 7
	return &ChannelConf{false, 0, 0, &bandwidth, nil, &spread_factor, nil}
}

func GetFSKChannel() *ChannelConf {
	var bandwidth uint32 = 0
	return &ChannelConf{false, 0, 0, &bandwidth, nil, nil, nil}
}
