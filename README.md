# warehouse-chatroom 聊天室

## gitlab 私服配置
## 部署

### 环境准备:
1. Ubuntu 16.04 + DOCKER 最新环境
2. 配置 Nginx 反向代理
    ```bash
    mkdir -p /etc/nginx/conf.d
    
    cat << EOF
    server {
      listen 80;
      server_name micro-web-test.srgow.com;
      location / {
        proxy_pass http://{主机IP}:8082;
      }
    }
    > /etc/nginx/conf.d/micro-web-test.conf
    
    cat << EOF
    server {
      listen 80;
      server_name gateway-test.srgow.com;
      location / {
        proxy_pass http://{主机IP}:8089;
      }
    }
    > /etc/nginx/conf.d/gateway-test.conf
    ``` 
3. 运行 nginx docker 容器
    ```bash
    docker run -d -p 80:80 --name nginx --restart unless-stopped -v /etc/nginx/conf/conf.d:/etc/nginx/conf.d nginx:1.14
    ```
4. 解压部署包
    ```bash
    tar -zxvf chatroom.tar chatroom/path
    cd chatroom/path
    cp tool/gcctl /usr/local/bin/
    ```
5. 配置访问主机hosts文件 (windows)
    ```text
    {聊天室主机IP} gateway-test.srgow.com
    {聊天室主机IP} micro-web-test.srgow.com
    ```
### 运行 聊天室容器
```bash
cd chatroom/path
gcctl package
gcctl deploy
```

## 访问聊天室 (windows)
1. http://micro-web-test.srgow.com/call 可直接调用api
2. http://micro-web-test.srgow.com/registry 可查看API调用报文格式
3. com.test365.warehouse.alydnh.api 为 外部访问api gateway
3.1 可使用 rest 方式访问 api gateway 如: post http://micro-web-test.srgow.com/api/member/login
4. com.test365.warehouse.alydnh.room 为 内部聊天室房间相关接口

### 登录
1. 进入网址: http://micro-web-test.srgow.com/call
2. Service: com.test365.warehouse.alydnh.api
3. Endpoint: Member.Login
4. Request
    ```json
    {"name":"alydnh"}
    ```
5. 返回：
    ```json
    {
      "result": {
        "success": true
      },
      "token": "登录token"
    }
    ```
5. api 调用地址: http://micro-web-test.srgow.com/api/member/login

### 创建房间
1. 进入网址: http://micro-web-test.srgow.com/call
2. Service: com.test365.warehouse.alydnh.room
3. Endpoint: RoomService.Create
4. Request:
    ```json
    {
      "result": {
        "success": true
      },
      "room": {
        "id": "fcf602901ef1a0fba0f8533181d0cfb8",
        "memberNames": [
          "alydnh",
          "realgang",
          "tracy"
        ]
      }
    }
    ```
5. 返回:
    ```json
    {
      "result": {
        "success": true
      },
      "room": {
        "id": "fcf602901ef1a0fba0f8533181d0cfb8",
        "memberNames": [
          "alydnh",
          "realgang",
          "tracy"
        ]
      }
    }
    ```
  
### 发送消息
1. 进入网址: http://micro-web-test.srgow.com/call
2. Service: com.test365.warehouse.alydnh.api
3. Endpoint: Chat.SendMessage
4. Request:
    ```json
    {
      "token": "dedf58dc7c31a810ec978a40188da753",
      "request": {
        "messages": [
          {
            "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
            "labels": {
              "a": "b"
            },
            "body": {
              // 数据扩展，限制字符串转换
            }
          }
        ]
      }
    }
    ```
5. 返回:
    ```json
    {
      "result": {
        "success": true
      }
    }
    ```

### 接收消息
1. 进入网址: http://micro-web-test.srgow.com/call
2. Service: com.test365.warehouse.alydnh.api
3. Endpoint: Chat.ReceiveMessage
4. Request:
    ```json
    {
      "token":"7c25c3590335ed8fc66229fb8ddc8ebd",
      "consumeMessageIDs":["1487a81eb698aa11f8b41dab398e082c"], //消费已收到消息，否则会再次收到前次消息
      "capacity": 10 // 容量
    }
    ```
5. 返回:
    ```json
    {
      "result": {
        "success": true
      },
      "messages": [
        {
          "sender": "alydnh",
          "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
          "labels": {
            "a": "b"
          },
          "id": "ec95a0b82a80a67a0c585574530b9af5",
          "timestamp": "20200803141947"
        },
        {
          "sender": "alydnh",
          "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
          "labels": {
            "a": "b"
          },
          "id": "9419a6384f26ad10dd808798300e762f",
          "timestamp": "20200803141947"
        },
        {
          "sender": "alydnh",
          "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
          "labels": {
            "a": "b"
          },
          "id": "2f41f49f3776e47c1e9044c254dfb349",
          "timestamp": "20200803141948"
        },
        {
          "sender": "alydnh",
          "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
          "labels": {
            "a": "b"
          },
          "id": "dfaba19ee9d61f13c52bd5d67d4921dd",
          "timestamp": "20200803141949"
        },
        {
          "sender": "alydnh",
          "roomID": "fcf602901ef1a0fba0f8533181d0cfb8",
          "labels": {
            "a": "b"
          },
          "id": "5619c41ea1b54ea9cc8ee0600740af7f",
          "timestamp": "20200803141949"
        }
      ]
    }
    ```