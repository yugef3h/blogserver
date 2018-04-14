var express = require('express');
var path = require('path');
//var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');

var index = require('./routes/index');
var users = require('./routes/users');
//var session=require('express-session');

var morgan = require("morgan");
var app = express();

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(morgan("dev"));
//跨域1  后期删
//缺少options方法
//改动了headers!!!!未测试
//另外如何避免预检请求的发生  判断  options
app.all('*', function(req, res, next) {
    //res.header("Access-Control-Allow-Origin", '*');
    res.header("Access-Control-Allow-Origin", "http://localhost:8080"); //http://localhost:8080为了跨域保持session，所以指定地址，生产环境改*
    res.header('Access-Control-Allow-Methods', 'PUT, GET, POST, DELETE, OPTIONS');
    res.header("Access-Control-Allow-Headers", "X-Requested-With,Content-Type, Authorization,Content-Length, ");
    res.header('Access-Control-Allow-Credentials', true);//可以带cookies
    res.header('Access-Control-Max-Age','1728000');// 预检请求有效期20天
    res.header("X-Powered-By", '3.2.1')
    if(req.method == "OPTIONS") {
        res.sendStatus(200);/*让options请求快速返回或者204无缓存*/
    } else {
        next();
    }
});
/*where use*/
/*app.use(session({
  secret:'yugeblog930114',
  name:'ygblog',
  cookie:{maxAge:60*1000*60*24},
  resave:false,             // 每次请求都重新设置session
  saveUninitialized:true
}));*/

// 验证用户登录
/*app.use(function(req, res, next){

  //后台请求
  if(req.session.username){ //表示已经登录后台
    next();
  }else if( req.url.indexOf("login") >=0 || req.url.indexOf("logout") >= 0){
    //登入，登出不需要登录
    next();
  }else{
    //next(); //TODO:这里是调试的时候打开的，以后需要删掉
    res.end('{"redirect":"true"}');

  };

});*/
// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
//
app.use('/public',express.static('public'));
// express展示
app.use('/', index);
app.use('/users', users);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
