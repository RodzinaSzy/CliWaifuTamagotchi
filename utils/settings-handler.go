package utils

import (
    "os"
    "fmt"
    "encoding/json"
    "path/filepath"
)

// ==============================
// SETTINGS STRUCT
// ==============================
type KeyBindings struct {
    Encourage      string `json:"encourage"`
    Gift           string `json:"gift"`
    DressUp        string `json:"dressup"`
    PoseMode       string `json:"poseMode"`
    BackgroundMode string `json:"backgroundMode"`
    SwapGender     string `json:"swapGender"`
    Quit           string `json:"quit"`
}

type Settings struct {
    Name           string      `json:"name"`
    DefaultMessage string      `json:"defaultMessage"`
    VimNavigation  bool        `json:"vimNavigation"`
    AvatarType     string      `json:"avatarType"`
    Keys           KeyBindings `json:"keys"`
}

var cachedSettings *Settings

// ==============================
// DEFAULT SETTINGS
// ==============================
func DefaultSettings() *Settings {
    return &Settings{
        Name:           "Waifu",
        DefaultMessage: "...",
        VimNavigation:  false,
        AvatarType:     "waifu",
        Keys: KeyBindings{
            Encourage:      "1",
            Gift:           "2",
            DressUp:        "3",
            PoseMode:       "4",
            BackgroundMode: "b",
            SwapGender:     "s",
            Quit:           "q",
        },
    }
}

// ==============================
// FILE HANDLING
// ==============================
func CreateSettingsFile() error {
    configDir := filepath.Join(os.Getenv("HOME"), ".config", "cliwaifutamagotchi")
    if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
        return fmt.Errorf("failed to create config directory: %w", err)
    }

    settingsPath := filepath.Join(configDir, "settings.json")
    if _, err := os.Stat(settingsPath); err == nil {
        return nil
    }

    file, err := os.Create(settingsPath)
    if err != nil {
        return fmt.Errorf("failed to create settings file: %w", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(DefaultSettings()); err != nil {
        return fmt.Errorf("failed to write settings file: %w", err)
    }

    return nil
}

// LoadSettings loads settings from file (or default if missing), cached for session
func LoadSettings() (*Settings, error) {
    if cachedSettings != nil {
        return cachedSettings, nil
    }

    configDir := filepath.Join(os.Getenv("HOME"), ".config", "cliwaifutamagotchi")
    settingsPath := filepath.Join(configDir, "settings.json")

    if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
        if err := CreateSettingsFile(); err != nil {
            return nil, err
        }
    }

    file, err := os.Open(settingsPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open settings file: %w", err)
    }
    defer file.Close()

    var s Settings
    if err := json.NewDecoder(file).Decode(&s); err != nil {
        s = *DefaultSettings()
    }

    cachedSettings = &s
    return cachedSettings, nil
}

// ==============================
// GLOBAL ASCII BASE PATH
// ==============================
func GetBasePath() string {
    s, err := LoadSettings()
    if err != nil {
        // fallback
        return "ascii-arts/waifu"
    }
    if s.AvatarType == "waifu" {
        return "ascii-arts/waifu"
    }
    if s.AvatarType == "husbando" {
        return "ascii-arts/husbando"
    }
    // fallback
    return "ascii-arts/waifu"
}