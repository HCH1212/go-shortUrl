// Code generated by Kitex v0.9.1. DO NOT EDIT.

package shorturlservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	example "gocode/kitex_gen/kitex/example"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"register": kitex.NewMethodInfo(
		registerHandler,
		newShortUrlServiceRegisterArgs,
		newShortUrlServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"login": kitex.NewMethodInfo(
		loginHandler,
		newShortUrlServiceLoginArgs,
		newShortUrlServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"writeShortUrl": kitex.NewMethodInfo(
		writeShortUrlHandler,
		newShortUrlServiceWriteShortUrlArgs,
		newShortUrlServiceWriteShortUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"redirect": kitex.NewMethodInfo(
		redirectHandler,
		newShortUrlServiceRedirectArgs,
		newShortUrlServiceRedirectResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"deleteShortUrl": kitex.NewMethodInfo(
		deleteShortUrlHandler,
		newShortUrlServiceDeleteShortUrlArgs,
		newShortUrlServiceDeleteShortUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"changeShortUrl": kitex.NewMethodInfo(
		changeShortUrlHandler,
		newShortUrlServiceChangeShortUrlArgs,
		newShortUrlServiceChangeShortUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"showShortUrl": kitex.NewMethodInfo(
		showShortUrlHandler,
		newShortUrlServiceShowShortUrlArgs,
		newShortUrlServiceShowShortUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"rateShortUrl": kitex.NewMethodInfo(
		rateShortUrlHandler,
		newShortUrlServiceRateShortUrlArgs,
		newShortUrlServiceRateShortUrlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	shortUrlServiceServiceInfo                = NewServiceInfo()
	shortUrlServiceServiceInfoForClient       = NewServiceInfoForClient()
	shortUrlServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return shortUrlServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return shortUrlServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return shortUrlServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ShortUrlService"
	handlerType := (*example.ShortUrlService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "example",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceRegisterArgs)
	realResult := result.(*example.ShortUrlServiceRegisterResult)
	success, err := handler.(example.ShortUrlService).Register(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortUrlServiceRegisterArgs() interface{} {
	return example.NewShortUrlServiceRegisterArgs()
}

func newShortUrlServiceRegisterResult() interface{} {
	return example.NewShortUrlServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceLoginArgs)
	realResult := result.(*example.ShortUrlServiceLoginResult)
	success, err := handler.(example.ShortUrlService).Login(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortUrlServiceLoginArgs() interface{} {
	return example.NewShortUrlServiceLoginArgs()
}

func newShortUrlServiceLoginResult() interface{} {
	return example.NewShortUrlServiceLoginResult()
}

func writeShortUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceWriteShortUrlArgs)
	realResult := result.(*example.ShortUrlServiceWriteShortUrlResult)
	success, err := handler.(example.ShortUrlService).WriteShortUrl(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortUrlServiceWriteShortUrlArgs() interface{} {
	return example.NewShortUrlServiceWriteShortUrlArgs()
}

func newShortUrlServiceWriteShortUrlResult() interface{} {
	return example.NewShortUrlServiceWriteShortUrlResult()
}

func redirectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceRedirectArgs)

	err := handler.(example.ShortUrlService).Redirect(ctx, realArg.Request)
	if err != nil {
		return err
	}

	return nil
}
func newShortUrlServiceRedirectArgs() interface{} {
	return example.NewShortUrlServiceRedirectArgs()
}

func newShortUrlServiceRedirectResult() interface{} {
	return example.NewShortUrlServiceRedirectResult()
}

func deleteShortUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceDeleteShortUrlArgs)

	err := handler.(example.ShortUrlService).DeleteShortUrl(ctx, realArg.Request)
	if err != nil {
		return err
	}

	return nil
}
func newShortUrlServiceDeleteShortUrlArgs() interface{} {
	return example.NewShortUrlServiceDeleteShortUrlArgs()
}

func newShortUrlServiceDeleteShortUrlResult() interface{} {
	return example.NewShortUrlServiceDeleteShortUrlResult()
}

func changeShortUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceChangeShortUrlArgs)

	err := handler.(example.ShortUrlService).ChangeShortUrl(ctx, realArg.Request)
	if err != nil {
		return err
	}

	return nil
}
func newShortUrlServiceChangeShortUrlArgs() interface{} {
	return example.NewShortUrlServiceChangeShortUrlArgs()
}

func newShortUrlServiceChangeShortUrlResult() interface{} {
	return example.NewShortUrlServiceChangeShortUrlResult()
}

func showShortUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ShortUrlServiceShowShortUrlArgs)
	realResult := result.(*example.ShortUrlServiceShowShortUrlResult)
	success, err := handler.(example.ShortUrlService).ShowShortUrl(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortUrlServiceShowShortUrlArgs() interface{} {
	return example.NewShortUrlServiceShowShortUrlArgs()
}

func newShortUrlServiceShowShortUrlResult() interface{} {
	return example.NewShortUrlServiceShowShortUrlResult()
}

func rateShortUrlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	_ = arg.(*example.ShortUrlServiceRateShortUrlArgs)
	realResult := result.(*example.ShortUrlServiceRateShortUrlResult)
	success, err := handler.(example.ShortUrlService).RateShortUrl(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortUrlServiceRateShortUrlArgs() interface{} {
	return example.NewShortUrlServiceRateShortUrlArgs()
}

func newShortUrlServiceRateShortUrlResult() interface{} {
	return example.NewShortUrlServiceRateShortUrlResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, request *example.RegisterRequest) (r *example.RegisterResponse, err error) {
	var _args example.ShortUrlServiceRegisterArgs
	_args.Request = request
	var _result example.ShortUrlServiceRegisterResult
	if err = p.c.Call(ctx, "register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, request *example.LoginRequest) (r *example.LoginResponse, err error) {
	var _args example.ShortUrlServiceLoginArgs
	_args.Request = request
	var _result example.ShortUrlServiceLoginResult
	if err = p.c.Call(ctx, "login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) WriteShortUrl(ctx context.Context, request *example.ShortUrlRequest) (r *example.ShortUrlResponse, err error) {
	var _args example.ShortUrlServiceWriteShortUrlArgs
	_args.Request = request
	var _result example.ShortUrlServiceWriteShortUrlResult
	if err = p.c.Call(ctx, "writeShortUrl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Redirect(ctx context.Context, request *example.RedirectRequest) (err error) {
	var _args example.ShortUrlServiceRedirectArgs
	_args.Request = request
	var _result example.ShortUrlServiceRedirectResult
	if err = p.c.Call(ctx, "redirect", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) DeleteShortUrl(ctx context.Context, request *example.DeleteShortUrlRequest) (err error) {
	var _args example.ShortUrlServiceDeleteShortUrlArgs
	_args.Request = request
	var _result example.ShortUrlServiceDeleteShortUrlResult
	if err = p.c.Call(ctx, "deleteShortUrl", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) ChangeShortUrl(ctx context.Context, request *example.ChangeShortUrlRequest) (err error) {
	var _args example.ShortUrlServiceChangeShortUrlArgs
	_args.Request = request
	var _result example.ShortUrlServiceChangeShortUrlResult
	if err = p.c.Call(ctx, "changeShortUrl", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) ShowShortUrl(ctx context.Context, request *example.ShowShortUrlRequest) (r *example.ShowShortUrlResponse, err error) {
	var _args example.ShortUrlServiceShowShortUrlArgs
	_args.Request = request
	var _result example.ShortUrlServiceShowShortUrlResult
	if err = p.c.Call(ctx, "showShortUrl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RateShortUrl(ctx context.Context) (r *example.RateShortUrlResponse, err error) {
	var _args example.ShortUrlServiceRateShortUrlArgs
	var _result example.ShortUrlServiceRateShortUrlResult
	if err = p.c.Call(ctx, "rateShortUrl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}