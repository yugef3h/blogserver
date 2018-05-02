let cheerio = require('cheerio');
let request = require('request');
const express = require('express')
const app = express()
const superagent = require('superagent')
require('superagent-charset')(superagent)
function getHtmlByUrl(href, callback) {
  let reg = /keyword/;
  if (reg.test(href)) {
    request(href.replace(/([\u4e00-\u9fa5])/g, (str) => encodeURIComponent(str)), function (err, response, body) {
      if (!err && response.statusCode == 200) {
        //body即为目标页面的html
        let $ = cheerio.load(body);
        let content = [];
        $('.result-game-item-title-link').each((i, v) => {
          let $v = $(v)
          content.push($v.attr('title') + '+' + $v.attr('href'))
        })
        let obj = {
          meny: content.join('-')
        }
        //console.log(obj);
        callback(obj)
      } else {
        console.log('get page error url => ' + href);
      }
    });
  }
  else {
    superagent.get(href)
      .charset('gbk')
      .end(function (err, res) {
        const $ = cheerio.load(res.text, {decodeEntities: false});
        let content = [];
        $('#list dd').each((i, v) => {
          let $v = $(v)
          content.push($v.find('a').text() + '+' + $v.find('a').attr('href'))
        })
        let obj = {
          author: $('#info p').first().text(),
          name: $('#info h1').text(),
          content: content.join('-'),
          //urls: urls.join('-')
        }
        //console.log(obj)
        callback(obj)
      })
  }
}

//let url = 'https://www.zwdu.com/search.php?keyword=强人';
//URL编码，只转换中文
//getHtmlByUrl( url);
module.exports = getHtmlByUrl;

//可参考同目录下其他文件