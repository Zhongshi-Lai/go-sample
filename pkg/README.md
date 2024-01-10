# 项目目录下的pkg

如果你把代码包放在根目录的pkg下，其他项目(指别的git引用你这个git)是可以直接导入pkg下的代码包的，即这里的代码包是开放的
与之相对,别的项目无法引用你项目的internal下面的任何逻辑,在编译期会报错
当然你的项目本身也可以直接访问自己的pkg。但是如果你要把代码放在pkg下，还想需要三思而后行吧，有没必要这样做，毕竟internal目录是最好的方式保护你的代码并且被go编译器强制校验internal的代码包不可分享的。
如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，这样做是合适的，如果你自己或公司的某一个项目，个人的经验，基本上用不上pkg