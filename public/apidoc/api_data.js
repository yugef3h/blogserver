define({ "api": [
  {
    "type": "POST",
    "url": "/users/login",
    "title": "",
    "description": "<p>登录验证</p>",
    "name": "login",
    "parameter": {
      "fields": {
        "path参数": [
          {
            "group": "path参数",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": ""
          },
          {
            "group": "path参数",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": ""
          }
        ]
      }
    },
    "sampleRequest": [
      {
        "url": "http://localhost:3000/users/login"
      }
    ],
    "group": "User",
    "version": "1.0.0",
    "filename": "routes/users.js",
    "groupTitle": "User"
  }
] });
