'use strict';

angular.module('mentorMeApp')
  .controller('BecomeMentorCtrl', function ($scope) {
    $scope.message = 'Hello';

    $scope.addMentor = function() {
      if($scope.newMentor === {}) {
        return;
      }
      $http.post('/api/mentor', $scope.newMentor);
      $scope.newMentor = {};
    };
  });
