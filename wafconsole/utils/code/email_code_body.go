package code

func GenerateEmailBody(code string) string {
	return `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>验证码邮件</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
       .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
       .header {
            text-align: center;
            color: #333;
        }
       .content {
            margin-top: 20px;
        }
       .code {
            font-size: 24px;
            font-weight: bold;
            color: #007BFF;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1 class="header">验证码邮件</h1>
        <div class="content">
            <p>尊敬的用户，您好！</p>
            <p>您的验证码是：</p>
            <p class="code">` + code + `</p>
            <p>该验证码 5 分钟内有效，请尽快使用。</p>
        </div>
    </div>
</body>
</html>
`
}
