package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Resume struct {
	Name        string
	Email       string
	Phone       string
	Summary     string
	WorkHistory []WorkExperience
	Education   []Education
	Skills      []string
}

type WorkExperience struct {
	Company          string
	Position         string
	StartDate        string
	EndDate          string
	Responsibilities []string
	Highlights       []string
}

type Education struct {
	Institution string
	Degree      string
	StartDate   string
	EndDate     string
}

func main() {
	resume := Resume{
		Name:    "Justis Gipson",
		Email:   "justis710@gmail.com",
		Phone:   "317-605-4164",
		Summary: "DevOps Engineer with a passion for automation and cloud technologies. I have experience with AWS and GCP. I have a strong background in Linux and MacOS systems administration. I am a quick learner and I am always looking to expand my skillset.",
		WorkHistory: []WorkExperience{
			{
				Company:   "Ellipsis Education",
				Position:  "DevOps Engineer",
				StartDate: "2020-01-01",
				EndDate:   "Present",
				Highlights: []string{
					"Joined as a Junior Developer to enhance technology scalability",
					"Led migration of the organization's infrastructure to AWS, resulting in a ~25% reduction in costs",
					"Managed and updated a key educational application, leading to a successful relaunch in Spring 2022",
					"Responsibilities include overseeing cloud infrastructure, ensuring effective monitoring, optimizing continuous integration and deployment processes, developing new features, managing database updates, and maintaining robust security protocols for the organization and applications",
					"Worked with Node, TypeScript, Python, Knex, Postgres, CircleCI, Github Actions, Elastic Beanstalk, Cloudfront, AWS Lambda, Eventbridge, RDS, CloudWatch, GuardDuty, Systems Manager, Grafana, Docker, Cloudflare, GCP",
				},
			},
			{
				Company:   "Zylo",
				Position:  "Associate Software Engineer",
				StartDate: "2020-01-01",
				EndDate:   "2020-12-31",
				Highlights: []string{
					"Contributed to the software development team of a SaaS management firm",
					"Played a pivotal role in streamlining dashboards for software spend and license management, leading to significant cost savings for the companies implementing the end product",
					"Worked with React.JS, Node.js, AWS, Docker",
				},
			},
			{
				Company:   "HikerFeed",
				Position:  "Software Engineer",
				StartDate: "2019-01-01",
				EndDate:   "2019-12-31",
				Highlights: []string{
					"Part of a team developing a unique platform for the backpacking community",
					"Involved in various stages of development",
					"Worked with Vue.JS/Nuxt and PHP/Laravel and Azure",
				},
			},
		},
		Education: []Education{
			{
				Institution: "Eleven Fifty Academy",
				Degree:      "Software Development",
				StartDate:   "2018-10-01",
				EndDate:     "2019-03-12",
			},
		},
		Skills: []string{
			"AWS",
			"GCP",
			"Linux",
			"Python",
			"Node.js",
			"TypeScript",
			"Postgres",
			"Docker",
			"Go",
			"Grafana",
			"Prometheus",
			"Github Actions",
		},
	}

	displayResume(resume)
}

func displayResume(resume Resume) {
	boldWhite := color.New(color.FgWhite, color.Bold)
	boldBlue := color.New(color.FgBlue, color.Bold)
	boldGreen := color.New(color.FgGreen, color.Bold)
	boldYellow := color.New(color.FgYellow, color.Bold)
	boldMagenta := color.New(color.FgMagenta, color.Bold)

	boldWhite.Println("Resume - Interactive CLI")
	boldWhite.Println("------------------------")

	boldBlue.Println("1. Personal Information")
	boldBlue.Println("2. Work History")
	boldBlue.Println("3. Education")
	boldBlue.Println("4. Skills")
	boldBlue.Println("5. Exit")

	for {
		boldWhite.Print("\nEnter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			boldGreen.Printf("Name: %s\n", resume.Name)
			boldGreen.Printf("Email: %s\n", resume.Email)
			boldGreen.Printf("Phone: %s\n", resume.Phone)
			boldGreen.Printf("Summary: %s\n", resume.Summary)
		case 2:
			boldYellow.Println("Work History:")
			for _, work := range resume.WorkHistory {
				boldMagenta.Printf("Company: %s\n", work.Company)
				boldMagenta.Printf("Position: %s\n", work.Position)
				boldMagenta.Printf("Duration: %s - %s\n", work.StartDate, work.EndDate)
				boldYellow.Println("Highlights:")
				for _, highlight := range work.Highlights {
					fmt.Printf("- %s\n", highlight)
				}
				fmt.Println()
			}
		case 3:
			boldYellow.Println("Education:")
			for _, edu := range resume.Education {
				boldMagenta.Printf("Institution: %s\n", edu.Institution)
				boldMagenta.Printf("Degree: %s\n", edu.Degree)
				boldMagenta.Printf("Duration: %s - %s\n", edu.StartDate, edu.EndDate)
				fmt.Println()
			}
		case 4:
			boldYellow.Println("Skills:")
			for _, skill := range resume.Skills {
				fmt.Printf("- %s\n", skill)
			}
		case 5:
			boldWhite.Println("Exiting...")
			os.Exit(0)
		default:
			color.Red("Invalid choice. Please try again.")
		}
	}
}
