package simhash

import (
	"github.com/yanyiwu/gojieba"
	"sync"
)

type GoJieba struct {
	C *gojieba.Jieba
}

var GJB *GoJieba
var one sync.Once

func NewGoJieba() (*GoJieba) {
	one.Do(func() {
		GJB = &GoJieba{
			C: gojieba.NewJieba(),
			//equals with x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
		}
	})
	return GJB
}

func (this *GoJieba) Close() {
	this.C.Free()
}

func (this *GoJieba) AddWords(words []string) {
	for _, word := range words {
		this.C.AddWord(word)
	}
}

func (this *GoJieba) JiebaCut(rawStr string, useHmm bool, cutAll bool) (words []string) {
	if cutAll {
		words = jiebaCutAll(this.C, &rawStr)
	} else {
		words = jiebaCut(this.C, &rawStr, useHmm)
	}

	return
}

func (this *GoJieba) JiebaCutWithFrequency(rawStr string, useHmm bool, cutAll bool) (wordsFreqs map[string]int) {
	wordsFreqs = make(map[string]int)
	if cutAll {
		words := jiebaCutAll(this.C, &rawStr)
		for _, word := range words {
			freq := wordsFreqs[word]
			wordsFreqs[word] = freq + 1
		}
	} else {
		words := jiebaCut(this.C, &rawStr, useHmm)
		for _, word := range words {
			freq := wordsFreqs[word]
			wordsFreqs[word] = freq + 1
		}
	}

	return
}

func (this *GoJieba) JiebaCutForSearch(rawStr string, useHmm bool) {
	jiebaCut4Search(this.C, &rawStr, useHmm)

}

func jiebaCutAll(x *gojieba.Jieba, rawStr *string) (words []string) {
	words = x.CutAll(*rawStr)
	return
}

func jiebaCut(x *gojieba.Jieba, rawStr *string, useHmm bool) (words []string) {
	words = x.Cut(*rawStr, useHmm)
	return
}

func jiebaCut4Search(x *gojieba.Jieba, rawStr *string, useHmm bool) (words []string) {
	words = x.CutForSearch(*rawStr, useHmm)
	return
}