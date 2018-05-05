create table `yuge`.`book1titles`
(
  `id` int unsigned not null auto_increment,
  author varchar(100) not null,
  novelname varchar(100) not null,
  title varchar(100) not null,
  url varchar(100) not null,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB CHARSET=utf8;



alter table book1titles AUTO_INCREMENT=1; #恢复主键自增



create table `yuge`.`map`  # sqlserver 中带有空格的列名 [stud ents]
	(
		`id` int unsigned not null auto_increment,
		tablename varchar(100) not null,      -- 注释
		novelname varchar(100) not null,    # 注释

    PRIMARY KEY (`id`)
	)
ENGINE = InnoDB CHARSET=utf8;