<template lang="pug">
  b-card(no-body)
    b-card-body
      h4 {{ value.title }}
      p {{ value.description }}
      p.card-text
        b-button.ml-2(
          v-if="$store.state.name !== value.username"
          :disabled="$store.state.name === ''"
          variant="outline-primary"
          size="sm"
          v-on:click="$emit('showmodal', value)") new request
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
}

@Component({})
export default class WantedListItem extends Vue {
  @Prop() value: WantedValue;
}
</script>
