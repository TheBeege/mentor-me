'use strict';

angular.module('mentorMeApp')
  .config(function ($stateProvider) {
    $stateProvider
      .state('no-permissions', {
        url: '/no-permissions',
        templateUrl: 'app/no-permissions/no-permissions.html',
        controller: 'NoPermissionsCtrl'
      });
  });
