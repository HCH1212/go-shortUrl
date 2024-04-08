// 短链实现
package longtoshort

import (
	"fmt"
	"strings"
)

func spilt(longUrl string) []string {
	return strings.Split(longUrl, "") //将长链接按字符分割成字符串数组
}

// id自增长算法实现短链
func autoShortUrl(id int, longUrl string) string {
	return GetString62(Encode62(id, longUrl), longUrl) //根据id和长链接生成短链接
}

func Encode62(id int, longUrl string) []int {
	tempE := []int{} //存储编码结果的临时数组

	for id > 0 {
		tempE = append(tempE, id%62) //取余数作为编码的一部分
		id /= 62                     //除以62，继续编码
	}
	return tempE //返回编码结果
}

func GetString62(indexA []int, longUrl string) string {
	res := "" //存储最终的短链接

	for _, val := range indexA {
		res += spilt(longUrl)[val] //根据编码结果从长链接中取出对应的字符，拼接成短链接
	}
	return reverseString(res) //反转短链接并返回
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s) //将字符串转换为rune数组
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from] //交换前后字符
	}
	return string(runes) //将rune数组转换回字符串并返回
}

func To(id int, longUrl string) string {
	fmt.Println(autoShortUrl(id, longUrl))
	return autoShortUrl(id, longUrl)
}
