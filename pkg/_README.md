说明：
外部应用程序可以使用的库代码(例如 /pkg/mypubliclib)。
其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思。
如果涉及到接口变动，请考虑是否需要升级版本。

注意，internal 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。
/pkg 目录仍然是一种很好的方式，可以显式地表示该目录中的代码对于其他人来说是安全使用的好方法。

Examples:

https://github.com/prometheus/prometheus/tree/master/pkg
https://github.com/istio/istio/tree/master/pkg
https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
https://github.com/google/gvisor/tree/master/pkg
https://github.com/google/syzkaller/tree/master/pkg
https://github.com/kubernetes/kubernetes/tree/master/pkg
https://github.com/kubernetes/kops/tree/master/pkg
https://github.com/grafana/grafana/tree/master/pkg
