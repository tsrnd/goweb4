<div class="row">
  <div class="span5">         
    <h4 class="title"><span class="text"><strong>Login</strong> Form</span></h4>
    {{.flash.error}}
    {{.flash.success}}
    <form action="/login" method="post">
      {{.xsrfdata}}
      <fieldset>
        <div class="control-group">
          <label class="control-label">Username</label>
          <div class="controls">
            <input type="text" placeholder="Enter your username" id="username" name="username" class="input-xlarge">
          </div>
        </div>
        <div class="control-group">
          <label class="control-label">Username</label>
          <div class="controls">
            <input type="password" placeholder="Enter your password" id="password" name="password" class="input-xlarge">
          </div>
        </div>
        <div class="control-group">
          <input tabindex="3" class="btn btn-inverse large" type="submit" value="Sign into your account">
          <hr>
          <p class="reset">Recover your <a tabindex="4" href="#" title="Recover your username or password">username or password</a></p>
        </div>
      </fieldset>
    </form>       
  </div>
