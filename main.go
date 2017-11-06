package main

import (
	"strings"
	"strconv"
	"fmt"
	"log"
	"os"
	"github.com/ramin0/chatbot"
	"github.com/ryanbradynd05/go-tmdb"
)
import _ "github.com/joho/godotenv/autoload"
  
func chatbotProcess(session chatbot.Session, message string) (string, error) {
	_, historyFound := session["history"]
	if !historyFound { //if he only answered the first question (what's your name)
		session["history"] = map[int]string{}
		session["state"] = 0
		session["queryMap"]= make(map[string]string)

		history, _  := session["history"].(map[int]string)
		state,_ := session["state"].(int)

		history[state]=message;
		state=2
		session["history"] = history
		session["state"] =state
		output:=fmt.Sprint("Hello ",message,", how old are you?")
		return output,nil

	}	
	history, _  := session["history"].(map[int]string)
	state,_ := session["state"].(int)
	log.Printf("%d",state)
	Query,_ := session["queryMap"].(map[string]string)
	output:="";
	history[state]=strings.ToLower(message);
	//what to do at every stage
	switch state{
		case 1:// mood/ feeling
			switch history[state]{
				case "happy" ,"excited","joyful","cheerful","good":
					Query["with_genres"] ="35&10749|35|10402|10751"
					//state=3
				case "sad","depressed":
					Query["with_genres"] ="18"
					//state=3					
				case "energitic","hyper","adventurous":
					Query["with_genres"] ="9648|53|9648|10752|28|12"
					//state=3
				case "loved","love","lovely":
					Query["with_genres"] ="10749"
					//state=3
				default:
					state =1
					output=fmt.Sprintf("Please choose one of the following feelings/moods \n Happy, Excited,Sad,Depressed,energitic,Hyper,Adventurous,Loved ")
					return output,nil
			}
		switch (history[4]){
			case "":
				output=fmt.Sprintf("Do you have any prefered language other than English ?")
				state=3
			default:
				state =5
				output=outputf(session)
			
			
		}
			
			
			
		case 2:// age
			age ,_ := strconv.Atoi(message)
			if(age>=18){
				Query["include_adult"]="true"
			}else{
				Query["include_adult"]="false"
			}
		state=1;
		output=fmt.Sprintf("What is your mood/how are you feeling rn")
		case 3: // to check if he likes other languages
			switch history[state]{
				case "yes":
						Query["with_original_language"]="en"
				state=4
					output=fmt.Sprintf("What other language do you like?")
					
				case "no":
					Query["with_original_language"]="en"
					state =5

					output= outputf(session);
				default:
					state= 3
					output=fmt.Sprintf("Please answer yes or no")
			}
		case 4://to ask about the language
			switch history[state]{
				case "english":
					Query["with_original_language"]="en"
				state =5

					output=outputf(session)
				case "arabic":
					Query["with_original_language"]+="|ar"	
				state =5

					output=outputf(session)
				case "chinese":
					Query["with_original_language"]+="|zh"
				state =5

					output=outputf(session)
				case	"french":
					Query["with_original_language"]+="|fr"	
				state =5

					output=outputf(session)
				case	"german":
					Query["with_original_language"]+="|de"	
				state =5

					output=outputf(session)
				case	"hindi":
					Query["with_original_language"]+="|hi"	
				state =5

					output=outputf(session)
				case	"italian":
					Query["with_original_language"]+="|it"	
				state =5

					output=outputf(session)
				case "japanese":
					Query["with_original_language"]+="|ja"	
				state =5

					output=outputf(session)
				case "latin":
					Query["with_original_language"]+="|la"
				state =5

					output=outputf(session)
				default:
					state=4;
					output=fmt.Sprintf("These are the languages we support english,arabic,chinese,french,german,hindi,italian,japanese,latin please choose one of them ?")
			}
		case 5:  //to check wether he liked the recommendations
			switch history[state]{
				case "yes":				
					state=7
					output=fmt.Sprintf("Thank you for using our MovieGuru to terminate type Bye")
				case "no":
					state =6
					output=fmt.Sprintf("Do you want to change mood/language/both?")
				default:
					state= 3
					output=fmt.Sprintf("Please answer yes or no")
			}
		case 6:
			switch history[state]{
				case "mood":				
					state=1;
					output=fmt.Sprintf("What is your mood/how are you feeling rn")
					
				case "language":
					state=4;
					output=fmt.Sprintf("what language do you want to add?")
				case "both":
					state=1;
					history[3]=""
					output=fmt.Sprintf("What is your mood/how are you feeling rn")
				default:	
				state=6;
				output=fmt.Sprintf("Please answer what you want to change (mood/language/both)")
			}
		case 7:
			switch history[state]{
		case "bye":// session destroy automatically through handle chat
				//chatbot.Destroy();
		default:
			state =8
			output=fmt.Sprintf("do you want anything else?")
		}
		case 8: //if he didnt say bye after case 5
			switch history[state]{
			case "yes":
				state =6 //sssssssss
			case"no":
				state =7
				output=fmt.Sprintf("Thank you for using our MovieGuru to terminate type Bye")
			default:
				state=8
				output=fmt.Sprintf("Please answer yes or no")

		}
			//fightClubInfo,_ := TMDb.GetMovieInfo(457497, nil)
			//output= fmt.Sprintf(fightClubInfo.OriginalLanguage)
		//output=fmt.Sprintf("Do you have a favourit actor/actress ?")
		
		
			
		default:
			output=fmt.Sprintf("nothing",state)
	}

	session["history"] = history
	session["state"] =state
	session["queryMap"]=Query
	//TMDb:=tmdb.Init("925d983ccfd7a1c5928c6b7ef7b6b692")
	//fightClubInfo, _ := TMDb.GetMovieInfo(550, nil)
	return output ,nil
}
func outputf(session chatbot.Session) (string){
	Query,_ := session["queryMap"].(map[string]string)
	state,_ := session["state"].(int)
	output:="";
	var o string
	
	Query["sort_by"]="vote_average.desc,popularity.desc"
	Query["page"]="1";
	TMDb:=tmdb.Init("925d983ccfd7a1c5928c6b7ef7b6b692")
	Movies, _ := TMDb.DiscoverMovie(Query)
	for i := 0; i < 6; i++ {
		o += Movies.Results[i].OriginalTitle +", "
	}
	o+= "\n" 
	o+="Are my recommendations okay?(yes/no)"
	//log.Printf("%d",state)
	session["queryMap"]=Query
	session["state"] =state
	
	output= fmt.Sprintf(o)
	return output

}
func main(){
	chatbot.WelcomeMessage = "Hello I'm the MovieGuru, What is your name?"
	chatbot.ProcessFunc(chatbotProcess)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}