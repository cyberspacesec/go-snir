package runner

// Options contains all the configuration options for the runner
type Options struct {
	// API server options
	API struct {
		Host string // API服务监听地址
		Port int    // API服务监听端口
	}
	// Logging options
	Logging struct {
		Debug   bool // 是否启用调试日志
		Silence bool // 是否静默日志输出
	}

	// Chrome options
	Chrome struct {
		Path      string // Chrome可执行文件路径
		UserAgent string // 自定义User-Agent
		Proxy     string // 代理服务器地址
		Timeout   int    // 页面加载超时时间（秒）
		Delay     int    // 截图前等待时间（秒）
		WindowX   int    // 窗口宽度
		WindowY   int    // 窗口高度
		WSS       string // WebSocket服务器地址
		Headless  bool   // 是否使用无头模式
	}

	// Scan options
	Scan struct {
		Driver            string // 使用的驱动（chromedp或gorod）
		Threads           int    // 并发线程数
		ScreenshotPath    string // 截图保存路径
		ScreenshotFormat  string // 截图格式（jpeg或png）
		ScreenshotQuality int    // 截图质量（仅对JPEG有效）
		ScreenshotSkipSave bool   // 是否跳过保存截图
		SaveHTML          bool   // 是否保存HTML内容
		SaveHeaders       bool   // 是否保存HTTP头
		SaveConsole       bool   // 是否保存控制台日志
		SaveCookies       bool   // 是否保存Cookie
		SaveNetwork       bool   // 是否保存网络请求日志
		HTTP              bool   // 是否使用HTTP协议
		HTTPS             bool   // 是否使用HTTPS协议
		Ports             []int  // 扫描的端口列表
		Timeout           int    // 扫描超时时间（秒）
		MaxRetries        int    // 最大重试次数
		JavaScript        string // 要在页面上执行的JavaScript代码
		JavaScriptFile    string // 包含JavaScript代码的文件路径
		FilePath          string // URL文件路径，用于批量扫描
	}

	// Writer options
	Writer struct {
		Db       bool   // 是否写入数据库
		DbURI    string // 数据库连接URI
		DbDebug  bool   // 是否启用数据库调试
		Jsonl    bool   // 是否写入JSONL文件
		JsonlFile string // JSONL文件路径
		Csv      bool   // 是否写入CSV文件
		CsvFile  string // CSV文件路径
		Stdout   bool   // 是否输出到标准输出
	}

	// Report options
	Report struct {
		OutputPath string // 报告输出路径
		Format     string // 报告格式 (html, json, csv)
		Port       int    // Web服务器端口
		Host       string // Web服务器主机地址
		InputFile  string // 输入文件路径，用于生成报告
	}
}