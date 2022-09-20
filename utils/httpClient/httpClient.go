package httpClient

import (
	"context"
	"github.com/imroc/req/v3"
	"github.com/lizhanfei/easygolib/log"
	"time"
)

func NewHttpClient(conf *Client, isLog bool, logger log.Zlog) *HttpClient {
	if conf.Timeout <= 0 {
		conf.Timeout = time.Second * 2
	}
	if conf.Retry <= 0 {
		conf.Retry = 2
	}
	reqClient := req.C().SetBaseURL(conf.Domain).
		SetCommonRetryCount(conf.Retry).
		//SetTimeout(conf.Timeout).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if isLog {
				if resp.IsSuccess() {
					logger.Infof(resp.Request.Context(), "req success, cost:%d, respBody:%s", resp.TotalTime(), string(resp.Bytes()))
				} else {
					logger.Errorf(resp.Request.Context(), "req fail, cost:%d, httpCode:%d, respBody:%s", resp.TotalTime(), resp.StatusCode, string(resp.Bytes()))
				}
			}
			return nil
		})

	return &HttpClient{
		Client:  reqClient,
		Name:    conf.Service,
		timeOut: conf.Timeout,
	}
}

type HttpClient struct {
	*req.Client
	Name     string
	isLogged bool
	timeOut  time.Duration
	ctx      context.Context
	logger   log.Zlog
}

func (this *HttpClient) Get(ctx context.Context, path string, header map[string]string) (*req.Response, error) {
	req := this.R().SetContext(ctx)
	if nil != header {
		for k, v := range header {
			req.SetHeader(k, v)
		}
	}
	return req.Get(path)
}

func (this *HttpClient) Post(ctx context.Context, path string, header map[string]string, data interface{}) (*req.Response, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(this.timeOut))
	defer cancel()
	req := this.R().SetContext(ctx)
	for k, v := range header {
		req.SetHeader(k, v)
	}
	req.SetBody(data)
	return req.Post(path)
}
