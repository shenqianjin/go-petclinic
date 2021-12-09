说明：
本项目的主干。

不要在这个目录中放置太多代码。 通常只是一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。

Examples:

https://github.com/heptio/ark/tree/master/cmd (just a really small main function with everything else in packages)
https://github.com/moby/moby/tree/master/cmd
https://github.com/prometheus/prometheus/tree/master/cmd
https://github.com/influxdata/influxdb/tree/master/cmd
https://github.com/kubernetes/kubernetes/tree/master/cmd
https://github.com/dapr/dapr/tree/master/cmd
https://github.com/ethereum/go-ethereum/tree/master/cmd