package block_chain_test

// type block struct {
// 	data 		string
// 	hash 		string
// 	prevHash 	string
// }

// func Test1() {
// 	genesisBlock := block{"Genesis Block", "", ""}

// 	// byte slice vs string
// 	// for _, aByte := range "Genesis Block" {
// 	// 	fmt.Printf("%b\n",aByte)
// 	// }

// 	// 이는 에러 발생!
// 	// sha256.Sum256이 반환하는 값은 [32]byte이기 때문에 string 타입인 block.hash로는 받을 수 없다.
// 	// genesisBlock.hash = sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
// 	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
// 	fmt.Println(hash)
// 	/*
// 	[137 235 10 192 49 166 61 36 33 205 5 162 251 228 31 62 163 95 92 55 18 202 131 156 191 107 133 196 238 7 183 163]
// 	*/

// 	// 대부분의 hash는 16진수로 되어있다. 이 말은 base16으로 인코딩 해주면 된다는 것 이다.
// 	hexHash := fmt.Sprintf("%x", hash)
// 	genesisBlock.hash = hexHash
// 	fmt.Println(genesisBlock.hash)
// 	/*
// 	89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
// 	*/
// }

// type blockchain struct {
// 	blocks []block
// }

// func (b *blockchain) getLastHash() string {
// 	if len(b.blocks) > 0 {
// 		return b.blocks[len(b.blocks) - 1].hash
// 	}
// 	return ""
// }

// func (b *blockchain) addBlock(data string) {
// 	newBlock := block{data, "", ""}

// 	// 1. previous Hash parse
// 	newBlock.prevHash = b.getLastHash()

// 	// 2. set hash = hash(data + prevHash)
// 	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
// 	newBlock.hash = fmt.Sprintf("%x", hash)

// 	// 3. adding block in chain
// 	b.blocks = append(b.blocks, newBlock)
// }

// func (b *blockchain) listBlocks() {
// 	for _, block := range b.blocks {
// 		fmt.Printf("Data: %s\n",block.data)
// 		fmt.Printf("Hash: %s\n",block.hash)
// 		fmt.Printf("Prev Hash: %s\n\n",block.prevHash)
// 	}
// }

// func Test2() {
// 	chain := blockchain{}
// 	chain.addBlock("Genesis Block")
// 	chain.addBlock("Second Block")
// 	chain.addBlock("Third Block")
// 	chain.listBlocks()

// 	// genesisBlock := block{"Genesis Block", "", ""}
// 	// hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
// 	// hexHash := fmt.Sprintf("%x", hash)
// 	// genesisBlock.hash = hexHash

// 	// // 이전 블록의 hash 값으로 이루어진 연결고리가 바로 블록체인
// 	// secodeBlocks := block{"Second Block", "", genesisBlock.hash}

// }

func Test() {
	// chain := blockchain.GetBlockchain()

	// chain.AddBlock("Second Block")
	// chain.AddBlock("Third Block")
	// chain.AddBlock("Fourth Block")

	// for _, block := range chain.AllBlocks() {
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %s\n", block.Hash)
	// 	fmt.Printf("Prev Hash: %s\n\n", block.PrevHash)
	// }
}
