package main

import (
	"fmt"
	"sort"
	"strconv"
)

type RecordInfo struct {
	M3u8Url		string
	StartTime	int
	EndTime		int
}

func main() {
	m := map[int]*RecordInfo {
		1658255400 : {M3u8Url:"http://video.bd.cn/test/d47d60f8ea7bfe7fa787a73d36.m3u8",StartTime:1658255400,EndTime:1658257200},
		1658253600 : {M3u8Url:"http://video.bd.cn/test/1afc4cbbdd6527148f831f0b8b.m3u8",StartTime:1658253600,EndTime:1658255400},
		1658251800 : {M3u8Url:"http://video.bd.cn/test/ae9928f1a092271e8032730626.m3u8",StartTime:1658251800,EndTime:1658253600},
		1658249999 : {M3u8Url:"http://video.bd.cn/test/931164b85da1112c7383f6d7c9.m3u8",StartTime:1658249999,EndTime:1658251800},
		1658248199 : {M3u8Url:"http://video.bd.cn/test/6bc373924d5da88eb8fa244ead.m3u8",StartTime:1658248199,EndTime:1658249999},
		1658246399 : {M3u8Url:"http://video.bd.cn/test/f4480ec7832f9dceb9f340cde4.m3u8",StartTime:1658246399,EndTime:1658248199},
	}
	//StartTime升序输出，首先定义一个切片用来存放value
	var time []int
	for _,item := range m{
		time = append(time,item.StartTime)
	}
	sort.Slice(time,func(i,j int)bool{
		return time[i] > time[j]
	})
	fmt.Println(time)

	//排序好了之后在根据切片中的数据进行遍历
	for _, item := range time{
		fmt.Println("排序后的结果是:" + m[item].M3u8Url + " " + strconv.Itoa(m[item].StartTime) + " " + strconv.Itoa(m[item].EndTime))
	}
}

