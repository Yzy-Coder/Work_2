package main

import "fmt"

type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}

//发布视频函数，是作者的方法，入参为作者和视频名字,返回一个视频结构体
func (Author Author) PostVideo (VideoName string) Video {

	NewVideo := Video{Author,VideoName,"这是一个测试视频结构体",10,0,0,0,0}

	return NewVideo
}

type Video struct {

	Author						//作者信息
	Name 			string 		//视频名字
	Introduce 		string 		//视频介绍
	Length			int 		//视频长度
	ThumbsUpNum		int 		//点赞数
	CoinNum			int 		//投币数
	CollectionsNum 	int			//收藏数
	TCCMixNum		int 		//一键三连次数

}

// 视频结构体的方法——点赞,视频点赞数+1
func (Video *Video) ThumbsUp	()  {

	Video.ThumbsUpNum++
	fmt.Println("点赞成功")
}

// 视频结构体的方法——投币,视频投币数+1
func (Video *Video) InsertCoin ()  {

	Video.CoinNum++
	fmt.Println("投币成功")
}

// 视频结构体的方法—-收藏,视频收藏数+1
func (Video *Video) Collect ()  {

	Video.CollectionsNum++
	fmt.Println("收藏成功")
}

// // 视频结构体的方法——一键三连,视频点赞、投币、收藏数都+1，同时一键三连数+1
func (Video *Video) TCCMix ()  {

	Video.CollectionsNum++
	Video.ThumbsUpNum++
	Video.CoinNum++
	Video.TCCMixNum++
	fmt.Println("一键三连")
}



func main() {

	Yzy := Author{"Yzy",true,":D","咸鱼一条",8} //初始化作者结构体
	TestVideo := Yzy.PostVideo("刺客信条混剪")	//测试发布视频方法

	fmt.Printf("%T\n",TestVideo)	//确认TestVideo的类型

	TestVideo.ThumbsUp()	//	测试点赞方法
	TestVideo.Collect()		// 	测试收藏方法
	TestVideo.InsertCoin()	//	测试投币方法
	TestVideo.TCCMix()		//	测试一键三连方法

	fmt.Println("TestVideo作者详情为", TestVideo.Author)
	fmt.Println("TestVideo的点赞数为",TestVideo.ThumbsUpNum,"投币数为",TestVideo.CoinNum,"收藏数为",TestVideo.CollectionsNum,"一键三连数为",TestVideo.TCCMixNum)
}