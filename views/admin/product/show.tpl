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
        <h3 class="box-title">Product</h3>

        <div class="box-tools pull-right">
          <button type="button" class="btn btn-box-tool" data-widget="collapse" data-toggle="tooltip" title="Collapse">
            <i class="fa fa-minus"></i></button>
          <button type="button" class="btn btn-box-tool" data-widget="remove" data-toggle="tooltip" title="Remove">
            <i class="fa fa-times"></i></button>
        </div>
      </div>
      <div class="box-body">
        <div class="row">
            <div class="col-xs-12">
                <div class="box box-primary">
                    <div class="box-header">
                        <h3 class="box-title">{{.Product.Name}}</h3>
                    </div>
                    <div class="box-body">
                        <div class="col-md-4 text-center">
                            {{range .Product.Images}}
                            <img alt="{{.Name}}" title="{{.Name}}" class="img-circle img-thumbnail isTooltip"
                                 src="/{{.URL}}"
                                 data-original-title="Usuario">
                            {{end}}
                        </div>
                        <div class="col-md-8">
                            <div class="table-responsive">
                                <table class="table table-condensed table-responsive table-user-information has-description">
                                    <tbody>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-asterisk text-primary"></span>
                                                Identification
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                                {{.Product.ID}}
                                        </td>
                                    </tr>
                                    <tr>
                                            <td>
                                                <strong>
                                                    <span class="glyphicon glyphicon-asterisk text-primary"></span>
                                                    Name
                                                </strong>
                                            </td>
                                            <td class="text-primary">
                                                    {{.Product.Name}}
                                            </td>
                                        </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-bookmark text-primary"></span>
                                                Size
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                            <a href="#">
                                                {{.Product.Size}}
                                            </a>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-bookmark text-primary"></span>
                                                Color
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                            <a href="#">
                                                {{.Product.Color}}
                                            </a>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-usd text-primary"></span>
                                            Price
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                            {{.Product.Price}}
                                        </td>
                                    </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-tasks text-primary"></span>
                                                In Stock
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                            {{.Product.InStock}}
                                        </td>
                                    </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="glyphicon glyphicon-tasks text-primary"></span>
                                                Group Name
                                            </strong>
                                        </td>
                                        <td class="text-primary">
                                            {{.Product.ProductGroup.Name}}
                                        </td>
                                    </tr>
                                    <tr>
                                        <td>
                                            <strong>
                                                <span class="text-primary"></span>
                                                Action
                                            </strong>
                                        </td>
                                        <td>
                                            <a href="#"><span
                                                        class="btn btn-sm btn-primary">Edit</span></a>
                                            <a href="#"><span
                                                        class="btn btn-sm btn-danger">Go To List</span></a>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
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
