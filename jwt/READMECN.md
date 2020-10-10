## jwt 
包含生成token，注销token等函数

'user_jwt:' + uuid 和 payload_id + 过期时间 作为 key-value保存在redis数据库中，返回token的内容含有用户uuid和用户类型信息和jwt的uuid等信息
