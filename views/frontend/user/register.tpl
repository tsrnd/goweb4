<div class="span7">         
  <h4 class="title"><span class="text"><strong>Register</strong> Form</span></h4>
  <form action="/registerHandler" method="post" class="form-stacked">
    <fieldset>
      <div class="control-group">
        {{ if .RegisterInfo.Errors.Name }}
        <p class="error">{{ .RegisterInfo.Errors.Name }}</p>
        {{ end }}
        <label class="control-label">Username</label>
        <div class="controls">
          <input type="text" name="userName" placeholder="Enter your username" class="input-xlarge">
        </div>
      </div>
      <div class="control-group">
        {{ if .RegisterInfo.Errors.Email }}
        <p class="error text-danger">{{ .RegisterInfo.Errors.Email }}</p>
        {{ end }}
        <label class="control-label">Email address:</label>
        <div class="controls">
          <input type="text" name="email" placeholder="Enter your email" class="input-xlarge">
        </div>
      </div>
      <div class="control-group">
        {{ if .RegisterInfo.Errors.Password }}
        <p class="error">{{ .RegisterInfo.Errors.Password }}</p>
        {{ end }}
        <label class="control-label">Password:</label>
        <div class="controls">
          <input type="password" name="password" placeholder="Enter your password" class="input-xlarge">
        </div>
      </div>  
      <div class="control-group">
        {{ if .RegisterInfo.Errors.Phone }}
        <p class="error">{{ .RegisterInfo.Errors.Phone }}</p>
        {{ end }}
        <label class="control-label">Phone Number:</label>
        <div class="controls">
          <input type="text" name="phoneNumber" placeholder="Enter your phone number" class="input-xlarge">
        </div>
      </div>
      <div class="control-group">
        {{ if .RegisterInfo.Errors.Address }}
        <p class="error">{{ .RegisterInfo.Errors.Address }}</p>
        {{ end }}
        <label class="control-label">Address:</label>
        <div class="controls">
          <input type="text" name="address" placeholder="Enter your address" class="input-xlarge">
        </div>
      </div>                                    
      <div class="control-group">
        <p>Now that we know who you are. I'm not a mistake! In a comic, you know how you can tell who the arch-villain's going to be?</p>
      </div>
      <hr>
      <div class="actions"><input tabindex="9" class="btn btn-inverse large" type="submit" value="Create your account"></div>
    </fieldset>
  </form>         
</div>
</div>
