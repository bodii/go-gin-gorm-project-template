# use supervisor document

## explain
use supervisor to manage the process

see: http://supervisord.org/

### supervisor主要是两个命令：
#### 1. supervisord是主进程命令，用于启动supervisor
```shell
# 启动
supervisord -c /etc/supervisord.conf
```

```shell
# 开机启动服务
vim /lib/systemd/system/supervisord.service
```
在supervisord.service中添加以下内容：
```service
# supervisord service for systemd (CentOS 7.0+)
# by ET-CS (https://github.com/ET-CS)
[Unit]
Description=Supervisor daemon
[Service]
Type=forking
ExecStart=/usr/bin/supervisord
ExecStop=/usr/bin/supervisorctl $OPTIONS shutdown
ExecReload=/usr/bin/supervisorctl $OPTIONS reload
KillMode=process
Restart=on-failure
RestartSec=42s
[Install]
WantedBy=multi-user.target
```

```shell
# 将服务脚本添加到systemctl自启动服务
systemctl enable supervisord.service
```

#### 2. supervisorctl是管理命令，用于管理supervisor中的应用程序
```shell
supervisorctl status                    # 查看所有应用程序状态
supervisorctl stop program_name         # 停止某个应用程序
supervisorctl start program_name        # 启动某个应用程序
supervisorctl restart program_name      # 重启某个应用程序
supervisorctl stop all                  # 停止所有应用程序
supervisorctl reread                    # 读取有更新（增加）的配置文件，不会启动新添加的程序
supervisorctl update                    # 更新配置文件变化并重启变化的服务
supervisorctl reload                    # 更新配置并重启supervisor
supervisorctl shutdown                  # 停止supervisor
supervisorctl tail program_name stdout  # 查看某个应用程序的输出
```


### create program config
`example_project.ini` in `/etc/supervisord.d/`
```ini
; 应用程序的路径，以main为例子
command={$project_path}/main             
; 应用程序的运行目录。如果应用程序涉及相对路径的文件，这个一定要修改好
directory={$project_path}              
; start at supervisord start (default: true)
autostart=true                
; 设置得太低可能会因为没来得及启动被判启动失败
startsecs=5                   
; max # of serial start failures when starting (default 3)
startretries=3                
; when to restart if exited after running (def: unexpected)
autorestart=true        
; send stop signal to the UNIX process group (default false)
stopasgroup=true             
; SIGKILL the UNIX process group (def false)
killasgroup=true 
```

#### or use `air` run golang project
> air can monitor changes in project files.
see: https://github.com/cosmtrek/air
```ini
; 应用程序的路径
command=air -c .air.toml              
; 应用程序的运行目录。如果应用程序涉及相对路径的文件，这个一定要修改好
directory={$project_path}              
; start at supervisord start (default: true)
autostart=true                
; 设置得太低可能会因为没来得及启动被判启动失败
startsecs=5                   
; max # of serial start failures when starting (default 3)
startretries=3                
; when to restart if exited after running (def: unexpected)
autorestart=true        
; send stop signal to the UNIX process group (default false)
stopasgroup=true             
; SIGKILL the UNIX process group (def false)
killasgroup=true 
```

### more setting
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/12.3.md
https://blog.csdn.net/xyang81/article/details/51555473
https://www.missshi.cn/api/view/blog/5aafcf405b925d681e000000
https://blog.csdn.net/u012724150/article/details/54616600
https://www.jianshu.com/p/3c70e12b656a
https://blog.csdn.net/sinat_21302587/article/details/76836283
http://wonse.info/centos7-supervisor-auto-start.html
https://studygolang.com/articles/23750