<template>
  <div id="app" class="bg-grey-lighter">
    <header-component></header-component>
    <router-view/>
    <footer-component></footer-component>
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import FooterComponent from "./components/FooterComponent.vue";
export default {
  name: "app",
  components: {
    HeaderComponent,
    FooterComponent
  },
  data() {
    return {
      error: ""
    };
  }
};
</script>

<style lang="scss">
</style>
