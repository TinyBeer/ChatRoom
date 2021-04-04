<template>
  <div class="login">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >

        <b-card title="登录">
          <b-form>
            <b-form-group label="ID">
              <b-form-input
                v-model="$v.user.id.$model"
                type="number"
                placeholder="请输入ID"
                :state="validateState('id')"
              > </b-form-input>
              <b-form-invalid-feedback :state="validateState('id')">
                ID不能超过5位数
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="请输入密码"
                :state="validateState('password')"
              > </b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码应为6-10位
              </b-form-invalid-feedback>
            </b-form-group>
          </b-form>

          <b-form-group>
            <b-button
              @click="login"
              variant="outline-primary"
              block
            >登录</b-button>
          </b-form-group>
        </b-card>

      </b-col>
    </b-row>
  </div>
</template>

<script>
import { minLength, required, maxLength } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        id: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      id: {
        required,
        maxLength: maxLength(5),

      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(10),
      },
    },
  },

  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      this.$v.user.$touch();
      // 验证数据
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求api
      this.userLogin(this.user).then(() => {
        // 跳转到主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      console.log('login');
    },
  },
};
</script>

<style lang="scss" scoped >
</style>
