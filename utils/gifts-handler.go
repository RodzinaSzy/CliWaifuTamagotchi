package utils

import (
    "fmt"
    "os"
    "encoding/json"
    "path/filepath"
)

// ==============================
// GIFTS STRUCT
// ==============================
type Gift struct {
    Name        string `json:"name"`
    Happiness   int    `json:"happiness"`
}

type GiftsFile struct {
    Gifts []Gift `json:"gifts"`
}

var cachedGifts *GiftsFile

// ==============================
// DEFAULT GIFTS
// ==============================
func DefaultGifts() *GiftsFile {
    return &GiftsFile{
        Gifts: []Gift{
            {Name: "Chocolate Bar", Happiness: 5},
			{Name: "Flower Bouquet", Happiness: 10},
            {Name: "Plushie", Happiness: 15},
			{Name: "Perfume", Happiness: 15},
            {Name: "Necklace", Happiness: 20},
            {Name: "Cute Sticker Pack", Happiness: 3},
			{Name: "Sketchbook", Happiness: 3},
        },
    }
}

// ==============================
// FILE CREATION
// ==============================
func CreateGiftsFile() error {
    configDir := filepath.Join(os.Getenv("HOME"), ".config", "cliwaifutamagotchi")
    if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
        return fmt.Errorf("failed to create config directory: %w", err)
    }

    giftsPath := filepath.Join(configDir, "gifts.json")

    if _, err := os.Stat(giftsPath); err == nil {
        return nil
    }

    file, err := os.Create(giftsPath)
    if err != nil {
        return fmt.Errorf("failed to create gifts file: %w", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(DefaultGifts()); err != nil {
        return fmt.Errorf("failed to write default gifts: %w", err)
    }

    return nil
}

// ==============================
// LOAD GIFTS
// ==============================
func LoadGifts() (*GiftsFile, error) {
    if cachedGifts != nil {
        return cachedGifts, nil
    }

    configDir := filepath.Join(os.Getenv("HOME"), ".config", "cliwaifutamagotchi")
    giftsPath := filepath.Join(configDir, "gifts.json")

    if _, err := os.Stat(giftsPath); os.IsNotExist(err) {
        if err := CreateGiftsFile(); err != nil {
            return nil, err
        }
    }

    file, err := os.Open(giftsPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open gifts file: %w", err)
    }
    defer file.Close()

    var gf GiftsFile
    if err := json.NewDecoder(file).Decode(&gf); err != nil || len(gf.Gifts) == 0 {
        // fallback to default if JSON broken or gifts list empty
        gf = *DefaultGifts()
    }

    cachedGifts = &gf
    return cachedGifts, nil
}
