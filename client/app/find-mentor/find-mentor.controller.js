'use strict';

// TODO: use ui.select (angular-ui/ui-select) to do better tagging stuff
angular.module('mentorMeApp')
  .controller('FindMentorCtrl', function ($scope, $http, socket) {

    $http.get('/api/mentors').success(function(mentors) {
      $scope.mentors = mentors;
      console.log("mentors:");
      console.log($scope.mentors);
      //socket.syncUpdates('thing', $scope.mentors);
    });

    $scope.updateMentorList = function() {
      console.log("tag inputs:");
      console.log($scope.tagInputs.replace(" ","").split(","));
      $http({
        url: '/api/mentors/bytag',
        method: 'GET',
        params: {tags: $scope.tagInputs.replace(" ","").split(",")}
      }).success(function(mentors) {
        $scope.mentors = mentors;
      });
    };
  });
