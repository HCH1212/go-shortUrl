# go-shortUrl
golang实现用户登录注册,接口自定义, 实现短链生成、删除、修改、查看（用户短链、短链排行）， rpc 框架kitex的使用

主要使用gin框架，mysql作数据库，redis实现缓存，短链实现算法是id自增长，用了一点logrus的日志功能，感觉最难的还是kitex的使用

附上参考资料
https://blog.csdn.net/qq_39149275/article/details/118968366?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522171177393716800222876437%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=171177393716800222876437&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-1-118968366-null-null.142^v100^pc_search_result_base9&utm_term=golang%20%E5%AE%9E%E7%8E%B0%E7%9F%AD%E9%93%BE%E7%B3%BB%E7%BB%9F&spm=1018.2226.3001.4187
https://blog.csdn.net/weixin_34362991/article/details/89620012?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_utm_term~default-4-89620012-blog-88986486.235^v43^pc_blog_bottom_relevance_base8&spm=1001.2101.3001.4242.3&utm_relevant_index=7
https://blog.csdn.net/qq_43716830/article/details/128731243?ops_request_misc=&request_id=&biz_id=102&utm_term=kitex&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-9-128731243.nonecase&spm=1018.2226.3001.4187
https://www.cloudwego.io/zh/docs/kitex/
https://blog.csdn.net/weixin_57209831/article/details/128913767
