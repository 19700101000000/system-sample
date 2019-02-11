<template lang="pug">
  b-card.my-2
    div(slot="header")
      b-link(:href="'/user/' + value.username") {{ value.username }} 
      template(v-if="value.evalparam") (Rate : {{ userRate }} / 5)
      template(v-else) (No rating.)
    div(slot="footer") {{ value.datetime }}
      b-button.ml-2(v-if="isSignin" variant="outline-secondary" size="sm") favorites({{ value.favorite }})
      span.ml-2(v-else) favorites({{ value.favorite }})
      b-button.ml-2(v-if="isSignin" variant="outline-secondary" size="sm") comments({{ value.comment }})
      span.ml-2(v-else) comments({{ value.comment }})
    div(style="height: 512px; overflow: hidden;" v-on:click="showModal")
      b-img(
        fluid
        :src="value.imagepath")
</template>
<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface CardValue {
  username:   string,
  evalparam:  number | null,
  evalsum:    number,
  favorite:   number,
  comment:    number,
  number:     number,
  imagepath:  string,
  datetime:   string,
}

@Component({})
export default class CardImage extends Vue {
  @Prop() value!: CardValue;

  public get isSignin(): boolean {
    return this.$store.state.name !== "";
  }

  public get userRate(): number {
    if (this.value.evalparam && this.value.evalsum && this.value.evalparam > 0) {
      return Math.round(this.value.evalsum / this.value.evalparam * 10) / 10
    }
    return 0
  }

  public showModal(e: Event) {
    this.$emit("showmodal", this.value);
  }
}
</script>
