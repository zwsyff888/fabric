package models

import (
	// "fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MblockData struct {
	Data []*TransData
}

type TransData struct {
	Txid    string
	ChainID string
	// Time    int64 //暂时去掉time的获取
}

func NewMblockData() *MblockData {
	return &MblockData{}
}

func (m *MblockData) initMblockData(input *pb.MblockData) {
	// fmt.Println("chenqiao: input.Data ", input.Data)
	// m.Data = make([]string, len(input.Data))
	// copy(m.Data, input.Data)
	// fmt.Println("chenqiao: m.Data ", m.Data)
	length := len(input.Datas)

	// m.Data = make([]*TransData{}, length)

	for i := 0; i < length; i++ {
		tmp := &TransData{}
		tmp.Txid = input.Datas[i].Txid
		// fmt.Println("@@@@chenqiao chainid: ", input.Datas[i].ChainID)
		tmp.ChainID = input.Datas[i].ChainID
		// if input.Datas[i].Time != nil {
		// tmp.Time = int64(input.Datas[i].Time.Seconds)
		// } else {
		// tmp.Time = 0
		// }

		m.Data = append(m.Data, tmp)
	}
}
