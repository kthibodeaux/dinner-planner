package config

import (
	"cmp"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Planner         PlannerConfig `toml:"planner"`
	RecipeDirectory string        `toml:"recipes"`
	StartDate       string
	Web             WebConfig `toml:"web"`
}

type PlannerConfig struct {
	Color          string    `toml:"color"`
	FirstDayOfWeek string    `toml:"first_day_of_week"`
	Keys           KeyConfig `toml:"keys"`
}

type KeyConfig struct {
	MainView string `toml:"main_view"`
	Focus    string `toml:"focus"`
	Help     string `toml:"help"`
	Recipes  string `toml:"recipes"`
	Down     string `toml:"down"`
	Up       string `toml:"up"`
	Day1     string `toml:"day_1"`
	Day2     string `toml:"day_2"`
	Day3     string `toml:"day_3"`
	Day4     string `toml:"day_4"`
	Day5     string `toml:"day_5"`
	Day6     string `toml:"day_6"`
	Day7     string `toml:"day_7"`
	Quit     string `toml:"quit"`
}

type WebConfig struct {
	Port string `toml:"port"`
}

const (
	defaultWebPort = "8080"
)

func LoadConfig() (*Config, error) {
	config := &Config{}

	var (
		flagConfigFilePath  string
		flagRecipeDirectory string
		flagStartDate       string
		flagWebPort         string
	)

	defaultConfigFilePath, err := defaultConfigFilePath()
	if err != nil {
		return nil, err
	}

	flag.StringVar(&flagConfigFilePath, "config-file", defaultConfigFilePath, "Path to the configuration file")
	flag.StringVar(&flagRecipeDirectory, "recipes", "", "Path to the recipes")
	flag.StringVar(&flagStartDate, "date", "", "Start date for the week")
	flag.StringVar(&flagWebPort, "port", "", "Port to serve the web application on")

	flag.Parse()

	if fileExists(flagConfigFilePath) {
		log.Println("Loading config file:", flagConfigFilePath)
		data, err := os.ReadFile(flagConfigFilePath)
		if err != nil {
			return nil, err
		}
		if err := toml.Unmarshal(data, config); err != nil {
			return nil, err
		}
	} else {
		log.Println("Config file does not exist:", flagConfigFilePath)
	}

	config.RecipeDirectory = expandHome(config.RecipeDirectory)

	if flagRecipeDirectory != "" {
		config.RecipeDirectory = flagRecipeDirectory
	}
	config.StartDate = cmp.Or(flagStartDate, config.StartDate)
	if flagWebPort != "" {
		config.Web.Port = flagWebPort
	}

	if config.Planner.Color == "" {
		config.Planner.Color = "#3A5B7E"
	}
	if config.Planner.FirstDayOfWeek == "" {
		config.Planner.FirstDayOfWeek = "sunday"
	}
	if config.RecipeDirectory == "" {
		config.RecipeDirectory = expandHome("~/recipes")
	}
	if config.StartDate == "" {
		config.StartDate, err = startOfWeek(time.Now(), config.Planner.FirstDayOfWeek)
		if err != nil {
			return nil, err
		}
	}
	if config.Web.Port == "" {
		config.Web.Port = defaultWebPort
	}

	if config.Web.Port[0] != ':' {
		config.Web.Port = ":" + config.Web.Port
	}

	if config.Planner.Keys.Quit == "" {
		config.Planner.Keys.Quit = "ctrl+c"
	}
	if config.Planner.Keys.MainView == "" {
		config.Planner.Keys.MainView = "q"
	}
	if config.Planner.Keys.Focus == "" {
		config.Planner.Keys.Focus = "f"
	}
	if config.Planner.Keys.Help == "" {
		config.Planner.Keys.Help = "h"
	}
	if config.Planner.Keys.Recipes == "" {
		config.Planner.Keys.Recipes = "0"
	}
	if config.Planner.Keys.Down == "" {
		config.Planner.Keys.Down = "j"
	}
	if config.Planner.Keys.Up == "" {
		config.Planner.Keys.Up = "k"
	}
	if config.Planner.Keys.Day1 == "" {
		config.Planner.Keys.Day1 = "1"
	}
	if config.Planner.Keys.Day2 == "" {
		config.Planner.Keys.Day2 = "2"
	}
	if config.Planner.Keys.Day3 == "" {
		config.Planner.Keys.Day3 = "3"
	}
	if config.Planner.Keys.Day4 == "" {
		config.Planner.Keys.Day4 = "4"
	}
	if config.Planner.Keys.Day5 == "" {
		config.Planner.Keys.Day5 = "5"
	}
	if config.Planner.Keys.Day6 == "" {
		config.Planner.Keys.Day6 = "6"
	}
	if config.Planner.Keys.Day7 == "" {
		config.Planner.Keys.Day7 = "7"
	}

	return config, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func defaultConfigFilePath() (string, error) {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		xdgConfigHome = filepath.Join(homePath, ".config")
	}
	return filepath.Join(xdgConfigHome, "dinner-planner.toml"), nil
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

func startOfWeek(currentDate time.Time, weekdayName string) (string, error) {
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
		return "", fmt.Errorf("invalid weekday name: %s", weekdayName)
	}

	offset := (int(currentDate.Weekday()) - int(targetDay) + 7) % 7
	date := currentDate.AddDate(0, 0, -offset)

	return date.Format("2006-01-02"), nil
}

func (c *Config) DayKeyMap() map[int]string {
	return map[int]string{
		0: c.Planner.Keys.Day1,
		1: c.Planner.Keys.Day2,
		2: c.Planner.Keys.Day3,
		3: c.Planner.Keys.Day4,
		4: c.Planner.Keys.Day5,
		5: c.Planner.Keys.Day6,
		6: c.Planner.Keys.Day7,
	}
}
