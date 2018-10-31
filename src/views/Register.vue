<template>
  <div class="register">
    <h1 class="is-size-2 has-text-centered"><strong>Create Account</strong></h1>
    <span>{{ error }}</span>
    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <label class="label">Username</label>
        <input class="input" type="text">
      </div>
      <div class="control">
        <label class="label">Password</label>
        <input class="input" type="password">
      </div>
    </div>
    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <label class="label">Email</label>
        <input class="input" type="email">
      </div>
      <div class="control">
        <label class="label">Confirm Email</label>
        <input class="input" type="email">
      </div>
    </div>
    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <label class="checkbox">
          <input type="checkbox" id="checkbox" v-model="checked">
          I agree to the <a href="#">terms and conditions</a>
        </label>
      </div>
    </div>
    <div class="field is-grouped is-grouped-centered">
      <div class="control">
        <button class="button is-primary is-medium">Continue</button>
      </div>
      <div class="control">
        <button class="button is-danger is-medium is-outlined">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "register",
  components: {},
  data() {
    return {
      error: ""
    };
  }
};
</script>

<style scoped lang="scss">
.register {
  padding: 12rem;
}

h1 {
  padding-bottom: 1rem;
}
</style>
