package main

import (
	"net/http"
	"log"
	"fmt"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	//"regexp"
	"strings"
)

type BasketballMatchInfo struct {
	matchId string
	matchVS string
	matchTime string
	matchWeekTime string
	leagueName string
	homeTeamName string
	guestTeamName string
	homeWin string
	guestWin string
	homeConcedeWin string
	concedeScore string
	guestConcedeWin string
	bigScore string
	finalScore string
	smallScore string

}

type MatchInfo struct {
	matchId string
	matchVS string
	matchTime string
	matchWeekTime string
	leagueName string
	homeTeamName string
	guestTeamName string
	homeWin string
	draw string
	guestWin string
	concede string
	homeConcedeWin string
	drawConcede string
	guestConcedeWin string

}
func Crawler(url *string) {

	client := http.Client{}
	request, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatal("error:", err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("error:", err)
	}
	//content,_ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))
	if response.StatusCode == http.StatusOK {

		token := html.NewTokenizer(response.Body)
		for {
			tn := token.Next()
			if tn == html.ErrorToken {
				return
			}
			name, hasAttr := token.TagName()
			if string(name) == "div" && hasAttr {

				key, val, moreAttr := token.TagAttr()

				if string(key) == "class" && string(val) == "section" {//regexp.MustCompile("https://book.douban.com/subject/([0-9])").MatchString(string(val))
					matchInfo := new(MatchInfo);
					hasmore := moreAttr;
					for hasmore {
						key,value,moreAttr := token.TagAttr();
						if string(key) == "match"{
							//fmt.Println(string(val), "---", string(value))
							matchInfo.matchVS = string(value)
						}
						if string(key) == "match_time"{
							matchInfo.matchTime = string(value)
						}
						if string(key) == "match_week"{
							matchInfo.matchWeekTime = string(value)
						}
						if string(key) == "league_val"{
							matchInfo.leagueName = string(value)
						}
						if string(key) == "match_id"{
							matchInfo.matchId = string(value)
						}

						hasmore = moreAttr;
					}
					fmt.Println(matchInfo)
				}
			}
		}

	}
	defer response.Body.Close()

}

func getValue(s *goquery.Selection)string{
	val,exists:= s.Attr("value")
	if exists{
		text := s.Text();
		return strings.Join([]string{val,text},":")
	}else{
		text := s.Text();
		return text
	}
}
func htmlHunter_football(url *string){
	client := http.Client{}
	request, err := http.NewRequest("GET",*url,nil);
	if err != nil{
		log.Fatal(err)
	}
	response,err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	doc,err := goquery.NewDocumentFromReader(response.Body);
	if err != nil{
		log.Fatal(err);
	}
	doc.Find("div.section").Each(func(i int,s *goquery.Selection){
		matchInfo := new(MatchInfo)
		matchId,_ := s.Attr("match_id")
		matchVs,_ := s.Attr("match")
		matchTime,_ := s.Attr("match_time")
		matchWeek,_ := s.Attr("match_week")
		leagueName,_ := s.Attr("league_val")
		matchInfo.matchId = matchId
		matchInfo.matchVS = matchVs
		matchInfo.matchTime = matchTime
		matchInfo.matchWeekTime =matchWeek
		matchInfo.leagueName = leagueName
		s.Find("table.saishi tr td").Each(func(i int ,s *goquery.Selection){
			switch i {
				case 3:
					homeTeamName := getValue(s)
					matchInfo.homeTeamName = homeTeamName
				case 4:
					guestTeamName := getValue(s)
					matchInfo.guestTeamName = guestTeamName
				case 5:
					matchInfo.homeWin = getValue(s)
				case 6:
					matchInfo.draw = getValue(s)
				case 7:
					matchInfo.guestWin = getValue(s)
				case 8:
					matchInfo.concede = getValue(s)
				case 9:
					matchInfo.homeConcedeWin = getValue(s)
				case 10:
					matchInfo.drawConcede = getValue(s)
				case 11:
					matchInfo.guestConcedeWin = getValue(s)
				}

		})
		fmt.Println(matchInfo)
	});

	defer response.Body.Close();
}
func htmlHunter_basketball(url *string){
	client := http.Client{}
	request, err := http.NewRequest("GET",*url,nil);
	if err != nil{
		log.Fatal(err)
	}
	response,err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	doc,err := goquery.NewDocumentFromReader(response.Body);
	if err != nil{
		log.Fatal(err);
	}
	doc.Find("div.section").Each(func(i int,s *goquery.Selection){
		matchInfo := new(BasketballMatchInfo)
		matchId,_ := s.Attr("match_id")
		matchVs,_ := s.Attr("match")
		matchTime,_ := s.Attr("match_time")
		matchWeek,_ := s.Attr("match_week")
		leagueName,_ := s.Attr("league_val")
		matchInfo.matchId = matchId
		matchInfo.matchVS = matchVs
		matchInfo.matchTime = matchTime
		matchInfo.matchWeekTime =matchWeek
		matchInfo.leagueName = leagueName
		s.Find("table.saishi tr td").Each(func(i int ,s *goquery.Selection) {
			switch i {
				case 3:
					homeTeamName := getValue(s)
					matchInfo.homeTeamName = homeTeamName
				case 4:
					guestTeamName := getValue(s)
					matchInfo.guestTeamName = guestTeamName
				case 5:
					matchInfo.guestWin = getValue(s)
				case 6:
					matchInfo.homeWin = getValue(s)
				case 7:
					matchInfo.guestConcedeWin = getValue(s)
				case 8:
					matchInfo.concedeScore = getValue(s)
				case 9:
					matchInfo.homeConcedeWin = getValue(s)
				case 10:
					matchInfo.bigScore = getValue(s)
				case 11:
					matchInfo.finalScore = getValue(s)
				case 12:
					matchInfo.smallScore = getValue(s)

			}
		})
		fmt.Println(matchInfo)
	})

	defer response.Body.Close()
}
func main() {
	football_url := "http://www.lottery.gov.cn/football/counter.jspx"
	basketball_url := "http://www.lottery.gov.cn/basketball/counter.jspx"
	//足球赛果
	//football_result_url := "http://www.lottery.gov.cn/football/result.jspx"
	//篮球赛果
	//basketball_result_url := "http://www.lottery.gov.cn/basketball/result.jspx"
	//Crawler(&url)
	/**
	 * http://info.sporttery.cn/football/info/fb_match_hhad.php?m=88581
	 * http://www.lottery.gov.cn/football/match_hhad.jspx?mid=88580
	 */
	htmlHunter_football(&football_url)
	htmlHunter_basketball(&basketball_url)
}




