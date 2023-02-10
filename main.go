package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 1})
	res, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("ping 出错: ", err)
	}
	fmt.Println(res) // pong

	// set
	res, err = rdb.Set("name", "wangbucuo", 3*time.Second).Result()
	if err != nil {
		fmt.Println("设置数据失败, err:", err)
	}
	fmt.Println(res)

	// get
	res, err = rdb.Get("name").Result()
	if err != nil {
		fmt.Println("获取数据失败, err:", err)
	}
	fmt.Println(res)

	//SetNX:key不存在时才设置（新增操作）
	rdb.SetNX("name", 19, 0) // name存在，不会修改

	//SetXX:key存在时才设置（修改操作）
	rdb.SetXX("name", "pyy", 0)
	rdb.SetXX("hobby", "football", 0) // 不会新增成功

	// strLen
	l, _ := rdb.StrLen("name").Result()
	fmt.Println(l)

	// 哨兵
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{MasterName: "Master", SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"}})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接出错: ", err)
		return
	}

	// 集群
	rdb2 := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"}})
	_, err = rdb2.Ping().Result()
	if err != nil {
		fmt.Println("连接集群出错， err: ", err)
		return
	}
}
