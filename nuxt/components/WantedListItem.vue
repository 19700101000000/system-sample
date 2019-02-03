<template lang="pug">
  b-card(
    :border-variant="cardVariant"
    no-body)
    b-card-body
      h4 {{ value.title }}
      p {{ value.description }}
      p.card-text
        b-button(
          v-if="$store.state.name !== value.username"
          :disabled="$store.state.name === ''"
          variant="outline-primary"
          size="sm"
          v-on:click="$emit('showmodal', value)") new request
        b-button(
          v-else
          variant="outline-primary"
          size="sm") edit
        span.ml-2 JPY {{ value.price }}
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface WantedValue {
  username:    string,
  number:      number,
  title:       string,
  description: string,
  price:       number,
  alive:       boolean,
}

@Component({})
export default class WantedListItem extends Vue {
  @Prop() value: WantedValue;

  public get cardVariant(): string {
    if (this.value.alive) {
      return "success";
    }
    return "danger";
  }
}
</script>
