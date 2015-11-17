'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('find-mentor', {
        url: '/find-mentor',
        templateUrl: 'app/find-mentor/find-mentor.html',
        controller: 'FindMentorCtrl'
      });
  });
