'use strict';

describe('Controller: FindMentorCtrl', function () {

  // load the controller's module
  beforeEach(module('mentorMeApp'));

  var FindMentorCtrl, scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    FindMentorCtrl = $controller('FindMentorCtrl', {
      $scope: scope
    });
  }));

  it('should ...', function () {
    expect(1).toEqual(1);
  });
});
