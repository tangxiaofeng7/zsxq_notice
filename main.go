package main

import (
	"fmt"
	"strconv"
	"time"
	"zsxq_notice/model"

	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Knowledge struct {
	Id          int
	Message     string
	Name        string
	Create_time string
}

func main() {
	//企业微信机器人Key
	key := fmt.Sprintf("%s", g.Cfg().Get("wechat.key"))
	robot := model.Robot{Key: key}

	//知识星球账号Token
	zsxq_access_token := fmt.Sprintf("%s", g.Cfg().Get("zsxq_access_token.token"))

	//计划任务
	_, err := gcron.AddSingleton("* * * * * *", func() {
		g.Log().Print("doing")
		arr := g.Cfg().Get("zsxq_group").([]interface{})
		for e := 0; e < len(arr); e++ {
			for _, b := range arr[e].(map[string]interface{}) {
				Cyber(b.(string), zsxq_access_token, robot)
			}
		}

		time_value, _ := strconv.Atoi(g.Cfg().Get("time.value").(string))
		time.Sleep(time.Duration(time_value) * time.Second)

	})

	if err != nil {
		panic(err)
	}
	select {}
}

func Cyber(group_id string, zsxq_access_token string, robot model.Robot) {
	c := g.Client()
	c.SetHeaderRaw(`
	Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"
	X-Version: 2.22.0
	X-Signature: 0b717e06f750fc27b447d83392fb3ca67db1a18a
	Sec-Ch-Ua-Mobile: ?0
	User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53
	Accept: application/json, text/plain, */*
	X-Timestamp: 1653445775
	X-Request-Id: c08fab25b-fb9a-d0de-fbae-d62185bfbb4
	Sec-Ch-Ua-Platform: "macOS"
	Origin: https://wx.zsxq.com
	Sec-Fetch-Site: same-site
	Sec-Fetch-Mode: cors
	Sec-Fetch-Dest: empty
	Referer: https://wx.zsxq.com/
	Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6
`)
	c.SetCookie("abtest_env", "product")
	c.SetCookie("zsxq_access_token", zsxq_access_token)

	if r, e := c.Get(fmt.Sprintf("https://api.zsxq.com/v2/groups/%s/topics?scope=all&count=5", group_id)); e != nil {
		panic(e)
	} else {
		defer r.Close()
		jsonObject, _ := gjson.DecodeToJson(r.ReadAllString())
		array := jsonObject.Get("resp_data.topics").Array()

		for i := 0; i < len(array); i++ {
			Item := array[i].(map[string]interface{})
			create_time := Item["create_time"]
			time1 := fmt.Sprintf("%s", create_time)
			fmt.Println("第", i+1, "篇\n发布时间", strings.Replace(time1[:16], "T", " ", -1))

			a := Item["talk"]
			b, _ := gjson.DecodeToJson(a)
			//星球文章内容
			// fmt.Println("内容", b.Get("text").String())

			//判断数据库中是否存在
			abc := b.Get("text").String()
			rs := []rune(abc)
			//作者
			author := b.Get("owner.name").String()
			fmt.Println("作者", author)

			list, err := g.DB().GetAll("select * from knowledge where message=?", string(rs[:10]))
			if err != nil {
				g.Log().Header(false).Fatal(err)
			}
			if list == nil {
				fmt.Println("list is nil")
				r, err := g.DB().Insert("knowledge", gdb.Map{
					"message":     string(rs[:10]),
					"name":        author,
					"create_time": strings.Replace(time1[:16], "T", " ", -1),
				})
				if err != nil {
					g.Log().Header(false).Fatal(err)
				} else {
					fmt.Println("插入成功", r)
				}
				//星球文章附件
				array_file := b.Get("files").Array()
				for j := 0; j < len(array_file); j++ {
					Item_file := array_file[j].(map[string]interface{})
					fmt.Println("附件", Item_file["name"].(string))
					//企业微信人通知
					s := fmt.Sprintf(`
				# <font color="info">知识星球发现新的内容</font>
				> <font color="info">内容</font>: %s
				> <font color="info">作者</font>: %s
				> <font color="info">附件</font>: %s
				> <font color="info">发布时间</font>: %s
	`, b.Get("text").String(), author, Item_file["name"].(string), strings.Replace(time1[:16], "T", " ", -1))
					res, _ := robot.SendMarkdown(s)
					if res.ErrorCode != 0 {
						fmt.Println(res.ErrorMessage)
					}
				}
			} else {
				fmt.Println("数据库中已存在")
			}
		}
	}
}
