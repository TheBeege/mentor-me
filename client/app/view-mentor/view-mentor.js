'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('view-mentor', {
        url: '/view-mentor?mentorID',
        templateUrl: 'app/view-mentor/view-mentor.html',
        controller: 'ViewMentorCtrl'
      });
  });
