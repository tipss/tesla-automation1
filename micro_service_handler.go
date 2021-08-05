package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

/*
 * Handle time zone request from web URL
 * Read and output the current time in destination Zone
*/
func zoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()
	m := make(map[string]interface{})
	m["time"] = time.Now().UTC()
	if tz := q.Get("tz"); tz != "" {
		if loc, err := time.LoadLocation(tz); err == nil {
			m["time"] = time.Now().In(loc)
		} else {
			m["error"] = "unknown timezone"
			m["time"] = nil
		}
	}
	json.NewEncoder(w).Encode(m)
	fmt.Println("Exit Zone Handler\n")
}

/*
 * Handle frr config reuquests
 * Read and output the current FRR router config and send output to caller
*/
func frrDeviceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()
	m := make(map[string]interface{})
	m["frr"] = "FRR Command output:"
	if tz := q.Get("cmd"); tz != "" {
		app := "vtysh"
		arg0 := " "
		tz, _ = url.PathUnescape(string(tz))
		one := strings.Replace(tz, "-c", "\" -c \"", -1)
		//one = one + "\""
		fmt.Println(one[1:])
		cmd := exec.Command(app, arg0, one[1:])
		out, err := cmd.Output()
		fmt.Printf("vtysh output is :%s\n", string(out))
		if err != nil {
			m["error"] = "unknown command!"
			//log.Fatal(err)
		} else {
			fmt.Printf("Input %s\n", tz)
			m["frr"] = string(out)
		}
	}
	json.NewEncoder(w).Encode(m)
	fmt.Println("Exit frrDeviceHandler")
}

func main() {
	/* Various Handlers */
	
	/* Time Zone Handler */
	http.HandleFunc("/time", zoneHandler)
	
	/* FRR config handler */
	http.HandleFunc("/frr", frrDeviceHandler)
	/* Run server at port 8000 */
	log.Fatal(http.ListenAndServe(":8000", nil))
}
