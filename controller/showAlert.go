package controller

import "net/http"

// 弹窗提示
func showAlert(w http.ResponseWriter, message string, redirectURL string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>提示</title>
	</head>
	<body>
		<script type="text/javascript">
			// 弹出提示框
			alert("` + message + `");
			// 跳转到指定页面
			window.location.href = "` + redirectURL + `";
		</script>
	</body>
	</html>
	`
	// 写入响应中
	_, _ = w.Write([]byte(html))
}
