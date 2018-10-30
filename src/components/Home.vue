<template>
  <div class="home">
    <section class="intro has-text-centered">
      <div class="container">
        <p class="is-size-4"><strong>{{ tagline }}</strong></p>
        <p class="is-size-5">
          <i>knest</i> leverages state-of-the-art techniques in machine learning and<br>
          computer vision in order to find and crop birds in your photographs.<br>
          <strong>You upload the photos, we give you the bird(s).</strong>
        </p>
        <div class="field is-grouped">
          <p class="control">
            <a class="button is-primary is-medium">sign up</a>
          </p>
          <p class="control">
            <a class="button is-medium">learn more</a>
          </p>
        </div>
      </div>
    </section>
    <section class="photo_comparison">
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <figure class="image is-320x240">
              <img src="../assets/7070.jpg" alt="tiny bird in tree" width="360" height="240">
            </figure>
            <p class="is-size-6">Original Photo</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <figure class="image is-320x240">
              <img src="../assets/7070_zoom.jpg" alt="zoomed image of tiny bird in tree" width="360" height="240">
            </figure>
            <p class="is-size-6">Processed Photo</p>
          </div>
        </div>
      </nav>
    </section>
  </div>
</template>

<script>
export default {
  name: "Home",
  props: {
    tagline: String
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

.intro {
  padding: 8rem;
}

.photos {
  display: flex;
  align-items: center;
}

p {
  padding: 0.5rem;
}

h1 {
  margin: 10px 0 0;
}
</style>
