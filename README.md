# 🥘 Dinner Planner

**Dinner Planner** is a local-first, terminal-based meal planning and recipe management tool built for home cooks who love structure, flexibility, and staying off the cloud. Recipes are stored as plain TOML files, editable in any text editor, and the app provides a powerful TUI for planning meals, generating shopping lists, and organizing your cooking week—all without a backend.

Optionally, the app can generate static HTML cookbooks and printable meal plans to be viewed on any device or stuck to your fridge.

---

## 💾 Local TOML-Based Recipe Storage

- Recipes live as individual `.toml` files in a directory on your machine. Perfect for using Git for cloud backups.
- Each recipe includes metadata like prep/cook time, ingredients, steps, and tags.
- Recipes can reference other recipes (e.g. “BBQ Chicken” can optionally include “Habanero BBQ Sauce”).
- Fully offline and version-control-friendly.

---

## 🧠 Terminal Meal Planning (TUI)

- The app uses a [gocui](https://github.com/jroimartin/gocui)-based terminal interface to:
  - Search by name, tag, or time constraints.
  - Filter meals (e.g. “chicken under 1 hour”).
  - Build and save weekly meal plans.
- Tracks when meals were last cooked.
- Handles prep scheduling (e.g., “make pizza dough the day before pizza night”).

---

## 🛒 Smart Shopping Lists

- Automatically generate grocery lists from your selected meal plan.
- Combines ingredients across all recipes.
- Can exclude always-available items (like water or salt).
- Includes ingredients from nested dependencies when desired.

---

## 🧩 Optional Dependencies Between Recipes

- Recipes can depend on others (e.g., a sauce or side).
- Dependencies are optional per plan—sometimes you want to go the extra mile, sometimes not.
- Prep and ingredients from dependencies are integrated into the plan and shopping list when selected.

---

## 🖨️ HTML Output (Optional)

- **Fridge-Ready Weekly Meal Plans**: Generates a printable HTML sheet that includes each day’s meals, prep tasks, and notes.
- **Cookbook Site Generator**: Creates a basic, navigable static site of all your recipes. Can be viewed locally or hosted (e.g., in an NGINX Docker container on your home server).

---

## ❤️ Philosophy

Dinner Planner is for people who like **full control**, **plain text**, and **staying local**. There’s no backend, no syncing, no cloud - just your recipes, your meal plans, and your kitchen. It’s tech that gets out of your way, supports your routines, and leaves you with something you can print, host, or just keep in your repo.
