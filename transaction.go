package main

import "math/big"

type Transaction struct {
	Timestamp        big.Int `json:"timestamp"`         // 交易时间
	BlockNumber      big.Int `json:"block_number"`      // 区块号
	TokenValue       big.Int `json:"token_value"`       // 代币数量
	Gas              big.Int `json:"gas"`               // gas
	GasPrice         big.Int `json:"gas_price"`         // gas price
	UsedGas          big.Int `json:"used_gas"`          // used gas
	Value            big.Int `json:"value"`             // eth number
	Hash             string  `json:"hash"`              // tx id
	Nonce            string  `json:"nonce"`             // tx nonce
	BlockHash        string  `json:"block_hash"`        // tx blockHash
	TransactionIndex big.Int `json:"transaction_index"` // tx idx in block
	LogIndex         big.Int `json:"log_index"`
	InternalIndex    string  `json:"internal_index"` //字符串处理，　默认是空, 第一层0, 第二层0_0,
	OpCode           string  `json:"op_code"`
	From             string  `json:"from"`       //  发起者
	To               string  `json:"to"`         // 接受者（合约地址）
	AddrToken        string  `json:"addr_token"` //  代币地址
	TokenType        uint64  `json:"token_type"` // 类型 1 表示是代币 0 表示Eth
	Input            string  `json:"input"`
	Err              string  `json:"err"`    //如果出错　显示错误信息
	Status           uint64  `json:"status"` //1 (success) or 0 (failure) or 2(pending), 99未知
}