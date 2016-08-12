package main

import (
	"net/http"
	"fmt"
	"io"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/redis.v4"
)

var cacheWriter *CacheWriter
var usingCache bool

type CacheWriter struct {
	ttl time.Duration 
	defaultWriter http.ResponseWriter 
	cacheClient *redis.Client
	reader *http.Request
}

func (c *CacheWriter) Write(b []byte) (int, error){
	url := c.reader.URL.String()
	if c.cacheClient != nil {
		fmt.Printf("Setting cache \n")	
		c.cacheClient.Set(url, b, c.ttl)
	}
	c.defaultWriter.Write(b)

	return 0, nil
}

func DecorateCacheWritter(c *viper.Viper, w http.ResponseWriter, r *http.Request) (io.Writer) {
	cacheEnabled := c.GetBool("enabled")
	if cacheEnabled && cacheWriter == nil {
		address := c.GetString("address")
		port := c.GetString("port")
		password := c.GetString("password")
		hostname := fmt.Sprintf("%s:%s", address, port)
		ttl := time.Duration(c.GetInt("ttl"))

		redisClient := redis.NewClient(&redis.Options{
	        Addr: hostname,
	        Password: password,
	    }) 

	    cacheWriter = &CacheWriter{
			ttl : ttl,
			defaultWriter: w,
			cacheClient: redisClient,
		}
	}

	if cacheEnabled { 
		cacheWriter.reader = r
		return cacheWriter	
	}

	return w
}

func WriteFromCache(c *viper.Viper, w http.ResponseWriter, r *http.Request) bool {
	url := r.URL.String()
	fmt.Printf(url)
	chartInCache, err := cacheWriter.cacheClient.Get(url).Result()
	fmt.Printf("%s", chartInCache)
	if err == nil {
		fmt.Printf("Getting from cache\n")
		w.Write([]byte(chartInCache))

		return true
	}

	return false;
}