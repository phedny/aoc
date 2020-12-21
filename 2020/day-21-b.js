const fs = require('fs');

const ingredientsPerAllergen = {};
const allIngredients = fs.readFileSync('./inputs/21/real.txt', { encoding: 'utf8' })
	.split('\n')
	.map(line => line.match(/([a-z ]+) \(contains ([a-z, ]+)/))
	.map(([_, ingredients, allergens]) => [ingredients.split(' '), allergens.split(', ')])
	.flatMap(([ingredients, allergens]) => {
		allergens.forEach(allergen => {
			if (ingredientsPerAllergen.hasOwnProperty(allergen)) {
				ingredientsPerAllergen[allergen].push(ingredients);
			} else {
				ingredientsPerAllergen[allergen] = [ingredients];
			}
		});
		return ingredients;
	});

const ingredientForAllergen = {};
while (Object.keys(ingredientForAllergen).length < Object.keys(ingredientsPerAllergen).length) {
	Object.entries(ingredientsPerAllergen).forEach(([allergen, ingredients]) => {
		const ingredientOptions = [...new Set(ingredients.flat())]
			.filter(ingredient => !Object.values(ingredientForAllergen).includes(ingredient))
			.filter(ingredient => ingredients.every(list => list.includes(ingredient)));
		if (ingredientOptions.length === 1) {
			ingredientForAllergen[allergen] = ingredientOptions[0];
		}
	});
}

const result = Object.keys(ingredientForAllergen)
	.sort((a, b) => a.localeCompare(b))
	.map(allergen => ingredientForAllergen[allergen])
	.join(',');

console.log(result);
