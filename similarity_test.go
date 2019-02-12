package simhash

import (
	"testing"
	"fmt"
)

func TestSimHashSimilar(t *testing.T) {
	g := NewGoJieba()
	srcStr := RemoveHtml("关于区块链和数字货币的关系，很多人或多或少都存在疑惑。简单来说，区块链是比特币的底层运用，而比特币只是区块链的一个小应用而已。" +
		"数字货币即虚拟货币，最早的数字货币诞生于2009年，其发明者中本聪为了应对经济危机对于实体货币经济的冲击。比特币是最早的数字货币，后来出现了以太币、火币以及莱特币等虚拟货币，这些虚拟货币是不能用来交易的。" +
		"狭义来讲，区块链是一种按照时间顺序将数据区块以顺序相连的方式组合成的一种链式数据结构， 并以密码学方式保证的不可篡改和不可伪造的分布式账本。" +
		"广义来讲，区块链技术是利用块链式数据结构来验证与存储数据、利用分布式节点共识算法来生成和更新数据、利用密码学的方式保证数据传输和访问的安全、利用由自动化脚本代码组成的智能合约来编程和操作数据的一种全新的分布式基础架构与计算方式。")
	dstStr := RemoveHtml("区块链技术为我们的信息防伪与数据追踪提供了革新手段。区块链中的数据区块顺序相连构成了一个不可篡改的数据链条，时间戳为所有的交易行为贴上了一套不讲课伪造的真是标签，这对于人们在现实生活中打击假冒伪劣产品大有裨益； " +
		"市场分析指出，整体而言，区块链技术目前在十大金融领域显示出应用前景，分别是资产证券化、保险、供应链金融、场外市场、资产托管、大宗商品交易、风险信息共享机制、贸易融资、银团贷款、股权交易交割。" +
		"这些金融场景有三大共性：参与节点多、验真成本高、交易流程长，而区块链的分布式记账、不可篡改、内置合约等特性可以为这些金融业务中的痛点提供解决方案。" +
		"传统的工业互联网模式是由一个中心化的机构收集和管理所有的数据信息，容易产生因设备生命周期和安全等方面的缺陷引起的数据丢失、篡改等问题。区块链技术可以在无需任何信任单个节点的同时构建整个网络的信任共识，从而很好的解决目前工业互联网技术领域的一些缺陷，让物与物之间能够实现更好的连接.")
	srcWordsWeight := g.C.ExtractWithWeight(srcStr, 30)
	dstWordsWeight := g.C.ExtractWithWeight(dstStr, 30)
	fmt.Printf("srcWordsWeight: %v\n", srcWordsWeight)
	fmt.Printf("dstWordsWeight: %v\n", dstWordsWeight)

	srcWords := make([]WordWeight, len(srcWordsWeight))
	dstWords := make([]WordWeight, len(dstWordsWeight))
	for i, ww := range srcWordsWeight {
		word := WordWeight{Word: ww.Word, Weight: ww.Weight}
		srcWords[i] = word
	}
	for i, ww := range dstWordsWeight {
		word := WordWeight{Word: ww.Word, Weight: ww.Weight}
		dstWords[i] = word
	}
	fmt.Printf("srcWords:%v\n", srcWords)
	fmt.Printf("dstWords:%v\n", dstWords)

	distance, err := SimHashSimilar(srcWords, dstWords)
	if err != nil {
		t.Errorf("failed: %v", err.Error())
	}

	t.Logf("SimHashSimilar distance: %v", distance)
}
