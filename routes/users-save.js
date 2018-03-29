var express = require('express');
var router = express.Router();
var sql = require('./db.js');
var crypto = require('crypto');
var jwt = require("jsonwebtoken");
var path=require('path');
var ueditor=require('ueditor');

router.use('/public', express.static(__dirname + '/public'));

//login
router.post('/login', function (req, res, next) {
    var secret = 'learnRestApiwithNickjs';
    var md5 = crypto.createHash('md5');
    var md5pass = md5.update(req.body.password).digest('base64');
    sql('select * from users where UserName=?', [req.body.username], function (err, data) {
        // console.log(data);
        if (data.length === 0) {
            res.end('{"err":"抱歉，系统中并无该用户，如有需要，请向管理员申请"}');//replace:res.send();return;
        } else if (data[0].Password !== md5pass) {
            res.end('{"err":"密码不正确"}');
        } else if (data.length !== 0 && data[0].Password === md5pass) {
            //创建token
            var token= jwt.sign({username: req.body.username}, secret, {
                'expiresIn': 10080 // 设置过期时间
            });
            //req.session.username = req.body.username; //自定义req.session.username,存session
            //req.session.password = md5pass;
            //json格式返回token
            res.json({
                success:true,
                message:'Enjoy your token!',
                token:token
            });
            //res.end('{"success":"true"}');
        }
    });
});
// delete_users
router.post('/delete_u', function(req, res, next) {
    sql('delete from users where Id=?', [req.body.Id],function(err,data){
        if(data.length==0){
            res.end('{"err":"抱歉，删除失败"}');
        }else{
            var obj = {
                success:"删除成功"
            };
            var str = JSON.stringify(obj);
            res.end(str);
        }

    });
});
// add_users
router.post('/add_u', function(req, res, next) {
    //console.log(req.body);
    var md5 = crypto.createHash('md5');
    req.body.password = md5.update(req.body.password).digest('base64');
    sql('insert into users (Id,ComTelephone,Password,UserName,CompanyId,IsAdmin) values (0,?,?,?,0,0)',
        [req.body.phone,req.body.password,req.body.name],function(err,data){
            console.log(data);
            if(data.length==0){
                res.end('{"err":"抱歉，添加失败"}');
            }else{
                res.end('{"success":"添加成功"}');
            }
        });
});
// update_users
router.post('/update_u', function(req, res, next) {
    var name = ''+req.body.name
    //console.log(name)
    sql('update users set UserName=?,ComTelephone=?  where Id=?;',
        [name,req.body.phone,req.body.Id],function(err,data){
            if( data.length==0 ){
                res.end('{"err":"抱歉，修改失败"}');
            }else{
                res.end('{"success":"修改成功"}');
            }
        });
});
//logout
router.post('/logout', function (req, res, next) {
    //req.session.username = ""; //清除session
    //req.session.password = "";
    res.end('{"success":"true"}');
});
//alterpass   这个接口出现两个错误，已修正 1.crypto多次加密 2.返回data类型RowDataPacket
router.post('/alterpass', function(req, res, next) {
    //console.log(req.body)
    var username = jwt.decode(req.body.token).username
    // crypto md5多次加密 不能 设置变量 var md5 = crypto.createHash('md5') !!!!错误
    var oldpass = crypto.createHash('md5').update(req.body.oldpass).digest('base64');
    var newpass = crypto.createHash('md5').update(req.body.newpass).digest('base64');
    sql('select Password from users where UserName=?',[username],function(err,data){
        console.log(data[0].Password===oldpass)
        if(data[0].Password===oldpass){
            sql('update users set Password=?  where UserName=?',
                [newpass,username],function (err,data){
                    if( data.length==0 ){
                        res.end('{"err":"抱歉，修改失败"}');
                    }else{
                        res.end('{"success":"修改成功"}');
                    }
                })
        }
    })

});
//adminlist
router.post('/adminList', function (req, res, next) {
    //console.log(req.body);
    //req.route.path = "/page"; //修改path来设定 对 数据库的操作?
    var page = req.body.page || 1;
    var rows = req.body.rows || 5;
    sql('select * from users order by id limit ?,?',[(page-1)*5,rows],function (err, data) {
        sql('select count(id) as count from users',function (err,count) {
            //console.log(data)
            var obj = {
                data:data,
                total:count,
                success:"成功"
            };
            var str = JSON.stringify(obj);
            res.send(str);
        })
    });
});
//res the newest articles

router.post('/article', function (req, res, next) {
    //console.log(req.body);
    //res.header("Access-Control-Allow-Origin", "*");
    //req.route.path = "/page"; //修改path来设定 对 数据库的操作
    var page = req.body.page || 1;
    var rows = req.body.rows;
    sql('select * from article order by id desc limit ?,?',[(page-1)*rows,rows],function (err, data,count) {
        sql('select count(id) as count from users',function (err,count) {
            var obj = {
                data:data,
                total:count,
                success:"成功"
            };
            var str = JSON.stringify(obj);
            res.send(str);
        })
    });
});
// delete_article
router.post('/delete_a', function(req, res, next) {
    sql('delete from article where id=?', [req.body.id],function(err,data){
        if(data.length==0){
            res.end('{"err":"抱歉，删除失败"}');
        }else{
            var obj = {
                success:"删除成功"
            };
            var str = JSON.stringify(obj);
            res.end(str);
        }

    });
});
//res the newest cases
router.post('/case', function (req, res, next) {
    //console.log(req.body);
    //req.route.path = "/page"; //修改path来设定 对 数据库的操作
    var page = req.body.page || 1;
    var rows = req.body.rows;
    sql('select * from cases order by id desc limit ?,?',[(page-1)*rows,rows],function (err, data,count) {
        var obj = {
            data:data,
            total:count,
            success:"成功"
        };
        var str = JSON.stringify(obj);
        res.send(str);
    });
});
//_details_articles
router.get('/articles/:id.html', function (req, res, next) {
    console.log(req.params.id);
    //req.route.path = "/page"; //修改path来设定 对 数据库的操作?
    sql('select * from article where id=?',[req.params.id], (err, data) => {
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
    });
});
//_details_cases
router.get('/cases/:id.html', function (req, res, next) {
    //console.log(req.params.id);
    //req.route.path = "/page"; //修改path来设定 对 数据库的操作?
    sql('select * from cases where id=?',[req.params.id], (err, data) => {
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
    });
});
//essage
router.post('/essay',(req,res) => {
    var title = req.body.title,
        tag='web',
        author='admin',
        content=req.body.content,
        des = '暂无',
        img = ['bg2.jpg','3e.jpg','2r.jpg','1k.jpg','21.jpg','024.jpg','36.jpg','93.jpg','201.jpg','a0.jpg'],
        index = Math.floor(Math.random()*5)+1,
        dir = img[index],
        createtime=new Date().toLocaleString().substring(0,10);//去掉substring可显示具体时间
    //console.log(content);
    sql('insert into article (id,title,tag,author,content,CreateTime,ArticleImg,see,chart,des) values (0,?,?,?,?,?,?,0,0,?)',
        [title,tag,author,content,createtime,dir,des],(err,data)=>{
            if(err){//err是必须的
                res.status(404).end(`警告，来不及说再见，部分内容已消失在外太空...

                                     也许多点击几次    发布   按钮，会发生奇迹...嗯，千真万确！

                `);
                return;
            }
            res.json({
                result:'success!'
            })
        });
})
//ueditor
router.use("/ueditor/ue", ueditor(path.join(process.cwd(), 'public'), function (req, res, next) {
    //客户端上传文件设置
    let imgDir = '/img/ueditor/';//  /pic/img/...后面也是
    let ActionType = req.query.action;
    if (ActionType === 'uploadimage' || ActionType === 'uploadfile' || ActionType === 'uploadvideo') {
        let file_url = imgDir;//默认图片上传地址
        /*其他上传格式的地址*/
        if (ActionType === 'uploadfile') {
            file_url = '/file/ueditor/'; //附件
        }
        if (ActionType === 'uploadvideo') {
            file_url = '/video/ueditor/'; //视频
        }
        res.ue_up(file_url); //你只要输入要保存的地址 。保存操作交给ueditor来做
        res.setHeader('Content-Type', 'text/html');
    }
    //  客户端发起图片列表请求
    else if(req.query.action === 'listimage'){
        let dir_url = '/img/ueditor/';
        res.ue_list(dir_url); // 客户端会列出 dir_url 目录下的所有图片
    }else if(req.query.action === 'listfile'){
        let dir_url = '/file/ueditor/';
        res.ue_list(dir_url); // 客户端会列出 dir_url 目录下的所有图片
    }
    // 客户端发起其它请求
    else {
        // console.log('config.json')
        res.setHeader('Content-Type', 'application/json');
        res.redirect('/ueditor/nodejs/config.json');
    }
}));

module.exports = router;

