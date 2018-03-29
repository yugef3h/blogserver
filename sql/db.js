var mysql = require('mysql');
var $conf = require('./config');
var $sql = require('./sql');
var crypto = require('crypto');
var jwt = require("jsonwebtoken");

var pool = mysql.createPool($conf.mysql);
//向前台返回JSON方法的简单封装
var jsonWrite = function (res,ret) {
    if(typeof ret === 'undefined') {
        res.json({
            'code': '1',
            'msg': '操作失败'
        })
    } else {
        res.json(ret)
    }
};
module.exports = {
    //添加电影
    addMovie: function (req, res, next) {
        pool.getConnection(function (err, connection) {
            //获取前台页面传来的参数
            //建立连接，向表中插入值
            var params = req.query || req.body
            var data = {
                name: params.name,
                rating: params.rating,
                imgUrl: params.imgUrl,
                description: params.description,
                kind: params.kind
            }
            connection.query($sql.insert, data, function (err, result) {
                if(result) {
                    result = {
                        code: 200,
                        msg: '增加成功'
                    }
                }
                //以json形式，把操作结果返回给前台界面
                jsonWrite(res, result);
                //释放连接
                connection.release();
            })
        })
    },
    //根据种类查询电影
    queryByKind: function (req, res, next) {
        var kind = req.query.kind || req.body.kind //为了拼凑正确的sql语句，这里要转下整数,由于kind是中文的话获取是乱码所以这里做一个转化，QAQ我才不会说是因为改了很多次编码格式都不成功才出此下策，
        kind == 'isShow' ? kind = '正在热映' : kind
        kind == 'willShow' ? kind = '即将上映' : kind
        var count = req.query.count ? parseInt(req.query.count) : 10 //返回的数量
        var start = req.query.start ? parseInt(req.query.start) : 0 //从第几个开始返回从0开始计数
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryByKind, [kind, start, count], function (err, result) {
                var newRes = {}
                newRes.result = result
                newRes.count = count //返回的字段增加count
                newRes.start = start //返回的字段增加start
                jsonWrite(res, newRes);
                connection.release();
            })
        })
    },
    //根据电影名称来更新电影的信息
    updateByName: function (req, res, next) {
        var params = req.query || req.body
        var data = {
            name: params.name,
            rating: params.rating,
            imgUrl: params.imgUrl,
            description: params.description,
            kind: params.kind
        }
        pool.getConnection(function (err, connection) {
            connection.query($sql.updateByName, [data, params.name], function (err ,result) {
                if(result) {
                    result = {
                        code: 200,
                        msg: '修改成功'
                    }
                }
                jsonWrite(res, result);
                connection.release();
            })
        })
    },
    //根据名称模糊查询
    queryByName: function (req, res, next) {
        var name = req.query.name || req.body.name
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryByName + "'%"+ name + "%'", function (err, result) {
                jsonWrite(res, result)
            })
        })
    },
    //登陆
    login: function (req, res, next) {
        var secret = 'learnRestApiwithNickjs';
        var md5 = crypto.createHash('md5');
        var password = req.body.password;
        var username = req.body.username;
        var md5pass = md5.update(password).digest('base64');
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryLogin,[username],function (err, data) {
                if (data.length === 0) {
                    res.end('{"err":"抱歉，系统中并无该用户，如有需要，请向管理员申请"}');//replace:res.send();return;
                } else if (data[0].Password !== md5pass) {
                    res.end('{"err":"密码不正确"}');
                } else if (data.length !== 0 && data[0].Password === md5pass) {
                    //创建token
                    var token= jwt.sign({username: username}, secret, {
                        'expiresIn': 10080 // 设置过期时间
                    });
                    //req.session.username = req.body.username; //自定义req.session.username,存session
                    //req.session.password = md5pass;
                    //json格式返回token
                    //待引入jsonwrite封装
                    res.json({
                        success:true,
                        message:'Enjoy your token!',
                        token:token
                    });
                }
                connection.release();
            })
        })
    },
    //查找article
    article: function (req, res, next){
        var page = req.body.page ? parseInt(req.body.page) : 1;
        //未转换
        var rows = req.body.rows;
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryArticle, [(page-1)*rows,rows],function (err, data, count) {
                connection.query($sql.queryId,function (err, count) {
                    var obj = {
                        data:data,
                        total:count,
                        success:"成功"
                    };
                    var str = JSON.stringify(obj);
                    res.send(str);
                });
                connection.release();
            })
        });
    },
    //查找case
    case: function (req, res, next) {
        var page = req.body.page ? parseInt(req.body.page):1;
        //待处理整数值
        var rows = req.body.rows;
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryCase,[(page-1)*rows,rows],function (err, data,count){
                var obj = {
                    data:data,
                    total:count,
                    success:"成功"
                };
                var str = JSON.stringify(obj);
                res.send(str);
                connection.release();
            })
        })
    },
    //删除用户
    delUsers: function (req, res, next) {
        pool.getConnection(function (err, connection) {
            connection.query($sql.delUsers, [req.body.Id], function (err, data) {
                if(data.length==0){
                    res.end('{"err":"抱歉，删除失败"}');
                }else{
                    var obj = {
                        success:"删除成功"
                    };
                    var str = JSON.stringify(obj);
                    res.end(str);
                }
                connection.release();
            })
        })
    },
    //添加用户
    addUsers: function (req, res, next) {
        var md5 = crypto.createHash('md5');
        req.body.password = md5.update(req.body.password).digest('base64');
        pool.getConnection(function (err, connection) {
            connection.query($sql.addUsers, [req.body.phone,req.body.password,req.body.name], function (err, data) {
                if(data.length==0){
                    res.end('{"err":"抱歉，添加失败"}');
                }else{
                    res.end('{"success":"添加成功"}');
                }
                connection.release();
            })
        })
    },
    //更新用户信息
    putUsers: function (req, res, next) {
        var name = ''+req.body.name;
        pool.getConnection(function (err, connection) {
            connection.query($sql.putUsers, [name,req.body.phone,req.body.Id], function (err, data) {
                if( data.length==0 ){
                    res.end('{"err":"抱歉，修改失败"}');
                }else{
                    res.end('{"success":"修改成功"}');
                }
                connection.release();
            })
        })
    },
    //改密
    putPass: function (req, res, next) {
        var username = jwt.decode(req.body.token).username
        // crypto md5多次加密 不能 设置变量 var md5 = crypto.createHash('md5') !!!!错误
        var oldpass = crypto.createHash('md5').update(req.body.oldpass).digest('base64');
        var newpass = crypto.createHash('md5').update(req.body.newpass).digest('base64');
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryPass, [username], function (err, data) {
                //console.log(data[0].Password===oldpass)
                if(data[0].Password===oldpass){
                    connection.query($sql.putPass, [newpass,username],function (err,data){
                            if( data.length==0 ){
                                res.end('{"err":"抱歉，修改失败"}');
                            }else{
                                res.end('{"success":"修改成功"}');
                            }
                    })
                }
                connection.release();
            })
        })
    },
    //管理员列表
    admin: function (req, res, next) {
        var page = req.body.page? parseInt(req.body.page) : 1;
        var rows = req.body.rows? parseInt(req.body.rows) : 5;
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryUsers, [(page-1)*5,rows], function (err, data) {
                    connection.query($sql.queryId, function (err,count){
                        //console.log(data)
                        var obj = {
                            data:data,
                            total:count,
                            success:"成功"
                        };
                        var str = JSON.stringify(obj);
                        res.send(str);
                    })
                    connection.release();
            })
        })
    },
    //删除文章
    delArticle: function (req, res, next) {
        pool.getConnection(function (err, connection) {
            connection.query($sql.delArticle, [req.body.id], function (err, data) {
                if(data.length==0){
                    res.end('{"err":"抱歉，删除失败"}');
                }else{
                    var obj = {
                        success:"删除成功"
                    };
                    var str = JSON.stringify(obj);
                    res.end(str);
                }
                connection.release();
            })
        })
    },
    //???
    essay: function (req, res, next) {
        var title = req.body.title,
            tag='web',
            author='admin',
            content=req.body.content,
            des = '暂无',
            img = ['bg2.jpg','3e.jpg','2r.jpg','1k.jpg','21.jpg','024.jpg','36.jpg','93.jpg','201.jpg','a0.jpg'],
            index = Math.floor(Math.random()*5)+1,
            dir = img[index],
            createtime=new Date().toLocaleString().substring(0,10);//去掉substring可显示具体时间
        pool.getConnection(function (err, connection) {
            connection.query($sql.queryEssage, [title,tag,author,content,createtime,dir,des], function (err, data) {
                if(err){//err是必须的
                    res.status(404).end(`警告，来不及说再见，部分内容已消失在外太空...`);
                    return;
                }
                res.json({
                    result:'success!'
                });
                connection.release();
            })
        })
    },
    //文章详情
    artId: function (req, res, next) {
        //console.log(req.params.id);
        //req.route.path = "/page"; //修改path来设定 对 数据库的操作?
        pool.getConnection(function (err, connection) {
            connection.query($sql.artId, [req.params.id], function (err, data) {
                if(data.length == 0){
                    //返回状态码，避免被爬虫错误收录
                    res.status(404).end('err');
                    return;
                }
                var obj = {
                    data:data,
                    success:"成功"
                };
                var str = JSON.stringify(obj);
                res.send(str);
                connection.release();
            })
        })
    },
    //案例详情
    caseId: function (req, res, next){
        //console.log(req.params.id);
        pool.getConnection(function (err, connection) {
            connection.query($sql.caseId, [req.params.id], function (err, data) {
                if(data.length == 0){
                    //返回状态码，避免被爬虫错误收录
                    res.status(404).end('err');
                    return;
                }
                var obj = {
                    data:data,
                    success:"成功"
                };
                var str = JSON.stringify(obj);
                res.send(str);
                connection.release();
            })
        })
    }
};
