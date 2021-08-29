SELECT  usr.name as usr,parent.name as parent
	FROM public.user as usr left join
	public.user as parent on usr.parent=parent.id order by usr.name asc