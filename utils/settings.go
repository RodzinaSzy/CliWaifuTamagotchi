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
type Settings struct {
	Name           string `json:"name"`
	DefaultMessage string `json:"defaultMessage"`
}

var cachedSettings *Settings

// ==============================
// DEFAULT SETTINGS
// ==============================
func DefaultSettings() *Settings {
	return &Settings{
		Name:           "Waifu",
		DefaultMessage: "...",
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
		return nil, fmt.Errorf("failed to decode settings file: %w", err)
	}

	cachedSettings = &s
	return cachedSettings, nil
}
