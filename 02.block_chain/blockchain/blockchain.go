package blockchain

type block struct {
	data 		string
	hash 		string
	prevHash 	string 
}

type blockchain struct {
	blocks []block
}

// 1. singletone variable
var b *blockchain

// 2. control function
func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}