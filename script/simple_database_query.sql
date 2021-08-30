SELECT  usr.user_name as usr,parent.user_name as parent
	FROM user as usr left join
	user as parent on usr.parent=parent.id order by usr.user_name asc