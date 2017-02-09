'use strict'
var smsApp = angular.module('smsApp', ['ngRoute']);

smsApp.config(function($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'assets/template/main.html',
            controller: 'smsController'
        });

});

