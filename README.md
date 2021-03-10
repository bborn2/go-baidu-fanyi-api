# go-baidu-fanyi-api
百度通用翻译api接口的golang实现

参照官网的python实现 https://fanyi-api.baidu.com/doc/21

暂时未考虑参数检查，返回判断等健壮性问题

使用方法：

import baidufanyiapi "github.com/bborn2/go-baidu-fanyi-api"

ret := baidufanyiapi.Translate("yourappid", "yourappkey", "en", "zh", "Hello")
