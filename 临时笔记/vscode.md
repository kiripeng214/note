# Launch 配置
```
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Remtoe",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "port": 18121, // 容器内部的dlv端口
            "host": "127.0.0.1", // 如果是本地Docker
            "showLog": true,
        },
        {
            "name": "Local",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "buildFlags": "-gcflags=-N",
             "dlvFlags": [
                "--check-go-version=false"
            ],
            "args":["serve","--config","/Users/zp/go/src/workspaces/job-backend/config/.application-vscode.yaml"]
        }
    ],
}
```