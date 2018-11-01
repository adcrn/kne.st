<template>
  <div class="login">
    <section class="hero is-light is-fullheight-with-navbar has-text-centered">
      <div class="hero-body">
        <div class="container">
          <h1 class="is-size-3 has-text-centered"><strong>Login</strong></h1>
          <span class="error">{{ error }}</span>
          <div class="field is-grouped is-grouped-centered">
            <div class="control">
              <label class="label">Username</label>
              <input class="input" type="text">
            </div>
          </div>
          <div class="field is-grouped is-grouped-centered">
            <div class="control">
              <label class="label">Password</label>
              <input class="input" type="password">
            </div>
          </div>
          <div class="buttons field is-grouped is-grouped-centered">
            <div class="control">
              <button class="button is-primary is-medium">Login</button>
            </div>
            <div class="control">
              <button class="button is-danger is-medium is-outlined">Cancel</button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "login",
  components: {},
  data() {
    return {
      error: ""
    };
  }
};
</script>

<style scoped lang="scss">
.login {
}

h1 {
  padding-bottom: 0.5rem;
}

.buttons {
  padding: 1rem;
}
</style>
