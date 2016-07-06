package main
import (
	"flag"
	"github.com/golang/glog"
	"github.com/openedinc/opened-go"
  "bufio"
)

func init() {
  flag.Parse()
}  


//read in a file

//read first line
//continue to loop

func main() {
	glog.V(1).Infof("My OpenEd Resource Find Script")
token,err := opened.GetToken ("","","","")


  stdsheet := os.Open("filename")
  for i := 1; i <len(file; i++ {


  query_params:=make(map[string]string)
  std :=bufio.NewScanner(os.Stdin)
  query_params["standard"]="std"

  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources,%+v",err)    
  } 
  glog.V(1).Infof("%d results returned",len(results.Resources))
 }
}