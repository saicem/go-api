# iwut-api-gin

## 部署

### swagger

在项目根目录下运行 `swag init` 生成 swagger 必要文件

#### 建立数据库

- 数据库名 iwut_log
- 用户名 iwut
- 密码 129028

## Api 使用

#### 上传日志

时间的解析目前只能接受以下格式

- 2000-10-10T12:12:12Z
- 2000-10-10T12:12:12+08:00

#### 事件

| 对象        | 事件     | 附加信息                                |
| ----------- | -------- | --------------------------------------- |
| Application | StartUp  |                                         |
|             | Exit     |                                         |
| Course      | Import   |                                         |
|             | Clear    |                                         |
|             | Add      |                                         |
|             | Remove   |                                         |
|             | BgChange |                                         |
|             | BgClear  |                                         |
| Room        | Query    | Week,Day,SectionStart,SectionEnd,Region |
| Library     | Query    | Keyword,SearchWay                       |
|             | Detail   | Id                                      |
| News        | Detail   | Id                                      |
|             | Query    | Tag                                     |