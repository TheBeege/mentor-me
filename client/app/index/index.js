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
