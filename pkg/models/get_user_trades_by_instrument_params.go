package models

type GetUserTradesByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	StartSeq       int64  `json:"start_seq,omitempty"`
	EndSeq         int64  `json:"end_seq,omitempty"`
	Count          int    `json:"count,omitempty"`
	IncludeOld     *bool  `json:"include_old,omitempty"`
	Sorting        string `json:"sorting,omitempty"`
}
