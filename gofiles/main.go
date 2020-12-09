package main

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	mysql, _        = newMysql()
	tagInt          = ""
	categoryQueries = []string{"scary games", "horror games", "action games", "action adventure games", "role playing games", "simulation games", "strategy games", "puzzle games", "casual games", "shooting games", "platform games", "fighting games", "beat up games", "ninja games", "robber games", "stealth games", "survival games", "battle royale games", "rhythm games", "text adventure games", "visual novel games", "interactive games", "mmorpg games", "roguelike games", "tactical games", "sandbox games", "first person party games", "jrpg games", "monster tamer games", "construction games", "life simulation games", "vehicle simulation games", "artillery games", "auto chess games", "auto battler games", "moba games", "multiplayer online battle arena games", "rts games", "real time strategy games", "real time tactics games", "tower defense games", "turn based strategy games", "turn based games", "war game", "machine game", "dating games", "gta games", "racing games", "sport games", "board games", "competitive games", "card games", "logic games", "party games", "trivia games", "typing games", "art game", "educational games", "exercise games", "cross genre games", "open world games", "science fiction games", "sifi games"}
	maxGorutines    = 10
)

func main() {
	maxLimiter := make(chan int, maxGorutines)
	for _, query := range categoryQueries {
		maxLimiter <- 1
		go func(search string) {
			saveEmailsFromSearch(search)
			<-maxLimiter
		}(query)
	}
}

func saveEmailsFromSearch(search string) {
	HTMLStr, err := getHTMLStrNodeAppStoreScraper("https://play.google.com/store/search?c=apps&q=" + url.QueryEscape(search))
	if err != nil {
		errLogger.Printf(err.Error())
	}
	appStrAry, err := getAppURLStrAryFromHTMLStr(&HTMLStr)
	if err != nil {
		errLogger.Printf(err.Error())
	}
	noDupAppStrAry := tools.getStrAry.noDuplicate(&appStrAry)

	for _, appStr := range noDupAppStrAry {
		resStr, err := scraper.getStr.results("https://play.google.com" + appStr)
		if err != nil {
			errLogger.Printf(err.Error())
		}
		resStr, err = tools.getStr.regex(&resStr, `mailto:[^"]+`)
		if err != nil {
			errLogger.Printf(err.Error())
		}
		resStr, err = tools.getStr.after(resStr, `mailto:`)
		if err != nil {
			errLogger.Printf(err.Error())
		}
		logger.Printf("%v", resStr)
		err = saveEmail(resStr)
		if err != nil {
			errLogger.Printf(err.Error())
		}
	}
}

func getHTMLStrNodeAppStoreScraper(url string) (requestStr string, err error) {
	request, err := scraper.getRequest(
		`https://s07qp3qvmg.execute-api.us-east-2.amazonaws.com/default/nodeAppStoreScraper`,
		`GET`,
	)
	if err != nil {
		return "", err
	}
	scraper.setRequest.header(request, "url", url)
	requestStr, err = scraper.getStr.requestResults(request)
	if err != nil {
		return "", err
	}
	return requestStr, err
}

func getAppURLStrAryFromHTMLStr(HTMLStr *string) (AppURLStrAry []string, err error) {
	AppURLStrAry, err = tools.getStrAry.regex(HTMLStr, `\/store\/apps\/details\?id=[^"]+`)
	if err != nil {
		return nil, err
	}
	return
}

func saveEmail(emailStr string) (err error) {
	err = mysql.execute.query(`INSERT INTO steamScraper.emails(email) VALUES("` + emailStr + `")`)
	if err != nil && !strings.Contains(err.Error(), `Duplicate entry`) {
		logger.Printf(err.Error())
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
