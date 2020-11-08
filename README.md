# authorization
##权限验证表: 
基本由以下几部分组成
```
用户 角色 权限 用户角色 角色权限

```
用户
```json
  "user_id" : 用户id
  "username" : 用户名
  "password" : 密码
  "salt" : 盐值
  "nick_name" : 昵称
  "email" : 邮箱
  "mobile" : 电话
  "status" : 状态
  "create_user_id" : 创建此的用户id
  "created_at" : 创建时间
  "updated_at" : 更新时间
  "deleted_at" : 删除时间

```
角色
```json
  "role_id" : 角色id
  "role_name" : 角色名
  "remark" : 标志
  "create_user_id" : 创建此的用户id
  "created_at" : 创建时间
  "updated_at" : 更新时间
  "deleted_at" : 删除时间
```

菜单
```json
    "menu_id" : 菜单id
    "parent_id" : 父菜单id
    "name" : 名称
    "url" : 路径
    "perms" : 可以通过的权限标记
    "type" : 菜单级别
    "icon" : 标志
    "order_num" : 顺序
```

用户角色
```json
  "id" : id
  "user_id" : 用户id
  "role_id" : 角色id
```
角色权限
```json
    "id" : id
    "role_id" : 角色id
    "menu_id" : 菜单id
```
token
```json
  "user_id" : 用户id
  "token" : token
  "expire_time" : 失效时间
  "update_time" : 更新时间
```
##迭代顺序
```
- java spring security 
    -> Spring Could Security OAuth2.0  
    -> golang jwt-go

+ postgresql -> tidb 
```