<template>
  <div class="home bg-grey-lighter py-8">
    <section class="landing h-screen ml-auto mr-auto w-5/6 max-w-lg">
      <div class="intro flex items-center justify-center text-center h-full">
        <div>
          <h1 class="tagline font-sans text-5xl font-bold text-center leading-none tracking-tight mb-4">
            You upload the photos,<br>we give you the bird(s)
          </h1>
          <h2 class="blurb text-2xl opacity-75 text-blue-darker font-sans font-medium text-center">
            knest automatically detects birds in your photos and creates <br>beautiful images, giving you more time to experience the world.
          </h2>
          <div class="flex flex-row items-center justify-center pt-8 mb-24">
             <p class="font-sans px-2"> 
              <router-link to="/register">
                <button class="btn btn-blue">Sign Up</button>
              </router-link>
            </p>
            <p class="font-sans px-2">
              <a href="#extra"><button class="btn btn-trans">Learn More</button></a>
            </p>
          </div>
          <image-slider></image-slider>
        </div>
      </div>
    </section>
    <section id="extra" class="h-screen ml-auto mr-auto flex flex-row">
      <p>test</p>
    </section>
  </div>
</template>

<script>
import ImageSlider from "../components/ImageSlider.vue";
export default {
  name: "home",
  components: {
    ImageSlider
  },
  data() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.btn {
  @apply ml-8 font-semibold text-xl py-2 px-4 border text-blue-darker rounded;
}

.btn-trans {
  @apply bg-transparent font-semibold text-xl py-2 px-4 border-blue rounded;
}

.btn-trans:hover {
  @apply bg-blue-darker text-white;
}

.btn-blue {
  @apply bg-blue-darker text-white border-blue;
}

.btn-blue:hover {
  @apply bg-transparent text-blue-darker;
}
</style>
