var cheerio = require('cheerio');
var request = require('request');
//var iconv = require('iconv-lite');
const express = require('express')
const app = express()
const superagent = require('superagent')
require('superagent-charset')(superagent)
const async = require('async');
function getHtmlByUrl(href,callback) {
    request(href.replace(/([\u4e00-\u9fa5])/g, (str) => encodeURIComponent(str) ), function(err, response, body) {
        if (!err && response.statusCode == 200) {
            //body即为目标页面的html
            var $ = cheerio.load(body);
            var v = $('.result-game-item-title-link').attr('href');
            superagent.get(v)
                .charset('gbk')
                .end(function (err, res) {
                    const $ = cheerio.load(res.text,{decodeEntities: false});
                    let content = [],id = 0
                    $('#list dd').each((i,v) => {
                        var $v = $(v)
                        content.push($v.find('a').text() + '+' + 'https://www.zwdu.com'+ $v.find('a').attr('href'))
                        ++ id
                    })
                    let obj = {
                        id: id,
                        name: $('#info h1').text(),
                        content: content.join('-'),
                        //urls: urls.join('-')
                    }
                    callback(obj)
                })
        } else {
            console.log('get page error url => ' + href);
        }
    });
}
//var url = 'https://www.zwdu.com/search.php?keyword=强人';
//URL编码，只转换中文
//getHtmlByUrl( url);
module.exports = getHtmlByUrl;
//可参考同目录下其他文件