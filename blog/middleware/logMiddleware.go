package middleware

import (
	"blog/tool"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	// JSON格式更适合日志分析系统
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置日志级别
	log.SetLevel(logrus.InfoLevel)
	// 同时输出到控制台和文件
	log.SetOutput(os.Stdout)
	// 如果需要同时输出到文件
	file, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(io.MultiWriter(os.Stdout, file))
	}
}

// CustomResponseWriter 自定义 ResponseWriter 用于捕获响应体:cite[1]:cite[3]:cite[7]
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer // 用于缓存响应体
}

// Write 重写 Write 方法，同时向原始 ResponseWriter 和缓冲区写入数据:cite[1]:cite[3]:cite[7]
func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMidderware() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		uId := tool.GenerateUID()
		c.Set("UID", uId)

		fmt.Println()
		// 打印请求头（根据需要可选）
		// for name, values := range c.Request.Header {
		//     for _, value := range values {
		//         fmt.Printf("%s: %s\n", name, value)
		//     }
		// }
		// 读取并打印请求体
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			var err error
			requestBodyBytes, err = io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println("读取请求体错误:", uId, err)
				c.AbortWithStatus(500)
				return
			}
			// 关键：将读取后的 body 重新写回，以便后续处理可以再次读取:cite[1]:cite[3]
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}
		var reqData string
		if len(requestBodyBytes) > 0 {
			if json.Valid(requestBodyBytes) {
				var prettyJSON bytes.Buffer
				if err := json.Indent(&prettyJSON, requestBodyBytes, "", "  "); err == nil {
					reqData = prettyJSON.String()

				} else {
					reqData = string(requestBodyBytes)
				}
			} else {
				reqData = string(requestBodyBytes)
			}
		}
		fmt.Printf("***********ReqStart************\nURL: %v Method: %v\nUID: %v\nTime: %v\nRawQuery: %v\nReqData: %s\n***********ReqEnd************\n",
			path, method, uId, tool.GetDate(), rawQuery, reqData)

		log.WithFields(logrus.Fields{
			"method": method,
			"path":   path,
			"nUID":   uId,
			"ip":     c.ClientIP(),
		}).Info(reqData)
		// 打印查询参数 (Query Parameters):cite[2]:cite[4]
		if len(c.Request.URL.Query()) > 0 {
			fmt.Println("查询参数:", uId)
			for key, values := range c.Request.URL.Query() {
				for _, value := range values {
					fmt.Printf("  %s: %s\n", key, value)
				}
			}
		}

		// 创建自定义 ResponseWriter 以捕获响应体:cite[1]:cite[3]:cite[7]
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw

		// 处理请求:cite[1]:cite[3]:cite[6]
		c.Next()
		// 计算耗时
		latency := time.Since(start)
		// 获取响应状态码
		statusCode := c.Writer.Status()
		// 获取响应体
		responseBody := crw.body.String()
		// 尝试美化打印 JSON 格式的响应体
		var rspData string
		if responseBody != "" {
			if json.Valid([]byte(responseBody)) {
				var prettyJSON bytes.Buffer
				if err := json.Indent(&prettyJSON, []byte(responseBody), "", "  "); err == nil {
					rspData = prettyJSON.String()
				} else {
					rspData = responseBody
				}
			} else {
				rspData = responseBody
			}
		}
		// 打印响应信息
		fmt.Printf("***********RspStart************\nURL :%v\nUID: %v\nRspStatus: %v\nConsumeTime: %v\nRspData: %s\n***********RspEnd************\n",
			path, uId, statusCode, latency, rspData)

		log.WithFields(logrus.Fields{
			"method":      method,
			"RspStatus":   statusCode,
			"ConsumeTime": latency,
			"UID":         uId,
		}).Info(rspData)
	}
}
