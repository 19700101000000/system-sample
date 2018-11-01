<template lang="pug">
div
  b-carousel#id(
    v-model='slide'
    controls
    indicators
    background='black'
    img-width='1280'
    img-height='320'
  )
    b-carousel-slide(v-for='img in images', img-blank)
      p {{ slide }} : {{ img }}
</template>

<script lang="ts">
import {
  Component,
  // Prop,
  Vue,
} from "nuxt-property-decorator"

@Component({})
export default class AdCarousel extends Vue {
  // @Prop() images
  slide: number = 0;
  images: string[] = [
    "Ad",
    "Ad",
    "Ad",
    "Ad",
    "Ad",
  ];
}
</script>
