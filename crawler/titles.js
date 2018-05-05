let cheerio = require('cheerio');
let request = require('request');
const express = require('express')
const async = require('async');
const app = express()
const superagent = require('superagent')
require('superagent-charset')(superagent)
//同名书籍
function getUrls(href, callback) {
  request(href.replace(/([\u4e00-\u9fa5])/g, (str) => encodeURIComponent(str)), function (err, response, body) {
    if (!err && response.statusCode == 200) {
      let $ = cheerio.load(body);

      let urls = [];
      $('.result-game-item-title-link').each((i, v) => {
        let $v = $(v)
        urls.push($v.attr('title') + '+' + $v.attr('href'))
      })
      let obj = {
        section: urls.join('-')
      }
      callback(obj)
    } else {
      console.log('get page error url => ' + href);
    }
  });
}
//保存章节目录
function getSections(href, callback) {

    superagent.get(href)
      .charset('gbk')
      .end(function (err, res) {
        const $ = cheerio.load(res.text, {decodeEntities: false});

        let urls = [];
        $('#list dd').each((i, v) => {
          let $v = $(v)
          urls.push($v.find('a').text() + '+' + $v.find('a').attr('href'))
        })

        async.mapLimit(urls, 5, function (url, cb) {
          let obj = {
            author: $('#info p').first().text(),
            novelname: $('#info h1').text(),
            section: url,
          }
          cb(null, obj)
        }, function (err, results) {
          callback(results)
        })

      })
}

//let url = 'https://www.zwdu.com/search.php?keyword=强人';
//URL编码，只转换中文
//getHtmlByUrl( url);
module.exports = {getUrls, getSections};

