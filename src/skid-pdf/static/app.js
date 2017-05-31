(function () {
    var app = angular.module('skidpdf', []);
    app.Root = '/pdf';

    app.controller('skidpdfcontrol', function ($scope, $rootScope, $http) {

        $scope.pdfRequest = {
            "url": "",
            "data": "",
            "grayscale": false,
            "landscape": false,
            "headers": {},
            "postParams": {}
        };

        $scope.basePdfRequest = {
            "url": "",
            "data": "",
            "grayscale": false,
            "landscape": false,
            "headers": {},
            "postParams": {}
        };

        $scope.addHeader = function () {};


        $scope.addPostData = function () {};

        $scope.makeRequest = function () {
            $http.get($scope.pdfRequest.url, {
                grayscale: $scope.pdfRequest.grayscale,
                landscape: $scope.pdfRequest.landscape
            }, null);
            console.log($scope.pdfRequest);
        };

        $scope.resetForm = function () {
            $scope.pdfRequest = $scope.basePdfRequest;
        };

        $scope.formSelected = "simpleGet";
    });

})();