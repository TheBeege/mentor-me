'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('index', {
        url: '/',
        templateUrl: 'app/index/index.html',
        controller: 'IndexCtrl'
      })
      .state('find_mentor', {
        url: '/find_mentor',
        templateUrl: 'app/find_mentor/find_mentor.html',
        controller: 'IndexCtrl'
      });
  });

