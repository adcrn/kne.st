<template>
    <div class="flex flex-row justify-center">
        <div class="pr-16">
            <transition-group name='fade' tag='div'>
                <div
                v-for="number in [currentNumber]"
                :key='number'
                >
                    <img class="shadow-md"
                    v-bind:src="currentImageOriginal.url"
                    v-bind:alt="currentImageOriginal.alt"
                    width="360"
                    height="240" 
                    >
                </div>
            </transition-group>
        </div>
        <div class="pl-16">
            <transition-group name='fade' tag='div'>
                <div
                v-for="number in [currentNumber]"
                :key='number'
                >
                    <img class="shadow-md"
                    v-bind:src="currentImageProcessed.url"
                    v-bind:alt="currentImageProcessed.alt"
                    width="360"
                    height="240"
                    >
                </div>
          </transition-group>
        </div>
    </div>
</template>

<script>
export default {
  name: "image-slider",
  components: {},
  data() {
    return {
      images_original: [
        {
          url: require("../assets/home/1.jpg"),
          alt: "image of tiny bird in tree"
        },
        {
          url: require("../assets/home/2.jpg"),
          alt: "image of chubby red bird in tree"
        },
        {
          url: require("../assets/home/4.jpg"),
          alt: "image of bird in tree with sky in background"
        },
        {
          url: require("../assets/home/5.jpg"),
          alt: "image of tiny bird hidden in tree branches"
        }
      ],
      images_processed: [
        {
          url: require("../assets/home/1_p.jpg"),
          alt: "zoomed in image of tiny bird in tree"
        },
        {
          url: require("../assets/home/2_p.jpg"),
          alt: "zoomed in image of chubby red bird in tree"
        },
        {
          url: require("../assets/home/4_p.jpg"),
          alt: "zoomed in image of bird in tree with sky in background"
        },
        {
          url: require("../assets/home/5_p.jpg"),
          alt: "zoomed in image of tiny bird hidden in tree branches"
        }
      ],
      currentNumber: 0,
      timer: null
    };
  },

  mounted: function() {
    this.startRotation();
  },

  methods: {
    startRotation: function() {
      this.timer = setInterval(this.next, 7500);
    },

    next: function() {
      this.currentNumber += 1;
    }
  },

  computed: {
    currentImageOriginal: function() {
      return this.images_original[
        Math.abs(this.currentNumber) % this.images_original.length
      ];
    },
    currentImageProcessed: function() {
      return this.images_processed[
        Math.abs(this.currentNumber) % this.images_processed.length
      ];
    }
  }
};
</script>

<style lang="scss">
.fade-enter-active,
.fade-leave-active {
  transition: all 0.8s ease;
  overflow: hidden;
  visibility: visible;
  opacity: 1;
  position: absolute;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
  visibility: hidden;
}
</style>
