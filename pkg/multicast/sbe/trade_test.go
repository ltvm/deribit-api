package sbe

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"

	"github.com/KyberNetwork/deribit-api/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeTrade(t *testing.T) {
	tests := []struct {
		event          []byte
		expectedOutput Trades
		expectedError  error
	}{
		// success case, 2 for Perpetual and 1 for Option
		{
			[]byte{
				0x04, 0x00, 0xea, 0x03, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48, 0x37, 0x03, 0x00,
				0x53, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9a, 0x99, 0x99, 0x99, 0x99, 0xe3, 0x99,
				0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xa6, 0x40, 0xc6, 0x17, 0xfd, 0x11, 0x83, 0x01, 0x00,
				0x00, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe2, 0x99, 0x40, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe7, 0x99,
				0x40, 0xf6, 0xbe, 0x4d, 0x06, 0x00, 0x00, 0x00, 0x00, 0x5b, 0xa2, 0x84, 0x08, 0x00, 0x00, 0x00,
				0x00, 0x01, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x85,
			},
			Trades{
				InstrumentId: 210760, // ETH-perpetual
				TradesList: []TradesTradesList{
					{
						Direction:     Direction.Buy,
						Price:         1656.9,
						Amount:        2817,
						TimestampMs:   1662454142918,
						MarkPrice:     1656.68,
						IndexPrice:    1657.93,
						TradeSeq:      105758454,
						TradeId:       142910043,
						TickDirection: TickDirection.ZeroPlus,
						Liquidation:   Liquidation.None,
						Iv:            math.NaN(),
						BlockTradeId:  0,
						ComboTradeId:  0,
					},
				},
			},
			nil,
		},
		{
			[]byte{
				0x04, 0x00, 0xea, 0x03, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48, 0x37, 0x03, 0x00,
				0x53, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x33, 0x33, 0x33, 0x33, 0xe3, 0x99,
				0x40, 0x00, 0x00, 0x00, 0x00, 0x80, 0xb3, 0xe2, 0x40, 0x18, 0x17, 0xfd, 0x11, 0x83, 0x01, 0x00,
				0x00, 0x71, 0x3d, 0x0a, 0xd7, 0xa3, 0xe2, 0x99, 0x40, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe7, 0x99,
				0x40, 0xee, 0xbe, 0x4d, 0x06, 0x00, 0x00, 0x00, 0x00, 0x53, 0xa2, 0x84, 0x08, 0x00, 0x00, 0x00,
				0x00, 0x02, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x33, 0x33, 0x33,
				0x33, 0xe3, 0x99, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x97, 0x40, 0x18, 0x17, 0xfd, 0x11,
				0x83, 0x01, 0x00, 0x00, 0x71, 0x3d, 0x0a, 0xd7, 0xa3, 0xe2, 0x99, 0x40, 0x1f, 0x85, 0xeb, 0x51,
				0xb8, 0xe7, 0x99, 0x40, 0xef, 0xbe, 0x4d, 0x06, 0x00, 0x00, 0x00, 0x00, 0x54, 0xa2, 0x84, 0x08,
				0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x85,
			},
			Trades{
				InstrumentId: 210760, // ETH-perpetual
				TradesList: []TradesTradesList{
					{
						Direction:     Direction.Buy,
						Price:         1656.8,
						Amount:        38300,
						TimestampMs:   1662454142744,
						MarkPrice:     1656.66,
						IndexPrice:    1657.93,
						TradeSeq:      105758446,
						TradeId:       142910035,
						TickDirection: TickDirection.Minus,
						Liquidation:   Liquidation.None,
						Iv:            math.NaN(),
						BlockTradeId:  0,
						ComboTradeId:  0,
					},
					{
						Direction:     Direction.Buy,
						Price:         1656.8,
						Amount:        1500,
						TimestampMs:   1662454142744,
						MarkPrice:     1656.66,
						IndexPrice:    1657.93,
						TradeSeq:      105758447,
						TradeId:       142910036,
						TickDirection: TickDirection.ZeroMinus,
						Liquidation:   Liquidation.None,
						Iv:            math.NaN(),
						BlockTradeId:  0,
						ComboTradeId:  0,
					},
				},
			},
			nil,
		},
		{
			[]byte{
				0x04, 0x00, 0xea, 0x03, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0b, 0x7c, 0x03, 0x00,
				0x53, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xfc, 0xa9, 0xf1, 0xd2, 0x4d, 0x62, 0x60,
				0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x40, 0x8e, 0x82, 0x35, 0x12, 0x83, 0x01, 0x00,
				0x00, 0x4e, 0x29, 0xaf, 0x95, 0xd0, 0x5d, 0x62, 0x3f, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe1, 0x99,
				0x40, 0x83, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0xbd, 0x84, 0x08, 0x00, 0x00, 0x00,
				0x00, 0x03, 0x00, 0x0a, 0xd7, 0xa3, 0x70, 0x3d, 0xc2, 0x60, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x85,
			},
			Trades{
				InstrumentId: 228363, // ETH-9SEP22-1350-P
				TradesList: []TradesTradesList{
					{
						Direction:     Direction.Sell,
						Price:         0.002,
						Amount:        200,
						TimestampMs:   1662457840270,
						MarkPrice:     0.002242,
						IndexPrice:    1656.43,
						TradeSeq:      1411,
						TradeId:       142916883,
						TickDirection: TickDirection.ZeroMinus,
						Liquidation:   Liquidation.None,
						Iv:            134.07,
						BlockTradeId:  0,
						ComboTradeId:  0,
					},
				},
			},
			nil,
		},
		// failure case
		{
			[]byte{
				0x04, 0x00, 0xea, 0x03, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0b, 0x7c,
			},
			Trades{},
			io.ErrUnexpectedEOF,
		},
		{
			[]byte{
				0x04, 0x00, 0xea, 0x03, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48, 0x37, 0x03, 0x00,
				0x53, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9a, 0x99, 0x99, 0x99, 0x99, 0xe3, 0x99,
				0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xa6, 0x40, 0xc6, 0x17, 0xfd, 0x11, 0x83, 0x01, 0x00,
				0x00, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe2, 0x99, 0x40, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0xe7, 0x99,
				0x40, 0xf6, 0xbe, 0x4d, 0x06, 0x00, 0x00, 0x00, 0x00, 0x5b, 0xa2, 0x84, 0x08, 0x00, 0x00, 0x00,
				0x00, 0x0a, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x85,
			},
			Trades{},
			ErrRangeCheck,
		},
	}

	marshaller := NewSbeGoMarshaller()
	tradePtr := reflect.TypeOf(&TradesTradesList{})

	for _, test := range tests {
		bufferData := bytes.NewBuffer(test.event)

		var header MessageHeader
		err := header.Decode(marshaller, bufferData)
		require.NoError(t, err)

		err = header.RangeCheck()
		require.NoError(t, err)

		var trades Trades
		err = trades.Decode(marshaller, bufferData, header.BlockLength, true)
		require.ErrorIs(t, err, test.expectedError)

		if err == nil {
			// replace NaN to 0 of output and expected output
			for i := 0; i < len(trades.TradesList); i++ {
				common.ReplaceNaNValueOfStruct(&trades.TradesList[i], tradePtr)
			}

			for i := 0; i < len(test.expectedOutput.TradesList); i++ {
				common.ReplaceNaNValueOfStruct(&test.expectedOutput.TradesList[i], tradePtr)
			}

			assert.Equal(t, trades, test.expectedOutput)
		}
	}
}
