// Code generated by protoc-gen-go.
// source: bftrader.proto
// DO NOT EDIT!

/*
Package bftrader is a generated protocol buffer package.

It is generated from these files:
	bftrader.proto

It has these top-level messages:
	BfVoid
	BfDailyInfo
	BfBidAskInfo
	BfOHLCInfo
	BfTickData
	BfTickPackData
	BfBarData
	BfTradeData
	BfOrderData
	BfPositionData
	BfAccountData
	BfErrorData
	BfLogData
	BfContractData
	BfKvData
	BfSubscribeReq
	BfOrderReq
	BfOrderResp
	BfCancelOrderReq
	BfConnectReq
	BfConnectResp
	BfGetContractReq
	BfGetTickPackReq
	BfGetBarReq
*/
package bftrader

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 方向常量
type BfDirection int32

const (
	BfDirection_DIRECTION_UNKNOWN BfDirection = 0
	BfDirection_DIRECTION_NONE    BfDirection = 1
	BfDirection_DIRECTION_LONG    BfDirection = 2
	BfDirection_DIRECTION_SHORT   BfDirection = 3
	BfDirection_DIRECTION_NET     BfDirection = 4
	BfDirection_DIRECTION_SELL    BfDirection = 5
)

var BfDirection_name = map[int32]string{
	0: "DIRECTION_UNKNOWN",
	1: "DIRECTION_NONE",
	2: "DIRECTION_LONG",
	3: "DIRECTION_SHORT",
	4: "DIRECTION_NET",
	5: "DIRECTION_SELL",
}
var BfDirection_value = map[string]int32{
	"DIRECTION_UNKNOWN": 0,
	"DIRECTION_NONE":    1,
	"DIRECTION_LONG":    2,
	"DIRECTION_SHORT":   3,
	"DIRECTION_NET":     4,
	"DIRECTION_SELL":    5,
}

func (x BfDirection) String() string {
	return proto.EnumName(BfDirection_name, int32(x))
}
func (BfDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 开平常量
type BfOffset int32

const (
	BfOffset_OFFSET_UNKNOWN        BfOffset = 0
	BfOffset_OFFSET_NONE           BfOffset = 1
	BfOffset_OFFSET_OPEN           BfOffset = 2
	BfOffset_OFFSET_CLOSE          BfOffset = 3
	BfOffset_OFFSET_CLOSETODAY     BfOffset = 4
	BfOffset_OFFSET_CLOSEYESTERDAY BfOffset = 5
)

var BfOffset_name = map[int32]string{
	0: "OFFSET_UNKNOWN",
	1: "OFFSET_NONE",
	2: "OFFSET_OPEN",
	3: "OFFSET_CLOSE",
	4: "OFFSET_CLOSETODAY",
	5: "OFFSET_CLOSEYESTERDAY",
}
var BfOffset_value = map[string]int32{
	"OFFSET_UNKNOWN":        0,
	"OFFSET_NONE":           1,
	"OFFSET_OPEN":           2,
	"OFFSET_CLOSE":          3,
	"OFFSET_CLOSETODAY":     4,
	"OFFSET_CLOSEYESTERDAY": 5,
}

func (x BfOffset) String() string {
	return proto.EnumName(BfOffset_name, int32(x))
}
func (BfOffset) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 状态常量
type BfStatus int32

const (
	BfStatus_STATUS_UNKNOWN    BfStatus = 0
	BfStatus_STATUS_NOTTRADED  BfStatus = 1
	BfStatus_STATUS_PARTTRADED BfStatus = 2
	BfStatus_STATUS_ALLTRADED  BfStatus = 3
	BfStatus_STATUS_CANCELLED  BfStatus = 4
)

var BfStatus_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_NOTTRADED",
	2: "STATUS_PARTTRADED",
	3: "STATUS_ALLTRADED",
	4: "STATUS_CANCELLED",
}
var BfStatus_value = map[string]int32{
	"STATUS_UNKNOWN":    0,
	"STATUS_NOTTRADED":  1,
	"STATUS_PARTTRADED": 2,
	"STATUS_ALLTRADED":  3,
	"STATUS_CANCELLED":  4,
}

func (x BfStatus) String() string {
	return proto.EnumName(BfStatus_name, int32(x))
}
func (BfStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 合约类型常量
type BfProduct int32

const (
	BfProduct_PRODUCT_UNKNOWN     BfProduct = 0
	BfProduct_PRODUCT_NONE        BfProduct = 1
	BfProduct_PRODUCT_EQUITY      BfProduct = 2
	BfProduct_PRODUCT_FUTURES     BfProduct = 3
	BfProduct_PRODUCT_OPTION      BfProduct = 4
	BfProduct_PRODUCT_INDEX       BfProduct = 5
	BfProduct_PRODUCT_COMBINATION BfProduct = 6
	BfProduct_PRODUCT_FOREX       BfProduct = 7
	BfProduct_PRODUCT_SPOT        BfProduct = 8
	BfProduct_PRODUCT_DEFER       BfProduct = 9
)

var BfProduct_name = map[int32]string{
	0: "PRODUCT_UNKNOWN",
	1: "PRODUCT_NONE",
	2: "PRODUCT_EQUITY",
	3: "PRODUCT_FUTURES",
	4: "PRODUCT_OPTION",
	5: "PRODUCT_INDEX",
	6: "PRODUCT_COMBINATION",
	7: "PRODUCT_FOREX",
	8: "PRODUCT_SPOT",
	9: "PRODUCT_DEFER",
}
var BfProduct_value = map[string]int32{
	"PRODUCT_UNKNOWN":     0,
	"PRODUCT_NONE":        1,
	"PRODUCT_EQUITY":      2,
	"PRODUCT_FUTURES":     3,
	"PRODUCT_OPTION":      4,
	"PRODUCT_INDEX":       5,
	"PRODUCT_COMBINATION": 6,
	"PRODUCT_FOREX":       7,
	"PRODUCT_SPOT":        8,
	"PRODUCT_DEFER":       9,
}

func (x BfProduct) String() string {
	return proto.EnumName(BfProduct_name, int32(x))
}
func (BfProduct) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 价格类型常量
type BfPriceType int32

const (
	BfPriceType_PRICETYPE_UNKONWN     BfPriceType = 0
	BfPriceType_PRICETYPE_LIMITPRICE  BfPriceType = 1
	BfPriceType_PRICETYPE_MARKETPRICE BfPriceType = 2
	BfPriceType_PRICETYPE_FAK         BfPriceType = 3
	BfPriceType_PRICETYPE_FOK         BfPriceType = 4
)

var BfPriceType_name = map[int32]string{
	0: "PRICETYPE_UNKONWN",
	1: "PRICETYPE_LIMITPRICE",
	2: "PRICETYPE_MARKETPRICE",
	3: "PRICETYPE_FAK",
	4: "PRICETYPE_FOK",
}
var BfPriceType_value = map[string]int32{
	"PRICETYPE_UNKONWN":     0,
	"PRICETYPE_LIMITPRICE":  1,
	"PRICETYPE_MARKETPRICE": 2,
	"PRICETYPE_FAK":         3,
	"PRICETYPE_FOK":         4,
}

func (x BfPriceType) String() string {
	return proto.EnumName(BfPriceType_name, int32(x))
}
func (BfPriceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 交易所类型
type BfExchange int32

const (
	BfExchange_EXCHANGE_UNKNOWN BfExchange = 0
	BfExchange_EXCHANGE_SSE     BfExchange = 1
	BfExchange_EXCHANGE_SZSE    BfExchange = 2
	BfExchange_EXCHANGE_CFFEX   BfExchange = 3
	BfExchange_EXCHANGE_SHFE    BfExchange = 4
	BfExchange_EXCHANGE_CZCE    BfExchange = 5
	BfExchange_EXCHANGE_DCE     BfExchange = 6
	BfExchange_EXCHANGE_SGE     BfExchange = 7
)

var BfExchange_name = map[int32]string{
	0: "EXCHANGE_UNKNOWN",
	1: "EXCHANGE_SSE",
	2: "EXCHANGE_SZSE",
	3: "EXCHANGE_CFFEX",
	4: "EXCHANGE_SHFE",
	5: "EXCHANGE_CZCE",
	6: "EXCHANGE_DCE",
	7: "EXCHANGE_SGE",
}
var BfExchange_value = map[string]int32{
	"EXCHANGE_UNKNOWN": 0,
	"EXCHANGE_SSE":     1,
	"EXCHANGE_SZSE":    2,
	"EXCHANGE_CFFEX":   3,
	"EXCHANGE_SHFE":    4,
	"EXCHANGE_CZCE":    5,
	"EXCHANGE_DCE":     6,
	"EXCHANGE_SGE":     7,
}

func (x BfExchange) String() string {
	return proto.EnumName(BfExchange_name, int32(x))
}
func (BfExchange) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Bar周期类型
type BfPeriod int32

const (
	BfPeriod_PERIOD_UNKNOWN BfPeriod = 0
	BfPeriod_PERIOD_M1      BfPeriod = 1
	BfPeriod_PERIOD_D1      BfPeriod = 2
)

var BfPeriod_name = map[int32]string{
	0: "PERIOD_UNKNOWN",
	1: "PERIOD_M1",
	2: "PERIOD_D1",
}
var BfPeriod_value = map[string]int32{
	"PERIOD_UNKNOWN": 0,
	"PERIOD_M1":      1,
	"PERIOD_D1":      2,
}

func (x BfPeriod) String() string {
	return proto.EnumName(BfPeriod_name, int32(x))
}
func (BfPeriod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 空参数
type BfVoid struct {
}

func (m *BfVoid) Reset()                    { *m = BfVoid{} }
func (m *BfVoid) String() string            { return proto.CompactTextString(m) }
func (*BfVoid) ProtoMessage()               {}
func (*BfVoid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ============================
// 常规行情数据类
type BfDailyInfo struct {
	OpenPrice     float64 `protobuf:"fixed64,1,opt,name=openPrice" json:"openPrice,omitempty"`
	HighPrice     float64 `protobuf:"fixed64,2,opt,name=highPrice" json:"highPrice,omitempty"`
	LowPrice      float64 `protobuf:"fixed64,3,opt,name=lowPrice" json:"lowPrice,omitempty"`
	PreClosePrice float64 `protobuf:"fixed64,4,opt,name=preClosePrice" json:"preClosePrice,omitempty"`
	UpperLimit    float64 `protobuf:"fixed64,5,opt,name=upperLimit" json:"upperLimit,omitempty"`
	LowerLimit    float64 `protobuf:"fixed64,6,opt,name=lowerLimit" json:"lowerLimit,omitempty"`
}

func (m *BfDailyInfo) Reset()                    { *m = BfDailyInfo{} }
func (m *BfDailyInfo) String() string            { return proto.CompactTextString(m) }
func (*BfDailyInfo) ProtoMessage()               {}
func (*BfDailyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// x档行情
type BfBidAskInfo struct {
	BidPrice  float64 `protobuf:"fixed64,1,opt,name=bidPrice" json:"bidPrice,omitempty"`
	AskPrice  float64 `protobuf:"fixed64,2,opt,name=askPrice" json:"askPrice,omitempty"`
	BidVolume int32   `protobuf:"varint,3,opt,name=bidVolume" json:"bidVolume,omitempty"`
	AskVolume int32   `protobuf:"varint,4,opt,name=askVolume" json:"askVolume,omitempty"`
}

func (m *BfBidAskInfo) Reset()                    { *m = BfBidAskInfo{} }
func (m *BfBidAskInfo) String() string            { return proto.CompactTextString(m) }
func (*BfBidAskInfo) ProtoMessage()               {}
func (*BfBidAskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// OHLC
type BfOHLCInfo struct {
	OpenPrice  float64 `protobuf:"fixed64,1,opt,name=openPrice" json:"openPrice,omitempty"`
	HighPrice  float64 `protobuf:"fixed64,2,opt,name=highPrice" json:"highPrice,omitempty"`
	LowPrice   float64 `protobuf:"fixed64,3,opt,name=lowPrice" json:"lowPrice,omitempty"`
	ClosePrice float64 `protobuf:"fixed64,4,opt,name=closePrice" json:"closePrice,omitempty"`
}

func (m *BfOHLCInfo) Reset()                    { *m = BfOHLCInfo{} }
func (m *BfOHLCInfo) String() string            { return proto.CompactTextString(m) }
func (*BfOHLCInfo) ProtoMessage()               {}
func (*BfOHLCInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Tick行情数据类
type BfTickData struct {
	// 代码相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 成交数据
	Date             string  `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
	Time             string  `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	LastPrice        float64 `protobuf:"fixed64,5,opt,name=lastPrice" json:"lastPrice,omitempty"`
	Volume           int32   `protobuf:"varint,6,opt,name=volume" json:"volume,omitempty"`
	OpenInterest     float64 `protobuf:"fixed64,7,opt,name=openInterest" json:"openInterest,omitempty"`
	LastVolume       int32   `protobuf:"varint,8,opt,name=lastVolume" json:"lastVolume,omitempty"`
	LastOpenInterest float64 `protobuf:"fixed64,9,opt,name=lastOpenInterest" json:"lastOpenInterest,omitempty"`
	// x档行情
	BidaskInfo []*BfBidAskInfo `protobuf:"bytes,10,rep,name=bidaskInfo" json:"bidaskInfo,omitempty"`
	// 常规行情
	DailyInfo *BfDailyInfo `protobuf:"bytes,11,opt,name=dailyInfo" json:"dailyInfo,omitempty"`
}

func (m *BfTickData) Reset()                    { *m = BfTickData{} }
func (m *BfTickData) String() string            { return proto.CompactTextString(m) }
func (*BfTickData) ProtoMessage()               {}
func (*BfTickData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BfTickData) GetBidaskInfo() []*BfBidAskInfo {
	if m != nil {
		return m.BidaskInfo
	}
	return nil
}

func (m *BfTickData) GetDailyInfo() *BfDailyInfo {
	if m != nil {
		return m.DailyInfo
	}
	return nil
}

// 1分钟的tick集合
type BfTickPackData struct {
	// 代码相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Date     string `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
	Time     string `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	// Tick集合
	Ticks []*BfTickData `protobuf:"bytes,5,rep,name=ticks" json:"ticks,omitempty"`
}

func (m *BfTickPackData) Reset()                    { *m = BfTickPackData{} }
func (m *BfTickPackData) String() string            { return proto.CompactTextString(m) }
func (*BfTickPackData) ProtoMessage()               {}
func (*BfTickPackData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BfTickPackData) GetTicks() []*BfTickData {
	if m != nil {
		return m.Ticks
	}
	return nil
}

// Bar行情数据类
type BfBarData struct {
	// 代码相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 周期
	Period BfPeriod `protobuf:"varint,3,opt,name=period,enum=bftrader.BfPeriod" json:"period,omitempty"`
	// 成交数据
	Date             string  `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
	Time             string  `protobuf:"bytes,5,opt,name=time" json:"time,omitempty"`
	Volume           int32   `protobuf:"varint,6,opt,name=volume" json:"volume,omitempty"`
	OpenInterest     float64 `protobuf:"fixed64,7,opt,name=openInterest" json:"openInterest,omitempty"`
	LastVolume       int32   `protobuf:"varint,8,opt,name=lastVolume" json:"lastVolume,omitempty"`
	LastOpenInterest float64 `protobuf:"fixed64,9,opt,name=lastOpenInterest" json:"lastOpenInterest,omitempty"`
	// OHLC
	OhlcInfo *BfOHLCInfo `protobuf:"bytes,10,opt,name=ohlcInfo" json:"ohlcInfo,omitempty"`
	// 常规行情
	DailyInfo *BfDailyInfo `protobuf:"bytes,11,opt,name=dailyInfo" json:"dailyInfo,omitempty"`
}

func (m *BfBarData) Reset()                    { *m = BfBarData{} }
func (m *BfBarData) String() string            { return proto.CompactTextString(m) }
func (*BfBarData) ProtoMessage()               {}
func (*BfBarData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BfBarData) GetOhlcInfo() *BfOHLCInfo {
	if m != nil {
		return m.OhlcInfo
	}
	return nil
}

func (m *BfBarData) GetDailyInfo() *BfDailyInfo {
	if m != nil {
		return m.DailyInfo
	}
	return nil
}

// 成交数据类
type BfTradeData struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	TradeId  string `protobuf:"bytes,3,opt,name=tradeId" json:"tradeId,omitempty"`
	OrderId  string `protobuf:"bytes,4,opt,name=orderId" json:"orderId,omitempty"`
	// 成交相关
	Direction BfDirection `protobuf:"varint,5,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset    BfOffset    `protobuf:"varint,6,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
	Price     float64     `protobuf:"fixed64,7,opt,name=price" json:"price,omitempty"`
	Volume    int32       `protobuf:"varint,8,opt,name=volume" json:"volume,omitempty"`
	TradeTime string      `protobuf:"bytes,9,opt,name=tradeTime" json:"tradeTime,omitempty"`
}

func (m *BfTradeData) Reset()                    { *m = BfTradeData{} }
func (m *BfTradeData) String() string            { return proto.CompactTextString(m) }
func (*BfTradeData) ProtoMessage()               {}
func (*BfTradeData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 订单数据类
type BfOrderData struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	OrderId  string `protobuf:"bytes,3,opt,name=orderId" json:"orderId,omitempty"`
	// 报单相关
	Direction    BfDirection `protobuf:"varint,4,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset       BfOffset    `protobuf:"varint,5,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
	Price        float64     `protobuf:"fixed64,6,opt,name=price" json:"price,omitempty"`
	TotalVolume  int32       `protobuf:"varint,7,opt,name=totalVolume" json:"totalVolume,omitempty"`
	TradedVolume int32       `protobuf:"varint,8,opt,name=tradedVolume" json:"tradedVolume,omitempty"`
	Status       BfStatus    `protobuf:"varint,9,opt,name=status,enum=bftrader.BfStatus" json:"status,omitempty"`
	OrderTime    string      `protobuf:"bytes,10,opt,name=orderTime" json:"orderTime,omitempty"`
	CancelTime   string      `protobuf:"bytes,11,opt,name=cancelTime" json:"cancelTime,omitempty"`
}

func (m *BfOrderData) Reset()                    { *m = BfOrderData{} }
func (m *BfOrderData) String() string            { return proto.CompactTextString(m) }
func (*BfOrderData) ProtoMessage()               {}
func (*BfOrderData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 持仓数据类
type BfPositionData struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 持仓相关
	Direction BfDirection `protobuf:"varint,3,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Position  int32       `protobuf:"varint,4,opt,name=position" json:"position,omitempty"`
	Frozen    int32       `protobuf:"varint,5,opt,name=frozen" json:"frozen,omitempty"`
	Price     float64     `protobuf:"fixed64,6,opt,name=price" json:"price,omitempty"`
	// 昨持仓
	YdPosition int32 `protobuf:"varint,7,opt,name=ydPosition" json:"ydPosition,omitempty"`
}

func (m *BfPositionData) Reset()                    { *m = BfPositionData{} }
func (m *BfPositionData) String() string            { return proto.CompactTextString(m) }
func (*BfPositionData) ProtoMessage()               {}
func (*BfPositionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 账户数据类
type BfAccountData struct {
	// 账号代码相关
	AccountId string `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	// 数值相关
	PreBalance     float64 `protobuf:"fixed64,2,opt,name=preBalance" json:"preBalance,omitempty"`
	Balance        float64 `protobuf:"fixed64,3,opt,name=balance" json:"balance,omitempty"`
	Available      float64 `protobuf:"fixed64,4,opt,name=available" json:"available,omitempty"`
	Commission     float64 `protobuf:"fixed64,5,opt,name=commission" json:"commission,omitempty"`
	Margin         float64 `protobuf:"fixed64,6,opt,name=margin" json:"margin,omitempty"`
	CloseProfit    float64 `protobuf:"fixed64,7,opt,name=closeProfit" json:"closeProfit,omitempty"`
	PositionProfit float64 `protobuf:"fixed64,8,opt,name=positionProfit" json:"positionProfit,omitempty"`
}

func (m *BfAccountData) Reset()                    { *m = BfAccountData{} }
func (m *BfAccountData) String() string            { return proto.CompactTextString(m) }
func (*BfAccountData) ProtoMessage()               {}
func (*BfAccountData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// 错误数据类
type BfErrorData struct {
	ErrorId        string `protobuf:"bytes,1,opt,name=errorId" json:"errorId,omitempty"`
	ErrorMsg       string `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
	AdditionalInfo string `protobuf:"bytes,3,opt,name=additionalInfo" json:"additionalInfo,omitempty"`
}

func (m *BfErrorData) Reset()                    { *m = BfErrorData{} }
func (m *BfErrorData) String() string            { return proto.CompactTextString(m) }
func (*BfErrorData) ProtoMessage()               {}
func (*BfErrorData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// 日志数据类
type BfLogData struct {
	LogTime    string `protobuf:"bytes,1,opt,name=logTime" json:"logTime,omitempty"`
	LogContent string `protobuf:"bytes,2,opt,name=logContent" json:"logContent,omitempty"`
}

func (m *BfLogData) Reset()                    { *m = BfLogData{} }
func (m *BfLogData) String() string            { return proto.CompactTextString(m) }
func (*BfLogData) ProtoMessage()               {}
func (*BfLogData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// 合约详细信息类
type BfContractData struct {
	// 代码编号相关
	Symbol         string    `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange       string    `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Name           string    `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ProductClass   BfProduct `protobuf:"varint,4,opt,name=productClass,enum=bftrader.BfProduct" json:"productClass,omitempty"`
	VolumeMultiple int32     `protobuf:"varint,5,opt,name=volumeMultiple" json:"volumeMultiple,omitempty"`
	PriceTick      float64   `protobuf:"fixed64,6,opt,name=priceTick" json:"priceTick,omitempty"`
}

func (m *BfContractData) Reset()                    { *m = BfContractData{} }
func (m *BfContractData) String() string            { return proto.CompactTextString(m) }
func (*BfContractData) ProtoMessage()               {}
func (*BfContractData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// K/V数据交换传入传出的数据类
type BfKvData struct {
	// kv对
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *BfKvData) Reset()                    { *m = BfKvData{} }
func (m *BfKvData) String() string            { return proto.CompactTextString(m) }
func (*BfKvData) ProtoMessage()               {}
func (*BfKvData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// 订阅行情时传入的对象类
type BfSubscribeReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
}

func (m *BfSubscribeReq) Reset()                    { *m = BfSubscribeReq{} }
func (m *BfSubscribeReq) String() string            { return proto.CompactTextString(m) }
func (*BfSubscribeReq) ProtoMessage()               {}
func (*BfSubscribeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// 发单时传入的对象类
type BfOrderReq struct {
	// 代码编号相关
	Symbol    string      `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange  string      `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Price     float64     `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
	Volume    int32       `protobuf:"varint,4,opt,name=volume" json:"volume,omitempty"`
	PriceType BfPriceType `protobuf:"varint,5,opt,name=priceType,enum=bftrader.BfPriceType" json:"priceType,omitempty"`
	Direction BfDirection `protobuf:"varint,6,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset    BfOffset    `protobuf:"varint,7,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
	Reason    string      `protobuf:"bytes,8,opt,name=reason" json:"reason,omitempty"`
}

func (m *BfOrderReq) Reset()                    { *m = BfOrderReq{} }
func (m *BfOrderReq) String() string            { return proto.CompactTextString(m) }
func (*BfOrderReq) ProtoMessage()               {}
func (*BfOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// 发单本地返回的对象类
type BfOrderResp struct {
	OrderId string `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
}

func (m *BfOrderResp) Reset()                    { *m = BfOrderResp{} }
func (m *BfOrderResp) String() string            { return proto.CompactTextString(m) }
func (*BfOrderResp) ProtoMessage()               {}
func (*BfOrderResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

// 撤单时传入的对象类
type BfCancelOrderReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	OrderId  string `protobuf:"bytes,3,opt,name=orderId" json:"orderId,omitempty"`
	Reason   string `protobuf:"bytes,4,opt,name=reason" json:"reason,omitempty"`
}

func (m *BfCancelOrderReq) Reset()                    { *m = BfCancelOrderReq{} }
func (m *BfCancelOrderReq) String() string            { return proto.CompactTextString(m) }
func (*BfCancelOrderReq) ProtoMessage()               {}
func (*BfCancelOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

// 连接时传入的对象类
type BfConnectReq struct {
	// Robot相关
	RobotId   string `protobuf:"bytes,1,opt,name=robotId" json:"robotId,omitempty"`
	RobotIp   string `protobuf:"bytes,2,opt,name=robotIp" json:"robotIp,omitempty"`
	RobotPort int32  `protobuf:"varint,3,opt,name=robotPort" json:"robotPort,omitempty"`
}

func (m *BfConnectReq) Reset()                    { *m = BfConnectReq{} }
func (m *BfConnectReq) String() string            { return proto.CompactTextString(m) }
func (*BfConnectReq) ProtoMessage()               {}
func (*BfConnectReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

// 连接时返回的对象类
type BfConnectResp struct {
	ExchangeOpened bool `protobuf:"varint,1,opt,name=exchangeOpened" json:"exchangeOpened,omitempty"`
}

func (m *BfConnectResp) Reset()                    { *m = BfConnectResp{} }
func (m *BfConnectResp) String() string            { return proto.CompactTextString(m) }
func (*BfConnectResp) ProtoMessage()               {}
func (*BfConnectResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

// 获取合约信息传入的对象类
type BfGetContractReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
}

func (m *BfGetContractReq) Reset()                    { *m = BfGetContractReq{} }
func (m *BfGetContractReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetContractReq) ProtoMessage()               {}
func (*BfGetContractReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

// 获取历史Tick数据的对象类（1分钟整存整取）
type BfGetTickPackReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	ToDate   string `protobuf:"bytes,3,opt,name=toDate" json:"toDate,omitempty"`
	ToTime   string `protobuf:"bytes,4,opt,name=toTime" json:"toTime,omitempty"`
	Count    int32  `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
}

func (m *BfGetTickPackReq) Reset()                    { *m = BfGetTickPackReq{} }
func (m *BfGetTickPackReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetTickPackReq) ProtoMessage()               {}
func (*BfGetTickPackReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

// 获取历史Bar数据的对象类
type BfGetBarReq struct {
	// 代码编号相关
	Symbol   string   `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string   `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Period   BfPeriod `protobuf:"varint,3,opt,name=period,enum=bftrader.BfPeriod" json:"period,omitempty"`
	ToDate   string   `protobuf:"bytes,4,opt,name=toDate" json:"toDate,omitempty"`
	ToTime   string   `protobuf:"bytes,5,opt,name=toTime" json:"toTime,omitempty"`
	Count    int32    `protobuf:"varint,6,opt,name=count" json:"count,omitempty"`
}

func (m *BfGetBarReq) Reset()                    { *m = BfGetBarReq{} }
func (m *BfGetBarReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetBarReq) ProtoMessage()               {}
func (*BfGetBarReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func init() {
	proto.RegisterType((*BfVoid)(nil), "bftrader.BfVoid")
	proto.RegisterType((*BfDailyInfo)(nil), "bftrader.BfDailyInfo")
	proto.RegisterType((*BfBidAskInfo)(nil), "bftrader.BfBidAskInfo")
	proto.RegisterType((*BfOHLCInfo)(nil), "bftrader.BfOHLCInfo")
	proto.RegisterType((*BfTickData)(nil), "bftrader.BfTickData")
	proto.RegisterType((*BfTickPackData)(nil), "bftrader.BfTickPackData")
	proto.RegisterType((*BfBarData)(nil), "bftrader.BfBarData")
	proto.RegisterType((*BfTradeData)(nil), "bftrader.BfTradeData")
	proto.RegisterType((*BfOrderData)(nil), "bftrader.BfOrderData")
	proto.RegisterType((*BfPositionData)(nil), "bftrader.BfPositionData")
	proto.RegisterType((*BfAccountData)(nil), "bftrader.BfAccountData")
	proto.RegisterType((*BfErrorData)(nil), "bftrader.BfErrorData")
	proto.RegisterType((*BfLogData)(nil), "bftrader.BfLogData")
	proto.RegisterType((*BfContractData)(nil), "bftrader.BfContractData")
	proto.RegisterType((*BfKvData)(nil), "bftrader.BfKvData")
	proto.RegisterType((*BfSubscribeReq)(nil), "bftrader.BfSubscribeReq")
	proto.RegisterType((*BfOrderReq)(nil), "bftrader.BfOrderReq")
	proto.RegisterType((*BfOrderResp)(nil), "bftrader.BfOrderResp")
	proto.RegisterType((*BfCancelOrderReq)(nil), "bftrader.BfCancelOrderReq")
	proto.RegisterType((*BfConnectReq)(nil), "bftrader.BfConnectReq")
	proto.RegisterType((*BfConnectResp)(nil), "bftrader.BfConnectResp")
	proto.RegisterType((*BfGetContractReq)(nil), "bftrader.BfGetContractReq")
	proto.RegisterType((*BfGetTickPackReq)(nil), "bftrader.BfGetTickPackReq")
	proto.RegisterType((*BfGetBarReq)(nil), "bftrader.BfGetBarReq")
	proto.RegisterEnum("bftrader.BfDirection", BfDirection_name, BfDirection_value)
	proto.RegisterEnum("bftrader.BfOffset", BfOffset_name, BfOffset_value)
	proto.RegisterEnum("bftrader.BfStatus", BfStatus_name, BfStatus_value)
	proto.RegisterEnum("bftrader.BfProduct", BfProduct_name, BfProduct_value)
	proto.RegisterEnum("bftrader.BfPriceType", BfPriceType_name, BfPriceType_value)
	proto.RegisterEnum("bftrader.BfExchange", BfExchange_name, BfExchange_value)
	proto.RegisterEnum("bftrader.BfPeriod", BfPeriod_name, BfPeriod_value)
}

var fileDescriptor0 = []byte{
	// 1672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x58, 0xcd, 0x6e, 0xdb, 0xd6,
	0x12, 0xbe, 0xfa, 0xb5, 0x74, 0xfc, 0x13, 0x86, 0x71, 0x72, 0x7d, 0x2f, 0x2e, 0x2e, 0x02, 0xa2,
	0x68, 0x03, 0x2d, 0x82, 0xc6, 0x01, 0x9a, 0x4d, 0x37, 0xfa, 0xa1, 0x6c, 0xc1, 0xb2, 0xc8, 0x1e,
	0xd1, 0x69, 0x9c, 0x4d, 0x4b, 0x49, 0x94, 0x4d, 0x98, 0x12, 0x55, 0x92, 0x72, 0xe2, 0x6e, 0x0a,
	0x04, 0x05, 0xda, 0x45, 0x1f, 0xa0, 0x8b, 0xae, 0x0a, 0xf4, 0x09, 0xfa, 0x18, 0x5d, 0x76, 0xd1,
	0x07, 0xe8, 0x4b, 0x14, 0xe8, 0xa6, 0x33, 0xe7, 0x87, 0x3c, 0x54, 0x9d, 0x36, 0x51, 0xd0, 0xa2,
	0x3b, 0xce, 0x37, 0xc3, 0x33, 0x33, 0xdf, 0xcc, 0x99, 0xa1, 0x44, 0x76, 0x46, 0xd3, 0x24, 0x72,
	0x27, 0x5e, 0x74, 0x7f, 0x11, 0x85, 0x49, 0xa8, 0xd7, 0xa4, 0x6c, 0xd4, 0x48, 0xb5, 0x35, 0x7d,
	0x1c, 0xfa, 0x13, 0xe3, 0x87, 0x02, 0xd9, 0x6c, 0x4d, 0x3b, 0xae, 0x1f, 0x5c, 0xf5, 0xe6, 0xd3,
	0x50, 0xff, 0x1f, 0xa9, 0x87, 0x0b, 0x6f, 0x6e, 0x47, 0xfe, 0xd8, 0xdb, 0x2b, 0xdc, 0x2d, 0xdc,
	0x2b, 0xd0, 0x0c, 0x40, 0xed, 0xb9, 0x7f, 0x76, 0xce, 0xb5, 0x45, 0xae, 0x4d, 0x01, 0xfd, 0xbf,
	0xa4, 0x16, 0x84, 0xcf, 0xb8, 0xb2, 0xc4, 0x94, 0xa9, 0xac, 0xbf, 0x45, 0xb6, 0x17, 0x91, 0xd7,
	0x0e, 0xc2, 0xd8, 0xe3, 0x06, 0x65, 0x66, 0x90, 0x07, 0xf5, 0xff, 0x13, 0xb2, 0x5c, 0x2c, 0xbc,
	0xa8, 0xef, 0xcf, 0xfc, 0x64, 0xaf, 0xc2, 0x4c, 0x14, 0x04, 0xf5, 0x70, 0xa2, 0xd4, 0x57, 0xb9,
	0x3e, 0x43, 0x8c, 0x17, 0x05, 0xb2, 0xd5, 0x9a, 0xb6, 0xfc, 0x49, 0x33, 0xbe, 0x60, 0xe9, 0x40,
	0x48, 0x23, 0x7f, 0xa2, 0x66, 0x93, 0xca, 0xa8, 0x73, 0xe3, 0x0b, 0x35, 0x97, 0x54, 0xc6, 0x44,
	0xc1, 0xee, 0x71, 0x18, 0x2c, 0x67, 0x3c, 0x97, 0x0a, 0xcd, 0x00, 0xd4, 0x82, 0xa5, 0xd0, 0x96,
	0xb9, 0x36, 0x05, 0x8c, 0xcf, 0x0b, 0x84, 0xb4, 0xa6, 0xd6, 0x61, 0xbf, 0xfd, 0x97, 0x32, 0x0a,
	0x5c, 0x8c, 0x57, 0xe9, 0x54, 0x10, 0xe3, 0xd7, 0x22, 0x86, 0xe1, 0xf8, 0xe3, 0x8b, 0x8e, 0x9b,
	0xb8, 0xfa, 0x1d, 0x52, 0x8d, 0xaf, 0x66, 0xa3, 0x30, 0x60, 0x31, 0xd4, 0xa9, 0x90, 0xd0, 0x85,
	0xf7, 0x7c, 0x7c, 0xee, 0xce, 0xcf, 0xb8, 0xff, 0x3a, 0x4d, 0x65, 0x5d, 0x27, 0xe5, 0x89, 0x9b,
	0x70, 0xd7, 0x75, 0xca, 0x9e, 0x11, 0x4b, 0x7c, 0x91, 0x36, 0x60, 0xf8, 0x8c, 0x49, 0x04, 0x6e,
	0x9c, 0xf0, 0x48, 0x78, 0xd5, 0x32, 0x00, 0x3d, 0x5f, 0x72, 0xaa, 0xaa, 0x8c, 0x2a, 0x21, 0xe9,
	0x06, 0xd9, 0x42, 0x1e, 0x7a, 0xf3, 0xc4, 0x8b, 0xbc, 0x38, 0xd9, 0xdb, 0x60, 0x2f, 0xe6, 0x30,
	0x56, 0x70, 0x38, 0x48, 0x50, 0x5d, 0x63, 0xef, 0x2b, 0x88, 0xde, 0x20, 0x1a, 0x4a, 0x96, 0x7a,
	0x4e, 0x9d, 0x9d, 0xf3, 0x3b, 0x5c, 0x7f, 0x8f, 0x10, 0x28, 0xa1, 0xcb, 0x3b, 0x63, 0x8f, 0xdc,
	0x2d, 0xdd, 0xdb, 0xdc, 0xbf, 0x73, 0x3f, 0xbd, 0x23, 0x6a, 0xdf, 0x50, 0xc5, 0x52, 0x7f, 0x48,
	0xea, 0x13, 0x79, 0x3f, 0xf6, 0x36, 0xe1, 0xf0, 0xcd, 0xfd, 0xdb, 0xea, 0x6b, 0xe9, 0xe5, 0xa1,
	0x99, 0x9d, 0xf1, 0x75, 0x81, 0xec, 0x70, 0xf6, 0x6d, 0xf7, 0x6f, 0xa8, 0x40, 0x83, 0x54, 0x12,
	0xf0, 0x15, 0x03, 0xfb, 0x98, 0xd6, 0xae, 0x1a, 0x9f, 0x6c, 0x01, 0xca, 0x4d, 0x8c, 0x17, 0x25,
	0x52, 0x87, 0x64, 0xdd, 0x68, 0xed, 0xa8, 0x1a, 0xa4, 0x0a, 0x57, 0xd2, 0x0f, 0x27, 0x2c, 0xae,
	0x9d, 0x7d, 0x5d, 0x75, 0x67, 0x33, 0x0d, 0x15, 0x16, 0x69, 0x06, 0xe5, 0x6b, 0x32, 0xa8, 0x28,
	0x19, 0xfc, 0x53, 0xba, 0xe4, 0x5d, 0x52, 0x0b, 0xcf, 0x83, 0xb1, 0xe8, 0x91, 0xc2, 0x2a, 0x99,
	0xf2, 0x5a, 0xd3, 0xd4, 0x6a, 0xbd, 0xfe, 0xf8, 0xb6, 0x88, 0x73, 0xd7, 0x41, 0x9b, 0xb5, 0xcb,
	0xb0, 0x47, 0x36, 0x98, 0x93, 0xde, 0x44, 0xf4, 0x87, 0x14, 0x51, 0x13, 0x46, 0xe0, 0x1d, 0x34,
	0x9c, 0x77, 0x29, 0xb2, 0x60, 0xfd, 0xc8, 0x1b, 0x27, 0x7e, 0x38, 0x67, 0xfc, 0xef, 0xac, 0x04,
	0x2b, 0x95, 0x34, 0xb3, 0xc3, 0x7a, 0x87, 0xd3, 0x69, 0xec, 0xf1, 0x91, 0xbb, 0x52, 0x6f, 0x8b,
	0x69, 0xa8, 0xb0, 0xd0, 0x77, 0x49, 0x65, 0xc1, 0xe6, 0x00, 0x2f, 0x14, 0x17, 0x94, 0xea, 0xd6,
	0x72, 0xd5, 0x85, 0xc9, 0xc1, 0x0e, 0x72, 0xb0, 0x1d, 0xea, 0x2c, 0xd4, 0x0c, 0xc0, 0x4e, 0x05,
	0x92, 0x2c, 0x0c, 0xfd, 0x4d, 0x48, 0x92, 0x54, 0x94, 0xfe, 0x80, 0x8a, 0xf2, 0x6b, 0x53, 0x51,
	0x79, 0x75, 0x2a, 0xaa, 0x2a, 0x15, 0x77, 0xc9, 0x66, 0x12, 0x26, 0x6e, 0x20, 0xba, 0x75, 0x83,
	0xf1, 0xa1, 0x42, 0xd8, 0xf2, 0xec, 0xc8, 0x49, 0xae, 0xa1, 0x73, 0x18, 0xc6, 0x11, 0x27, 0x6e,
	0xb2, 0x8c, 0x19, 0x6b, 0x2b, 0x71, 0x0c, 0x99, 0x86, 0x0a, 0x0b, 0xb6, 0x81, 0x30, 0x67, 0x46,
	0x32, 0xe1, 0x24, 0xa7, 0x00, 0xdb, 0x23, 0xee, 0x7c, 0xec, 0x05, 0x4c, 0xbd, 0xc9, 0xd4, 0x0a,
	0x62, 0xfc, 0xcc, 0x26, 0x99, 0x1d, 0xc6, 0x3e, 0x12, 0xb0, 0x76, 0x1d, 0x72, 0x6c, 0x97, 0x5e,
	0x91, 0x6d, 0x38, 0x70, 0x21, 0x1c, 0x8b, 0x3d, 0x9b, 0xca, 0x18, 0xc4, 0x34, 0x0a, 0x3f, 0xf5,
	0x78, 0x1b, 0x43, 0x4b, 0x71, 0xe9, 0x25, 0xac, 0x43, 0x96, 0x57, 0x13, 0x99, 0x84, 0x20, 0x5d,
	0x41, 0x8c, 0x2f, 0x8a, 0x64, 0xbb, 0x35, 0x6d, 0x8e, 0xc7, 0xe1, 0x72, 0x9e, 0xb0, 0x24, 0x71,
	0xc9, 0x73, 0x11, 0x5a, 0x87, 0xe7, 0x99, 0x01, 0x78, 0x1e, 0x7c, 0xba, 0xb4, 0xdc, 0x00, 0x89,
	0x12, 0x8b, 0x5b, 0x41, 0xb0, 0xed, 0x46, 0x42, 0xc9, 0x17, 0xb7, 0x14, 0xd9, 0xb9, 0x97, 0x30,
	0x07, 0xdc, 0x51, 0x20, 0xd7, 0x76, 0x06, 0xb0, 0x6a, 0x84, 0xb3, 0x99, 0x1f, 0xc7, 0xf2, 0x82,
	0xe2, 0x56, 0x4f, 0x11, 0xcc, 0x7a, 0xe6, 0x46, 0x67, 0xfe, 0x5c, 0xa4, 0x27, 0x24, 0xec, 0x2a,
	0xb1, 0xfb, 0xc3, 0xa9, 0x2f, 0xa7, 0xa4, 0x0a, 0xe9, 0x6f, 0x93, 0x1d, 0xc9, 0x9d, 0x30, 0xaa,
	0x31, 0xa3, 0x15, 0xd4, 0xb8, 0xc0, 0x3b, 0x67, 0x46, 0x51, 0xc8, 0xef, 0x1c, 0x24, 0xe2, 0xa1,
	0x90, 0x92, 0x20, 0x45, 0x56, 0x6d, 0x7c, 0x3c, 0x8e, 0xcf, 0xd2, 0x6a, 0x0b, 0x19, 0x9d, 0xb9,
	0x93, 0x09, 0x3b, 0xd6, 0x0d, 0xd8, 0x60, 0xe4, 0x97, 0x6f, 0x05, 0x35, 0x4c, 0x5c, 0x45, 0xfd,
	0xf0, 0x4c, 0xba, 0x0a, 0xc2, 0x33, 0xd6, 0x86, 0xc2, 0x95, 0x10, 0xf9, 0x77, 0xdf, 0x59, 0x3b,
	0x84, 0x21, 0x3d, 0x4f, 0x84, 0x33, 0x05, 0x31, 0x7e, 0x62, 0x3d, 0x8a, 0x52, 0xe4, 0x8e, 0x93,
	0x37, 0xd9, 0xb6, 0x73, 0x77, 0x96, 0x6e, 0x5b, 0x7c, 0xd6, 0x1f, 0x91, 0x2d, 0xf8, 0x7a, 0x9e,
	0x2c, 0xc7, 0x49, 0x1b, 0x76, 0x45, 0x2c, 0x06, 0xc5, 0xad, 0xdc, 0xc6, 0xe3, 0x7a, 0x9a, 0x33,
	0x44, 0x0a, 0xf8, 0x90, 0x3b, 0x5e, 0x06, 0x89, 0xbf, 0x08, 0x3c, 0xd1, 0xa7, 0x2b, 0x28, 0xf6,
	0x03, 0x6b, 0x51, 0x5c, 0xd3, 0xa2, 0xa8, 0x19, 0x60, 0xec, 0x93, 0x5a, 0x6b, 0x7a, 0x74, 0xc9,
	0x52, 0xd2, 0x48, 0xe9, 0xc2, 0xbb, 0x12, 0xf9, 0xe0, 0x23, 0xf6, 0xfa, 0xa5, 0x1b, 0x2c, 0x65,
	0x26, 0x5c, 0x30, 0x3a, 0x48, 0xc6, 0x70, 0x39, 0x8a, 0xc7, 0x91, 0x3f, 0xf2, 0xa8, 0xf7, 0xc9,
	0x3a, 0x64, 0x18, 0xdf, 0xb0, 0xef, 0x47, 0x36, 0x7c, 0xd7, 0x3c, 0x22, 0xbb, 0x8a, 0xa5, 0xeb,
	0x77, 0x41, 0x39, 0xb7, 0x0b, 0x1e, 0x4a, 0x22, 0xae, 0x16, 0xde, 0x75, 0xab, 0xc9, 0x96, 0x4a,
	0x9a, 0xd9, 0xe5, 0xc7, 0x4a, 0xf5, 0xb5, 0x87, 0xf8, 0xc6, 0x9f, 0x0e, 0x71, 0x88, 0x36, 0xf2,
	0xdc, 0x18, 0x4e, 0xaf, 0xf1, 0xbc, 0xb9, 0x64, 0xbc, 0x93, 0xae, 0x26, 0xea, 0xc5, 0x0b, 0x75,
	0xcd, 0x14, 0x72, 0x6b, 0xc6, 0x78, 0x4e, 0x34, 0x68, 0x4d, 0x36, 0x4f, 0xdf, 0x88, 0xcc, 0x97,
	0x2f, 0xb2, 0x2c, 0xc4, 0x72, 0x2e, 0xc4, 0x8f, 0xf1, 0xc7, 0x10, 0x5c, 0x8a, 0x39, 0xa4, 0x8d,
	0x5e, 0xe1, 0x84, 0x28, 0x1c, 0x85, 0xd9, 0x3c, 0x93, 0x62, 0xa6, 0x59, 0x08, 0xb7, 0x52, 0xc4,
	0xee, 0x64, 0x8f, 0x76, 0x18, 0x25, 0xf2, 0x87, 0x50, 0x0a, 0x18, 0x8f, 0x70, 0x68, 0xa6, 0x1e,
	0x80, 0x06, 0x68, 0x7a, 0x19, 0x30, 0x7e, 0x55, 0x79, 0xdc, 0x53, 0x8d, 0xae, 0xa0, 0x46, 0x17,
	0x49, 0x39, 0xf0, 0x12, 0x79, 0x65, 0xd7, 0x6d, 0xd2, 0xaf, 0x0a, 0xe2, 0x20, 0xf9, 0xa5, 0xbd,
	0x2e, 0xbb, 0xf0, 0x4e, 0x12, 0x76, 0xb2, 0x4f, 0x6d, 0x21, 0x71, 0xdc, 0xc9, 0x3e, 0xb7, 0x85,
	0x84, 0xad, 0xcd, 0x56, 0x81, 0xb8, 0xd4, 0x5c, 0x30, 0xbe, 0x67, 0xbf, 0xa6, 0x21, 0x1c, 0xf8,
	0xba, 0x5e, 0x37, 0x92, 0xd7, 0xf9, 0xb8, 0xce, 0xa2, 0x2e, 0xbf, 0x24, 0xea, 0xca, 0xf5, 0x51,
	0x57, 0x95, 0xa8, 0x1b, 0x5f, 0xf2, 0xff, 0x00, 0xd2, 0xeb, 0x71, 0x9b, 0xdc, 0xec, 0xf4, 0xa8,
	0xd9, 0x76, 0x7a, 0xd6, 0xe0, 0xa3, 0x93, 0xc1, 0xd1, 0xc0, 0xfa, 0x70, 0xa0, 0xfd, 0x0b, 0xa6,
	0xe3, 0x4e, 0x06, 0x0f, 0xac, 0x81, 0xa9, 0x15, 0xf2, 0x58, 0xdf, 0x1a, 0x1c, 0x68, 0x45, 0xfd,
	0x16, 0xb9, 0x91, 0x61, 0xc3, 0x43, 0x8b, 0x3a, 0x5a, 0x49, 0xbf, 0x49, 0xb6, 0x95, 0x97, 0x4d,
	0x47, 0x2b, 0xe7, 0xdf, 0x1d, 0x9a, 0xfd, 0xbe, 0x56, 0x69, 0xc0, 0x6f, 0xe7, 0x9a, 0xbc, 0x82,
	0x68, 0x60, 0x75, 0xbb, 0x43, 0xd3, 0x51, 0x82, 0xb8, 0x41, 0x36, 0x05, 0x26, 0x22, 0xc8, 0x00,
	0xcb, 0x36, 0x07, 0xe0, 0x5e, 0x23, 0x5b, 0x02, 0x68, 0xf7, 0xad, 0xa1, 0x09, 0xbe, 0x21, 0x1f,
	0x15, 0x71, 0xac, 0x4e, 0xf3, 0x14, 0xfc, 0xff, 0x87, 0xdc, 0x56, 0xe1, 0x53, 0x73, 0xe8, 0x98,
	0x14, 0x55, 0x95, 0xc6, 0x33, 0x8c, 0x82, 0x7f, 0x45, 0x61, 0x14, 0x43, 0xa7, 0xe9, 0x9c, 0x0c,
	0x95, 0x28, 0x76, 0x89, 0x26, 0xb0, 0x81, 0xe5, 0x38, 0xb4, 0xd9, 0x31, 0x3b, 0x10, 0x0a, 0xf8,
	0x11, 0xa8, 0xdd, 0xa4, 0x12, 0x2e, 0x2a, 0xc6, 0xcd, 0x7e, 0x5f, 0xa0, 0x25, 0x05, 0x6d, 0x37,
	0x07, 0x6d, 0xc8, 0x1e, 0xd0, 0x72, 0xe3, 0xc7, 0x02, 0x2e, 0x44, 0xb1, 0x50, 0x90, 0x49, 0x9b,
	0x5a, 0x9d, 0x93, 0xb6, 0xca, 0x00, 0xe4, 0x27, 0xc1, 0xac, 0x08, 0x12, 0x31, 0x3f, 0x38, 0xe9,
	0x39, 0xa7, 0xbc, 0x08, 0x12, 0xeb, 0x9e, 0x38, 0x27, 0xd4, 0x1c, 0x82, 0x4f, 0xc5, 0xd0, 0xb2,
	0x91, 0x76, 0x60, 0x01, 0x0a, 0x23, 0xb1, 0xde, 0xa0, 0x63, 0x3e, 0xd1, 0x2a, 0xfa, 0xbf, 0xc9,
	0x2d, 0x09, 0xb5, 0xad, 0xe3, 0x56, 0x6f, 0xd0, 0x64, 0xb6, 0x55, 0xd5, 0xb6, 0x6b, 0x51, 0xb0,
	0xdd, 0x50, 0xa3, 0x19, 0xda, 0x96, 0xa3, 0xd5, 0x54, 0xa3, 0x8e, 0xd9, 0x35, 0xa9, 0x56, 0x6f,
	0x7c, 0x86, 0xfd, 0x95, 0x8e, 0x6f, 0xe4, 0xc9, 0xa6, 0xbd, 0xb6, 0xe9, 0x9c, 0xda, 0x26, 0x26,
	0x66, 0x0d, 0x58, 0x62, 0x7b, 0x64, 0x37, 0x83, 0xfb, 0xbd, 0xe3, 0x9e, 0xc3, 0x44, 0x48, 0x10,
	0x2a, 0x95, 0x69, 0x8e, 0x9b, 0xf4, 0xc8, 0x14, 0xaa, 0x22, 0xf7, 0x26, 0x55, 0xdd, 0xe6, 0x11,
	0x6f, 0x35, 0x05, 0xb2, 0x8e, 0x80, 0xd6, 0xef, 0xd8, 0x5f, 0x32, 0x66, 0xb6, 0x97, 0x34, 0xf3,
	0x49, 0xfb, 0xb0, 0x39, 0x38, 0x30, 0xf3, 0xc4, 0xa6, 0xe8, 0x70, 0x88, 0x7e, 0xe1, 0xa4, 0x0c,
	0x79, 0x3a, 0x44, 0x7f, 0x40, 0x61, 0x0a, 0xb5, 0xbb, 0x5d, 0xe0, 0xa0, 0x94, 0x37, 0x3b, 0xec,
	0x9a, 0x9c, 0xd5, 0xcc, 0xec, 0x29, 0x44, 0x5a, 0xc9, 0x1d, 0xdf, 0x01, 0xa4, 0x9a, 0x77, 0x78,
	0x60, 0x6a, 0x1b, 0x8d, 0xf7, 0xb1, 0xef, 0x6c, 0xf9, 0xc3, 0x79, 0xc7, 0x36, 0x69, 0xcf, 0xea,
	0x28, 0x21, 0x6e, 0x93, 0xba, 0xc0, 0x8e, 0x1f, 0x40, 0x7c, 0x99, 0xd8, 0x79, 0xa0, 0x15, 0x5b,
	0xa5, 0x5f, 0x0a, 0x85, 0x51, 0x95, 0xfd, 0xd7, 0xf7, 0xf0, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x74, 0x46, 0xa3, 0xfd, 0x13, 0x00, 0x00,
}
