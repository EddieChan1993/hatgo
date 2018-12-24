package util
import (
	"net/url"
	"net/http"
	"strings"
	"log"
	"encoding/base64"
	"io/ioutil"
	"sync"
)

type request struct {
	url    string //url地址
	req    *http.Request //请求实例
	cli    *url.Values
	header map[string]string //请求头
	param  map[string]string //post提交参数
	sync.RWMutex
}

//构造request实例对象
func HttpCurl(url string)*request {
	if url=="" {
		log.Fatalln("Lack of request url")
	}
	return &request{
		url:url,
	}
}
//传入header
func (this *request) SetHeader(headers map[string]string)*request {
	this.header =headers
	return this
}
//传入请求参数，POST/GET
func (this *request) SetParms(postData map[string]string)*request {
	this.param =postData
	return this
}

//将参数加入请求中
func (this *request) initParams() *strings.Reader {
	for k,v:=range this.param {
		this.cli.Add(k,v)
	}
	return strings.NewReader(this.cli.Encode())

}

//post请求
func (this *request)Post()([]byte,error)  {
	return this.send(http.MethodPost)
}

//get请求
func (this *request)Get()([]byte,error)  {
	return this.send(http.MethodGet)
}

//将用户自定义请求头添加到http.Request实例
func (this *request) initHeaders(){
	for k, v := range this.header {
		this.req.Header.Set(k,v)
	}
}

//发送请求
func (this *request)send(method string) ([]byte,error){
	this.Lock()
	defer this.Unlock()

	this.cli=&url.Values{}
	req,err:=http.NewRequest(method,this.url,this.initParams())
	if err !=nil{
		return nil,err
	}

	this.req=req
	this.initHeaders()

	/**************处理响应数据***************/
	resp,err := http.DefaultClient.Do(req)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}


func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic "+base64.StdEncoding.EncodeToString([]byte(auth))
}