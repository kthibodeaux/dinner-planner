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
	Color             string    `toml:"color"`
	FirstDayOfWeek    string    `toml:"first_day_of_week"`
	IgnoreIngredients []string  `toml:"ignore_ingredients"`
	Keys              KeyConfig `toml:"keys"`
	ScrollOffset      int       `toml:"scroll_offset"`
	ScrollAmount      int       `toml:"scroll_amount"`
}

type KeyConfig struct {
	MainView           string `toml:"main_view"`
	Focus              string `toml:"focus"`
	Help               string `toml:"help"`
	ShoppingList       string `toml:"shopping_list"`
	Recipes            string `toml:"recipes"`
	Down               string `toml:"down"`
	Up                 string `toml:"up"`
	ScrollDown         string `toml:"scroll_down"`
	ScrollUp           string `toml:"scroll_up"`
	Day1               string `toml:"day_1"`
	Day2               string `toml:"day_2"`
	Day3               string `toml:"day_3"`
	Day4               string `toml:"day_4"`
	Day5               string `toml:"day_5"`
	Day6               string `toml:"day_6"`
	Day7               string `toml:"day_7"`
	Quit               string `toml:"quit"`
	ShoppingListToggle string `toml:"shopping_list_toggle"`
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

	defaultIfEmpty(&config.Planner.Color, "#3A5B7E")
	defaultIfEmpty(&config.Planner.FirstDayOfWeek, "sunday")
	defaultIfEmpty(&config.RecipeDirectory, expandHome("~/recipes"))
	if config.StartDate == "" {
		config.StartDate, err = startOfWeek(time.Now(), config.Planner.FirstDayOfWeek)
		if err != nil {
			return nil, err
		}
	}
	defaultIfEmpty(&config.Web.Port, defaultWebPort)
	if config.Web.Port[0] != ':' {
		config.Web.Port = ":" + config.Web.Port
	}

	defaultIntIfEmpty(&config.Planner.ScrollOffset, 3)
	defaultIntIfEmpty(&config.Planner.ScrollAmount, 15)

	if len(config.Planner.IgnoreIngredients) == 0 {
		config.Planner.IgnoreIngredients = append(config.Planner.IgnoreIngredients, "water")
	}

	defaultIfEmpty(&config.Planner.Keys.Quit, "ctrl+c")
	defaultIfEmpty(&config.Planner.Keys.MainView, "q")
	defaultIfEmpty(&config.Planner.Keys.Focus, "f")
	defaultIfEmpty(&config.Planner.Keys.Help, "h")
	defaultIfEmpty(&config.Planner.Keys.ShoppingList, "s")
	defaultIfEmpty(&config.Planner.Keys.Recipes, "0")
	defaultIfEmpty(&config.Planner.Keys.Down, "j")
	defaultIfEmpty(&config.Planner.Keys.Up, "k")
	defaultIfEmpty(&config.Planner.Keys.ScrollDown, "ctrl+d")
	defaultIfEmpty(&config.Planner.Keys.ScrollUp, "ctrl+u")
	defaultIfEmpty(&config.Planner.Keys.Day1, "1")
	defaultIfEmpty(&config.Planner.Keys.Day2, "2")
	defaultIfEmpty(&config.Planner.Keys.Day3, "3")
	defaultIfEmpty(&config.Planner.Keys.Day4, "4")
	defaultIfEmpty(&config.Planner.Keys.Day5, "5")
	defaultIfEmpty(&config.Planner.Keys.Day6, "6")
	defaultIfEmpty(&config.Planner.Keys.Day7, "7")
	defaultIfEmpty(&config.Planner.Keys.ShoppingListToggle, "t")

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

func defaultIfEmpty(val *string, def string) {
	if *val == "" {
		*val = def
	}
}

func defaultIntIfEmpty(val *int, def int) {
	if *val == 0 {
		*val = def
	}
}
