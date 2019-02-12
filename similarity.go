package simhash

import (
	"hash/fnv"
	"fmt"
	"strings"
)

const (
	SIMILAR_DISTANCE = 3
)

type WordWeight struct {
	Word   string
	Weight float64
}

func SimHashSimilar(srcWordWeighs, dstWordWeights []WordWeight) (distance int, err error) {

	srcFingerPrint, err := simhashFingerPrint(srcWordWeighs)
	if err != nil {
		return
	}
	fmt.Println("srcFingerPrint: ", srcFingerPrint)
	dstFingerPrint, err := simhashFingerPrint(dstWordWeights)
	if err != nil {
		return
	}
	fmt.Println("dstFingerPrint: ", dstFingerPrint)

	distance = hammingDistance(srcFingerPrint, dstFingerPrint)

	return
}

func simhashFingerPrint(wordWeights []WordWeight) (fingerPrint []string, err error) {
	binaryWeights := make([]float64, 32)
	for _, ww := range wordWeights {
		bitHash := strHashBitCode(ww.Word)
		weights := calcWithWeight(bitHash, ww.Weight) //binary每个元素与weight的乘积结果数组
		binaryWeights, err = sliceInnerPlus(binaryWeights, weights)
		//fmt.Printf("ww.Word:%v, bitHash:%v, ww.Weight:%v, binaryWeights: %v\n", ww.Word,bitHash, ww.Weight, binaryWeights)
		if err != nil {
			return
		}
	}
	fingerPrint = make([]string, 0)
	for _, b := range binaryWeights {
		if b > 0 { // bit 1
			fingerPrint = append(fingerPrint, "1")
		} else { // bit 0
			fingerPrint = append(fingerPrint, "0")
		}
	}

	return
}

func strHashBitCode(str string) string {
	h := fnv.New32a()
	h.Write([]byte(str))
	b := int64(h.Sum32())
	return fmt.Sprintf("%032b", b)
}

func calcWithWeight(bitHash string, weight float64) []float64 {
	bitHashs := strings.Split(bitHash, "")
	binarys := make([]float64, 0)

	for _, bit := range bitHashs {
		if bit == "0" {
			binarys = append(binarys, float64(-1)*weight)
		} else {
			binarys = append(binarys, float64(weight))
		}
	}

	return binarys
}

func sliceInnerPlus(arr1, arr2 [] float64) (dstArr []float64, err error) {
	dstArr = make([]float64, len(arr1), len(arr1))

	if arr1 == nil || arr2 == nil {
		err = fmt.Errorf("sliceInnerPlus array nil")
		return
	}
	if len(arr1) != len(arr2) {
		err = fmt.Errorf("sliceInnerPlus array Length NOT match, %v != %v", len(arr1), len(arr2))
		return
	}

	for i, v1 := range arr1 {
		dstArr[i] = v1 + arr2[i]
	}

	return
}

func hammingDistance(arr1, arr2 []string) int {
	count := 0
	for i, v1 := range arr1 {
		if v1 != arr2[i] {
			count++
		}
	}

	return count
}