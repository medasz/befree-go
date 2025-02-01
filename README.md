# befree-go
go版befree代理池

> 能夠在linux、windows、darwin多平台運行

```shell
[root@centos7 test]# ./befree-go -h
Usage of 
 _______             ________                                       ______            
|       \           |        \                                     /      \           
| $$$$$$$\  ______  | $$$$$$$$______    ______    ______          |  $$$$$$\  ______  
| $$__/ $$ /      \ | $$__   /      \  /      \  /      \  ______ | $$ __\$$ /      \ 
| $$    $$|  $$$$$$\| $$  \ |  $$$$$$\|  $$$$$$\|  $$$$$$\|      \| $$|    \|  $$$$$$\
| $$$$$$$\| $$    $$| $$$$$ | $$   \$$| $$    $$| $$    $$ \$$$$$$| $$ \$$$$| $$  | $$
| $$__/ $$| $$$$$$$$| $$    | $$      | $$$$$$$$| $$$$$$$$        | $$__| $$| $$__/ $$
| $$    $$ \$$     \| $$    | $$       \$$     \ \$$     \         \$$    $$ \$$    $$
 \$$$$$$$   \$$$$$$$ \$$     \$$        \$$$$$$$  \$$$$$$$          \$$$$$$   \$$$$$$ 
:
  -c string
    	Specify your custom clash.exe path
  -f string
    	Specify a contain subscribe file path (default "./aaa.txt")
  -p int
    	Specify a port number(http&socks5) (default 59981)
  -t string
    	Specify a link for speed testing(default:https://www.google.com)
  -y string
    	Specify a yourself clash yaml file (default "sectest.yaml")

```
## 待辦項
- [x] 完成-y和-c參數
- [x] 完成linux和windows兼容
- [ ] 完成-f參數
- [ ] 完成-p參數
- [ ] 完成-t參數
- [ ] 完成mac測試

## 注意
<span style="color: red">自帶的windows和linux版clash全是本人從網上下載，請自行甄別安全性，也可通過`-c`參數指定clash
