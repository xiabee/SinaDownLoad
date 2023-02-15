# SinaDownLoad
为了拯救挂掉的新浪微博图床，我写了个脚本批量下载图床里的图

* 当前进度：可以实现指定目录的批量下载



## 应用场景

* 源文件中引用了 `https://weibo.com` 等一系列图床中的图
* 源文件可进行按行读操作（如`.md`文件、`.txt`文件、各类代码文件等）



## 使用方法

```bash
git clone https://github.com/xiabee/SinaDownLoad.git
cd SinaDownLoad
go build
./SinaDownLoad
```



## 下载内容

* 程序会提示输入需要下载的内容的目录，默认目录为当前工作目录

* 程序会自动读取目录内的所有文件，找到图片并下载到指定文件夹
* 仅支持绝对路径

```go
[+] Directory to be processed: [/Users/xiabee/Desktop/GitHub/SinaDownLoad]：
```



## 目标位置

* 程序会提示输入目标位置，默认目录为当前目录下创建的 `Download` 目录

* 仅支持相对路径


```go
[+] Download location:[Download]：
```

