var cheerio = require('cheerio');
var request = require('request');
//var iconv = require('iconv-lite');
function getHtmlByUrl(href,callback) {
    request(href.replace(/([\u4e00-\u9fa5])/g, (str) => encodeURIComponent(str) ), function(err, response, body) {
        if (!err && response.statusCode == 200) {
            //body即为目标页面的html
            var $ = cheerio.load(body);
            var v = $('.result-game-item-title-link').attr('href');
            //console.log(v);
            request(v, function(err,response, body) {
                //var buf =  iconv.decode(body, 'gbk');
                if (!err && response.statusCode == 200) {
                    var $ = cheerio.load(body,{decodeEntities: false});
                    var urls = [];
                    var a = 1;
                    $('dd a').each(function (idx, element) {
                        var $element = $(element);
                        urls.push({
                            url: 'https://www.zwdu.com'+ $element.attr('href'),
                            title: '第 '+ a +' 章',
                            //title: $element.text()
                        })
                        a++;
                    })
                    callback(urls);
                }
            })
        } else {
            console.log('get page error url => ' + href);
        }
    });
}
//https://www.zwdu.com/search.php?keyword=择天记     ->    https://www.zwdu.com/search.php?keyword=%E6%8B%A9%E5%A4%A9%E8%AE%B0
//var url = 'https://www.zwdu.com/search.php?keyword=强人';
//URL编码，只转换中文
//getHtmlByUrl( url);
module.exports = getHtmlByUrl;
//可参考同目录下其他文件