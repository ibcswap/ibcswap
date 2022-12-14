package types

// IBC transfer events
const (
	EventTypeTimeout      = "timeout"
	EventTypePacket       = "fungible_token_packet"
	EventTypeSwap         = "ibc_swap"
	EventTypeChannelClose = "channel_closed"

	AttributeType          = "type"
	AttributeData          = "data"
	AttributeMemo          = "memo"
	AttributeKeyAmount     = "amount"
	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
	//AttributeKeyTraceHash       = "trace_hash"
	//AttributeKeySendingDenom    = "sending_token_denom"
	//AttributeKeySendingAmount   = "sending_token_amount"
	//AttributeKeyReceivingDenom  = "receiving_token_denom"
	//AttributeKeyReceivingAmount = "receiving_token_amount"
	//AttributeKeyCounterParty    = "expected_counterparty_address"
	//AttributeKeyRefundReceiver  = "refund_receiver"
	//	//AttributeKeyRefundDenom     = "refund_denom"
	//	//AttributeKeyRefundAmount    = "refund_amount"
)
