<template>
  <div class="login">
    <div class="flex h-screen items-center justify-center">
      <form class="w-full max-w-xs">
        <div class="md:flex md:items-center mb-6">
          <div class="md:w-1/3">
            <label class="block text-grey-darker font-bold md:text-right mb-1 md:mb-0 pr-4" for="inline-username">
              Username
            </label>
          </div>
          <div class="md:w-2/3">
            <input class="bg-white appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:border-blue" id="inline-full-name" type="text" placeholder="janedoe">
          </div>
        </div>
        <div class="md:flex md:items-center mb-6">
          <div class="md:w-1/3">
            <label class="block text-grey-darker font-bold md:text-right mb-1 md:mb-0 pr-4" for="inline-username">
              Password
            </label>
          </div>
          <div class="md:w-2/3">
            <input class="bg-white appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-blue" id="inline-username" type="password" placeholder="***********">
          </div>
        </div>
        <div class="md:flex md:items-center">
          <div class="md:w-1/3"></div>
          <div class="md:w-2/3">
            <button class="shadow bg-blue-darker hover:bg-blue-darkest focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded" type="button">
              Login
            </button>
          </div>
        </div>
      </form>
    </div>
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
</style>
