package mysql_table

//CREATE TABLE `admin` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//`username` varchar(32) NOT NULL DEFAULT '' COMMENT 'username',
//`password` varchar(100) NOT NULL DEFAULT '' COMMENT 'password',
//`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT 'nickname',
//`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT 'phone number',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'is used 1: yes -1: no',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',
//PRIMARY KEY (`id`),
//UNIQUE KEY `unique_username` (`username`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='administrator table';

func CreateAdminTableSql() (sql string) {
	sql = "CREATE TABLE `admin` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',"
	sql += "`username` varchar(32) NOT NULL DEFAULT '' COMMENT 'username',"
	sql += "`password` varchar(100) NOT NULL DEFAULT '' COMMENT 'password',"
	sql += "`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT 'nickname',"
	sql += "`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT 'phone number',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'is used 1: yes -1: no',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',"
	sql += "PRIMARY KEY (`id`),"
	sql += "UNIQUE KEY `unique_username` (`username`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='administrator table';"

	return
}

func CreateAdminTableDataSql() (sql string) {
	sql = "INSERT INTO `admin` (`id`, `username`, `password`, `nickname`, `mobile`, `created_user`) VALUES"
	sql += "(1, 'admin', 'f78382de80cf583cf854bbac0b6e796fbde36fe2739ca4ae072637010f179cb0', 'administrator', '13888888888', 'init');"

	return
}
