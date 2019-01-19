package models

import (
"fmt"
"encoding/json"
"net/http"
//"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
//_ "github.com/jinzhu/gorm/dialects/mysql"
_ "github.com/jinzhu/gorm/dialects/postgres"
_ "github.com/lib/pq"

"github.com/tidwall/gjson"

)


//from now on, i only to modify this file to add gtd review standard,that includes, detail struct and if control and totalscore and review struct instance data

//json is that it will be changed to this string json in db
//https://stackoverflow.com/questions/26327391/json-marshalstruct-returns
type Reviewdatadetail struct{
Totalscore    int `json:"totalscore"`
Averagescoreofhistory   int `json:"averagescoreofhistory"`
Patience      int  `json:"patience"`
Usebrain      int   `json:"usebrain"`
Battlewithlowerbrain int   `json:"battlewithlowerbrain"`
Learnnewthings int     `json:"learnnewthings"`
Makeuseofthingsuhavelearned int    `json:"makeuseofthingsuhavelearned"`
Difficultthings int  `json:"difficultthings"`
Challengethings int  `json:"challengethings"`
Threeminutes    int   `json:"threeminutes"`
Getlesson       int    `json:"getlesson"`
Learntechuse    int    `json:"learntechuse"` 
Thenumberoftasks_score  int    `json:"thenumberoftasks_score"`
Serviceforgoal_score  int    `json:"serviceforgoal_score"`
Onlystartatask int       `json:"onlystartatask_score" sql:"size:999999"`
Atomadifficulttask  int    `json:"atomadifficulttask"`
Alwaysprofit       int     `json:"alwaysprofit"` 
Markataskimmediately int   `json:"markataskimmediately"`
Doanimportantthingearly int  `json:"doanimportantthingearly"`
}

























func Reviewalgorithmjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  //fmt.Println(cookie1.Value)

  var reviewdays []Reviewofday
  db.Where("email =  ?", email).Order("date").Find(&reviewdays)
  
  c.JSON(200, gin.H{
      "reviewdata":reviewdays,
    })

}




func Reviewalgorithmjsonforios(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  email:= "yang756260386@gmail.com"
  //fmt.Println(cookie1.Value)

  var reviewdays []Reviewofday
  db.Where("email =  ?", email).Order("date").Find(&reviewdays)

  c.JSON(200, gin.H{
      "reviewdata":reviewdays,
    })

}










func Reviewforios(c *gin.Context) {


c.HTML(http.StatusOK, "reviewalgoforios.html",nil)


}

















func   Check_reviewdaylog(date string,email string){
/*
check if date row was created in reviewday table,if it is no,the function will create it

*/

var reviewday Reviewofday 
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewday)
if reviewday.Date!=date{
db.Create(&Reviewofday{Date: date,Email:email,Details:"no"})
}else{
fmt.Println("===========the record has been created in the past==========")
}
}








//compute total_scores of someday

func Compute_singleday(date string,email string) string{
//https://tour.golang.org/basics/10
fmt.Println("------------ i am here to compute the single day---------------------------")
var tasks []Tasks
//email := "yang756260386@gmail.com"
var brainuse_score,makeuseofthethingsuhavelearned_score,difficultthings_score,threeminutes_score,getlesson_score,learntechuse_score,battlewithlowerbrain_score,patience_score,learnnewthings_score int = 0,0,0,0,0,0,0,0,0 
var serviceforgoal_score,onlystartatask_score = 0,0
var atomadifficulttask_score,alwaysprofit_score = 0,0
var doanimportantthingearly_score,markataskimmediately_score = 0,0

var challengetag_score = 0
db.Where("Email= ?", email).Where("finishtime =  ?", date).Order("id desc").Find(&tasks)

var taskcount_score int

var countoffinishedtasks int

var countofgivenuptasks int


db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Not("status", []string{"unfinished","unfinish"}).Count(&countoffinishedtasks)


db.Table("tasks").Where("Email= ?", email).Where("finishtime =  ?", date).Where("status =?","giveup").Count(&countofgivenuptasks)




fmt.Println("-----=======-------++++++++++++++=-----======----------------")
fmt.Println(countofgivenuptasks)



taskcount_score =  2  * (countoffinishedtasks - countofgivenuptasks) +countofgivenuptasks*1

 for _,item :=range tasks{
fmt.Println("------------i had been into loop----------------")

var jsonoftasktags = item.Tasktags
if  challengetag := gjson.Get(jsonoftasktags, "challengetag").String();challengetag=="yes"{
challengetag_score  = challengetag_score + 5
}




var json = item.Reviewdatas
fmt.Println(json)
fmt.Println("------------i had been into loop----------------")
if  brainuse := gjson.Get(json, "brainuse").String();brainuse=="yes"{
fmt.Println(brainuse)
brainuse_score = brainuse_score +5
 } 
if  makeuseofthings := gjson.Get(json, "makeuseofthings").String();makeuseofthings=="yes"{
makeuseofthethingsuhavelearned_score = makeuseofthethingsuhavelearned_score + 5
 }


if  doanimportantthingearly := gjson.Get(json, "doanimportantthingearly").String();doanimportantthingearly =="yes"{
doanimportantthingearly_score = doanimportantthingearly_score + 10
 }
 

if  markataskimmediately := gjson.Get(json, "markataskimmediately").String();markataskimmediately =="yes"{
markataskimmediately_score = markataskimmediately_score + 1
 }





if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
alwaysprofit_score = alwaysprofit_score + 5
 }


if  alwaysprofit := gjson.Get(json, "alwaysprofit").String();alwaysprofit=="yes"{
alwaysprofit_score = alwaysprofit_score + 5
 }




if  learnnewthings := gjson.Get(json, "learnnewthings").String();learnnewthings=="yes"{
learnnewthings_score = learnnewthings_score +5
 }


if  serviceforgoal := gjson.Get(json, "serviceforgoal").String();serviceforgoal=="yes"{
serviceforgoal_score  = serviceforgoal_score  + 20
 }


if  onlystartatask := gjson.Get(json, "onlystartatask").String();onlystartatask=="yes"{
onlystartatask_score  = onlystartatask_score  + 10
 }





if  battlewithlowerbrain := gjson.Get(json, "battlewithlowerbrain").String();battlewithlowerbrain=="yes"{
battlewithlowerbrain_score = battlewithlowerbrain_score +5

 }


if  atomadifficulttask := gjson.Get(json, "atomadifficulttask").String();atomadifficulttask=="yes"{
atomadifficulttask_score = atomadifficulttask_score +5

 }






if  patience := gjson.Get(json, "patience").String();patience=="yes"{
patience_score = patience_score + 10

 }




if  difficultthings := gjson.Get(json, "difficultthings").String();difficultthings=="yes"{
difficultthings_score = difficultthings_score +10

 }



if  threeminutes := gjson.Get(json, "threeminutes").String();threeminutes=="yes"{
threeminutes_score = threeminutes_score +5

 }



if  getlesson:= gjson.Get(json, "getlesson").String();getlesson=="yes"{
getlesson_score = getlesson_score +5

 }



if  learntechuse := gjson.Get(json, "learntechuse").String();learntechuse=="yes"{
learntechuse_score = learntechuse_score +5

 }


}

total_score:=taskcount_score+doanimportantthingearly_score+atomadifficulttask_score+onlystartatask_score+markataskimmediately_score+challengetag_score + brainuse_score+alwaysprofit_score + makeuseofthethingsuhavelearned_score + battlewithlowerbrain_score + patience_score + learnnewthings_score+difficultthings_score+threeminutes_score+getlesson_score+learntechuse_score + serviceforgoal_score
review := &Reviewdatadetail{Totalscore: total_score,Challengethings:challengetag_score,Markataskimmediately:markataskimmediately_score,Doanimportantthingearly:doanimportantthingearly_score,Alwaysprofit:alwaysprofit_score,Atomadifficulttask:atomadifficulttask_score,Onlystartatask:onlystartatask_score,Thenumberoftasks_score:taskcount_score,Difficultthings:difficultthings_score,Threeminutes:threeminutes_score,Getlesson:getlesson_score,Learntechuse:learntechuse_score,Patience:patience_score,Serviceforgoal_score:serviceforgoal_score,Usebrain:brainuse_score,Battlewithlowerbrain:battlewithlowerbrain_score,Learnnewthings:learnnewthings_score,Makeuseofthingsuhavelearned:makeuseofthethingsuhavelearned_score}


//https://stackoverflow.com/questions/8270816/converting-go-struct-to-json

reviewjson, err := json.Marshal(review)
if err!=nil{
fmt.Println("----------there is an error ----------")
fmt.Println(err)
}

reviewstring := string(reviewjson)
fmt.Println("-------------i am pritning reviewstring---------------")
fmt.Printf("%+v\n", review)
//fmt.Println(reviewjson)
fmt.Println(reviewstring)
fmt.Println("-------------i am pritning reviewstring---------------")

var reviewday Reviewofday
db.Where("date =  ?", date).Where("email =  ?", email).Find(&reviewday)
db.Model(&reviewday).Update("Details", reviewstring)

return reviewstring
}






