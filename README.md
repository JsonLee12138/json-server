### app 模块
1. controller 处理请求入口参数处理
2. service 处理业务逻辑
3. repository 处理数据库操作
4. entity 处理数据库实体
5. dto 处理数据传输对象

在注入的时候需要以此注入
在controller的entry中注入service
在service的entry中注入repository
在repository的entry中注入entity