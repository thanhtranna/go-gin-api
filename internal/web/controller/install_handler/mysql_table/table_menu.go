package mysql_table

//CREATE TABLE `menu` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//`pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'parent id',
//`name` varchar(32) NOT NULL DEFAULT '' COMMENT 'menu name',
//`link` varchar(100) NOT NULL DEFAULT '' COMMENT 'link address',
//`icon` varchar(60) NOT NULL DEFAULT '' COMMENT 'icon',
//`level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'menu type 1: first level menu 2: second level menu',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'enable 1: yes -1: no',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='left menu bar table';

func CreateMenuTableSql() (sql string) {
	sql = "CREATE TABLE `menu` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',"
	sql += "`pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'parent id',"
	sql += "`name` varchar(32) NOT NULL DEFAULT '' COMMENT 'menu name',"
	sql += "`link` varchar(100) NOT NULL DEFAULT '' COMMENT 'link address',"
	sql += "`icon` varchar(60) NOT NULL DEFAULT '' COMMENT 'icon',"
	sql += "`level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'menu type 1: first level menu 2: second level menu',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'enable 1: yes -1: no',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT 'delete 1: yes -1: no',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'founder',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT 'updater',"
	sql += "PRIMARY KEY (`id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='left menu bar table';"

	return
}
