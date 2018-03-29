const mysql = require('mysql');
// 连接池有bug 暂时不用
/*var pool = mysql.createPool({
    host:'localhost',
    user:'root',
    password:'',
    database:'yuge',
    port:'3306'
});*/
/*
module.exports=function(sql,arr,callback){
    //start connect
    pool.getConnection(function(err,conn){
        if(err){
            console.log('[query] - :'+err);
            return;
        }else{
            console.log('[--------------------------dbConnected----------------------------]')
            conn.query(sql,arr,(err,data) => {
                callback && callback(err,data);
                conn.release();
            });

        }

    });
};*/

module.exports=function(sql,arr,callback){
    let conn = mysql.createConnection({
        host:'localhost',
        user:'root',
        password:'',
        port:'3306',
        database:'yuge'
    });
    //start connect
    conn.connect(function(err){
        if(err){
            console.log('[query] - :'+err);
            return;
        }
        //console.log('[dbConnected!]')
    });
    //执行sql语句
    conn.query(sql,arr,(err,data) => {
        callback && callback(err,data);
    });
    //end connect
    conn.end((err) => {
        if(err){
            return;
        }
        //console.log('[dbGoodbye!]');
    });
};

