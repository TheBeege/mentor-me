'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('index', {
        url: '/',
        templateUrl: 'app/index/index.html',
        controller: 'IndexCtrl'
      });
  });

/* in your javascript */
App.controller('home', function (page) {
  // this runs whenever a 'home' page is loaded
  // 'page' is the HTML app-page element
  $(page)
    .find('.app-button')
    .on('click', function () {
      console.log('button was clicked!');
    });
});