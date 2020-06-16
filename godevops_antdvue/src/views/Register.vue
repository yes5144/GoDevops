<template>
  <a-card title="Register" style="width: 360px;height:400px">
    <a-form :form="form" @submit="registerSubmit">
      <a-form-item v-bind="formItemLayout">
        <span slot="label">Nickname&nbsp;</span>
        <a-input
          v-decorator="[
          'name',
          {
            rules: [{ required: true, message: 'Please input your nickname!', whitespace: true }],
          },
        ]"
        />
      </a-form-item>

      <a-form-item v-bind="formItemLayout" label="Telephone">
        <a-input
          v-decorator="[
          'telephone',
          {
            rules: [{ required: true, message: 'Please input your phone number!' }],
          },
        ]"
          style="width: 100%"
        >
          <a-select
            slot="addonBefore"
            v-decorator="['prefix', { initialValue: '86' }]"
            style="width: 60px"
          >
            <a-select-option value="86">+86</a-select-option>
          </a-select>
        </a-input>
      </a-form-item>

      <a-form-item v-bind="formItemLayout" label="Password" has-feedback>
        <a-input
          v-decorator="[
          'password',
          {
            rules: [
              {
                required: true,
                message: 'Please input your password!',
              },
              {
                validator: validateToNextPassword,
              },
            ],
          },
        ]"
          type="password"
        />
      </a-form-item>

      <a-form-item v-bind="tailFormItemLayout">
        <a-checkbox v-decorator="['agreement', { valuePropName: 'checked' }]">
          I have read the
          <a href>agreement</a>
        </a-checkbox>
      </a-form-item>
      <a-form-item v-bind="tailFormItemLayout">
        <a-button type="primary" block>Register</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>
import { requestRegister } from "../api";
export default {
  data() {
    return {
      confirmDirty: false,
      formItemLayout: {
        labelCol: {
          xs: { span: 24 },
          sm: { span: 8 }
        },
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 16 }
        }
      },
      tailFormItemLayout: {
        wrapperCol: {
          xs: {
            span: 24,
            offset: 0
          },
          sm: {
            span: 16,
            offset: 8
          }
        }
      }
    };
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "register" });
  },
  methods: {
    registerSubmit(e) {
      e.preventDefault();
      this.form.validateFieldsAndScroll((err, values) => {
        if (!err) {
          console.log("Received values of form: ", values);
          requestRegister(values).then(res => {
            this.logining = false;
            // http code
            let { code, msg, data } = res;
            console.log(code, msg, data);
          });
        }
      });
    },
    validateToNextPassword(rule, value, callback) {
      const form = this.form;
      if (value && this.confirmDirty) {
        form.validateFields(["confirm"], { force: true });
      }
      callback();
    }
  }
};
</script>