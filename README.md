:house: Dont Miss 2 Man yen
====
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg?style=flat-square)](https://github.com/prettier/prettier)

Overview
## Description
Housing service for Money Forward employees

This service will provide apartments that qualify for housing allowance

![image](https://user-images.githubusercontent.com/21120045/71453382-56ba0800-27ce-11ea-8f83-324b11b7f602.png)

![image](https://user-images.githubusercontent.com/21120045/71453386-5a4d8f00-27ce-11ea-8de7-01ebf6869381.png)

## Features
* The eligible apartments will be displayed
* Data is updated automatically daily from [Suumo](https://suumo.jp/)
* Sortable list with search filter at homepage
* Sign in to save your favorite apartment for later viewing
* Support Japanese, English and Vietnamese

## How to find apartments that suits conditions

As stated in Money Forward's support policy Any official employee living in an apartment that meets either of the following two conditions will receive a pension
- The apartment is located within a radius of 1.5km from the office
- Travel to work by foot or by train takes less than 15 minutes

## Technologies used to solve above conditions

- The apartment is located within a radius of 1.5km from the office
    - With latitude and longitude values retrieved from suumo, we can calculate the direct distance from office using Haversine formula ([reference](https://en.wikipedia.org/wiki/Haversine_formula))

![image](https://user-images.githubusercontent.com/21120045/71454892-1d859600-27d6-11ea-94dd-c415671bc789.png)

- Travel to work by foot or by train takes less than 15 minutes
    - This problem can be easily solved using [Google Maps Distance Matrix API](https://developers.google.com/maps/documentation/distance-matrix/start).
    
![image](https://user-images.githubusercontent.com/21120045/71454949-4c9c0780-27d6-11ea-92fb-4f03de1b7759.jpg)

## Development Environment
- MacOS X Mojave 10.14

## Technology stack
***WebServer:***
- Nginx

***Backend:***
- Mysql 8.0.18
- Ruby 2.6.5
- Rails 6.0.1

***Frontend:*** 
- Node 13.1.0
- ExpressJS 
- ReactJS
- React Router 
- Redux
- Redux Saga
- Lodash

***Deploy:*** 
- Amazon Web Service EC2
## Cloud architecture

![image](https://user-images.githubusercontent.com/21120045/71454152-b31f2680-27d2-11ea-8a77-3752c0b9b98b.png)


