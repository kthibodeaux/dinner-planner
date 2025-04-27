# 🥘 Dinner Planner

**Dinner Planner** is a local-first, terminal-based meal planning and recipe management tool built for home cooks who love structure, flexibility, and staying off the cloud. Recipes are stored as plain TOML files, editable in any text editor, and the app provides a powerful TUI for planning meals and generating shopping lists - all without a backend.

---

## 💾 Local TOML-Based Recipe Storage

- Recipes live as individual `.toml` files in a directory on your machine. Perfect for using Git.
- Each recipe includes metadata like prep/cook time, ingredients, steps, and tags.
- Recipes can reference other recipes (e.g. “BBQ Chicken” can optionally include “Habanero BBQ Sauce”).
- Fully offline and version control friendly.
- Optionally run the web server `cmd/web` to serve recipes as a local web app.

---

## 🧠 Terminal Meal Planning (TUI)

- [bubbletea](https://github.com/charmbracelet/bubbletea)-based terminal interface to build a weekly dinner plan.

---

## ❤️ Philosophy

Dinner Planner is for people who like **full control**, **plain text**, and **staying local**. There's no backend, no syncing, no cloud - just your recipes, your meal plans, and your kitchen. It's tech that gets out of your way, supports your routines, and leaves you with something you can print, host, or just keep in your repo.
