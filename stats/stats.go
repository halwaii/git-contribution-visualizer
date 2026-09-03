package stats

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Stats(repos []string) {

	// initialize map
	// contribution contains date -> number of commits
	contributions := make(map[string]int)
	for _, repo := range repos{
	// 1) define command and its arguments
	// git log --format=%ad --date=short
		cmd := exec.Command("git","-C",repo,"log","--format=%ad","--date=short")

	// 2) run the command and capture output
		output, err := cmd.Output()
		if err != nil{
			// fmt.Println("error reading repository : ",err)
			continue
		}
	
	// what does split does ?
	// string(output) = "2026-8-8
	// 					 2026-4-6
	// 					 2026-8-3"
	// split where "\n" present 
	// dates = {"2026-8-8","2026-4-6","2026-8-3"}

		dates := strings.Split(string(output),"\n")

		for _, date := range dates{
			if date == ""{
				continue
			}
			// date -> key , commits -> values
			contributions[date]++
		}
		// 3) print output string
		// fmt.Println("repository : ", repo)
		// fmt.Println(string(output))
	}

	today := time.Now()

	// start from 364 days ago
	start := today.AddDate(0,0,-210)

	// move to previous sunday
	// loop till sunday is found
	for start.Weekday() != time.Sunday{
		start = start.AddDate(0,0,-1)
	}
	fmt.Println()
	fmt.Println("Git Contribution Graph")
	fmt.Println("----------------------")
	days := []string{"sun","mon","tue","wed","thu","fri","sat"}

	// print month
	fmt.Print("      ")
	currMonth := time.Month(0)

	for week:=0;week<=30;week++{
		date:= start.AddDate(0,0,week*7)
		month := date.Month()
		if month != currMonth{
			fmt.Printf("%-3s", date.Format("Jan"))
			currMonth = month
		} else {
			fmt.Print("   ")
		}
	}
	fmt.Println()
	// make 7 rows sunday to saturday
	for day:=0;day<7;day++{
		fmt.Printf("%s ",days[day])
		// make columns
		for week:=0;week<=30;week++{
			// date calculation
			date := start.AddDate(0,0,week*7+day)

			if date.After(today){
				fmt.Print("  ")
				continue
			}
			// convert date to string format
			dateStr := date.Format("2006-01-02")

			// get commits from map
			commits := contributions[dateStr]

			// get background color
			color := getColor(commits)

			// print
			if commits==0{
				fmt.Print(color)
				fmt.Print("- ")
			} else {
				fmt.Print(color)
				fmt.Printf("%2d", commits)
			}
			fmt.Print("\033[0m")
			fmt.Print(" ")
		}
		fmt.Println()
	}
	// old code

	// today := time.Now()

	// // check for commits in last 365 days
	// for i:=0;i<365;i++{
	// 	// 1) AddDate(year, month, day) => (0,0,-1) -> yesterday
	// 	date := today.AddDate(0,0,-i)

	// 	// 2) convert date to string format
	// 	// what happened on this day ?
	// 	dateStr := date.Format("2006-01-02")

	// 	// 3) check for number of commits on that particular day
	// 	commits := contributions[dateStr]
		
	// 	fmt.Println(dateStr, commits)
	// }
	// // fmt.Println(contributions)
}

// helper function to get background color for each day
func getColor(commits int) string{
	if commits == 0{
		return "\033[48;5;236m" // gray
	} else if commits <=2 {
		return "\033[48;5;255m"
	} else if commits <=4 {
		return "\033[48;5;220m"
	}else if commits <=5 {
		return "\033[48;5;46m"
	} else if commits <=7 {
		return "\033[48;5;208m"
	}
	return "\033[48;5;196m"

}