<template>
  <div id="app">
    <nav class="bg-grey-lighter" role="navigation" aria-label="main navigation">
      <div class="container mx-auto">
        <div class="flex justify-between items-center">
          <div>
            <router-link to="/">
              <img src="./assets/logo.png" alt="knest logo" width="64" height="28">
            </router-link>
          </div>
          <div class="flex text-lg font-semibold">
            <p class="px-2">
              <router-link to="/register">
                <button class="btn btn-blue">Sign Up</button>
              </router-link>
            </p>
            <p class="px-2">
              <router-link to="/login">
                <button class="btn btn-blue">Login</button>
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
.btn {
  @apply bg-transparent font-semibold py-2 px-4 border rounded;
}

.btn-blue {
  @apply text-blue-darker border-blue;
}

.btn-blue:hover {
  @apply bg-blue-darker text-white border-transparent;
}
</style>
