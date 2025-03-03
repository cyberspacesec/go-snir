package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/internal/islazy"
	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
)

// 报告模板
const htmlReportTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>网页截图扫描报告</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
            color: #333;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background-color: #fff;
            padding: 20px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
            border-radius: 5px;
        }
        h1 {
            color: #2c3e50;
            border-bottom: 2px solid #3498db;
            padding-bottom: 10px;
        }
        .summary {
            background-color: #f8f9fa;
            padding: 15px;
            border-radius: 5px;
            margin-bottom: 20px;
        }
        .screenshot-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
            gap: 20px;
        }
        .screenshot-item {
            border: 1px solid #ddd;
            border-radius: 5px;
            overflow: hidden;
            transition: transform 0.3s ease;
        }
        .screenshot-item:hover {
            transform: translateY(-5px);
            box-shadow: 0 5px 15px rgba(0,0,0,0.1);
        }
        .screenshot-img {
            width: 100%;
            height: 200px;
            object-fit: cover;
            border-bottom: 1px solid #eee;
        }
        .screenshot-info {
            padding: 15px;
        }
        .screenshot-title {
            font-weight: bold;
            margin-bottom: 10px;
            color: #2c3e50;
        }
        .screenshot-url {
            font-size: 0.9em;
            color: #3498db;
            word-break: break-all;
            margin-bottom: 10px;
        }
        .screenshot-meta {
            font-size: 0.8em;
            color: #7f8c8d;
        }
        .status-code {
            display: inline-block;
            padding: 3px 6px;
            border-radius: 3px;
            font-size: 0.8em;
            font-weight: bold;
        }
        .status-2xx {
            background-color: #2ecc71;
            color: white;
        }
        .status-3xx {
            background-color: #3498db;
            color: white;
        }
        .status-4xx {
            background-color: #f39c12;
            color: white;
        }
        .status-5xx {
            background-color: #e74c3c;
            color: white;
        }
        .status-0 {
            background-color: #95a5a6;
            color: white;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>网页截图扫描报告</h1>
        <div class="summary">
            <p><strong>生成时间:</strong> {{.GeneratedAt}}</p>
            <p><strong>总计截图:</strong> {{len .Results}}</p>
        </div>
        
        <div class="screenshot-grid">
            {{range .Results}}
            <div class="screenshot-item">
                {{if .Screenshot}}
                <img src="{{.Screenshot}}" alt="{{.Title}}" class="screenshot-img">
                {{else}}
                <div class="screenshot-img" style="background-color: #eee; display: flex; align-items: center; justify-content: center;">
                    <span>无截图</span>
                </div>
                {{end}}
                <div class="screenshot-info">
                    <div class="screenshot-title">{{if .Title}}{{.Title}}{{else}}无标题{{end}}</div>
                    <div class="screenshot-url">{{.URL}}</div>
                    <div class="screenshot-meta">
                        <span class="status-code status-{{.StatusCodeClass}}">{{.ResponseCode}}</span>
                        <span>{{.ProbedAt.Format "2006-01-02 15:04:05"}}</span>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
    </div>
</body>
</html>`

// 报告数据结构
type ReportData struct {
	GeneratedAt string
	Results     []ReportResult
}

// 报告结果项
type ReportResult struct {
	URL            string
	Title          string
	Screenshot     string
	ResponseCode   int
	StatusCodeClass string
	ProbedAt       time.Time
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "生成HTML报告",
	Long:  "根据扫描结果生成HTML格式的报告",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 检查输入文件
		if opts.Report.InputFile == "" {
			return fmt.Errorf("请使用 --input 参数指定JSONL结果文件")
		}

		// 检查输入文件是否存在
		if !islazy.FileExists(opts.Report.InputFile) {
			return fmt.Errorf("输入文件不存在: %s", opts.Report.InputFile)
		}

		// 读取JSONL文件
		log.Info("读取结果文件", "file", opts.Report.InputFile)
		results, err := readJSONLResults(opts.Report.InputFile)
		if err != nil {
			return fmt.Errorf("读取结果文件失败: %v", err)
		}

		if len(results) == 0 {
			return fmt.Errorf("结果文件中没有有效的记录")
		}

		log.Info("读取到结果记录", "count", len(results))

		// 准备报告数据
		reportData := ReportData{
			GeneratedAt: time.Now().Format("2006-01-02 15:04:05"),
			Results:     make([]ReportResult, 0, len(results)),
		}

		// 处理每个结果
		for _, result := range results {
			// 获取状态码类别
			statusClass := "0"
			if result.ResponseCode >= 200 && result.ResponseCode < 300 {
				statusClass = "2xx"
			} else if result.ResponseCode >= 300 && result.ResponseCode < 400 {
				statusClass = "3xx"
			} else if result.ResponseCode >= 400 && result.ResponseCode < 500 {
				statusClass = "4xx"
			} else if result.ResponseCode >= 500 && result.ResponseCode < 600 {
				statusClass = "5xx"
			}

			// 处理截图路径，使其相对于报告文件
			screenshotPath := result.Screenshot
			if screenshotPath != "" {
				// 如果是绝对路径，转换为相对路径
				if filepath.IsAbs(screenshotPath) {
					relPath, err := filepath.Rel(filepath.Dir(opts.Report.OutputPath), screenshotPath)
					if err == nil {
						screenshotPath = relPath
					}
				}
			}

			reportData.Results = append(reportData.Results, ReportResult{
				URL:            result.URL,
				Title:          result.Title,
				Screenshot:     screenshotPath,
				ResponseCode:   result.ResponseCode,
				StatusCodeClass: statusClass,
				ProbedAt:       result.ProbedAt,
			})
		}

		// 确保输出目录存在
		outputDir := filepath.Dir(opts.Report.OutputPath)
		if _, err := islazy.CreateDir(outputDir); err != nil {
			return fmt.Errorf("创建输出目录失败: %v", err)
		}

		// 创建输出文件
		outputFile, err := os.Create(opts.Report.OutputPath)
		if err != nil {
			return fmt.Errorf("创建输出文件失败: %v", err)
		}
		defer outputFile.Close()

		// 解析模板
		tmpl, err := template.New("report").Parse(htmlReportTemplate)
		if err != nil {
			return fmt.Errorf("解析报告模板失败: %v", err)
		}

		// 执行模板
		if err := tmpl.Execute(outputFile, reportData); err != nil {
			return fmt.Errorf("生成报告失败: %v", err)
		}

		log.Info("HTML报告生成成功", "path", opts.Report.OutputPath)
		return nil
	},
}

// 从JSONL文件读取结果
func readJSONLResults(filePath string) ([]*models.Result, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	var results []*models.Result
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// 解析JSON
		var result models.Result
		if err := json.Unmarshal([]byte(line), &result); err != nil {
			log.Error("解析JSON行失败", "error", err, "line", line)
			continue
		}

		results = append(results, &result)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func init() {
	reportCmd.AddCommand(htmlCmd)

	// 添加HTML报告相关选项
	htmlCmd.Flags().StringVar(&opts.Report.InputFile, "input", "", "JSONL格式的结果文件路径")
	htmlCmd.Flags().StringVar(&opts.Report.OutputPath, "output", "report.html", "HTML报告输出路径")
	htmlCmd.MarkFlagRequired("input")

	log.Debug("已注册html报告命令")
}