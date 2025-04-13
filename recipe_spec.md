# TOML Recipe Format Specification

This document outlines the structure of TOML files used by the recipe management app. Each file represents a single recipe.

---

## Top-Level Fields

These fields are declared at the top level of each TOML file:

| Field                 | Type               | Required | Description |
|----------------------|--------------------|----------|-------------|
| `name`               | `string`           | y        | The recipe's display name. |
| `cookbook_category`  | `string`           | n        | High-level cookbook grouping (e.g., `Entrees`, `Desserts`). |
| `source`             | `string`           | n        | URL or book reference for the recipe. |
| `tags`               | `array of strings` | n        | Arbitrary tags like `"chicken"`, `"quick"`, etc. |
| `dependencies`       | `array of strings` | n        | Other recipes this one depends on (file names without extensions). |

---

## Parts

Define parts using the `[parts.<name>]` table. If a recipe has a single part, use `[parts.default]`.

| Field            | Type               | Required | Description |
|------------------|--------------------|----------|-------------|
| `name`           | `string`           | y        | Display name for this part. |
| `prep_time`      | `int`              | n        | Preparation time in minutes for this part. |
| `cook_time`      | `int`              | n        | Cooking time in minutes for this part. |
| `ingredients`    | `array of strings` | n        | List of ingredients specific to this part. Each string can optionally include a quantity and unit, separated by `;`. |
| `steps`          | `array of strings` | n        | Step-by-step instructions specific to this part. |

### Example:
```toml
[parts.chicken]
name = "Chicken"
ingredients = [
  "2 8oz; Chicken Breast",
  "2 tsp; Olive Oil"
]
steps = [
  "Marinate the chicken...",
  "Bake at 450°F for 12–15 minutes..."
]

[parts.tortilla]
cook_time = 10
prep_time = 10
name = "Tortilla"
ingredients = [
  "16 oz; Flour",
  "2 tsp; Canola Oil"
]
steps = [
  "Combine ingredients...",
]
```
