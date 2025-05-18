# 📋 测试用例设计

| 编号  | 测试场景               | 请求体内容                                  | 预期响应                                          | 实际结果 | 状态 |
| ----- | ---------------------- | ------------------------------------------- | ------------------------------------------------- | -------- | ---- |
| TC001 | 登录成功               | Username+ password 都正确                   | 状态码 200，返回 token                            | 通过     | ✅    |
| TC002 | 账号错误               | password  正确，Username缺失                | 状态码 401，error 为 Invalid username or password | 通过     | ✅    |
| TC003 | 密码错误               | Username 正确，password 缺失                | 状态码 401，error 为 Invalid username or password | 通过     | ✅    |
| TC004 | 注册                   | Username+ password 都正确                   | 状态码 200，返回 User registered successfully     | 通过     | ✅    |
| TC005 | 注册                   | Username 缺失，password 缺失                | 状态码 400，error                                 | 通过     | ✅    |
| TC006 | 取所单个文章的详细信息 |                                             | 状态码 200，返回文章的详细信息                    | 通过     | ✅    |
| TC007 | 创建文章               | Title正确，Content正确                      | 状态码 400，返回文章的详细信息                    | 通过     | ✅    |
| TC008 | 取所有文章列表详细信息 |                                             | 状态码 400，返回文章列表的详细信息                | 通过     | ✅    |
| TC009 | 更新文章               | ID正确,Title正确，Content正确，非创建者修改 | 状态码 403，error 为 No permission                | 通过     | ✅    |
| TC010 | 更新文章               | ID正确,Title正确，Content正确，创建者修改   | 状态码 200，返回文章的详细信息                    | 通过     | ✅    |
| TC011 | 删除文章               | ID正确,非创建者删除                         | 状态码 403，error 为 No permission                | 通过     | ✅    |
| TC012 | 删除文章               | ID正确,创建者删除                           | 状态码 200，返回 Post deleted                     | 通过     | ✅    |
| TC013 | 获取文章的所有评论列表 | 文章ID正确                                  | 状态码 200，返回评论列表信息                      | 通过     | ✅    |
| TC014 | 新增评论               | 文章ID正确，Content正确                     | 状态码 200，返回评论列表信息                      | 通过     | ✅    |
| TC015 | 新增评论               | 文章ID正确                                  | 状态码 400，error                                 | 通过     | ✅    |



