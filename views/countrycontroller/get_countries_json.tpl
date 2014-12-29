<!DOCTYPE html>
<html>
<head>
	<title>Beego</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<script type="text/javascript" src="/static/js/jquery.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="/static/js/angular.min.js"></script>
	<script type="text/javascript" src="/static/js/angular-ui.js"></script>	
	<script type="text/javascript" src="/static/js/app.js"></script>	
		

	<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.css"/>
</head>

<body>
	<header class="hero-unit" style="color:black;background-color:#A9F16C">
		<div class="container">
			<div class="row">
				<div class="hero-text">
					<h1>Welcome {{.user}}! on JSON details page!</h1>			    
					<div ng-app="billApp" ng-view></div>
				</div>
			</div>
		</div>
	</header>
	<script type="text/ng-template" id="/list.html">
		<table class="table table-striped">		
			<thead>		</thead>
			<tbody>
				<tr ng-repeat="bill in bills">
					<td>{{bill.id}}</td>
				</tr></tbody>
		</table>			
		</script>

		 <script type="text/ng-template" id="/detail.html">
<form autocomplete="off" id="myForm" name="myForm" class="form-horizontal" role="form">
        <fieldset>
        </fieldset>          
          <legend>Country
          </legend>
          </form>

		</script>

	</body>
	</html>
