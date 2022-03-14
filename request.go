package ghttp

import (
	"errors"
	"net/http"
)

import (
	"github.com/mozillazg/request"
)

/* ================================================================================
 * Http网络请求
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	 * Http请求接口
	 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
	IHttpRequest interface {
		Get(url string) (IHttpResponse, error)
		Post(url string) (IHttpResponse, error)
		Put(url string) (IHttpResponse, error)
		Delete(url string) (IHttpResponse, error)

		SetUserAgent(userAgent string)
		SetHeaders(headers map[string]string)
		SetParams(params map[string]string)
		SetCookies(cookies map[string]string)
		SetJson(json interface{})
		SetData(data map[string]string)
		SetFiles(files FormFiles)
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Http请求数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type HttpRequest struct {
	userAgent  string
	headers    map[string]string
	params     map[string]string
	cookies    map[string]string
	json       interface{}
	data       map[string]string
	files      FormFiles
	request    *request.Request
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化Http请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewHttpRequest() IHttpRequest {
	httpRequest := &HttpRequest{}
	httpRequest.request = request.NewRequest(new(http.Client))

	return httpRequest
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Get请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) Get(url string) (IHttpResponse, error) {
	if len(url) == 0 {
		return nil, errors.New("argument url error")
	}

	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if len(s.params) > 0 {
		s.request.Params = s.params
	}

	resp, err := s.request.Get(url)
	defer resp.Body.Close()

	httpResponse := new(HttpResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Post请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) Post(url string) (IHttpResponse, error) {
	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if s.json != nil {
		s.request.Json = s.json
	}

	if s.data != nil {
		s.request.Data = s.data
	}

	if len(s.files) > 0 {
		fileFields := make([]request.FileField, 0)
		for _, file := range s.files {
			fileField := request.FileField{
				FieldName: file.FieldName, FileName: file.FileName, File: file.Datas,
			}
			fileFields = append(fileFields, fileField)
		}
		s.request.Files = fileFields
	}

	resp, err := s.request.Post(url)
	defer resp.Body.Close()

	httpResponse := new(HttpResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Put请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) Put(url string) (IHttpResponse, error) {
	if len(url) == 0 {
		return nil, errors.New("argument url error")
	}

	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if len(s.params) > 0 {
		s.request.Params = s.params
	}

	resp, err := s.request.Put(url)
	defer resp.Body.Close()

	httpResponse := new(HttpResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Delete请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) Delete(url string) (IHttpResponse, error) {
	if len(url) == 0 {
		return nil, errors.New("argument url error")
	}

	//http头
	if len(s.headers) == 0 {
		s.headers = make(map[string]string, 0)
	}

	//用户代理
	s.headers["User-Agent"] = s.getUserAgent()
	s.request.Headers = s.headers

	if len(s.cookies) > 0 {
		s.request.Cookies = s.cookies
	}

	if len(s.params) > 0 {
		s.request.Params = s.params
	}

	resp, err := s.request.Delete(url)
	defer resp.Body.Close()

	httpResponse := new(HttpResponse)
	httpResponse.Header = resp.Header

	if data, err := resp.Content(); err == nil {
		httpResponse.Data = data
	}

	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Status = resp.Status

	return httpResponse, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置用户代理http头
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetUserAgent(userAgent string) {
	s.userAgent = userAgent
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置头
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetHeaders(headers map[string]string) {
	s.headers = headers
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置参数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetParams(params map[string]string) {
	s.params = params
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Cookie
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetCookies(cookies map[string]string) {
	s.cookies = cookies
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置Json数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetJson(json interface{}) {
	s.json = json
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置字典数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetData(data map[string]string) {
	s.data = data
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置文件数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *HttpRequest) SetFiles(files FormFiles) {
	s.files = files
}

func (s *HttpRequest) getUserAgent() string {
	if len(s.userAgent) == 0 {
		GetUserAgent()
	}

	return s.userAgent 
}
