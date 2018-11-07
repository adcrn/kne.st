<template>
  <div class="home bg-grey-lighter">
    <div class="landing h-full mx-auto flex flex-col">
      <div class="intro flex flex-col justify-center mt-32">
        <p class="tagline font-sans text-4xl font-bold text-center">
          You upload the photos, we give you the bird(s).
        </p>
        <p class="blurb text-2xl font-sans font-medium text-center my-8">
          Knest leverages state-of-the-art techniques<br>
          in machine learning and computer vision<br>
          in order to find and crop birds in your photographs.
        </p>
        <div class="flex justify-center py-4">
          <p class="font-sans px-2">
            <router-link to="/register">
              <button class="btn">Sign Up</button>
            </router-link>
          </p>
          <p class="font-sans px-2">
            <a href="#extra"><button class="btn-trans btn-blue">Learn More</button></a>
          </p>
        </div>
      </div>
      <div class="flex flex-row justify-center py-32">
        <div class="px-32">
            <img class="shadow-md" src="../assets/home/1.jpg" alt="tiny bird in tree" width="360" height="240">
        </div>
        <div class="px-32">
            <img class="shadow-md" src="../assets/home/1_p.jpg" alt="zoomed image of tiny bird in tree" width="360" height="240">
        </div>
      </div>
    </div>
    <div id="extra" class="h-screen mx-auto flex flex-row">
      <p>test</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "home",
  components: {},
  data() {
    return {
      images_original: [],
      images_processed: []
    };
  },
  methods: {},
  created() {
    (this.images_original = [
      {
        url: "../assets/home/1.jpg",
        alt: "image of tiny bird in tree"
      }
    ]),
      (this.images_processed = [
        {
          url: "../assets/home/1_p.jpg",
          alt: "zoomed in image of tiny bird in tree"
        }
      ]);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.btn {
  @apply bg-blue-darker font-semibold text-lg text-white py-2 px-4 border rounded;
}

.btn-trans {
  @apply bg-transparent font-semibold text-lg py-2 px-4 border rounded;
}

.btn-blue {
  @apply text-blue-darker border-blue;
}

.btn-blue:hover {
  @apply bg-blue-darker text-white border-transparent;
}
</style>
