package demo

import ( "fmt" ) 

func ItpRun() {
  date := fmt.Sprintf("%d-%d-%d", 2023, 1, 28)
  time := fmt.Sprintf("%d:%d:%d", 8, 30, 0)
  datetime := fmt.Sprintf("%s,%s", 
                    "1/28/2023", "08:00:00")
  println(date); println(time); println(datetime)
}