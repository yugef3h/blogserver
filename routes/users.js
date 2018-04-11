var express = require('express');
var router = express.Router();
var crypto = require('crypto');
var jwt = require("jsonwebtoken");
var path=require('path');
var ueditor=require('ueditor');
var db = require('../sql/db');
var ztj = require('../crawler/ztj');

router.use('/public', express.static(__dirname + '/public'));
/**
 * 获得某个用户
 * @api {POST} /users/login
 * @apiDescription 登录验证
 * @apiName login
 * @apiParam (path参数) {String} username
 * @apiParam (path参数) {String} password
 * @apiSampleRequest /users/login
 * @apiGroup User
 * @apiVersion 1.0.0
 */
router.post('/login', function (req, res, next) {// router.all
    db.login(req, res, next);
});
// delete_users
router.post('/delete_u', function(req, res, next) {
    db.delUsers(req, res, next);
});
// add_users
router.post('/add_u', function(req, res, next) {
    db.addUsers(req, res, next);
});
// update_users
router.post('/update_u', function(req, res, next) {
    db.putUsers(req, res, next);
});
//logout
router.post('/logout', function (req, res, next) {
    //req.session.username = ""; //清除session
    //req.session.password = "";
    res.end('{"success":"true"}');
});
//putPass   这个接口出现两个错误，已修正 1.crypto多次加密 2.返回data类型RowDataPacket
router.post('/alterpass', function(req, res, next) {
    db.putPass(req, res, next);
});
//adminlist
router.post('/adminList', function (req, res, next) {
    db.admin(req, res, next);
});
//the newest articles
router.post('/article', function (req, res, next) {
    db.article(req, res ,next);
});
// delete_article
router.post('/delete_a', function(req, res, next) {
    db.delArticle(req, res, next)
});
//res the newest cases
router.post('/case', function (req, res, next) {
    db.case(req, res, next);
});
//_details_articles
router.get('/articles/:id.html', function (req, res, next) {
    db.artId(req, res, next);
});
//_details_cases
router.get('/cases/:id.html', function (req, res, next) {
    db.caseId(req, res, next);
});
//essage
router.post('/essay',function (req,res,next)  {
    db.essay(req,res,next);
})
//novel
router.post('/novel',function(req,res,next) {
    db.novelKey(req,res,next);
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

