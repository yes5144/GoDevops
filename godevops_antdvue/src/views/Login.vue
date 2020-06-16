<template>
  <a-card
    title="Login"
    style="width: 360px; height:360px; left:50%; top:50%; position:relative; margin-left:-180px; margin-top:150px; "
  >
    <a-form
      id="components-form-demo-normal-login"
      :form="form"
      class="login-form"
      @submit="handleSubmit"
    >
      <a-form-item>
        <a-input
          v-decorator="[
          'telephone',
          { rules: [{ required: true, message: 'Please input your telephone!' }] },
        ]"
          placeholder="Telephone"
        >
          <a-icon slot="prefix" type="telephone" style="color: rgba(0,0,0,.25)" />
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-input
          v-decorator="[
          'password',
          { rules: [{ required: true, message: 'Please input your Password!' }] },
        ]"
          type="password"
          placeholder="Password"
        >
          <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)" />
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-checkbox
          v-decorator="[
          'remember',
          {
            valuePropName: 'checked',
            initialValue: true,
          },
        ]"
        >Remember me</a-checkbox>
        <a class="login-form-forgot" href>Forgot password</a>
        <a-button type="primary" html-type="submit" class="login-form-button">Log in</a-button>Or
        <a href="/register">register now!</a>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>
import { requestLogin } from "../api";
export default {
  data() {
    return {
      logining: false,
      ruleForm: {
        account: null,
        checkPass: null
      }
    };
  },

  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "normal_login" });
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          console.log("Received values of form: ", values);
          let loginParams = {
            telephone: values.telephone,
            password: values.password
          };
          console.log("after:", loginParams);
          requestLogin(loginParams).then(res => {
            this.logining = false;
            // http code
            let { code, msg, data } = res;
            console.log(code, msg, data);

            if (!data.token) {
              this.$message({
                message: msg,
                type: "error"
              });
            } else {
              localStorage.token = data.token;
              this.$store.commit("LOGIN", {
                token: data.token
              });
              // let redirect=
              let redirect = decodeURIComponent("/hello");
              console.log(redirect);

              this.$router.push({ path: redirect });
            }
          });
        } else {
          console.log("error submit");
          return false;
        }
      });
    }
  }
};
</script>
<style>
#components-form-demo-normal-login .login-form {
  max-width: 300px;
}
#components-form-demo-normal-login .login-form-forgot {
  float: right;
}
#components-form-demo-normal-login .login-form-button {
  width: 100%;
}
</style>