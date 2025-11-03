// 这是一个示例配置文件，展示如何配置OSS
// 实际使用时，请通过环境变量设置这些值，不要直接在config.go中硬编码

package config

// 示例：如果使用环境变量，config.go应该是这样：
var ExampleConfig = Config{
	OSS: OSSConfig{
		Endpoint:        "oss-cn-beijing.aliyuncs.com", // 从环境变量OSS_ENDPOINT读取
		AccessKeyID:     "",                            // 从环境变量OSS_ACCESS_KEY_ID读取
		AccessKeySecret: "",                            // 从环境变量OSS_ACCESS_KEY_SECRET读取
		BucketName:      "zcc-mjh",                     // 从环境变量OSS_BUCKET_NAME读取
		Domain:          "",                            // 从环境变量OSS_DOMAIN读取
	},
}

// 设置环境变量的方法：
// Windows PowerShell:
//   $env:OSS_ENDPOINT="oss-cn-beijing.aliyuncs.com"
//   $env:OSS_ACCESS_KEY_ID="your-access-key-id"
//   $env:OSS_ACCESS_KEY_SECRET="your-access-key-secret"
//   $env:OSS_BUCKET_NAME="zcc-mjh"
//   $env:OSS_DOMAIN=""
//
// Windows CMD:
//   set OSS_ENDPOINT=oss-cn-beijing.aliyuncs.com
//   set OSS_ACCESS_KEY_ID=your-access-key-id
//   set OSS_ACCESS_KEY_SECRET=your-access-key-secret
//   set OSS_BUCKET_NAME=zcc-mjh
//   set OSS_DOMAIN=
//
// Linux/Mac:
//   export OSS_ENDPOINT="oss-cn-beijing.aliyuncs.com"
//   export OSS_ACCESS_KEY_ID="your-access-key-id"
//   export OSS_ACCESS_KEY_SECRET="your-access-key-secret"
//   export OSS_BUCKET_NAME="zcc-mjh"
//   export OSS_DOMAIN=""
