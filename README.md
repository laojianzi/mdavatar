# MDAvatar

<p align="center">
  <a href="#">
    <img alt="MDAvatar" width="128" height="128" src="https://github.com/laojianzi/mdavatar/blob/main/mdavatar.png?raw=true">
  </a>
</p>
<p align="center">
  <b>MDAvatar</b> 可以根据字符串生成单字符头像，并且可以高度自定义，支持生成中文头像
</p>

## Features

- 自定义头像 text 处理方式
- 自定义图片 size
- 自定义字体
  - 自带默认字体，字体颜色为白色，暂不支持自定义字体颜色
  - 当字体设置支持中文时，头像 text 也可以渲染中文
- 自定义头像背景颜色列表
  - 自带默认颜色列表 Material Design Colors
  - 当没有设置 background 时，默认是随机获取颜色列表中的一种 RGBA
- 自定义背景
  - 当设置 background 后，不会启用随机颜色作为 background
  - background 也是一种 RGBA
  
## Installation

- Require `go` version >= `1.13`
- Require `go mod` enable
  
```bash
$ go get -u github.com/laojianzi/mdavatar
```

## Examples

```go
package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/laojianzi/mdavatar"
)

func main() {
	avatar, err := mdavatar.New("MDAvatar").Build()
	if err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("out-%d.png", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := png.Encode(file, avatar); err != nil {
		log.Fatal(err)
	}
}
```

## TODO

- [ ] 支持 cli (生成 png/jpg)
- [ ] 支持自定义形状 (圆形、椭圆形、方形 ...)
- [ ] 支持多种返回形式 (HTTP、Base64、WriteToFile ...)

## Prior Art

项目参考了一些现有的思路或者实现等

- [mdclub](https://github.com/zdhxiong/mdclub/tree/master/src/Vendor)
- [Michael Okoko](https://blog.logrocket.com/working-with-go-images/)

## License

This project is licensed under the MIT License.

License can be found [here](https://github.com/laojianzi/mdavatar/blob/master/LICENSE).
