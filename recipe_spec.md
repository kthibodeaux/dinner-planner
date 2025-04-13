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
| `include_in_cookbook`| `bool`             | n        | Whether to include this recipe in cookbook exports. |
| `tags`               | `array of strings` | n        | Arbitrary tags like `"chicken"`, `"quick"`, etc. |
| `dependencies`       | `array of strings` | n        | Other recipes this one depends on (file names without extensions). |

---

## Parts

Define parts using the `[parts.<name>]` table. If a recipe has a single part, use `[parts.default]`.

| Field            | Type               | Required | Description |
|------------------|--------------------|----------|-------------|
| `name`           | `string`           | y        | Display name for this part. |
| `prep_time`      | `TimeUnit`         | n        | Preparation time for this part. |
| `cook_time`      | `TimeUnit`         | n        | Cooking time for this part. |
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
prep_time.amount = 10
prep_time.unit = "minutes"
cook_time.amount = 15
cook_time.unit = "minutes"
name = "Tortilla"
ingredients = [
  "16 oz; Flour",
  "2 tsp; Canola Oil"
]
steps = [
  "Combine ingredients...",
]
```

### TimeUnit Format

`prep_time` and `cook_time` are expressed as subtables using the following structure:

| Field    | Type     | Required | Description                  |
|----------|----------|----------|------------------------------|
| `amount` | `integer`| y        | Numerical time value.        |
| `unit`   | `string` | y        | Time unit (`"minutes"`, `"hours"`, etc). |

### Example:

```toml
prep_time.amount = 15
prep_time.unit = "minutes"
```
