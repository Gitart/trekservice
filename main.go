package main

import (
	     "fmt"              
         "net/http"         
         "io/ioutil"      
         "os"
         "time"
         "log"
         r "github.com/dancannon/gorethink"          // Rethinkdb Driver
)

// Declaration inetrfaces & structure 
type Mst map[string]interface{}                    // Map - string - interface

//  Gloabal Variables
var sessionArray []*r.Session


/********************************************************************************************************************************
  Инициализация подключения 
*********************************************************************************************************************************/
func Dbini() {
    
   // Инициализация подключения к 
   IpPort        := "10.10.10.10:28015"         //    
   DatabaseKey   := "XXXXX"
   DatabaseName  := "System"

   // Сессия подключения
   session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: DatabaseName, AuthKey:DatabaseKey})   
     
   // Обработка ошибок
   if err != nil {
      log.Println(err)
      return
   }
   sessionArray=append(sessionArray, session)
}


/
/  Запуск сервиса с параметрами
/
func main() {
     // go Dloopp()                                // Запуск паралелльного процесса
     Ip:=os.Getenv("IPPORT")
     fmt.Println("Control Port: " + Ip + "\n\n")
     Dbini()                                       // Инициализация подключения
     nmTicker()                                    // Запуск тикера
   
}


/
/	 TICKER
/	 
/	 Выполняется постоянно в цикле 
/   (если не включена go func() {}())
/
/	 Необходим для мониторинга состояния сервиса!!!
/	 Параметр time.Sleep(time.Second * 500)
/	 Определяет количество интераций

func nmTicker(){

   // Частота запуска в секундах time.Second * 1
   // ticker := time.NewTicker(time.Minute *1)
   // ticker := time.NewTicker(time.Millisecond * 1 )
   IpPort :=os.Getenv("IPPORT") //"http://10.0.3.24:5555/"
   Notify :=os.Getenv("NOTIFY")

   // var Inter time.Duration
   // It,_:= time.ParseDuration("5.5s") 
   // os.Getenv("INTERVAL")
   // It,_:=fmt.Printf("%f", 5)

   Interval := os.Getenv("INTERVAL")+"s"         // Cекунды чтения  
   Inter, _ := time.ParseDuration(Interval)      // Интервал сканирования в секундах

   // Частота проверкит в секундах
   // ticker := time.NewTicker(time.Second * 5)
   ticker := time.NewTicker(Inter)
  
   // i:=0 
   // go func() - Данная возможность позволяет выйти из цикла по количеству итераций кратным указанным в Sleep 
    go func() {
        
          // for t := range ticker.C {
          for range ticker.C {
        	  // i++
            // fmt.Println("Проверка ", i, t)
            //r.DB("System").Table("Log").Insert(Mst{"Descr":"Проверка вставки в скоростном режиме", "Id":i}).RunWrite(sessionArray[0])
            // checktrack()
            if Checkapi(IpPort){
               log.Println(Notify)
             } 

             // Если необходимо считывать с тела 
             // else{
             // checktrack() 
             // }
          }
    }()

   // Количество повторений кратное 10 = 50 раз если не будет этой строки цикл будет вечным
   time.Sleep(time.Second * 10000)

   // time.Sleep(time.Millisecond * 500)
   ticker.Stop()
   fmt.Println("Ticker stopped")
}


//  Control porst
func Checkapi(url string) bool {
     _, err := http.Get(url)

     if err!=nil{
        return true
      } 
        return false
}


//  Preview
func checktrack() {
    response, err := http.Get("http://localhost:1111/")

    if err != nil {
        //fmt.Printf("%s", err)
        //os.Exit(1)
        fmt.Println("Связь порвалась..... Сервису плохо..")
    } else {
        defer response.Body.Close()
       
        contents, err := ioutil.ReadAll(response.Body)
       
        if err != nil {
           fmt.Printf("%s", err)
           os.Exit(1)
        }

        fmt.Printf("%s\n", string(contents))

        // if string(contents)=="Ok"{
        //    fmt.Printf("%s\n", "OKKKEY connect!")
        //  }
    }
  }



//  Add new ticker
func nmTickert(){
    timeChan := time.NewTimer(time.Second).C
    tickChan := time.NewTicker(time.Millisecond * 400).C
    doneChan := make(chan bool)

    go func() {
        time.Sleep(time.Second * 2)
        doneChan <- true
    }()

    nmTicker()

    for {
        select {
        case <- timeChan:   fmt.Println("Timer expired")
        case <- tickChan:   fmt.Println("Ticker ticked")
        case <- doneChan:   fmt.Println("Done")
        return
      }
    }
}

// Sleep
func Dloopp(){
   for{
        time.Sleep(time.Second * 5)
        fmt.Println("-->>Ok")
   }
}
