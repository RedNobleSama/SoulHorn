/**
* @Author: oreki
* @Date: 2021/6/7 23:39
* @Email: a912550157@gmail.com
 */

package gredis

import (
	"SoulHorn/utils"
	"context"
	"encoding/json"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"net"
	"time"
)

//var RedisClient *redis.Client

// InitRedis 初始化redis
func InitRedis(dbNum int) *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		//连接信息
		Network:  "tcp",               //网络类型, tcp 或者 unix, 默认tcp
		Addr:     utils.RedisHost,     //ip:port
		Username: utils.RedisUsername, //用户名, 使用指定用户名验证当前连接
		Password: utils.RedisPassword, //密码,
		DB:       dbNum,               //连接后选中的redis数据库index

		//命令执行失败时的重试策略
		MaxRetries:      3,                      //命令执行失败时最大重试次数，默认3次重试。
		MinRetryBackoff: 8 * time.Millisecond,   //每次重试最小间隔时间，默认8ms，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次重试最大时间间隔，默认512ms，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒
		ReadTimeout:  3 * time.Second, //读超时，默认3秒，-1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认与读超时相等

		//连接池容量、闲置连接数量、闲置连接检查
		PoolSize:           16,                               //连接池最大Socket连接数，默认为10倍CPU数量，10 * runtime.NumCPU()
		MinIdleConns:       8,                                //启动阶段创建指定数量的Idle连接，并长期维持Idle状态的连接数不少于指定数量。
		MaxConnAge:         0 * time.Second,                  //连接存活时长，超过指定时长则关闭连接。默认为0，不关闭旧连接。
		PoolTimeout:        4 * time.Second,                  //当所有连接都处于繁忙状态时，客户端等待可用连接的最大等待时长。默认为读超时+1秒
		IdleTimeout:        time.Duration(utils.IdleTimeout), //关闭闲置连接时间，默认5分钟，-1表示取消闲置超时检查
		IdleCheckFrequency: 1 * time.Minute,                  //闲置连接检查周期，默认为1分钟；-1表示不做检查，只在客户端获取连接时对闲置连接进行处理。

		//自定义连接函数
		Dialer: func(ctx context.Context,
			network string,
			addr string,
		) (net.Conn, error) {
			netDialer := net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 60 * time.Minute,
			}
			return netDialer.Dial(network, addr)
		},

		//钩子函数，建立新连接时调用
		OnConnect: func(ctx context.Context,
			cn *redis.Conn,
		) error {
			fmt.Println("Redis Conn =", cn)
			return nil
		},
	})

	//ctx := context.Background()
	//pong, err := RedisClient.Ping(ctx).Result()
	//fmt.Println("Redis Ping =", pong, err)
	return RedisClient
}

// Cache 缓存数据到db0
func Cache(data interface{}, dbNum int) {
	var ctx = context.Background()
	r := InitRedis(dbNum)
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = r.Set(ctx, "moose-go", marshal, 10*time.Minute).Result()
	if err != nil {
		fmt.Println("失败")
	}
	fmt.Println("成功")
}

//
//// GetCache 从缓存中获取数据
//func GetCache(data interface{}) ( ,bool) {
//	var ctx = context.Background()
//	marshal, err := RedisClient.Get(ctx, "moose-go").Result()
//	if err != nil {
//		fmt.Println("获取结果失败，无结果")
//		return false
//	}
//	err = json.Unmarshal([]byte(marshal), &data)
//	if err != nil {
//		fmt.Println("转义失败")
//	}
//	return data, true
//}
