package main

import ( 
	"encoding/json"

   "fmt"

   ps "github.com/mitchellh/go-ps"

   "github.com/StackExchange/wmi"
)

 
type ProcessList struct {
	Id int      `json:"id"`
	Name  string  `json:"name"`
}
 
type OperatingSystem struct {
	 Name string 
}

func getProcessList() interface{} {
 	 response := make([]ProcessList,40)
  	processList, err := ps.Processes()
	if err != nil {
		fmt.Println("ps.Processes() Failed, are you using windows?")
			return  "response"
	}
	 
 	for x := range processList {
			var process ps.Process
			process = processList[x]
			// Push only 40 items in the list
			if x <40{
				response[x] = ProcessList{Id:process.Pid(), Name:process.Executable()}; 			
		}			 
	}

	fmt.Println(response)
	res, err := json.Marshal(&response)

	fmt.Println(string(res))

	if err != nil {
		fmt.Println(err.Error()) 		
	}
 	return string(res)
}

func getOSinfo () string{
	var osInfo []OperatingSystem
	err := wmi.Query("Select * from Win32_OperatingSystem", &osInfo)
	if err != nil {
		fmt.Printf(err.Error())
		return err.Error()
	}
	return (osInfo[0].Name)

}