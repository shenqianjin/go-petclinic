说明：
私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。
请注意，这个布局模式是由 Go 编译器本身执行的。有关更多细节，请参阅Go 1.4 release notes 。
注意，你并不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录。
