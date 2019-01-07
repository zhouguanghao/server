package redis

import (
	"github.com/davyxu/golog"
	"github.com/go-redis/redis"
	"server/config"
	"time"
)

var (
	Redisdb *redis.Client
	log = golog.New("redis")
)

type Person struct {
	Name string
	Age int
}

func SetValue(key, value string) error {
	Redisdb.Pipeline().Select(5)
	err := Redisdb.Set(key, value, 0).Err()
	if err != nil {
		log.Debugln("Redis set failed")
		return err
	}
	return nil
}

func GetValue(key string) (string, error) {
	v, err := Redisdb.Get(key).Result()
	return v, err
}

func init() {
	log.Debugln("Redis init.")
	cfg := new(config.Config)
	cfg.InitConfig("server.config")
	redis_addr := cfg.Read("redis", "address")
	redis_passwd := cfg.Read("redis", "password")
	Redisdb = redis.NewClient(&redis.Options{
		Addr: redis_addr,
		Password: redis_passwd,
		DialTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize: 10,
		PoolTimeout: 30 * time.Second,
		//DB: 5,
	})
	//defer Redisdb.Close()

	_,err := Redisdb.Ping().Result()
	if err != nil {
		panic(err)
	}

	//person := &Person{Name: "zhou", Age:18}
	//log.Debugln(person)
	//b,err := json.Marshal(person)
	//log.Debugln(b, err, string(b))
	//err = Redisdb.Set("person", string(b), 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//v,err := Redisdb.Get("person").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(v)
	//pe1 := new(Person)
	//err = json.Unmarshal([]byte(v), pe1)
	//fmt.Println(*pe1, pe1.Name, pe1.Age)

}

