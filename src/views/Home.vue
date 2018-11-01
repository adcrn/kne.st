<template>
  <div class="home">
    <section class="hero is-light is-fullheight-with-navbar has-text-centered">
      <div class="hero-body">
        <div class="container">
          <div class="intro">
          <p class="is-size-4"><strong>An end-to-end solution that selects and enhances bird photographs</strong></p>
          <p class="is-size-5">
            <i>knest</i> leverages state-of-the-art techniques in machine learning and<br>
            computer vision in order to find and crop birds in your photographs.<br>
            <strong>You upload the photos, we give you the bird(s).</strong>
          </p>
          </div>
          <div class="field is-grouped">
            <p class="control">
              <router-link to="/register">
                <a class="button is-primary is-medium">sign up</a>
              </router-link>
            </p>
            <p class="control">
              <a class="button is-medium">learn more</a>
            </p>
          </div>
        </div>
        <nav class="level">
          <div class="level-item has-text-centered">
            <div>
              <figure class="image is-320x240">
                <img src="../assets/home/1.jpg" alt="tiny bird in tree" width="360" height="240">
              </figure>
              <p class="is-size-6">Original Photo</p>
            </div>
          </div>
          <div class="level-item has-text-centered">
            <div>
              <figure class="image is-320x240">
                <img src="../assets/home/1_p.jpg" alt="zoomed image of tiny bird in tree" width="360" height="240">
              </figure>
              <p class="is-size-6">Processed Photo</p>
            </div>
          </div>
        </nav>
      </div>
    </section>
  </div>
</template>

<script>
// @ is an alias to /src
//import Home from "@/components/Home.vue";

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
.container {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.hero-body {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.intro {
  padding-top: 10rem;
}

.level-item {
  padding: 6rem;
}

p {
  padding: 0.5rem;
}

h1 {
  margin: 10px 0 0;
}
</style>
