var countryApp = angular.module('countryApp',['ngRoute','restangular']);

/**
*Create route function 
*/

countryApp.config(function($routeProvider) {
	$routeProvider
	.when('/', {templateUrl: 'list.html',
		controller: 'countryController'
	})
	.when('/add', {
		templateUrl: '/detail.html',
		controller: 'countryEditController',
	})
	.when('/edit/:id', {
		templateUrl: '/detail.html',
		controller: 'countryEditController',		 

	}
	)
	.otherwise({redirectTo:'/'});
})


//create controllers to hold controllers 
//var Controllers = {};

/**
*create controller
*/


countryController = countryApp.controller('countryController',function($scope,$http,Restangular,$location, $routeParams,$route)
{

	$scope.headers = [
	// { title: 'S.No.',       value: 'id' },
	{ title: 'Country Name',       value: 'Name' },
	{ title: 'Continent', value: 'Continent' },
	{ title: 'Region',     value: 'Region'},
	{ title: 'Population',     value: 'Population'},
	{ title: 'Edit'},	
	];
	$scope.newdata = {};
	$scope.search = {};
	// $scope.alert = {title: 'Holy guacamole!', content: 'Best check yo self, you\'re not looking too good.', type: 'info'};
	//intialize status of save button on details page 
	
	$scope.sortvalue= '=';
	$scope.sortedby= '=';

	//$scope.name="";
	
	//$scope.names=[{name:'user1',city:'city45'},{name:'user12',city:'city98'},{name:'user98',city:'city2'}];	
	//UserFactory.getnames().success(function(data) {$scope.names = data;});	
	//RestangularProvider.setBaseUrl('')			
	
	//$scope.insurers=data.data;	
	//if (data.recordcount) {$scope.recordcount=data.recordcount; $scope.numpages = $scope.recordcount/10};		
$scope.add_new = function(){
			$location.path('/add')
		}	
	$scope.fetchResult = function () {
		$scope.showloading = true;
		$scope.insurers= {};

		return Restangular.all('json/all').getList($scope.search).then(function (data) {		

			if (data) 
			{				
				$scope.countries = data;
				// console.log($scope.countries)
				$scope.records_per_page=10;
				$scope.numPages = (data.recordcount/$scope.records_per_page)-1;
				$scope.recordcount = data.recordcount;		  
				$scope.showloading = false;  	  

			}
		});	
	};
	
	// no more need using ng-change
	// $scope.$watchCollection('search.pageNumber', function(newVal, oldVal) {
 //            $scope.search.pageNumber = newVal;    	
	// 		$scope.fetchResult();
 //          });

$scope.selectPage = function(page) {
	$scope.search.pageNumber = page;    	
	$scope.fetchResult();
} 
	//$scope.$watch("search.searchText", function() {_.debounce(function() {$scope.selectPage(); console.log('changed')},100)}, true);
	
	$scope.onSort = function (sortedBy, sortDir) {
		console.debug(sortedBy) ;
		console.debug(sortDir) ;
		$scope.search.sortDir = sortDir;
		$scope.search.sortedBy = sortedBy;      
		$scope.search.pageNumber = 1;    
		$scope.sortedby= sortedBy;
		$scope.sortvalue=sortedBy;
		$scope.fetchResult().then(function () {
      //The request fires correctly but sometimes the ui doesn't update, that's a fix
      $scope.search.pageNumber = 1;
  });

	};

	$scope.selectPage(1);


});

countryApp.config(function(RestangularProvider) {
 RestangularProvider.setRestangularFields({
      id: "Id"  });
})
countryEditController = countryApp.controller('countryEditController',function($scope,$http,Restangular,$location, $routeParams,$route)
{

	//$scope.alerts = [{ type: 'danger', msg: 'Oh snap! Change a few things up and try submitting again.' }];
	//$scope.bill = data;


	$scope.op = [{"name":"associates cons client","id":3},{"name":"company 2","id":4},{"name":"AANAND SEAMLESS TUBES PVT. LTD.","id":6},{"name":"Aarna Clothing Pvt. Ltd.","id":7},{"name":"AANAND SEAMLESS TUBES PVT. LTD.","id":8},{"name":"Aarna Clothing Pvt. Ltd.","id":9},{"name":"AARNA CLOTHING PVT. LTD.","id":10},{"name":"ACCURATE BILLING SERVICES PVT. LTD.","id":11},{"name":"ALPHA PAPER CONTAINERS PVT.LTD.","id":12},{"name":"Arvind Accele Ltd.(Khatraj Project)","id":13},{"name":"Arvind Ltd.(Vadinar Project)","id":14},{"name":"ASTRAL POLY-TECHNIK LTD.","id":15},{"name":"AVI TECHNO CAST PVT. LTD.","id":16},{"name":"BALAJI COIR PVT. LTD.","id":17},{"name":"BHAVNA ENGINEERING CO.","id":18}];
	
	$scope.tasksheaders = ['Task','Particulars', 'Rate','Qty','Amount',''];
	data = [];
	
	$scope.tasks_deleted = [];
	if($routeParams.id)	
	{
		console.debug($routeParams.id);
		Restangular.one('json', $route.current.params.id).get().then(function (data)
			{$scope.country = data;
		//$scope.clientname={"name":data.client,"id":data.clientid};
		//$scope.clients = data.clients;delete $scope.bill.clients;
		console.log($scope.country)
	});	
	}

	$scope.deletedisabled=(typeof $route.current.params.id === 'undefined');	

	 // if the id is not set means we are addding hence no delete	
	

	$scope.savecurrent = function(event){
		//console.log(event.target.button('what'));	
		
		// alert('saving');
		// event.target.button('loading');
		console.log($scope.myForm.$error);

		if ($scope.myForm.$valid)
		{		
			var $btn = $(event.target);	
			$btn.button('loading') ;		
		//$scope.bill.clientid = $scope.clientname.id;
		// delete  $scope.country.client;
		// delete  $scope.bill.address;
		// delete  $scope.bill.code;
		console.log("Saving "+$scope.country.Id)	
		// if ($scope.country.Id)
		// {		
		// 	$scope.country.post().then(function(){$scope.alerts=[{ type: 'success', msg: 'Record Saved!!' }];
		// 		$location.path('/')
		// 	});
		// }
		// else
			//baseNames.post($scope.insurer).then(function(){$location.path('/')});
		Restangular.all('json').post($scope.country).then(function(data){
			$scope.country.id = data;
			console.log(data);
			$scope.alerts=[{ type: 'success', msg: 'Record Added!!' }];									
			// $location.path('/edit/'+$scope.country.id);				      				
			$location.path('/');				      				
		});
	}
	else
	{			
			// $scope.alert={ type: 'danger', msg: 'Record Deleted!!' };			
			// event.preventDefault();
			$scope.alerts = [{ type: 'warning', msg: 'Please fill out the form correctly!!' }];
		}		
	}	
	

		$scope.deletecurrent = function(){				
			$scope.country.post('DELETE').then(function(){$scope.alerts=[{ type: 'danger', msg: 'Record Deleted!!' }];
				$location.path('/')});
		}

		

		$scope.est_names = function(est_name) {
			return $http.get("/json&q="+est_name).then(function(response){
				return response.data;
			});
		}

		
		$scope.loading = function($event) {		
			var $btn = $($event.target);
			$btn.button('loading') ;
			console.log($btn);		
		}
		
	});


angular.module('countryApp').directive('sortBy', function () {
	return {

		template : 
		'<a style="cursor: pointer;" ng-click="sort(sortvalue)">'+
		'<span ng-transclude></span>'+ 
		'<span ng-show="sortedby == sortvalue">'+
		'<b ng-class="{true: \'glyphicon glyphicon-arrow-up\', false: \'glyphicon glyphicon-arrow-down\'}[sortdir == \'asc\']"></b></span>'+
		'</a>',
		restrict: 'E',
		transclude: true,
		replace: true,
		scope: {
			sortdir: '=',
			sortedby: '=',
			sortvalue: '@',
			onsort: '='
		},
		link: function (scope, element, attrs) {
			scope.sort = function () {
				if (scope.sortvalue)
				{
					if (scope.sortedby == scope.sortvalue)
						scope.sortdir = scope.sortdir == 'asc' ? 'desc' : 'asc';
					else {
						scope.sortedby = scope.sortvalue;
						scope.sortdir = 'asc';
					}
					scope.onsort(scope.sortedby, scope.sortdir);
				}
			}
		}
	};
});
