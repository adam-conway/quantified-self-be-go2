# README

This application is a back-end API that serves up data for a JavaScript front end, found [here](https://github.com/turingschool/qs-frontend-starter).

This API has been deployed to heroku [here](https://radiant-bastion-95270.herokuapp.com/). The link for the application accessing this API can be found below:

[Quantified Self](http://warlike-hour.surge.sh/)<br>

This API has the following end-points:

## Foods
`GET /api/v1/foods`<br>
Returns all foods

`GET /api/v1/foods/:id`<br>
Returns a single food

`POST /api/v1/foods`<br>
Creates a new food

`DELETE /api/v1/foods/:id`<br>
Deletes a food

`PATCH /api/v1/foods/:id`<br>
Updates a food

## Meal Endpoints
`GET /api/v1/meals`<br>
Returns all meals (and foods)

`GET /api/v1/meals/:meal_id/foods`<br>
Returns all foods for a single meal

`POST /api/v1/meals/:meal_id/foods/:id`<br>
Adds a food to a meal

`DELETE /api/v1/meals/:meal_id/foods/:id`<br>
Removes a food from a meal
