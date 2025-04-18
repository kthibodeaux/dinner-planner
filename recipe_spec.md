# TOML Recipe Format Specification

This document outlines the structure of TOML files used by the recipe management app. Each file represents a single recipe.

---

## Top-Level Fields

These fields are declared at the top level of each TOML file:

| Field                 | Type               | Required | Description |
|----------------------|--------------------|----------|-------------|
| `name`               | `string`                | y        | The recipe's display name. |
| `cookbook_category`  | `string`                | n        | High-level cookbook grouping (e.g., `Entrees`, `Desserts`). |
| `source`             | `string`                | n        | URL or book reference for the recipe. |
| `tags`               | `array of strings`      | n        | Arbitrary tags like `"chicken"`, `"quick"`, etc. |
| `dependencies`       | `array of dependencies` | n        | Other recipes this one depends on. |

---

## Parts

| Field            | Type               | Required | Description |
|------------------|--------------------|----------|-------------|
| `name`           | `string`           | y        | Display name for this part. |
| `prep_time`      | `int`              | n        | Preparation time in minutes for this part. |
| `cook_time`      | `int`              | n        | Cooking time in minutes for this part. |
| `ingredients`    | `array of strings` | n        | List of ingredients specific to this part. Each string can optionally include a quantity and unit, separated by `;`. |
| `steps`          | `array of strings` | n        | Step-by-step instructions specific to this part. |

### Example:
```toml
[[parts]]
name = "Chicken"
ingredients = [
  "2 8oz; Chicken Breast",
  "2 tsp; Olive Oil"
]
steps = [
  "Marinate the chicken...",
  "Bake at 450°F for 12–15 minutes..."
]

[[parts]]
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

---

## Dependencies

| Field            | Type               | Required | Description |
|------------------|--------------------|----------|-------------|
| `recipe_id`      | `string`           | y        | Filename of the dependency (without extension). |
| `is_required`    | `bool`             | n        | If the dependency is a requirement for the recipe. Default false |

### Example:
```toml
name = "Chicken Fried Steak"

[[dependencies]]
recipe_id = "white-pepper-gravy"
is_required = true

[[dependencies]]
recipe_id = "homemade-bread-crumbs"

```
[[parts]]
ingredients = [
  "2; Cube Steak",
  "Bread Crumbs",
]
```
