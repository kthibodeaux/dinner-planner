package recipe

import (
	"testing"
)

func TestNewRecipe(t *testing.T) {
	part := &Part{}
	recipe := NewRecipe("Pasta", "Main Course", "Italian Cookbook", []Part{*part})

	if recipe.Name != "Pasta" {
		t.Errorf("Expected name 'Pasta', got '%s'", recipe.Name)
	}

	if recipe.CookbookCategory != "Main Course" {
		t.Errorf("Expected cookbookCategory 'Main Course', got '%s'", recipe.CookbookCategory)
	}

	if recipe.Source != "Italian Cookbook" {
		t.Errorf("Expected source 'Italian Cookbook', got '%s'", recipe.Source)
	}

	if len(recipe.Parts) != 1 {
		t.Errorf("Expected 1 part, got %d", len(recipe.Parts))
	}
}

func TestAddPart(t *testing.T) {
	recipe := &Recipe{}

	newPart := Part{
		CookTime:    TimeUnit{Amount: 30, Unit: "minutes"},
		Ingredients: []string{"1/2 cup; greek yogurt", "1/2 cup; coconut milk", "1/2 tsp; salt"},
		Name:        "Garnish",
		PrepTime:    TimeUnit{Amount: 1, Unit: "hours"},
		Steps:       []string{"Mix yogurt, coconut milk, and salt for garnish."},
	}

	recipe.AddPart(newPart)

	if len(recipe.Parts) != 1 {
		t.Fatalf("Expected 1 part, got %d", len(recipe.Parts))
	}

	if recipe.Parts[0].Name != newPart.Name {
		t.Errorf("Expected part name %s, got %s", newPart.Name, recipe.Parts[0].Name)
	}

	for i, ingredient := range newPart.Ingredients {
		if recipe.Parts[0].Ingredients[i] != ingredient {
			t.Errorf("Expected ingredient %s, got %s", ingredient, recipe.Parts[0].Ingredients[i])
		}
	}

	for i, step := range newPart.Steps {
		if recipe.Parts[0].Steps[i] != step {
			t.Errorf("Expected step %s, got %s", step, recipe.Parts[0].Steps[i])
		}
	}
}
