外部测试应用程序和测试数据。随时根据需要构建/test目录。对于较大的项目，有一个数据子目录更好一些。例如，
如果需要Go忽略目录中的内容， 则可以使用/test/data或/test/testdata这样的目录名字。
请注意，Go还将忽略以“.”或“_”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。

更多样例查看/test。

Examples:

https://github.com/openshift/origin/tree/master/test (test data is in the /testdata subdirectory)