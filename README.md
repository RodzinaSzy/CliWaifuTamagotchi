# 🫂 CliWaifuTamagotchi

Preview:
![Result](result.gif)
---
![Reactions](reactions.jpg)

![Repo size](https://img.shields.io/github/repo-size/HenryLoM/CliWaifuTamagotchi?color=lightgrey)
![Commits](https://img.shields.io/github/commit-activity/t/HenryLoM/CliWaifuTamagotchi/main?color=blue)
![Last commit](https://img.shields.io/github/last-commit/HenryLoM/CliWaifuTamagotchi?color=informational)
![License](https://img.shields.io/github/license/HenryLoM/CliWaifuTamagotchi?color=orange)

## 📑 Table of Contents
- [✨ Overview](#-overview)
- [🎬 Launching Process](#-launching-process)
- [🎨 Customization](#-customization)
- [📂 Project Structure](#-project-structure)
- [⚙️ Core Scripts](#-core-scripts)
    - [main.go](#maingo)
    - [utils/app-utils.go](#utilsapp-utilsgo)
    - [utils/commands-utils.go](#utilscommands-utilsgo)
    - [utils/happiness-utils.go](#utilshappiness-utilsgo)
    - [utils/palette-handler.go](#utilspalette-handlergo)
    - [utils/settings-handler.go](#utilssettings-handlergo)
    - [utils/encouragements-handler.go](#utilsencouragements-handlergo)
- [📜 Notes & Error handling](#-notes--error-handling)
- [🛐 Special thanks](#-special-thanks)

---

## ✨ Overview
CliWaifuTamagotchi is a **terminal-based tamagotchi** that:

- Renders **ASCII expressions and clothes**.
- Provides a small set of **interactions**: Encourage, Dress Up, Background Mode, Quit.
- Uses a **persistent color palette** stored in `~/.config/cliwaifutamagotchi/palette.json`.
- Uses **persistent detail settings** stored in `~/.config/cliwaifutamagotchi/settings.json`.
- Has minimal UI built using **`tview` and `tcell`**.
- Has **Vim-style navigation**: Use `h`, `j`, `k`, `l` keys for intuitive navigation and selection (Must be enabled in **settings.json**).

No tons of loops - only one function that repeats itself every 5 seconds. Everything handles and updates according to it.

---

## 🎬 Launching Process

<details>
  <summary><b>Brew</b> (macOS)</summary>

  1. **Install**

  ```bash
  brew install HenryLoM/CliWaifuTamagotchi/cliwt
  ```

  2. **Run**

  ```bash
  cliwt
  ```

  ---

</details>

<details>
  <summary><b>AUR</b> (Arch)</summary>

  1. **Install**

  ```bash
  yay -S cliwt
  ```
  or
  ```bash
  paru -S cliwt
  ```

  2. **Run**

  ```bash
  cliwt
  ```

  ---

</details>

<details>
  <summary><b>Git</b> (Source code)</summary>

  1. **Clone repository**

  ```bash
  git clone https://github.com/HenryLoM/CliWaifuTamagotchi.git
  cd CliWaifuTamagotchi
  ```

  2. **Build app yourself, then run**

  ```bash
  go build -o cliwt
  ./cliwt
  ```

  - **Or run directly for development**

  ```bash
  go run main.go
  ```

  ---

</details>

> **💡 Notes**
>
> * First run creates `~/.config/cliwaifutamagotchi/` directory and `palette.json`, `settings.json` files in it on its own if missing.
> * On macOS, ensure your terminal supports **true color** for best rendering.

---

## 🎨 Customization

1. Palette<br>
JSON file is in `~/.config/cliwaifutamagotchi/` ; Named `palette.json`<br>
JSON file's structure:
```
{
  "background": "#1e1e2e",
  "foreground": "#cdd6f4",
  "border": "#cba6f7",
  "accent": "#eba0ac",
  "title": "#b4befe"
}
```
> Note: default palette is Catppuchin (Mocha).

2. Settings<br>
JSON file is in `~/.config/cliwaifutamagotchi/` ; Named `settings.json`<br>
JSON file's structure:
```
{
  "name": "Waifu",
  "defaultMessage": "...",
  "vimNavigation": false,
  "keys": {
    "encourage": "l",
    "dressup": "2",
    "backgroundMode": "b",
    "quit": "q"
  }
}
```
> Note: try to avoid key overrides when using `"vimNavigation": true`.

---

## 📂 Project Structure

```
CliWaifuTamagotchi/
│
├── README.md
├── LICENSE
├── .gitignore
├── result.gif
├── reactions.jpg
├── go.mod
├── go.sum
├── main.go                             # Main file that launches the project
│
└── utils/
    │
    ├── ascii-arts/
    │   ├── clothes/                    # ASCII bodies
    │   └── expressions/                # ASCII heads
    │
    ├── assets/
    │   └── words-of-encouragement.txt  # List of lines for Encouragement function
    │
    ├── app-utils.go                    # Main helpers
    ├── commands-utils.go               # Functions for the Action Space
    ├── happiness-utils.go              # Happiness scoring system
    ├── palette-handler.go              # Handling palette out of the file
    ├── encouragements-handler.go       # Handling encouragements out of the file
    └── settings-handler.go             # Handling settings out of the file
```

---

## ⚙️ Core Scripts

### **main.go**

* Loads ASCII **head, blink frames, and body**.
* Displays **actions menu**: Encourage, Dress Up, Quit.
* Handles **user input** (keys and navigation).
* Queues UI updates safely using `app.QueueUpdateDraw` via `UIEventsChan` that keeps UI changes in order.

### **utils/app-utils.go**

* Helper functions for **loading ASCII files**.
* Manages **UI rendering** and **widget updates**.

### **utils/commands-utils.go**

* Implements **interactions logic**:

  * `Encourage`: random encouraging phrase + happy frame.
  * `DressUp`: swaps body/outfit based on selection.
  * `BackgroundMode`: fills the TUI with Waifu, removing all of the odd elements.
* Caches **clothes in memory** to reduce disk reads.

### **utils/happiness-utils.go**

* Handles the bar and changes emotions of the avatar.
* Handles the happiness scores.

### **utils/palette-handler.go**

* Loads palette from `~/.config/cliwaifutamagotchi/palette.json`.
* Creates **default palette** if missing.
* Provides **color application** helpers.

### **utils/settings-handler.go**

* Loads settings from `~/.config/cliwaifutamagotchi/settings.json`.
* Creates **default settings** if missing.

### **utils/encouragements-handler.go**

* Loads settings from `~/.config/cliwaifutamagotchi/words-of-encouragement.txt`.
* Restores **default encouragements** from relative directory if missing.

---

## 📜 Notes & Error handling

#### **Errors:**
* See errors you can't explain? Try to remove `~/.config/cliwaifutamagotchi/` directory (you can do a backup). If it worked, it means customization was updated and there was a conflict.
* If error remains, leave the issue, we could solve it together!

#### **Warning:**
* Missing/malformed ASCII files may cause a wrong output; handle carefully if modifying assets inside the structure.

#### **Read, if you want to contribute**
* Project lives only because there are people who use it. Let's make sure we do it for people, not to recieve anotehr achivement to out profiles.
* Keep the code clean and constructive.
* Two main goals of the project:
  * As cusomizable TUI as we can get.
  * As lightweight tool as we can create, since the project assumes users leave it on the background.

#### **Future plans:**
* More interactions (feeding, timed events, stats).
* Save selected outfit and preferences.
* Unit tests and error handling improvements.
* Husbando version ("Swap Mode").
* Maybe "Pose Mode" - loop animation or specific pose to select and have on the background.
* Maybe handle stderr so Waifu reacts to the errors you get during your work.

---

## 🛐 Special thanks

- **[sutemo](https://sutemo.itch.io/)** — for the [amazing sprites](https://sutemo.itch.io/female-character) you can see in the project.
- **[mininit](https://github.com/mininit)** — for embedding all assets and enabling a clean, pure build process.
- **[Ali Medhat](https://github.com/Alimedhat000)** — for adding Vim-style navigation.
- **[Isaac Hesslegrave](https://github.com/HeadedBranch)** — for implementing Arch Linux support via `yay` and `paru`.

---

⤴︎ Return to the [📑 Table of Contents](#-table-of-contents) ⤴︎