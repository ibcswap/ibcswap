package types

import (
	"crypto/sha256"
	"github.com/gogo/protobuf/proto"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

func NewLimitOrder(msg *MsgMakeSwapRequest, channelId string) LimitOrder {
	maker := NewMakerFromMsg(msg)
	buf, _ := proto.Marshal(maker)
	id := Hash(buf).String()
	var takers []*SwapTaker
	return LimitOrder{
		Id:                id,
		Maker:             maker,
		Status:            Status_INITIAL,
		FillStatus:        FillStatus_NONE_FILL,
		ChannelId:         channelId,
		Takers:            takers,
		CancelTimestamp:   0,
		CompleteTimestamp: 0,
	}
}

func NewOTCOrder(msg *MsgMakeSwapRequest, channelId string) OTCOrder {
	maker := NewMakerFromMsg(msg)
	buf, _ := proto.Marshal(maker)
	id := Hash(buf).String()
	return OTCOrder{
		Id:                id,
		Maker:             maker,
		Status:            Status_INITIAL,
		ChannelId:         channelId,
		Takers:            nil,
		CancelTimestamp:   0,
		CompleteTimestamp: 0,
	}
}

func NewMakerFromMsg(msg *MsgMakeSwapRequest) *SwapMaker {
	return &SwapMaker{
		SourcePort:            msg.SourcePort,
		SourceChannel:         msg.SourceChannel,
		SellToken:             msg.SellToken,
		BuyToken:              msg.BuyToken,
		MakerAddress:          msg.MakerAddress,
		MakerReceivingAddress: msg.MakerReceivingAddress,
		DesiredTaker:          msg.DesiredTaker,
		CreateTimestamp:       msg.CreateTimestamp,
	}
}

func NewTakerFromMsg(msg *MsgTakeSwapRequest) *SwapTaker {
	return &SwapTaker{
		OrderId:               msg.OrderId,
		SellToken:             msg.SellToken,
		TakerAddress:          msg.TakerAddress,
		TakerReceivingAddress: msg.TakerReceivingAddress,
		CreateTimestamp:       msg.CreateTimestamp,
	}
}

func Hash(content []byte) tmbytes.HexBytes {
	hash := sha256.Sum256(content)
	return hash[:]
}
