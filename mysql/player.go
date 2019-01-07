package mysql

import (
	"encoding/json"
	"server/redis"
)

type Player struct {
	Id	 int32
	Account	 string
	Password string
	State	 int32
}

func (p *Player)CreatePlayer() {
	_, err := DB.Exec("insert into player (account, password, state) values (?, ?, ?)", p.Account, p.Password, p.State)
	if err != nil {
		log.Debugln("insert into data err", err)
	}
}

func (p *Player)GetPlayerData(name string)  {
	DB.QueryRow("select * from player where name = ?", name).Scan(&p.Id, &p.Account, &p.Password, &p.State)
	p.SetPlayer()
}

func (p *Player)UpdatePlayerPassword() {
	_, err := DB.Exec("undate player set password = ? where account = ?", p.Password, p.Account)
	if err != nil {
		log.Errorf("undate player: %s password fail", p.Account, err)
	}
}

func (p *Player)UpdatePlayerStat() {
	_, err := DB.Exec("update player set state = ? where account = ?", p.State, p.Account)
	if err != nil {
		log.Errorf("update player: %s state fail", p.Account, err)
	}
}

func (p *Player)SetPlayer() {
	v,_ := json.Marshal(p)
	redis.Redisdb.Set(p.Account, string(v), 0)
}

func (p *Player)GetPlayer() bool {
	v,err := redis.Redisdb.Get(p.Account).Result()
	if err == nil {
		json.Unmarshal([]byte(v), p)
		return true
	}

	return false
}

func GetAllPlayerInfo() {
	rows, err := DB.Query("select * from player")
	defer rows.Close()
	if err != nil {
		log.Debugln("Get all player info err:", err)
		panic(err)
	}

	for rows.Next() {
		p := new(Player)
		err := rows.Scan(&p.Id, &p.Account, &p.Password, &p.State)
		if err != nil {
			log.Debugln("Scan row err", err)
			panic(err)
		}

		p.SetPlayer()
	}
}
