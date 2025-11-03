package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lccxxo/blog/config"
)

// InitOSS 初始化OSS（验证配置）
func InitOSS() error {
	cfg := config.AppConfig.OSS

	// 检查配置是否为空
	if cfg.Endpoint == "" || cfg.AccessKeyID == "" || cfg.AccessKeySecret == "" || cfg.BucketName == "" {
		return fmt.Errorf("OSS配置不完整，请检查配置")
	}

	return nil
}

// UploadToOSS 上传文件到OSS
// file: 文件内容（io.Reader）
// filename: 原始文件名
// fileType: 文件类型，如 "image", "video"
// 返回：文件URL和错误
func UploadToOSS(file io.Reader, filename, fileType string) (string, error) {
	cfg := config.AppConfig.OSS

	// 读取文件内容到字节数组
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return "", fmt.Errorf("读取文件内容失败: %v", err)
	}
	fileContent := buf.Bytes()

	// 获取文件扩展名
	ext := strings.ToLower(filepath.Ext(filename))

	// 生成唯一文件名：fileType/年/月/日/UUID.扩展名
	dateDir := time.Now().Format("2006/01/02")
	uuidStr := generateUUID()
	objectKey := fmt.Sprintf("%s/%s/%s%s", fileType, dateDir, uuidStr, ext)

	// 构建OSS URL
	ossURL := fmt.Sprintf("https://%s.%s/%s", cfg.BucketName, cfg.Endpoint, objectKey)

	// 生成签名
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	contentType := getContentType(ext)

	// 构建签名字符串（OSS签名规范）
	stringToSign := fmt.Sprintf("PUT\n\n%s\n%s\n/%s/%s", contentType, date, cfg.BucketName, objectKey)

	// 计算签名
	mac := hmac.New(sha1.New, []byte(cfg.AccessKeySecret))
	mac.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authorization := fmt.Sprintf("OSS %s:%s", cfg.AccessKeyID, signature)

	// 创建HTTP请求
	req, err := http.NewRequest("PUT", ossURL, bytes.NewReader(fileContent))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Date", date)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(fileContent)))

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("上传文件到OSS失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("上传文件到OSS失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	// 生成访问URL
	var fileURL string
	if cfg.Domain != "" {
		// 如果配置了自定义域名，使用自定义域名
		fileURL = fmt.Sprintf("%s/%s", strings.TrimRight(cfg.Domain, "/"), objectKey)
	} else {
		// 否则使用OSS默认域名
		endpointHost := strings.TrimPrefix(strings.TrimPrefix(cfg.Endpoint, "https://"), "http://")
		fileURL = fmt.Sprintf("https://%s.%s/%s", cfg.BucketName, endpointHost, objectKey)
	}

	return fileURL, nil
}

// getContentType 根据文件扩展名获取Content-Type
func getContentType(ext string) string {
	contentTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".webp": "image/webp",
		".svg":  "image/svg+xml",
		".mp4":  "video/mp4",
		".avi":  "video/x-msvideo",
		".mov":  "video/quicktime",
		".wmv":  "video/x-ms-wmv",
		".flv":  "video/x-flv",
		".mkv":  "video/x-matroska",
		".webm": "video/webm",
	}

	if ct, ok := contentTypes[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}

// DeleteFromOSS 从OSS删除文件
// url: 文件的完整URL或对象键
func DeleteFromOSS(url string) error {
	cfg := config.AppConfig.OSS

	// 从URL中提取对象键
	objectKey := extractObjectKeyFromURL(url)
	if objectKey == "" {
		return fmt.Errorf("无法从URL中提取对象键")
	}

	// 构建OSS URL
	endpointHost := strings.TrimPrefix(strings.TrimPrefix(cfg.Endpoint, "https://"), "http://")
	ossURL := fmt.Sprintf("https://%s.%s/%s", cfg.BucketName, endpointHost, objectKey)

	// 生成签名
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	stringToSign := fmt.Sprintf("DELETE\n\n\n%s\n/%s/%s", date, cfg.BucketName, objectKey)

	// 计算签名
	mac := hmac.New(sha1.New, []byte(cfg.AccessKeySecret))
	mac.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authorization := fmt.Sprintf("OSS %s:%s", cfg.AccessKeyID, signature)

	// 创建HTTP请求
	req, err := http.NewRequest("DELETE", ossURL, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Date", date)
	req.Header.Set("Authorization", authorization)

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("从OSS删除文件失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态（204 No Content表示删除成功）
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("从OSS删除文件失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	return nil
}

// extractObjectKeyFromURL 从URL中提取对象键
func extractObjectKeyFromURL(url string) string {
	// 如果URL包含域名，提取路径部分
	if strings.Contains(url, "://") {
		// 移除协议和域名部分
		parts := strings.SplitN(url, "://", 2)
		if len(parts) == 2 {
			path := strings.SplitN(parts[1], "/", 2)
			if len(path) == 2 {
				return path[1]
			}
		}
	}

	// 如果已经是对象键格式（包含 image/ 或 video/）
	if strings.HasPrefix(url, "image/") || strings.HasPrefix(url, "video/") {
		return url
	}

	// 如果是旧的上传路径格式 /uploads/...，提取文件名并构造对象键
	if strings.HasPrefix(url, "/uploads/") {
		parts := strings.Split(url, "/uploads/")
		if len(parts) == 2 {
			// 假设旧格式为 /uploads/images/2025/11/03/filename.jpg
			// 转换为 image/2025/11/03/filename.jpg
			if strings.HasPrefix(parts[1], "images/") {
				return strings.Replace(parts[1], "images/", "image/", 1)
			}
			if strings.HasPrefix(parts[1], "videos/") {
				return strings.Replace(parts[1], "videos/", "video/", 1)
			}
		}
	}

	return ""
}

// generateUUID 生成UUID
func generateUUID() string {
	return uuid.New().String()
}

// UploadFile 上传本地文件到OSS（用于从本地文件系统上传）
func UploadFile(localPath, fileType string) (string, error) {
	file, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	filename := filepath.Base(localPath)
	return UploadToOSS(file, filename, fileType)
}
