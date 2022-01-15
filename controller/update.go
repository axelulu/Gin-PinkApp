package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"pinkacg/dao/mysql"
	"pinkacg/logic"
	"pinkacg/pkg/snowflake"
	"regexp"
	"strconv"
	"strings"
)

// UpdateHandle 获取app更新
func UpdateHandle(c *gin.Context) {
	// 逻辑处理
	version, err := logic.GetUpdate()
	if err != nil {
		zap.L().Error("logic.GetUpdate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, version)
}

// GetDouYinUrl2 ----------------------------------------------一下均为自定义脚本-----------------------------------
type GetDouYinUrl2 struct {
	Cover string `json:"cover" form:"cover" binding:"required"`
}
type GetDouYinUrl struct {
	Url string `json:"url" form:"url" binding:"required"`
}
type StuRead struct {
	Video []map[string][]map[string]string `json:"video"`
}

func GetDouYinPostUrlHandle(c *gin.Context) {
	// 6830,6880,7427,7428,7769,7770,10041,11526,11696,11887,12061,12643,14731,14732
	for i := 48191; i <= 52890; i++ {
		post, err := mysql.GetPostByCId(int64(i))
		if err != nil {
			print(err.Error())
			return
		}
		print("文章id号码：" + strconv.Itoa(i))
		stu := StuRead{}
		_ = json.Unmarshal([]byte(post.Video), &stu)
		url := (stu.Video)[0]["list"][0]["url"]
		//if !strings.HasPrefix(url, "https://www.acfun.cn/v") {
		//	print("跳转了")
		//	continue
		//}
		client := &http.Client{}
		//提交请求
		reqest, err := http.NewRequest("GET", "https://tenapi.cn/bilivideo/?url="+url, nil)

		//增加header选项
		reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
		reqest.Header.Add("Connection", "keep-alive")
		reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
		if err != nil {
			panic(err)
		}
		var getDouYinUrl GetDouYinUrl2
		//处理返回结果
		response, _ := client.Do(reqest)
		bodyByte, _ := ioutil.ReadAll(response.Body)
		body := string(bodyByte)
		json.Unmarshal([]byte(body), &getDouYinUrl)
		print(getDouYinUrl.Cover)
		defer response.Body.Close()

		mysql.UpdatePostByPostId(post.PostId, "cover", getDouYinUrl.Cover)
	}
}

func GetDouYinUrlHandle(c *gin.Context) {
	// 1. 获取请求参数及校验
	p := new(GetDouYinUrl)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	client := &http.Client{}
	//生成要访问的url
	url := p.Url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	body := string(bodyByte)
	compile := regexp.MustCompile("window.pageInfo =(.*?);")

	submatch := compile.FindAllSubmatch([]byte(body), -1)

	compiles := regexp.MustCompile("\\\\\\\"url\\\\\\\":\\\\\\\"(.*?)\\\\\\\",\\\\\\\"")

	submatchs := compiles.FindAllSubmatch(submatch[0][0], -1)
	qweee := strings.Split(string(submatchs[0][0]), "\\\"url\\\":\\\"")
	qweee2 := strings.Split(qweee[1], "\\\",\\\"\"\\\\\\\"url")
	qweee3 := strings.Split(qweee2[0], "\\\",")
	defer response.Body.Close()
	ResponseSuccess(c, qweee3[0])
}

type beautyI struct {
	Title   string
	Video   string
	Cover   string
	Star    string
	Content string
	//User_link   string
	//User_avatar string
	//User        string
	//User_time   string
}

var zzmx []beautyI

func ShellHandle(c *gin.Context) {
	//打开文件
	file, err := os.Open("acgmh.json")
	if err != nil {
		fmt.Println("打开文件失败，错误：", err)
		return
	}
	//声明文件读取在函数结束时关闭
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&zzmx)
	if err == nil {
		//fmt.Printf("json数据读取成功：%#v", zzmx)
	} else {
		fmt.Println("json数据读取失败，错误：", err)
	}
	for _, v := range zzmx {
		print("标题：" + v.Title)
		postId := snowflake.GenID()
		//c := map[string]interface{}{
		//	"video": []interface{}{
		//		map[string]interface{}{
		//			"name": v.Title,
		//			"list": []map[string]string{0: {"url": v.Video, "name": v.Title}},
		//		},
		//	},
		//}
		//videos, err := json.Marshal(c)
		if err != nil {
			return
		}
		_, err = mysql.CreatePost(postId, 2079499053372018688, "post", 1, v.Title, v.Cover, v.Content, "")
		if err != nil {
			zap.L().Error("logic.GetUpdate failed", zap.Error(err))
			return
		}
	}
	ResponseSuccess(c, "成功")
}
