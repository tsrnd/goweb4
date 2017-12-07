<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{ .PageTitle }}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="">
  <!--[if ie]><meta content='IE=8' http-equiv='X-UA-Compatible'/><![endif]-->
  <!-- bootstrap -->
  <link href="/static/css/bootstrap/css/bootstrap.min.css" rel="stylesheet">      
  <link href="/static/css/bootstrap/css/bootstrap-responsive.min.css" rel="stylesheet">
  
  <link href="/static/css/themes/css/bootstrappage.css" rel="stylesheet"/>
  
  <!-- global styles -->
  <link href="/static/css/themes/css/flexslider.css" rel="stylesheet"/>
  <link href="/static/css/themes/css/main.css" rel="stylesheet"/>

  <!-- scripts -->
  <script src="/static/css/themes/js/jquery-1.7.2.min.js"></script>
  <script src="/static/css/bootstrap/js/bootstrap.min.js"></script>				
  <script src="/static/css/themes/js/superfish.js"></script>	
  <script src="/static/css/themes/js/jquery.scrolltotop.js"></script>
  <!--[if lt IE 9]>			
      <script src="ht../../tp://html5shim.googlecode.com/svn/trunk/html5.js"></script>
      <script src="js/respond.min.js"></script>
  <![endif]-->
</head>
<body>		
  <div id="top-bar" class="container">
    <div class="row">
      <div class="span4">
        <form method="POST" class="search_form">
          <input type="text" class="input-block-level search-query" Placeholder="eg. T-sirt">
        </form>
      </div>
      <div class="span8">
        <div class="account pull-right">
          <ul class="user-menu">				
            {{if .Name}}
              <li><a href="/userProfile">My Account</a></li>
            {{end}}
            <li><a href="/cart">Your Cart
                <sup id="js-cart-count" class="badge" style="background-color: red;margin-left: 1px"></sup></a></li>
            <li><a href="/checkout">Checkout</a></li>	
            {{if .Name}}				
              <li><a href="#" data-toggle="collapse" data-target="#demo">{{ .Name }}</a></li>
              <li id="demo" class="collapse"> <a href="/logout">Logout</a></li>
            {{else}}
              <li><a href="/login">Login</a></li>
            {{end}}
          </ul>
        </div>
      </div>
    </div>
  </div>
  <div id="wrapper" class="container">
      <section class="navbar main-menu">
        <div class="navbar-inner main-menu">				
          <a href="/" class="logo pull-left"><img src="/static/css/themes/images/logo.png" class="site_logo" alt=""></a>
          <nav id="menu" class="pull-right">
            <ul>
                {{ range $val := .ProductGroup }}
                  <li><a href="/group/{{ $val.ID }}">{{ $val.Name }}</a>					
                {{ end }}
            </ul>
          </nav>
        </div>
      </section>