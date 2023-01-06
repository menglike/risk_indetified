package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


func main(){
	g := gin.Default()
	//加载 模版目录
	g.LoadHTMLGlob("./template/*")
	//加载静态文件目录
	g.StaticFS("./static/",http.Dir("./static"))


	g.GET("/",func(c *gin.Context){

		c.HTML(200,"index.html"	,gin.H{

		})
	})

	g.POST("/post_form",func(c *gin.Context){
		ip := c.DefaultPostForm("ip","")
		email := c.DefaultPostForm("email","")
		phone := c.DefaultPostForm("phone","")
		addr := c.DefaultPostForm("addr","")
		url := ""
		if ip!= ""{
			url  =  "https://api.liangmlk.cn/risk_limit_ip.php?ip="+ip
		}
		if email!="" {
			url  =  "https://api.liangmlk.cn/risk_limit_ip.php?email="+email
		}
		if phone!="" {
			url  =  "https://api.liangmlk.cn/risk_limit_ip.php?mobile="+phone
		}
		if addr!="" {
			url  =  "https://api.liangmlk.cn/risk_limit_ip.php?address="+addr
		}
		client := &http.Client{}
		req,_ := http.NewRequest("GET",url,nil)

		resp,err := client.Do(req)
		if err!=nil{
			panic(err)
		}
		s,_ := ioutil.ReadAll(resp.Body)
		var result map[string] interface{}
		err = json.Unmarshal(s,&result)
		if err!=nil{
			panic(err)
		}

		c.JSON(200,result)
	})
	g.Run(":9999")
}
