## jwt
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/jwt/READMECN.md)
generate token and revoke token function

'user_jwt:' + uuid, and payload_id + expired time, as key-value saved in redis. Token, the return of the function, contains user's uuid user's type and jwt-uuid.

