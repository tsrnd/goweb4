  <!-- Content Wrapper. Contains page content -->
<div class="content-wrapper">
  <!-- Content Header (Page header) -->
  <section class="content-header">
    <h1>
      Pace page
      <small>Loading example</small>
    </h1>
    <ol class="breadcrumb">
      <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
      <li><a href="#">Examples</a></li>
      <li class="active">Pace page</li>
    </ol>
  </section>

  <!-- Main content -->
  <section class="content">
    <!-- Default box -->
    <div class="box">
      <div class="box-header with-border">
        <h3 class="box-title">Add Product</h3>
        <div class="box-tools pull-right">
          <button type="button" class="btn btn-box-tool" data-widget="collapse" data-toggle="tooltip" title="Collapse">
            <i class="fa fa-minus"></i></button>
          <button type="button" class="btn btn-box-tool" data-widget="remove" data-toggle="tooltip" title="Remove">
            <i class="fa fa-times"></i></button>
        </div>
      </div>
      <div class="box-body">
        <div class="row">
          <div class="col-md-6 col-md-offset-3">
            <form enctype="multipart/form-data" class="col-md-12" method="POST" action="/product">
              <div class="form-group col-md-12">
                <label for="name">Name</label>
                <input type="text" class="form-control" id="name" name="name" placeholder="Name.." >
                <small class="text-danger"></small>
              </div>
              <div class="clearfix"></div>
              <!-- size -->
              <div class="form-group col-md-4">
                <label for="size">Size</label>
                <select class="form-control" name="size" id="size">
                  <option value="XXL">XXL</option>
                  <option value="XL">XL</option>
                  <option value="L">L</option>
                  <option value="S">S</option>
                  <option value="XS">XS</option>
                </select>
              </div>
              <!-- color -->
              <div class="form-group col-md-4">
                <label for="color">Color</label>
                <select class="form-control" name="color" id="color">
                  <option value="blue" style="background-color:blue"></option>
                  <option value="red" style="background-color:red"></option>
                  <option value="yellow" style="background-color:yellow"></option>
                </select>
              </div>
              <!-- instock -->
              <div class="form-group col-md-12">
                  <label for="in_stock">Instock</label>
                  <input type="number" class="form-control" min=1 name="in_stock" id="in_stock" >
              </div>
              <div class="clearfix"></div>
              <!-- price -->
              <div class="form-group col-md-12">
                <label for="price">Price</label>
                <input type="number" name="price" id="price" min=1 class="form-control" >
              </div>
              <div class="clearfix"></div>
              <!-- group -->
              <div class="form-group col-md-12">
                <label for="group_id">Group</label>
                <select class="form-control" name="group_id" id="group_id">
                  <!-- 
                    {{range .}}
                      
                    {{end}} 
                  -->
                </select>
              </div>
              <!-- image -->
              <div class="form-group col-md-12">
                <label for="image">Image</label>
                <input type="file" name="image" id="image" class="form-control">
              </div>
              <div class="clearfix"></div>
          
              <div class="col-md-6">
                  <button type="submit" class="btn btn-primary pull-right">Submit</button>
              </div>
            </form>
          </div>
        </div> 
      </div>
      <!-- /.box-body -->
      <div class="box-footer">
        Goweb4
      </div>
      <!-- /.box-footer-->
    </div>
    <!-- /.box -->
  </section>
  <!-- /.content -->
</div>
<!-- /.content-wrapper -->
