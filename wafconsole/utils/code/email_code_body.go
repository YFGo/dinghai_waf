package code

import (
	"html/template"
	"strings"
	"time"
)

func GenerateEmailBody(code string) (string, error) {
	tmpl, err := template.New("email").Parse(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        /* 整体样式 */
        body {
            font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
            margin: 0;
            padding: 20px;
            background-color: #f8f9fa;
        }

        /* 容器 */
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: white;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            padding: 24px;
        }

        /* 标题 */
        .header {
            text-align: center;
            margin-bottom: 24px;
        }

        .header h1 {
            color: #005FCC; /* 品牌主色，可替换 */
            margin: 0;
            font-size: 24px;
        }

        /* 内容区域 */
        .content {
            margin-bottom: 24px;
            line-height: 1.6;
            color: #333;
        }

        /* 验证码框 */
        .code-box {
            background-color: #f8f9fa;
            border: 1px solid #e9ecef;
            border-radius: 4px;
            padding: 16px;
            text-align: center;
            margin: 20px 0;
        }

        .code-value {
            font-size: 20px;
            font-weight: 500;
            color: #005FCC;
            letter-spacing: 4px;
        }

        /* 页脚 */
        .footer {
            border-top: 1px solid #e9ecef;
            padding-top: 16px;
            text-align: center;
            color: #6c757d;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>定海</h1>
            <p>安全验证邮件</p>
        </div>

        <div class="content">
            <p>尊敬的用户：</p>
            <p>您正在进行注册操作，本次验证码为：</p>
            <div class="code-box">
                <span class="code-value">{{.Code}}</span>
            </div>
            <p>该验证码有效期为 <strong>5分钟</strong>，请及时完成验证。</p>
            <p>如非本人操作，请忽略本邮件。</p>
        </div>

        <div class="footer">
            <p>定海项目团队</p>
            <p>© {{.Year}} 定海项目 保留所有权利</p>
        </div>
    </div>
</body>
</html>`)
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	data := struct {
		Code string
		Year int
	}{
		Code: code,
		Year: time.Now().Year(),
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
