package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	version = "0.1.0"

	codereview = CodeReview{
		Vocabulars: []Vocabular{
			Vocabular{
				Base: Base{
					Name:        "changer",
					Description: "an individual unit of work altering what exists",
				},
			},
			Vocabular{
				Base: Base{
					Name:        "submission",
					Description: "a collection of changes",
				},
			},
			Vocabular{
				Base: Base{
					Name:        "submitter",
					Description: "the persion proposing the submission",
				},
			},
			Vocabular{
				Base: Base{
					Name:        "reviewer",
					Description: "the people evaluating the submission",
				},
			},
			Vocabular{
				Base: Base{
					Name:        "annotation",
					Description: "remarks of ratings bestowed upon the submission",
				},
			},
		},
		Methods: []Method{
			Method{
				Base: Base{
					Name:        "Team Review",
					Description: "before you release the next version, start a team meeting and review...",
					DescriptionLong: []string{
						"for a final review organize a team meeting and review / reverse engineer the project.",
						"the team meeting can be organized arround the checklist and it is recomendet to write a review report.",
						"TODO: send meeting email",
						"TODO: create meeting timetable",
						"TODO: report creator tool",
					},
				},
			},
			Method{
				Base: Base{
					Name:        "Pair Programming",
					Description: "this method is perfect for early reviews",
					DescriptionLong: []string{
						"Pair programming is an agile software development technique in which two programmers work together at one workstation.",
						"One, the driver, writes code while the other, the observer or navigator,[1] reviews each line of code as it is typed in.",
						"The two programmers switch roles frequently.",
						"TODO: report creator tool",
					},
				},
			},
		},
		Problems: []Problem{
			Problem{
				Base: Base{
					Name:        "Mental Model Synccronisation",
					Description: "check the architecture implementation",
				},
			},
			Problem{
				Base: Base{
					Name:        "Tribal knowledge development",
					Description: "'architecture oral history requires that the team is both willing and able to retell the stories and keep the oral history alive' Michael Keeling",
				},
			},
		},
		Audits: []Audit{
			Audit{
				Base: Base{
					Name:        "Application Architecture",
					Description: "",
				},
			},
		},
		Maintainability: []Maintain{
			Maintain{
				Base: Base{
					Name:        "Learnability",
					Description: "is it well documented",
				},
				Items: []string{
					"Patterns and Conventions",
					"Risks & Goals",
					"Common Vocabulary",
					"Teaching Moments",
				},
			},
			Maintain{
				Base: Base{
					Name:        "Understandable",
					Description: "how to getting started",
				},
				Items: []string{
					"point out problems",
					"enables elevator pitch",
				},
			},
			Maintain{
				Base: Base{
					Name:        "Serviceable",
					Description: "how to use",
				},
				Items: []string{
					"exposes addressable gotchas",
					"exposes end-user interaction points",
					"Establishes consensus on supported workflows",
				},
			},
		},
		Tools: []Tool{
			Tool{
				Base: Base{
					Name:        "Version Control",
					Description: "is the sourcecode under version control?",
					DescriptionLong: []string{
						"for versioning the sourcecode it is highly recomendet to use a VCS (Version Control System).",
						"the most popular implementations are git and mercurial-hg",
					},
				},
			},
			Tool{
				Base: Base{
					Name:        "Continuous Integration",
					Description: "is the sourcecode under continuous integration?",
					DescriptionLong: []string{
						"for an automatted review system you can use different types of ci services.",
						"for open source software you can use travis-ci.",
						"for closed source you can run your own ci server local and fire up a hook to crunch the latest changes",
						"the following software can be used to bring up a ci server",
						"- jenkins       https://jenkins.io/",
						"- strider-cd    https://github.com/Strider-CD/strider",
					},
				},
			},
			Tool{
				Base: Base{
					Name:        "Code Complexity",
					Description: "",
					DescriptionLong: []string{
						"the following parameters can be measured and used to detect hacky code:",
						"- Maintain",
						"- Difficulty",
						"- Lines of code                https://github.com/bytbox/sloc",
						"- Estimated numbers of bugs",
					},
				},
			},
		},
		Checklist: []Check{
			// "==> is the readme fine?"
			// "==> do you use linter"
			// "==> are tests imlemented"
			Check{
				Base: Base{
					Name: "Functionality",
				},
				Items: []string{
					"implement relevant specification(s)",
					"anticipate and address failure modes",
					"avoid (or adequately justifies) obvious performance issues",
					"provide appropriate user feedback",
					"expose logs, errors, and other runtime debugging details as needed",
				},
			},
			Check{
				Base: Base{
					Name: "Testing",
				},
				Items: []string{
					"secure behaviors described in relevant specification(s)",
					"test public interfaces",
					"cover anticipated failure modes",
					"benchmark recognized performance issues",
					"run green",
				},
			},
			Check{
				Base: Base{
					Name: "Legibilit",
				},
				Items: []string{
					"embrace style and patterns established by the project, team, and community",
					"consume existing libraries, methods, and types",
					"use suggestive names for functions, variables, and classes",
					"document public interfaces as appropriate",
				},
			},
			Check{
				Base: Base{
					Name: "Security",
				},
				Items: []string{
					"bound and sanitize inputs",
					"sanitize outputs, including analytics and logs",
					"recognize and address common vulnerabilities (SQL injection, XSS, buffer overruns, etc)",
					"isolate platform-/environment-specific functionality",
				},
			},
			Check{
				Base: Base{
					Name: "VCS Hygiene",
				},
				Items: []string{
					"provide clear, useful commit messages",
					"reference external issues, tickets, and bugs where they exist",
					"use incremental, atomic commits",
					"avoid binaries, logs, local configurations, environment-specific files, plaintext secrets, or other unwanted artifacts",
				},
			},
		},
		Feedback: []Feedback{
			Feedback{
				Base: Base{
					Name: "What/how format",
				},
				Items: []string{
					"What are some ways to make this code more readable?",
					"How can the code be optimized to parse the JSON only once?",
					"What happens when X happens (edge case)?",
				},
			},
			Feedback{
				Base: Base{
					Name: "I like/wonder/wish format",
				},
				Items: []string{
					"I like how this is flexible and makes it easy to test.",
					"I wonder what the advantage of looping through the array every time is.",
					"I wish the JSON can be parsed only once since it’s expensive to read the file every time.",
				},
			},
			Feedback{
				Base: Base{
					Name: "Avoid you/your:",
				},
				Items: []string{
					"Why did you do it this way? What is your reason for this?",
					"I wonder what you were thinking about.",
					"I wish you didn’t have to parse this JSON multiple times.",
				},
			},
			Feedback{
				Base: Base{
					Name: "Avoid yes/no questions:",
				},
				Items: []string{
					"Is there a better way to do this?",
					"Did you intend this? Was this intentional?",
					"Should this method be broken up?",
				},
			},
		},
	}
)

func usage() {
	fmt.Println("USAGE: codereview [flags] options")
	fmt.Println("")
	fmt.Println("FLAGS:")
	flag.PrintDefaults()
	fmt.Println("")
	fmt.Println("OPTIONS:")
	fmt.Println("  docs                print the whole docs")
	fmt.Println("  vocabular           the code review vocabular")
	fmt.Println("  methods             review methods")
	fmt.Println("  problems            problem solving helper")
	fmt.Println("  audits              audit guideline")
	fmt.Println("  maintainability     knowledge about maintability")
	fmt.Println("  checklist           a handy codereview checklist")
	fmt.Println("  feedback            how to give feedback")
	fmt.Println("  tools               the toolkit manager")
	fmt.Println("  links               a list of useful links")
	fmt.Println("  credits             print credits")
	fmt.Println("")
	fmt.Println("Copyright (c) 2016-2018, Paul Vollmer. All rights reserved")
	fmt.Println("check the credits for a list of all sources used for this tool.")
	fmt.Println("")
}

func main() {
	flagVersion := flag.Bool("v", false, "print the version and exit")
	flag.Usage = usage
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	flagArgs := flag.Args()
	flagArgsTotal := len(flagArgs)
	if flagArgsTotal != 0 {
		switch flagArgs[0] {
		case "docs":
			docs()
			break
		case "vocabular":
			codereview.PrintVocabulars()
			break
		case "methods":
			codereview.PrintMethods()
			break
		case "problems":
			codereview.PrintProblems()
			break
		case "audits":
			codereview.PrintAudits()
			break
		case "maintainability":
			codereview.PrintMaintainability()
			break
		case "checklist":
			codereview.PrintChecklist()
			break
		case "feedback":
			codereview.PrintFeedback()
			break
		case "tools":
			codereview.PrintTools()
			break
		case "links":
			codereview.PrintLinks()
			break
		case "credits":
			fmt.Println("codereview is build by knowledge of the following sources:")
			fmt.Println("  https://en.wikipedia.org/wiki/Code_review")
			fmt.Println("  https://www.youtube.com/watch?v=pJFM321_lAs")
			fmt.Println("  https://blog.fogcreek.com/increase-defect-detection-with-our-code-review-checklist-example/")
			fmt.Println("  https://github.com/rjz/code-review-checklist")
			fmt.Println("  https://medium.com/unpacking-trunk-club/designing-awesome-code-reviews-5a0d9cd867e3")
			break
		case "version":
			fmt.Println(version)
			break
		default:
			fmt.Println("unknown option")
		}
		os.Exit(0)
	}

	// default output
	banner()
	intro()
	usage()
}

type CodeReview struct {
	Vocabulars      []Vocabular
	Methods         []Method
	Problems        []Problem
	Audits          []Audit
	Maintainability []Maintain
	Tools           []Tool
	Checklist       []Check
	Feedback        []Feedback
}

type Base struct {
	Name            string
	Description     string
	DescriptionLong []string
}

func (b *Base) Print() {
	fmt.Printf("%-20s %s\n", strings.ToUpper(b.Name), b.Description)
	if len(b.DescriptionLong) != 0 {
		fmt.Println("")
		for _, v := range b.DescriptionLong {
			fmt.Printf("  %s\n", v)
		}
		fmt.Println("")
	}
}

type Vocabular struct {
	Base
}

type Method struct {
	Base
}

type Problem struct {
	Base
}

type Audit struct {
	Base
}

type Maintain struct {
	Base
	Items []string
}

type Tool struct {
	Base
}

type Check struct {
	Base
	Items []string
}

type Feedback struct {
	Base
	Items []string
}

func banner() {
	fmt.Println("C O D E   R E V I E W")
	fmt.Println("=====================")
	fmt.Println("")
}

func intro() {
	fmt.Println("  Code review is systematic examination (sometimes referred to as peer review) of computer source code.")
	fmt.Println("  It is intended to find mistakes overlooked in the initial development phase,")
	fmt.Println("  improving the overall quality of software. Reviews are done in various forms such as pair programming,")
	fmt.Println("  informal walkthroughs, and formal inspections. (source: wikipedia)")
	fmt.Println("")
}

func docs() {
	banner()
	intro()
	codereview.PrintVocabulars()
	codereview.PrintMethods()
	codereview.PrintProblems()
	codereview.PrintAudits()
	codereview.PrintMaintainability()
	codereview.PrintTools()
	codereview.PrintChecklist()
	codereview.PrintFeedback()
	codereview.PrintLinks()
	fmt.Println("")
	fmt.Println("|||")
	fmt.Println("VVV")
	fmt.Println("")
	fmt.Println("And don't forget...")
	fmt.Println("  'Given enough eyes, all bugs are shallow' ...Linus's Law")
	fmt.Println("")
}

func (c *CodeReview) PrintVocabulars() {
	fmt.Println("V O C A B U L A R")
	fmt.Println("-----------------\n")
	for _, v := range c.Vocabulars {
		v.Print()
	}
	fmt.Println("")
}

func (c *CodeReview) PrintMethods() {
	fmt.Println("M E T H O D S")
	fmt.Println("-------------\n")
	for _, v := range c.Methods {
		v.Print()
	}
	fmt.Println("")
}

func (c *CodeReview) PrintProblems() {
	fmt.Println("P R O B L E M S")
	fmt.Println("---------------\n")
	for _, v := range c.Problems {
		v.Print()
	}
	fmt.Println("")
}

func (c *CodeReview) PrintAudits() {
	fmt.Println("A U D I T S")
	fmt.Println("-----------\n")
	for _, v := range c.Audits {
		v.Print()
	}
	fmt.Println("")
}

func (c *CodeReview) PrintMaintainability() {
	fmt.Println("M A I N T A I N A B I L I T Y")
	fmt.Println("-----------------------------\n")
	for _, v := range c.Maintainability {
		v.Print()
		// fmt.Printf("%s (%s)\n", strings.ToUpper(v.Name), v.Description)
		for _, i := range v.Items {
			fmt.Println("  - " + i)
		}
	}
	fmt.Println("")
}

func (c *CodeReview) PrintTools() {
	fmt.Println("T O O L S")
	fmt.Println("---------\n")
	for _, v := range c.Tools {
		// fmt.Println(v.Name)
		// fmt.Println("  - " + v.Description)
		v.Print()
	}
	fmt.Println("")
}

func (c *CodeReview) PrintChecklist() {
	fmt.Println("C H E C K L I S T")
	fmt.Println("-----------------\n")
	fmt.Println("  guide for clean, merge-ready pull requests...")
	fmt.Println("")
	for _, v := range c.Checklist {
		v.Print()
		for _, i := range v.Items {
			fmt.Println("  - [ ] " + i)
		}
		fmt.Println("")
	}
	fmt.Println("source: https://github.com/rjz/code-review-checklist\n")
}

func (c *CodeReview) PrintFeedback() {
	fmt.Println("F E E D B A C K")
	fmt.Println("---------------\n")
	fmt.Println("  how to give feedback")
	fmt.Println("")
	for _, v := range c.Feedback {
		v.Print()
		for _, i := range v.Items {
			fmt.Println("  - " + i)
		}
		fmt.Println("")
	}
	fmt.Println("source: https://medium.com/unpacking-trunk-club/designing-awesome-code-reviews-5a0d9cd867e3\n")
}

func (c *CodeReview) PrintLinks() {
	fmt.Println("L I N K S")
	fmt.Println("---------\n")
	fmt.Println("- https://github.com/ryanmcdermott/code-review-tips")
	fmt.Println("- https://smartbear.com/learn/code-review/best-practices-for-peer-code-review/")
	fmt.Println("- https://blog.digitalocean.com/how-to-conduct-effective-code-reviews/")
	fmt.Println("")
}
