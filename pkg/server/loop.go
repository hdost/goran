package server

import (
"time"
"fmt"
)

func Loop(){
 for {
   time.Sleep(5 * time.Second);
   fmt.Println("Loop again");
 }
}
