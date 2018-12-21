package mysql

type Player struct {
	id	 int32
	account	 string
	password string
	state	 int32
}

var PlayerMap = map[string]*Player{}

func (p *Player)CreatePlayer() {
	_, err := DB.Exec("insert into player (account, password, state) values (?, ?, ?)", p.account, p.password, p.state)
	if err != nil {
		log.Debugln("insert into data err", err)
		return
	}
}

func (p *Player)GetPlayerData(name string)  {
	DB.QueryRow("select * from player where name = ?", name).Scan(&p.id, &p.account, &p.password, &p.state)
	if _,ok := PlayerMap[p.account]; ok {
		return
	} else {
		PlayerMap[p.account] = p
	}
	return
}

func GetAllPlayerInfo() {

	rows,err := DB.Query("select * from player")
	defer rows.Close()
	if err != nil {
		log.Debugln("Get all player info err:", err)
		return
	}

	for rows.Next() {
		p := new(Player)
		err := rows.Scan(&p.id, &p.account, &p.password, &p.state)
		if err != nil {
			log.Debugln("Scan all rows err", err)
			return
		}
		log.Debugln(p)
		PlayerMap[p.account] = p
	}
	log.Debugln(PlayerMap)
}
