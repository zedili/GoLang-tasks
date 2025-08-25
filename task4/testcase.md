1、注册
```bash
curl --location --request POST 'http://127.0.0.1:8090/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Username":"user1",
    "Password":"Password"
}'
```
响应
```json
{
    "data": null,
    "message": "注册成功"
}
```
2、登陆
```bash
curl --location --request POST 'http://127.0.0.1:8090/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Username":"user1",
    "Password":"Password"
}'
```
响应
```json
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA",
    "message": "登陆成功"
}
```
3、创建文章
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/post/create' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"文章11",
    "content":"文章11内容"
}
'
```
响应
```json
{
    "data": null,
    "message": "创建文章成功"
}
```
4、文章列表
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/post/listPost' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
}
'
```
响应
```json
{
    "data": [
        {
            "ID": 3,
            "Title": "post 1",
            "Content": "this is post 1",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 4,
            "Title": "post 2",
            "Content": "this is post 2",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 5,
            "Title": "post 1",
            "Content": "this is post 5",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 6,
            "Title": "post 2",
            "Content": "this is post 6",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 1,
            "Title": "post 1",
            "Content": "this is post 5",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 7
        },
        {
            "ID": 2,
            "Title": "post 2",
            "Content": "this is post 6",
            "UserID": 1,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 7,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 8,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 9,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 10,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 11,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 12,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 13,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 14,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 15,
            "Title": "",
            "Content": "",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 16,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 17,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 18,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 19,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 20,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 21,
            "Title": "文章1",
            "Content": "文章1内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        },
        {
            "ID": 22,
            "Title": "文章11",
            "Content": "文章11内容",
            "UserID": 3,
            "Comments": null,
            "CommentStatus": "",
            "CommentCount": 0
        }
    ],
    "message": "查询成功"
}
```
5、文章详情
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/post/getPostById' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":1
}
'
```
响应
```json
{
    "data": {
        "ID": 1,
        "Title": "post 1",
        "Content": "this is post 5",
        "UserID": 1,
        "Comments": null,
        "CommentStatus": "",
        "CommentCount": 7
    },
    "message": "成功"
}
```
6、删除文章
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/post/deleteById' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":22
}
'
```
响应
```json
{
    "data": null,
    "message": "删除成功"
}
```
7、创建评论
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/comment/create' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "postId":1,
    "content":"pp"
}
'
```
响应
```json
{
    "data": null,
    "message": "评论文章成功"
}
```
8、查看文章评论
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/comment/listComment' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "postId":1
}
'
```
响应
```json
{
    "data": [
        {
            "ID": 1,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 2,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 3,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 4,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 5,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 6,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 7,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 8,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 19,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 20,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 21,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 22,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 23,
            "Content": "1th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 24,
            "Content": "2th comment of post 1",
            "UserID": 1,
            "PostID": 1
        },
        {
            "ID": 25,
            "Content": "pp",
            "UserID": 3,
            "PostID": 1
        },
        {
            "ID": 26,
            "Content": "pp",
            "UserID": 3,
            "PostID": 1
        }
    ],
    "message": "查询成功"
}
```
9、删除评论
```bash
curl --location --request POST 'http://127.0.0.1:8090/user/comment/deleteById' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxOTg0MjYsImlkIjozLCJ1c2VybmFtZSI6InVzZXIxIn0.g7srHndkUQO-Z8VqhADtmkU4fWUZbPljaKYdEePR7EA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":25
}
'
```
响应
```bash
{
    "data": null,
    "message": "删除成功"
}
```
