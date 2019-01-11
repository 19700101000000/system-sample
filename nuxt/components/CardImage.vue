<template lang="pug">
  b-card.my-2(
    :header="value.username"
    :footer="value.datetime")
    b-img(
      thumnail
      fluid
      width="256"
      height="256" 
      :src="'/api/images/' + value.imagepath")
</template>
<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface CardValue {
  username:  string,
  imagepath: string,
  datetime:  string,
}

@Component({})
export default class CardImage extends Vue {
  @Prop() value!: CardValue;
}
</script>
