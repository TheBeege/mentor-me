'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('become-mentor', {
        url: '/become-mentor',
        templateUrl: 'app/become-mentor/become-mentor.html',
        controller: 'BecomeMentorCtrl'
      });
  });