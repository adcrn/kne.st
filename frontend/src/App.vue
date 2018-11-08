<template>
  <div id="app">
    <nav class="absolute pin-t pin-l w-full py-4">
      <div class="flex justify-between items-center px-8">
        <router-link to="/">
          <img src="./assets/logo.png" alt="knest logo" width="64" height="28">
        </router-link>
        <div class="">
          <div class="flex items-center">
            <a class="text-lg leading-normal text-blue hover:text-blue-dark no-underline ml-8" href="#">Features</a>
            <a class="text-lg leading-normal text-blue hover:text-blue-dark no-underline ml-8" href="#">Pricing</a>
            <router-link to="/register"><button class="btn btn-blue">Sign Up</button></router-link>
            <router-link to="/login"><button class="btn btn-blue">Login</button></router-link>
          </div>
        </div>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<style lang="scss">
.btn {
  @apply ml-8 font-semibold text-lg py-2 px-4 border text-blue-darker rounded;
}

.btn-trans {
  @apply bg-transparent font-semibold text-lg py-2 px-4 border rounded;
}

.btn-blue {
  @apply bg-blue-darker text-white border-blue;
}

.btn-blue:hover {
  @apply bg-transparent text-blue-darker;
}
</style>
