package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"time"
)

/**
 *区块结构体的定义
 */
type Block struct {
	Height		int64 	//区块高度
	TimeStamp	int64	//时间戳
	Hash		[]byte	//区块的hash
	Data		[]byte	//数据
	PreHash		[]byte	//上一个区块的hash
	Version		string	//版本号
	Nonce		int64	//随机数，用于pow工作量证明算法计算
}
/**
 *新建一个区块实例，并返回该区块
 */
func NewBlock(height int64, data []byte, prevHash []byte) (Block) {
	block := Block{
		Height:    height + 1,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PreHash:   prevHash,
		Version:   "0x01",
	}
	//调用util.SHA256Has进行Hash计算
	heightBytes,_ := util.IntToBytes(block.Height)
	timeBytes,_ := util.IntToBytes(block.TimeStamp)
	versionBytes := util.StringToBytes(block.Version)
	//bytes.Join用于[]byte的拼接
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeBytes,
		data,
		prevHash,
		versionBytes,
	},[]byte{})
	block.Hash = util.SHA256Hash(blockBytes)
	return block
}