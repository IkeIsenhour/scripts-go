package scripts

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func main() {

// 	url := "https://www.instagram.com/api/v1/friendships/1528663181/following/?count=200"
// 	xIgAppId := "936619743392459"
// 	cookie := "csrftoken=VOwQlVtbp2PdBfNkNlfRRsWEUjX52fcg; ps_l=1; ps_n=1; wd=1728x620; mid=ZspplwAEAAHKT9VpKhlT0sFWSL37; datr=l2nKZjqLpVLozJTMXM-LJrxx; ig_did=4153B836-A31C-404F-B949-5508EC7B4B6A; ig_nrcb=1; rur=\"VLL\0541528663181\0541756078806:01f7efc08fbd2332a766ac5c3e057407931ec5064b64f562a242513a7490b33eaccb1b63\"; ds_user_id=1528663181; sessionid=1528663181%3AmOqkb7EzbK2PZe%3A12%3AAYcCQvI_tEo8ZCPaF98jtLAQwGGrRmvrTqFGjx8MhA; shbid=\"11976\0541528663181\0541756077379:01f7ab0fb847d5cc639a0dbbef97b14fa2cf7f7531b6be5cdba3508c0111f7a5d77eff91\"; shbts=\"1724541379\0541528663181\0541756077379:01f720c0d4b3387f9279c5f851036b9b514070774aaacfa507ea1186e3418d9e9c8ab9c5\""

// 	c := &http.Client{
// 		Timeout: 30 * time.Second,
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		fmt.Printf("error formatting request: %s", err)
// 		os.Exit(1)
// 	}

// 	req.Header.Add("X-IG-App-ID", xIgAppId)
// 	req.Header.Add("Cookie", cookie)

// 	res, err := c.Do(req)
// 	if err != nil {
// 		fmt.Printf("error making request: %s", err)
// 		os.Exit(1)
// 	}

// 	bytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("error reading response body: %s", err)
// 	}

// 	fmt.Println(res.StatusCode)
// 	fmt.Println(string(bytes))
// }
