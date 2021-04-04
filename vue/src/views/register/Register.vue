<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >

        <b-card title="注册">
          <b-form>
            <b-form-group label="昵称">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="请输入昵称"
                :state="validateState('name')"
              > </b-form-input>
              <b-form-invalid-feedback :state="validateState('name')">
                昵称不能超过10位
              </b-form-invalid-feedback>
            </b-form-group>

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
              @click="register"
              variant="outline-primary"
              block
            >注册</b-button>
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
        name: '',
        id: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      name: {
        required,
        maxLength: maxLength(10),
      },
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
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userRegister(this.user).then(() => {
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
    },
  },
};
</script>

<style lang="scss" scoped >
</style>
