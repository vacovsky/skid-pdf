(function () {
    var app = angular.module('skidpdf', []);
    app.Root = '/pdf';

    app.controller('skidpdfcontrol', function ($scope, $http) {

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
            if ($scope.formSelected == "simpleGET") {
                $http.get("/pdf?uri=" + $scope.pdfRequest.url + "&grayscale=" + $scope.pdfRequest.grayscale + "&landscape=" + $scope.pdfRequest.landscape).then(function (response) {
                    var file = new Blob([response.data], {
                        type: 'application/pdf'
                    });
                    var fileURL = URL.createObjectURL(file);
                    var a = document.createElement('a');
                    a.href = fileURL;
                    a.target = '_blank';
                    a.download = $scope.formSelected + ".pdf";
                    document.body.appendChild(a);
                    a.click();
                });
            } else if ($scope.formSelected == "complexGET") {

            } else if ($scope.formSelected == "complexPOST") {

            }


            console.log($scope.pdfRequest);
        };

        $scope.resetForm = function () {
            $scope.pdfRequest = $scope.basePdfRequest;
        };

        $scope.formSelected = "simpleGet";
    });

})();