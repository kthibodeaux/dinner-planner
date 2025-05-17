package config

import (
	"cmp"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

const defaultWebPort = "8080"

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		cfg, err := loadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		instance = cfg
	})
	return instance
}

func loadConfig() (*Config, error) {
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

	// Defaults for keys
	k := &config.Planner.Keys
	if k.Quit == "" {
		k.Quit = "ctrl+c"
	}
	if k.MainView == "" {
		k.MainView = "q"
	}
	if k.Focus == "" {
		k.Focus = "f"
	}
	if k.Help == "" {
		k.Help = "h"
	}
	if k.Recipes == "" {
		k.Recipes = "0"
	}
	if k.Down == "" {
		k.Down = "j"
	}
	if k.Up == "" {
		k.Up = "k"
	}
	if k.Day1 == "" {
		k.Day1 = "1"
	}
	if k.Day2 == "" {
		k.Day2 = "2"
	}
	if k.Day3 == "" {
		k.Day3 = "3"
	}
	if k.Day4 == "" {
		k.Day4 = "4"
	}
	if k.Day5 == "" {
		k.Day5 = "5"
	}
	if k.Day6 == "" {
		k.Day6 = "6"
	}
	if k.Day7 == "" {
		k.Day7 = "7"
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
	weekdayMap := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}

	targetDay, ok := weekdayMap[strings.ToLower(weekdayName)]
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
