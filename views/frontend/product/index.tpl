<section class="header_text sub">
<img class="pageBanner" src="/static/css/themes/images/pageBanner.png" alt="New products" >
  <h4><span>New products</span></h4>
</section>
<section class="main-content">
  
  <div class="row">           
    <div class="span9">               
      <ul class="thumbnails listing-products">
        {{ range $product := .Products }}
          <li class="span3">
            <div class="product-box">
              <span class="sale_tag"></span>                        
              <a href="/detail/{{ $product.ID }}"><img alt="" src="/{{ if $product.Images }}{{ (index $product.Images 0).URL }}{{ else }}static/css/themes/images/ladies/6.jpg {{ end }}"></a><br/>
              <a href="/detail/{{ $product.ID }}" class="title">{{ $product.Name }}</a><br/>
              <a href="#" class="category">{{ $product.Size }}</a>
              <p class="price">{{ $product.Price }}$</p>
            </div>
          </li>
        {{ end }}
      </ul>               
      <hr>
      <div class="pagination pagination-small pagination-centered">
        <ul>
           
          <li><a href="?page={{ .Paginator.PrevPage }}">Prev</a></li>
          {{ range $page := .Paginator.PagePath}}
            <li><a href="?page={{ $page }}" class="">{{ $page }}</a>			
          {{ end }}     
          <li><a href="?page={{ .Paginator.NextPage }}">Next</a></li>       
        </ul>
      </div>
    </div>
    <div class="span3 col">
      <div class="block"> 
        <ul class="nav nav-list">
          <li class="nav-header">SUB CATEGORIES</li>
          <li><a href="products.html">Nullam semper elementum</a></li>
          <li class="active"><a href="products.html">Phasellus ultricies</a></li>
          <li><a href="products.html">Donec laoreet dui</a></li>
          <li><a href="products.html">Nullam semper elementum</a></li>
          <li><a href="products.html">Phasellus ultricies</a></li>
          <li><a href="products.html">Donec laoreet dui</a></li>
        </ul>
        <br/>
        <ul class="nav nav-list below">
          <li class="nav-header">MANUFACTURES</li>
          <li><a href="products.html">Adidas</a></li>
          <li><a href="products.html">Nike</a></li>
          <li><a href="products.html">Dunlop</a></li>
          <li><a href="products.html">Yamaha</a></li>
        </ul>
      </div>
      <div class="block">
        <h4 class="title">
          <span class="pull-left"><span class="text">Randomize</span></span>
          <span class="pull-right">
            <a class="left button" href="#myCarousel" data-slide="prev"></a><a class="right button" href="#myCarousel" data-slide="next"></a>
          </span>
        </h4>
        <div id="myCarousel" class="carousel slide">
          <div class="carousel-inner">
            <div class="active item">
              <ul class="thumbnails listing-products">
                <li class="span3">
                  <div class="product-box">
                    <span class="sale_tag"></span>                        
                    <img alt="" src="/static/css/themes/images/ladies/1.jpg"><br/>
                    <a href="product_detail.html" class="title">Fusce id molestie massa</a><br/>
                    <a href="#" class="category">Suspendisse aliquet</a>
                    <p class="price">$261</p>
                  </div>
                </li>
              </ul>
            </div>
            <div class="item">
              <ul class="thumbnails listing-products">
                <li class="span3">
                  <div class="product-box">                       
                    <img alt="" src="/static/css/themes/images/ladies/2.jpg"><br/>
                    <a href="product_detail.html" class="title">Tempor sem sodales</a><br/>
                    <a href="#" class="category">Urna nec lectus mollis</a>
                    <p class="price">$134</p>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="block">               
        <h4 class="title"><strong>Best</strong> Seller</h4>               
        <ul class="small-product">
          <li>
            <a href="#" title="Praesent tempor sem sodales">
              <img src="/static/css/themes/images/ladies/3.jpg" alt="Praesent tempor sem sodales">
            </a>
            <a href="#">Praesent tempor sem</a>
          </li>
          <li>
            <a href="#" title="Luctus quam ultrices rutrum">
              <img src="/static/css/themes/images/ladies/4.jpg" alt="Luctus quam ultrices rutrum">
            </a>
            <a href="#">Luctus quam ultrices rutrum</a>
          </li>
          <li>
            <a href="#" title="Fusce id molestie massa">
              <img src="themes/images/ladies/5.jpg" alt="Fusce id molestie massa">
            </a>
            <a href="#">Fusce id molestie massa</a>
          </li>   
        </ul>
      </div>
    </div>
  </div>
</section>