<template>
  <div id="app">
    <nav class="navbar has-shadow" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <router-link to="/">
          <img src="./assets/logo.png" alt="knest logo" width="64" height="28">
        </router-link>
      </div>
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="field is-grouped">
            <p class="control">
              <router-link to="/register">
                <a class="button is-primary"><strong>Sign up</strong></a>
              </router-link>
            </p>
            <p class="control">
              <router-link to="/login">
                <a class="button is-light">Login</a>
              </router-link>
            </p>
          </div>
        </div>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<style lang="scss">
#app {
  background-color: #f8f8f8;
}
.navbar {
  padding-left: 0.5rem;
  background-color: #f8f8f8;
}
</style>
