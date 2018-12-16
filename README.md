# Back-END

本项目实现了一个基本的新闻后台网站的功能：

+ 分类别展示新闻
+ 新闻分页
+ 用户登录注册（基于jwt）
+ 用户发表评论
+ 用户收藏新闻

新闻来源于聚合数据API，我们将json数据通过utils中封装好的数据库操作方法，把json数据保存到
boltdb数据库中。

API设计见[文档](https://github.com/PROJECT-PML/Back-END/blob/master/API_DESIGN.md)
项目设计思路见[文档](https://github.com/PROJECT-PML/FILE/blob/master/%E5%AE%9E%E9%AA%8C%E6%8A%A5%E5%91%8A.md)