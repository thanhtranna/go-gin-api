package mysql_table

//CREATE TABLE `authorized` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//`business_key` varchar(32) NOT NULL DEFAULT '' COMMENT 'caller key',
//`business_secret` varchar(60) NOT NULL DEFAULT '' COMMENT 'caller secret',
//`business_developer` varchar(60) NOT NULL DEFAULT '' COMMENT 'caller developer',
//`remark` varchar(255) NOT NULL DEFAULT '' COMMENT 'remarks',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'enable 1: yes -1: no',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',
//PRIMARY KEY (`id`),
//UNIQUE KEY `unique_business_key` (`business_key`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='authorized caller table';

func CreateAuthorizedTableSql() (sql string) {
	sql = "CREATE TABLE `authorized` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',"
	sql += "`business_key` varchar(32) NOT NULL DEFAULT '' COMMENT 'caller key',"
	sql += "`business_secret` varchar(60) NOT NULL DEFAULT '' COMMENT 'caller secret',"
	sql += "`business_developer` varchar(60) NOT NULL DEFAULT '' COMMENT 'caller developer',"
	sql += "`remark` varchar(255) NOT NULL DEFAULT '' COMMENT 'remarks',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'enable 1: yes -1: no',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',"
	sql += "PRIMARY KEY (`id`),"
	sql += "UNIQUE KEY `unique_business_key` (`business_key`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='caller table';"

	return
}

func CreateAuthorizedTableDataSql() (sql string) {
	sql = "INSERT INTO `authorized` (`id`, `business_key`, `business_secret`, `business_developer`, `remark`, `created_user`) VALUES (1, 'admin', '12878dd962115106db6d', 'administrator', 'Management panel', 'init');"

	return
}
