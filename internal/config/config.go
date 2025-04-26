package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type config struct {
	Keys            KeyConfig     `toml:"keys"`
	Planner         PlannerConfig `toml:"planner"`
	RecipeDirectory string        `toml:"recipes"`
	StartDate       string
	Web             WebConfig `toml:"web"`
}

type KeyConfig struct {
	Quit string `toml:"quit"`
}

type PlannerConfig struct {
	FirstDayOfWeek string `toml:"first_day_of_week"`
}

type WebConfig struct {
	Port string `toml:"port"`
}

var (
	flagConfigFilePath  string
	flagRecipeDirectory string
	flagStartDate       string
	flagWebPort         string
)

const (
	defaultWebPort = "8080"
)

func LoadConfig() *config {
	flag.StringVar(&flagConfigFilePath, "config-file", defaultConfigFilePath(), "Path to the configuration file")
	flag.StringVar(&flagRecipeDirectory, "recipes", "", "Path to the recipes")
	flag.StringVar(&flagStartDate, "date", "", "Start date for the week")
	flag.StringVar(&flagWebPort, "port", "", "Port to serve the web application on")

	flag.Parse()

	config := &config{}
	if fileExists(flagConfigFilePath) {
		log.Println("Loading config file:", flagConfigFilePath)
		data, err := os.ReadFile(flagConfigFilePath)
		if err != nil {
			log.Fatal(err)
		}
		if err := toml.Unmarshal(data, config); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Config file does not exist:", flagConfigFilePath)
	}

	config.RecipeDirectory = expandHome(config.RecipeDirectory)

	if flagRecipeDirectory != "" {
		config.RecipeDirectory = flagRecipeDirectory
	}
	if flagStartDate != "" {
		config.StartDate = flagStartDate
	}
	if flagWebPort != "" {
		config.Web.Port = flagWebPort
	}

	if config.Planner.FirstDayOfWeek == "" {
		config.Planner.FirstDayOfWeek = "sunday"
	}
	if config.RecipeDirectory == "" {
		config.RecipeDirectory = expandHome("~/recipes")
	}
	if config.StartDate == "" {
		config.StartDate = startOfWeek(time.Now(), config.Planner.FirstDayOfWeek)
	}
	if config.Web.Port == "" {
		config.Web.Port = defaultWebPort
	}

	if config.Web.Port[0] != ':' {
		config.Web.Port = ":" + config.Web.Port
	}

	if config.Keys.Quit == "" {
		config.Keys.Quit = "ctrl+c"
	}

	return config
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func defaultConfigFilePath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		xdgConfigHome = filepath.Join(homePath, ".config")
	}
	return filepath.Join(xdgConfigHome, "dinner-planner.toml")
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(home, path[1:])
	}
	return path
}

func startOfWeek(currentDate time.Time, weekdayName string) string {
	weekdayName = strings.ToLower(weekdayName)

	weekdayMap := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}

	targetDay, ok := weekdayMap[weekdayName]
	if !ok {
		log.Fatal("invalid weekday name:", weekdayName)
	}

	offset := (int(currentDate.Weekday()) - int(targetDay) + 7) % 7
	date := currentDate.AddDate(0, 0, -offset)

	return date.Format("2006-01-02")
}
