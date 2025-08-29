1.登录和注册接口不需要再请求头里携带token，其它接口需要在请求头添加 Authorization 格式为 Bearer +登录接口返回的token

2 接口 ,接口使用post 方式 json格式，响应业务码200 为业务处理成功，其它为失败
2.1 注册
http://localhost:8081/user/register
请求参数
{
 "username":"张三",
 "password":"123456",
 "email":"xasljkl@qq.com"
}
响应
{
    "code": 200,
    "msg": "操作成功",
    "data": null
}
2.2 登录
http://localhost:8081/user/login
请求参数
{
 "username":"张三",
 "password":"123456"
}
响应参数
// data里面的内容就为token
{
    "code": 200,
    "msg": "操作成功",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY1MjU1OTQsInVzZXJJZCI6NiwidXNlck5hbWUiOiLlvKDkuIkifQ.O3LcaxE4hpU9HqS6_uYVXOpvBrwfX8EHw0pMXS3Ux-Y"
}
2.3 文章提交
http://localhost:8081/post/add
{
 "title":"文章1",
 "content":"总之，int类型和uint类型之间的转换需要根据具体的情况而定，需要充分了解数据类型的特征和范围，才能避免出现不必要的错误"
}
公共响应

2.4 查询用户的文章列表
