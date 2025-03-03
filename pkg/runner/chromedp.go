package runner

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"

	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
)

// ChromeDP implements the Driver interface using chromedp
type ChromeDP struct {
	ctx    context.Context
	cancel context.CancelFunc
	opts   *Options
}

// NewChromeDP creates a new ChromeDP driver
func NewChromeDP(opts *Options) (*ChromeDP, error) {
	// 设置Chrome选项
	chromedpOpts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.DisableGPU,
	}

	// 根据配置设置无头模式
	if opts.Chrome.Headless {
		chromedpOpts = append(chromedpOpts, chromedp.Headless)
	}

	// 设置窗口大小
	chromedpOpts = append(chromedpOpts, chromedp.WindowSize(opts.Chrome.WindowX, opts.Chrome.WindowY))

	// 设置自定义User-Agent
	if opts.Chrome.UserAgent != "" {
		chromedpOpts = append(chromedpOpts, chromedp.UserAgent(opts.Chrome.UserAgent))
	}

	// 设置代理
	if opts.Chrome.Proxy != "" {
		chromedpOpts = append(chromedpOpts, chromedp.ProxyServer(opts.Chrome.Proxy))
	}

	// 设置Chrome路径
	if opts.Chrome.Path != "" {
		chromedpOpts = append(chromedpOpts, chromedp.ExecPath(opts.Chrome.Path))
	}

	// 创建Chrome上下文
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromedpOpts...)

	// 创建新的Chrome实例
	ctx, cancel = chromedp.NewContext(ctx)

	// 设置超时
	if opts.Chrome.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(opts.Chrome.Timeout)*time.Second)
	}

	return &ChromeDP{
		ctx:    ctx,
		cancel: cancel,
		opts:   opts,
	}, nil
}

// Witness implements the Driver interface
func (c *ChromeDP) Witness(target string, runner *Runner) (*models.Result, error) {
	result := &models.Result{
		URL:      target,
		ProbedAt: time.Now(),
	}

	// 创建网络事件监听器
	networkEvents := make(map[string]*models.NetworkLog)
	chromedp.ListenTarget(c.ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			networkEvents[e.RequestID.String()] = &models.NetworkLog{
				Type:   models.HTTP,
				URL:    e.Request.URL,
				Method: e.Request.Method,
			}
		case *network.EventResponseReceived:
			if nl, ok := networkEvents[e.RequestID.String()]; ok {
				nl.StatusCode = int(e.Response.Status)
				nl.ContentType = e.Response.MimeType
			}
		}
	})

	// 执行截图任务
	var buf []byte
	var htmlContent string
	var title string
	var responseCode int
	var cookies []*network.Cookie

	tasks := []chromedp.Action{
		network.Enable(),
		chromedp.Navigate(target),
	}

	// 添加延迟
	if c.opts.Chrome.Delay > 0 {
		tasks = append(tasks, chromedp.Sleep(time.Duration(c.opts.Chrome.Delay)*time.Second))
	}

	// 获取页面信息
	tasks = append(tasks, 
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取响应码
			var statusCode int
			for _, nl := range networkEvents {
				if nl.URL == target || strings.HasSuffix(target, nl.URL) {
					statusCode = nl.StatusCode
					break
				}
			}
			responseCode = statusCode
			return nil
		}),
		chromedp.Title(&title),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取HTML内容
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			htmlContent, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取Cookies
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
		chromedp.CaptureScreenshot(&buf),
	)

	// 执行任务
	err := chromedp.Run(c.ctx, tasks...)
	if err != nil {
		result.Failed = true
		result.FailedReason = err.Error()
		return result, err
	}

	// 填充结果
	result.Title = title
	result.ResponseCode = responseCode
	result.HTML = htmlContent

	// 保存截图
	if !c.opts.Scan.ScreenshotSkipSave {
		filename := fmt.Sprintf("%s_%s.%s", 
			strings.ReplaceAll(target, "/", "_"), 
			time.Now().Format("20060102150405"), 
			c.opts.Scan.ScreenshotFormat)
		filepath := filepath.Join(c.opts.Scan.ScreenshotPath, filename)
		
		err = ioutil.WriteFile(filepath, buf, 0644)
		if err != nil {
			log.Error("保存截图失败", "error", err)
		} else {
			result.Filename = filepath
			result.Screenshot = filepath
		}
	}

	// 保存Cookies
	if c.opts.Scan.SaveCookies && cookies != nil {
		for _, cookie := range cookies {
			result.Cookies = append(result.Cookies, models.Cookie{
				Name:   cookie.Name,
				Value:  cookie.Value,
				Domain: cookie.Domain,
				Path:   cookie.Path,
			})
		}
	}

	// 保存网络日志
	if c.opts.Scan.SaveNetwork {
		for _, nl := range networkEvents {
			result.Network = append(result.Network, *nl)
		}
	}

	return result, nil
}

// Close implements the Driver interface
func (c *ChromeDP) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}