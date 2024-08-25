package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	insta "github.com/IkeIsenhour/scripts-go/internal/instagram-follows"
)

func main() {
	start := time.Now()

	c := &http.Client{
		Timeout: 30 * time.Second,
	}
	accountId := "1528663181"
	xIgAppId := "936619743392459"
	cookie := "csrftoken=uyRb3C7e41pXkGYleVVVjhi88LbzgOSE; ps_l=1; ps_n=1; wd=1728x620; mid=ZspplwAEAAHKT9VpKhlT0sFWSL37; datr=l2nKZjqLpVLozJTMXM-LJrxx; ig_did=4153B836-A31C-404F-B949-5508EC7B4B6A; ig_nrcb=1; ds_user_id=1528663181; shbid=\"11976\0541528663181\0541756077379:01f7ab0fb847d5cc639a0dbbef97b14fa2cf7f7531b6be5cdba3508c0111f7a5d77eff91\"; shbts=\"1724541379\0541528663181\0541756077379:01f720c0d4b3387f9279c5f851036b9b514070774aaacfa507ea1186e3418d9e9c8ab9c5\"; rur=\"CCO\0541528663181\0541756149086:01f7f1a447c586ae2f582605c92597e0ac1730f34f124eb0371e911f7e6480aac9c6867c\"; sessionid=1528663181%3AqswXDhsN5wi6xQ%3A8%3AAYekA2Mei16d7Fz4RpH9DAWiqpE53WTEud7s9p8rgg"

	ir := insta.NewInstagramRequest(c, accountId, xIgAppId, cookie)

	l, err := insta.CompareFollowersAgainstFollowing(ir)
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for _, v := range l {
		fmt.Printf("%v. %s - %s\n", i, v.Username, v.FullName)
		i++
	}

	fmt.Println(time.Since(start))
	//11.49698325s
	//8.951331417s

}
