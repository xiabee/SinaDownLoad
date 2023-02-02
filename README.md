# SinaDownLoad
为了拯救挂掉的新浪微博图床，我写了个脚本批量下载图床里的图

* 当前进度：可以实现特定目录的批量下载
* BUG：目前下载完所有图片会 panic，过两天来处理



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

修改 `main.go` 中的 `path`：此目录为需要操作的根目录，本程序会自动读取目录内的所有文件，找到图片并下载到指定文件夹

```go
path := "/Users/xiabee/Desktop/GitHub/gitpage/source/_posts"
// change your path here
```



## 目标位置

* 程序会在当前目录中创建 `Download` 目录，图片下载在该目录中

* `main.go` 中的 `Download` 字符串可以自定义修改，修改后即为下载图片的目标位置（仅支持相对路径）

```go
errDown := lib.BatchDownLoad(urlList, "Download")
```

