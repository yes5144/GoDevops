<template>
  <div class="hello">
    <h1>{{ username }}</h1>
    <p>For a guide and recipes on how to configure / customize this project!</p>
  </div>
</template>

<script>
import { getUserInfo } from "../api";
export default {
  data() {
    return {
      username: this.$store.state.user.Name || "unknown"
    };
  },
  name: "HelloWorld",
  mounted() {
    this.UserInfo();
  },
  methods: {
    UserInfo() {
      getUserInfo().then(res => {
        console.log(res);
        let { code, msg, data } = res;
        if (!data.user) {
          this.$message({
            message: msg,
            code: code,
            type: "error"
          });
        } else {
          localStorage.user = data.user.Name;
          this.$store.commit("LOGIN", {
            user: data.user.Name
          });
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
