# 说明
>  ygser 接口文档: [待配置](https://github.com/yugef3h/blogserver)

>  如果对您对此项目有兴趣，可以点 "Star" 支持一下 谢谢！ ^_^

>  apidoc 生成 api 文档
```
1、npm install apidoc -g
2、在 package.json 中配置 apidoc
3、app.js 中设置生成文档的路径，若我放在 public/apidoc 中，则 app.use('/public',express.static('public'));
4、public 文件夹下 创建 apidoc
5、根目录下运行 apidoc -i routes -o public/apidoc，routes 为接口注释文件
```
>  如有问题请直接在 Issues 中提，或者您发现问题并有非常好的解决方案，欢迎 PR 👍

>  相关项目地址：[前端项目地址](http://www.blackatall.cn)  
