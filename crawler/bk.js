/**
 * Created by 12 on 2017/7/4.
 */
const fs = require('fs')
const cheerio = require('cheerio')
const mysql = require('mysql')
const eventproxy = require('eventproxy')
const express = require('express')
const app = express()
const superagent = require('superagent')
require('superagent-charset')(superagent)
const async = require('async');



function trim(str) {
  return str.replace(/(^\s*)|(\s*$)/g, '').replace(/&nbsp;/g, '')
}

//将Unicode转汉字
function reconvert(str) {
  str = str.replace(/(&#x)(\w{1,4});/gi, function ($0) {
    return String.fromCharCode(parseInt(escape($0).replace(/(%26%23x)(\w{1,4})(%3B)/g, "$2"), 16));
  });
  return str
}


function fetUrl(url, callback) {
  superagent.get(url)
    .charset('gbk')  //该网站编码为gbk，用到了superagent-charset
    .end(function (err, res) {
      let $ = cheerio.load(res.text)
      const arr = []
      const content = null
      console.log($("#content"))
      if ($("#content")[0]) {
        const content = reconvert($("#content").html())
        //分析结构后分割html
        const contentArr = content.split('<br><br>')
        contentArr.forEach(elem => {
          const data = trim(elem.toString())
          arr.push(data)
        })
        const obj = {
          previous: 'https://www.zwdu.com' + $('.bottem2 a').eq(0).attr('href'),
          next: 'https://www.zwdu.com' + $('.bottem2 a').eq(2).attr('href'),
          title: $('.bookname h1').text(),
          content: arr.join('-').slice(0, 20000)  //由于需要保存至mysql中，不支持直接保存数组，所以将数组拼接成字符串，取出时再分割字符串即可,mysql中varchar最大长度，可改为text类型
        }
        //console.log(obj)
        callback(obj)  //将obj传递给第四个参数中的results
      } else {
        callback({'err':'章节封顶'})
      }


    })
}

//fetUrl('https://www.zwdu.com/book/29944/10900483.html')
module.exports = fetUrl;

