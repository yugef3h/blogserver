const mysql = require('mysql');
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
    });
};

