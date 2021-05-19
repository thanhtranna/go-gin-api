package mysql_table

//CREATE TABLE `authorized_api` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//`business_key` varchar(30) NOT NULL DEFAULT '' COMMENT 'caller key',
//`method` varchar(30) NOT NULL DEFAULT '' COMMENT 'request method',
//`api` varchar(100) NOT NULL DEFAULT '' COMMENT 'request address',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='authorized interface address table';

func CreateAuthorizedAPITableSql() (sql string) {
	sql = "CREATE TABLE `authorized_api` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',"
	sql += "`business_key` varchar(30) NOT NULL DEFAULT '' COMMENT 'caller key',"
	sql += "`method` varchar(30) NOT NULL DEFAULT '' COMMENT 'request method',"
	sql += "`api` varchar(100) NOT NULL DEFAULT '' COMMENT 'request address',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',"
	sql += "PRIMARY KEY (`id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='authorized interface address table';"

	return
}

func CreateAuthorizedAPITableDataSql() (sql string) {
	sql = "INSERT INTO `authorized_api` (`id`, `business_key`, `method`, `api`,`created_user`) VALUES"
	sql += "(1, 'admin', 'GET', '/api/**', 'init'),"
	sql += "(2, 'admin', 'POST', '/api/**', 'init'),"
	sql += "(3, 'admin', 'PUT', '/api/**', 'init'),"
	sql += "(4, 'admin', 'DELETE', '/api/**', 'init'),"
	sql += "(5, 'admin', 'PATCH', '/api/**', 'init');"

	return
}
